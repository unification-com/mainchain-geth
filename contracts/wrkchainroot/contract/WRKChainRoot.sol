pragma solidity ^0.5.10;

import "./SafeMath.sol";

contract WRKChainRoot {

    // because it'd be daft not to
    using SafeMath for uint256;

    /**
    * ---------
    * CONSTANTS
    * ---------
    */

    uint constant MAX_AUTH_ORACLE_ADDRESSES = 100;
    bytes32 constant btsNull = "";

    /**
    * ------
    * EVENTS
    * ------
    */

    // RecordHeader - called when block hashes are recorded. Allows querying the chain for hashes
    // without requiring to store that actual data.
    event RecordHeader(
        bytes32  indexed _chainId,
        uint256  indexed _height,
        bytes32  indexed _hash,
        bytes32 _parentHash,
        bytes32 _receiptRoot,
        bytes32 _txRoot,
        bytes32 _stateRoot,
        uint    _timestamp,
        address _submittedBy
    );

    // RecordHeaderHistorical - allow WRKChains to submit historical data - i.e. block hashes for
    // block numbers lower than the last recent block submission. Front ends should also
    // monitor this event log, but if the same block number has already been recorded and the
    // RecordHeader already emitted, then its corresponding RecordHeaderHistorical should
    // be ignored.
    event RecordHeaderHistorical(
        bytes32  indexed _chainId,
        uint256  indexed _height,
        bytes32  indexed _hash,
        bytes32 _parentHash,
        bytes32 _receiptRoot,
        bytes32 _txRoot,
        bytes32 _stateRoot,
        uint    _timestamp,
        address _submittedBy
    );

    event RegisterWrkChain(
        bytes32  indexed _chainId,
        bytes32  _genesisHash,
        uint    _timestamp
    );

    event LogFallbackFunctionCalled(
        address indexed _from,
        uint256 _amount
    );

    event AuthoriseNewAddress(
        bytes32  indexed _chainId,
        address  indexed _authorisedBy,
        address _address
    );

    event RemoveAuthorisedAddress(
        bytes32  indexed _chainId,
        address  indexed _removedBy,
        address _address
    );

    event WRKChainOwnerChanged(
        bytes32  indexed _chainId,
        address _old,
        address _new
    );

    /**
    * ----------
    * STRUCTURES
    * ----------
    */

    /**
    * @dev Struct to hold an individual WRKChain's high level data, and latest block hash submission
    */
    struct Wrkchain {
        uint256 lastBlock; // Last block num submitted
        bytes32 genesisHash; // Genesis Block's hash
        bool isWrkchain; // used in require statements to check if WRKChain exists
        mapping(address => bool) authAddresses; // Current wallet addresses allowed to write hashes
        uint256 numAuthAddresses; // counter for number of authorised addresses. See MAX_AUTH_ORACLE_ADDRESSES
        address owner; // WRKChain owner address (the address that registered the WRKChain)
                       // Owner can add/remove authorised Oracle addresses.
        uint256 numBlocksSubmitted; // Counter for number of blocks submitted
        uint256 numHistoricalBlocksSubmitted; // Counter for number of historical blocks submitted
        bytes32 lastHash; // Latest block hash submitted
        bytes32 lastParentHash; // Latest Parent Block Hash submitted
        bytes32 lastReceiptRoot; // Latest Receipt merkle root submitted (can be repurposed for any arbitrary hash)
        bytes32 lastTxRoot; // Latest Tx merkle root submitted (can be repurposed for any arbitrary hash)
        bytes32 lastStateRoot; // Latest StateDb merkle root submitted (can be repurposed for any arbitrary hash)
        uint lastSubmissionTime;
        address lastSubmittedBy; // Latest address to submit the data
    }

    /**
    * ---------
    * VARIABLES
    * ---------
    */

    /*
    * Mapping of ChainID -> Wrkchain Struct
    */
    mapping(bytes32 => Wrkchain) public wrkchainList;

    /**
    * ---------
    * MODIFIERS
    * ---------
    */

    /*
    * @dev Modifier to ensure only current authorised WRKChain addresses can execute a function
    * @param _chainId the WRKChain network ID
    */
    modifier onlyAuth(bytes32 _chainId) {
        require(wrkchainList[_chainId].isWrkchain, "ChainID does not exist");
        require(wrkchainList[_chainId].authAddresses[msg.sender] == true || msg.sender == wrkchainList[_chainId].owner, "not authorised");
        _;
    }

    /*
    * @dev Modifier to ensure only the WRKChain owner can execute a function
    * @param _chainId the WRKChain network ID
    */
    modifier onlyWrkchainOwner(bytes32 _chainId) {
        require(wrkchainList[_chainId].isWrkchain, "ChainID does not exist");
        require(msg.sender == wrkchainList[_chainId].owner, "You are not the owner");
        _;
    }

    /*
    * @dev Modifier to ensure supplied Chain ID is not null
    * @param _chainId the WRKChain ID
    */
    modifier chainIdNotNull(bytes32 _chainId) {
        require(_chainId != btsNull, "Chain ID required");
        _;
    }

    /**
    * -------------
    * MAIN CONTRACT
    * -------------
    **/


    /*
    * @dev contract constructor
    */
    constructor() public { }

    /*
    * @dev fallback function for the contract. Log event so UND can be tracked and returned
    */
    function() payable external {
        //Log who sent, and how much. Not much more can be done.
        emit LogFallbackFunctionCalled(msg.sender, msg.value);
    }

    /*
    * @dev Register a new WRKChain.
    * @param _chainId - the WRKChain Network ID
    * @param _authAddresses - list of addresses authorised to submit hashes via recordHeader
    * @param _genesisHash - WRKChain's hashed Genesis block
    */
    function registerWrkChain(
        bytes32 _chainId,
        address[] memory _authAddresses,
        bytes32 _genesisHash) public chainIdNotNull(_chainId) {

        require(!wrkchainList[_chainId].isWrkchain, "Chain ID already exists");
        require(_genesisHash.length > 0 && _genesisHash != btsNull, "Genesis hash required");
        require(_authAddresses.length > 0, "Initial Authorised addresses required");
        require(_authAddresses.length <= 10, "Maximum 10 Authorised addresses allowed");

        wrkchainList[_chainId] = Wrkchain({
            lastBlock: 0,
            genesisHash: _genesisHash,
            isWrkchain: true,
            numAuthAddresses: _authAddresses.length,
            owner: msg.sender,
            numBlocksSubmitted: 0,
            numHistoricalBlocksSubmitted: 0,
            lastHash: _genesisHash,
            lastParentHash: 0x0,
            lastReceiptRoot: 0x0,
            lastTxRoot: 0x0,
            lastStateRoot: 0x0,
            lastSubmissionTime: block.timestamp,
            lastSubmittedBy: msg.sender
        });

        emit RegisterWrkChain(_chainId, _genesisHash, block.timestamp);

        for (uint i=0; i< _authAddresses.length; i++) {
            wrkchainList[_chainId].authAddresses[_authAddresses[i]] = true;
            emit AuthoriseNewAddress(_chainId, msg.sender, _authAddresses[i]);
        }

        // Check the owner is also authorised, in case their address wasn't included in _authAddresses
        if (!wrkchainList[_chainId].authAddresses[msg.sender]) {
            wrkchainList[_chainId].authAddresses[msg.sender] = true;
            wrkchainList[_chainId].numAuthAddresses.add(1);
            emit AuthoriseNewAddress(_chainId, msg.sender, msg.sender);
        }
    }

    /*
    * @dev Record a WRKChain's block hash(es).
    * A block should only be submitted once. Block data can't be overwritten, but
    * historical data can be submitted as secondary submissions. These are recorded
    * as RecordHeaderHistorical events. If the same block number has been recorded
    * using RecordHeader AND RecordHeaderHistorical, then the data in RecordHeaderHistorical
    * should be ignored by front end implementations.
    * @param _chainId - the WRKChain Network ID
    * @param _height - Block number being submitted
    * @param _hash - WRKChain's main header hash
    * @param _parentHash - optional Header hash of previous block
    * @param _receiptRoot - optional Merkle root of WRKChain's receipts (can be repurposed for any arbitrary hash)
    * @param _txRoot - optional Merkle root of WRKChain's Txs (can be repurposed for any arbitrary hash)
    * @param _stateRoot - optional Merkle root of WRKChain's state DB (can be repurposed for any arbitrary hash)
    */
    function recordHeader(
        bytes32 _chainId,
        uint256 _height,
        bytes32 _hash,
        bytes32 _parentHash,
        bytes32 _receiptRoot,
        bytes32 _txRoot,
        bytes32 _stateRoot) public onlyAuth(_chainId) chainIdNotNull(_chainId) {

        require(_height > 0, "Block number required");
        require(_hash.length > 0 && _hash != btsNull, "hash required");

        require(wrkchainList[_chainId].isWrkchain, "Chain ID does not exist");

        Wrkchain storage my_wrkchain = wrkchainList[_chainId];

        if(_height > wrkchainList[_chainId].lastBlock) {
            wrkchainList[_chainId].numBlocksSubmitted = wrkchainList[_chainId].numBlocksSubmitted.add(1);
            wrkchainList[_chainId].lastBlock = _height;
            wrkchainList[_chainId].lastHash = _hash;
            wrkchainList[_chainId].lastParentHash = _parentHash;
            wrkchainList[_chainId].lastReceiptRoot = _receiptRoot;
            wrkchainList[_chainId].lastTxRoot = _txRoot;
            wrkchainList[_chainId].lastStateRoot = _stateRoot;
            wrkchainList[_chainId].lastSubmissionTime = block.timestamp;
            wrkchainList[_chainId].lastSubmittedBy = msg.sender;

            emit RecordHeader(
                    _chainId,
                    _height,
                    _hash,
                    _parentHash,
                    _receiptRoot,
                    _txRoot,
                    _stateRoot,
                    block.timestamp,
                    msg.sender
            );
        } else {
            // submitting a block < the current highest submitted block number.
            // Allow block to be recorded, but as a secondary RecordHeaderHistorical event.
            // This offers a convenient method for WRKChains to "fill in" historical data
            // should they wish.
            wrkchainList[_chainId].numHistoricalBlocksSubmitted = wrkchainList[_chainId].numHistoricalBlocksSubmitted.add(1);
            emit RecordHeaderHistorical(
                    _chainId,
                    _height,
                    _hash,
                    _parentHash,
                    _receiptRoot,
                    _txRoot,
                    _stateRoot,
                    block.timestamp,
                    msg.sender
            );
        }
    }

    /**
    * ------------------------
    * WRKCHAIN ADMIN FUNCTIONS
    * ------------------------
    */

    /*
    * @dev Authorise a new wallet to write hashes for the WRKChain.
    *      Only WRKChain owner can all this
    * @param _chainID - WRKChain network ID
    * @param _newAuthAddresses - array of one or more new addresses that can write hashes
    */
    function addAuthAddresses(
        bytes32 _chainId,
        address[] memory _newAuthAddresses
    ) public onlyWrkchainOwner(_chainId) chainIdNotNull(_chainId) {
        require(_newAuthAddresses.length > 0, "Auth addresses required");
        require(_newAuthAddresses.length <= 10, "Max submission 10 addresses allowed");
        require(wrkchainList[_chainId].isWrkchain, "Chain ID does not exist");

        uint256 newTotal =  wrkchainList[_chainId].numAuthAddresses.add(_newAuthAddresses.length);
        require(newTotal <= MAX_AUTH_ORACLE_ADDRESSES, "Max 10 authorised addresses allowed");

        for (uint j=0; j< _newAuthAddresses.length; j++) {
            require(_newAuthAddresses[j] != address(0), "cannot authorise zero address");
            wrkchainList[_chainId].authAddresses[_newAuthAddresses[j]] = true;
            wrkchainList[_chainId].numAuthAddresses.add(1);
            emit AuthoriseNewAddress(_chainId, msg.sender, _newAuthAddresses[j]);
        }
    }

    /*
    * @dev Remove a wallet from the list of authorised addresses that can write hashes for the WRKChain.
    *      Only WRKChain owner can all this
    * @param _chainID - WRKChain network ID
    * @param _addressesToRemove - array of one or more new addresses to remove
    */
    function removeAuthAddresses(
        bytes32 _chainId,
        address[] memory _addressesToRemove
    ) public onlyWrkchainOwner(_chainId) chainIdNotNull(_chainId) {
        require(_addressesToRemove.length > 0, "Auth addresses required");
        require(_addressesToRemove.length <= 10, "Max submission 10 addresses allowed");
        require(wrkchainList[_chainId].isWrkchain, "Chain ID does not exist");

        require(_addressesToRemove.length < wrkchainList[_chainId].numAuthAddresses, "at least one remaining authorised address required");

        for (uint j=0; j< _addressesToRemove.length; j++) {
            require(_addressesToRemove[j] != address(0), "cannot authorise zero address");
            delete wrkchainList[_chainId].authAddresses[_addressesToRemove[j]];
            wrkchainList[_chainId].numAuthAddresses.sub(1);
            emit RemoveAuthorisedAddress(_chainId, msg.sender, _addressesToRemove[j]);
        }
    }

    /*
    * @dev Change WRKChain ownership to a new wallet address.
    *      Only WRKChain owner can all this
    * @param _chainID - WRKChain network ID
    * @param _newOwner - wallet address of new owner
    */
    function changeWrkchainOwner(
        bytes32 _chainId,
        address _newOwner
    ) public onlyWrkchainOwner(_chainId) chainIdNotNull(_chainId) {
        require(_newOwner != address(0), "New owner required, and should be non-zero address");
        require(wrkchainList[_chainId].isWrkchain, "Chain ID does not exist");

        Wrkchain storage my_wrkchain = wrkchainList[_chainId];

        wrkchainList[_chainId].owner = _newOwner;
        wrkchainList[_chainId].authAddresses[_newOwner] = true;

        emit WRKChainOwnerChanged(_chainId, msg.sender, _newOwner);
    }

    /**
    * -------
    * GETTERS
    * -------
    */

    //get the Genesis hash for a WRKChain
    function getGenesis(bytes32 _chainId) public chainIdNotNull(_chainId) view returns (bytes32 genesisHash_) {
        require(wrkchainList[_chainId].isWrkchain, "Chain ID does not exist");

        Wrkchain storage my_wrkchain = wrkchainList[_chainId];

        genesisHash_ = my_wrkchain.genesisHash;
    }

    // get last block number recorded for a WRKChain
    function getLastRecordedBlockNum(bytes32 _chainId) public chainIdNotNull(_chainId) view returns (uint256 lastBlock_) {
        require(wrkchainList[_chainId].isWrkchain, "Chain ID does not exist");

        Wrkchain storage my_wrkchain = wrkchainList[_chainId];

        lastBlock_ = my_wrkchain.lastBlock;
    }

    //get number of blocks submitted
    function getNumBlocksSubmitted(bytes32 _chainId) public chainIdNotNull(_chainId) view returns (uint256 numBlocksSubmitted_) {
        require(wrkchainList[_chainId].isWrkchain, "Chain ID does not exist");

        Wrkchain storage my_wrkchain = wrkchainList[_chainId];

        numBlocksSubmitted_ = my_wrkchain.numBlocksSubmitted;
    }

    // get latest block hashes
    function getLastBlock(bytes32 _chainId) public chainIdNotNull(_chainId) view returns(bytes32 lastHash_,
                                                                bytes32 lastParentHash_,
                                                                bytes32 lastReceiptRoot_,
                                                                bytes32 lastTxRoot_,
                                                                bytes32 lastStateRoot_,
                                                                uint lastSubmissionTime_,
                                                                address lastSubmittedBy_) {

        require(wrkchainList[_chainId].isWrkchain, "Chain ID does not exist");

        Wrkchain storage my_wrkchain = wrkchainList[_chainId];

        lastHash_ = my_wrkchain.lastHash;
        lastParentHash_ = my_wrkchain.lastParentHash;
        lastReceiptRoot_ = my_wrkchain.lastReceiptRoot;
        lastTxRoot_ = my_wrkchain.lastTxRoot;
        lastStateRoot_ = my_wrkchain.lastStateRoot;
        lastSubmissionTime_ = my_wrkchain.lastSubmissionTime;
        lastSubmittedBy_ = my_wrkchain.lastSubmittedBy;
    }

    /**
    * --------------
    * HELPER QUERIES
    * --------------
    */

    function isAuthorisedAddress(bytes32 _chainId, address _address) public chainIdNotNull(_chainId) view returns (bool isAuthorised_) {
        require(wrkchainList[_chainId].isWrkchain, "Chain ID does not exist");
        isAuthorised_ = wrkchainList[_chainId].authAddresses[_address];
    }

    function chainExists(bytes32 _chainId) public chainIdNotNull(_chainId) view returns (bool chainExists_) {
        chainExists_ = wrkchainList[_chainId].isWrkchain;
    }

}
