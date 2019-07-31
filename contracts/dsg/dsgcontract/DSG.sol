pragma solidity ^0.4.24;

contract DSGContract {
  address public x1 = 0xa13A71dfe5cD57F9b67ec6A54AD2Ae7537D7fc3b;
  address public x2 = 0xb47FD09F1d379Ce2c9BFF59D668cf0B7b994a2B7;
  address public x3 = 0xcA29F1275470DE81DE6E1Bb53a55228Da676E752;

  function rotate() external {
    if (msg.sender != 0xf30F951b0426f8Bf37ac71967407081358DF7a7B) return;

    x1 = 0xa13A71dfe5cD57F9b67ec6A54AD2Ae7537D7fc3b;
    x2 = 0xb47FD09F1d379Ce2c9BFF59D668cf0B7b994a2B7;
    x3 = 0xd71aD3263666e03004479B30CA840a26eaC5b763;
  }
}
