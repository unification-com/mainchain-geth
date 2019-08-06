pragma solidity ^0.5.10;

import "./SafeMath.sol";

contract Beacon {

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

    // RecordHeader - called when beacon hashes are recorded. Allows querying the chain for hashes
    // without requiring to store that actual data.
    event RecordHash(
        bytes32  indexed _beaconId,
        bytes32  indexed _hashId, // generated using beacon ID, hash Num and Timestamp
        bytes32  indexed _hash,
        uint256  _hashNum,
        uint     _timestamp,
        address _submittedBy
    );

    event RegisterBeacon(
        bytes32  indexed _beaconId,
        uint    _timestamp
    );

    event LogFallbackFunctionCalled(
        address indexed _from,
        uint256 _amount
    );

    event AuthoriseNewAddress(
        bytes32  indexed _beaconId,
        address  indexed _authorisedBy,
        address _address
    );

    event RemoveAuthorisedAddress(
        bytes32  indexed _beaconId,
        address  indexed _removedBy,
        address _address
    );

    event BeaconOwnerChanged(
        bytes32  indexed _beaconId,
        address _old,
        address _new
    );

    /**
    * ----------
    * STRUCTURES
    * ----------
    */

    /**
    * @dev Struct to hold an individual Beacon's high level data, and latest block hash submission
    */
    struct BeaconData {
        bool isBeacon; // used in require statements to check if Beacon exists
        mapping(address => bool) authAddresses; // Current wallet addresses allowed to write hashes
        uint256 numAuthAddresses; // counter for number of authorised addresses. See MAX_AUTH_ORACLE_ADDRESSES
        address owner; // Beacon owner address (the address that registered the Beacon)
        // Owner can add/remove authorised Oracle addresses.
        uint256 hashNumIdx; // Counter for number of hashes submitted
        bytes32 lastHash; // Latest beacon hash submitted
        bytes32 lastHashId; // Latest beacon generated hash ID
        uint lastSubmissionTime;
        address lastSubmittedBy; // Latest address to submit the data
    }

    /**
    * ---------
    * VARIABLES
    * ---------
    */

    /*
    * Mapping of BeaconID -> Beacon Struct
    */
    mapping(bytes32 => BeaconData) public beaconList;

    /**
    * ---------
    * MODIFIERS
    * ---------
    */

    /*
    * @dev Modifier to ensure only current authorised Beacon addresses can execute a function
    * @param _beaconId the Beacon's ID
    */
    modifier onlyAuth(bytes32 _beaconId) {
        require(beaconList[_beaconId].isBeacon, "BeaconID does not exist");
        require(beaconList[_beaconId].authAddresses[msg.sender] == true || msg.sender == beaconList[_beaconId].owner, "not authorised");
        _;
    }

    /*
    * @dev Modifier to ensure only the Beacon owner can execute a function
    * @param _beaconId the Beacon's ID
    */
    modifier onlyBeaconOwner(bytes32 _beaconId) {
        require(beaconList[_beaconId].isBeacon, "BeaconID does not exist");
        require(msg.sender == beaconList[_beaconId].owner, "You are not the owner");
        _;
    }

    /*
    * @dev Modifier to ensure supplied Chain ID is not null
    * @param _beaconId the Beacon's ID
    */
    modifier beaconIDNotNull(bytes32 _beaconId) {
        require(_beaconId != btsNull, "Beacon ID required");
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
    * @dev Register a new Beacon.
    * @param _beaconId - the Beacon ID
    * @param _authAddresses - list of addresses authorised to submit hashes via recordHeader
    */
    function registerBeacon(
        bytes32 _beaconId,
        address[] memory _authAddresses) public beaconIDNotNull(_beaconId) {

        require(!beaconList[_beaconId].isBeacon, "Beacon ID already exists");
        require(_authAddresses.length > 0, "Initial Authorised addresses required");
        require(_authAddresses.length <= 10, "Maximum 10 Authorised addresses allowed");

        beaconList[_beaconId] = BeaconData({
            isBeacon: true,
            numAuthAddresses: _authAddresses.length,
            owner: msg.sender,
            hashNumIdx: 0,
            lastHash: 0x0,
            lastHashId: 0x0,
            lastSubmissionTime: 0,
            lastSubmittedBy: address(0)
            });

        emit RegisterBeacon(_beaconId, block.timestamp);

        for (uint i=0; i< _authAddresses.length; i++) {
            beaconList[_beaconId].authAddresses[_authAddresses[i]] = true;
            emit AuthoriseNewAddress(_beaconId, msg.sender, _authAddresses[i]);
        }

        // Check the owner is also authorised, in case their address wasn't included in _authAddresses
        if (!beaconList[_beaconId].authAddresses[msg.sender]) {
            beaconList[_beaconId].authAddresses[msg.sender] = true;
            beaconList[_beaconId].numAuthAddresses.add(1);
            emit AuthoriseNewAddress(_beaconId, msg.sender, msg.sender);
        }
    }

    /*
    * @dev Record a Beacon's hash.
    * A hash should only be submitted once. Block data can't be overwritten, but
    * historical data can be submitted as secondary submissions. These are recorded
    * as RecordHeaderHistorical events. If the same block number has been recorded
    * using RecordHeader AND RecordHeaderHistorical, then the data in RecordHeaderHistorical
    * should be ignored by front end implementations.
    * @param _beaconId - the Beacon ID
    * @param _hash - hash being submitted for the Beacon - usually a Merkle root.
    */
    function recordHash(
        bytes32 _beaconId,
        bytes32 _hash) public onlyAuth(_beaconId) beaconIDNotNull(_beaconId) {

        require(_hash.length > 0 && _hash != btsNull, "hash required");

        require(beaconList[_beaconId].isBeacon, "Beacon ID does not exist");

        uint timestamp = block.timestamp;

        uint256 hashNumIdx = beaconList[_beaconId].hashNumIdx.add(1);

        bytes32 hashId = keccak256(abi.encode(_beaconId, hashNumIdx, timestamp));

        beaconList[_beaconId].hashNumIdx = hashNumIdx;
        beaconList[_beaconId].lastHash = _hash;
        beaconList[_beaconId].lastHashId = hashId;
        beaconList[_beaconId].lastSubmissionTime = timestamp;
        beaconList[_beaconId].lastSubmittedBy = msg.sender;

        emit RecordHash(
            _beaconId,
             hashId,
            _hash,
            hashNumIdx,
            timestamp,
            msg.sender
        );
    }

    /**
    * ----------------------
    * BEACON ADMIN FUNCTIONS
    * ----------------------
    */

    /*
    * @dev Authorise a new wallet to write hashes for the Beacon.
    *      Only Beacon owner can all this
    * @param _beaconId - Beacon ID
    * @param _newAuthAddresses - array of one or more new addresses that can write hashes
    */
    function addAuthAddresses(
        bytes32 _beaconId,
        address[] memory _newAuthAddresses
    ) public onlyBeaconOwner(_beaconId) beaconIDNotNull(_beaconId) {
        require(_newAuthAddresses.length > 0, "Auth addresses required");
        require(_newAuthAddresses.length <= 10, "Max submission 10 addresses allowed");
        require(beaconList[_beaconId].isBeacon, "Chain ID does not exist");

        uint256 newTotal =  beaconList[_beaconId].numAuthAddresses.add(_newAuthAddresses.length);
        require(newTotal <= MAX_AUTH_ORACLE_ADDRESSES, "Max 10 authorised addresses allowed");

        for (uint j=0; j< _newAuthAddresses.length; j++) {
            require(_newAuthAddresses[j] != address(0), "cannot authorise zero address");
            beaconList[_beaconId].authAddresses[_newAuthAddresses[j]] = true;
            beaconList[_beaconId].numAuthAddresses.add(1);
            emit AuthoriseNewAddress(_beaconId, msg.sender, _newAuthAddresses[j]);
        }
    }

    /*
    * @dev Remove a wallet from the list of authorised addresses that can write hashes for the Beacon.
    *      Only Beacon owner can all this
    * @param _beaconId - Beacon ID
    * @param _addressesToRemove - array of one or more new addresses to remove
    */
    function removeAuthAddresses(
        bytes32 _beaconId,
        address[] memory _addressesToRemove
    ) public onlyBeaconOwner(_beaconId) beaconIDNotNull(_beaconId) {
        require(_addressesToRemove.length > 0, "Auth addresses required");
        require(_addressesToRemove.length <= 10, "Max submission 10 addresses allowed");
        require(beaconList[_beaconId].isBeacon, "Chain ID does not exist");

        require(_addressesToRemove.length < beaconList[_beaconId].numAuthAddresses, "at least one remaining authorised address required");

        for (uint j=0; j< _addressesToRemove.length; j++) {
            require(_addressesToRemove[j] != address(0), "cannot authorise zero address");
            delete beaconList[_beaconId].authAddresses[_addressesToRemove[j]];
            beaconList[_beaconId].numAuthAddresses.sub(1);
            emit RemoveAuthorisedAddress(_beaconId, msg.sender, _addressesToRemove[j]);
        }
    }

    /*
    * @dev Change Beacon ownership to a new wallet address.
    *      Only Beacon owner can all this
    * @param _beaconId - Beacon ID
    * @param _newOwner - wallet address of new owner
    */
    function changeBeaconOwner(
        bytes32 _beaconId,
        address _newOwner
    ) public onlyBeaconOwner(_beaconId) beaconIDNotNull(_beaconId) {
        require(_newOwner != address(0), "New owner required, and should be non-zero address");
        require(beaconList[_beaconId].isBeacon, "Chain ID does not exist");

        beaconList[_beaconId].owner = _newOwner;
        emit BeaconOwnerChanged(_beaconId, msg.sender, _newOwner);

        // Check the new owner is also authorised
        if (!beaconList[_beaconId].authAddresses[_newOwner]) {
            beaconList[_beaconId].authAddresses[_newOwner] = true;
            beaconList[_beaconId].numAuthAddresses.add(1);
            emit AuthoriseNewAddress(_beaconId, msg.sender, _newOwner);
        }
    }

    /**
    * -------
    * GETTERS
    * -------
    */

    // get latest hashes
    function getLastHash(bytes32 _beaconId) public beaconIDNotNull(_beaconId) view returns(
        uint hashNumIdx_,
        bytes32 lastHashId_,
        bytes32 lastHash_,
        uint lastSubmissionTime_,
        address lastSubmittedBy_) {

        require(beaconList[_beaconId].isBeacon, "Chain ID does not exist");

        BeaconData storage my_beacon = beaconList[_beaconId];

        hashNumIdx_ = my_beacon.hashNumIdx;
        lastHashId_ = my_beacon.lastHashId;
        lastHash_ = my_beacon.lastHash;
        lastSubmissionTime_ = my_beacon.lastSubmissionTime;
        lastSubmittedBy_ = my_beacon.lastSubmittedBy;
    }

    /**
    * --------------
    * HELPER QUERIES
    * --------------
    */

    function isAuthorisedAddress(bytes32 _beaconId, address _address) public beaconIDNotNull(_beaconId) view returns (bool isAuthorised_) {
        require(beaconList[_beaconId].isBeacon, "Chain ID does not exist");
        isAuthorised_ = beaconList[_beaconId].authAddresses[_address];
    }

    function beaconExists(bytes32 _beaconId) public beaconIDNotNull(_beaconId) view returns (bool beaconExists_) {
        beaconExists_ = beaconList[_beaconId].isBeacon;
    }

}
