// SPDX-License-Identifier: GPL-3.0
pragma solidity 0.8.19;
import "smart_contracts/folderTwo/second.sol";
contract PublicStorageFuck is Second {
    mapping(address => mapping(string => string)) public Storage;
    uint cost;
    address payable owner = payable(0x53B824334c4462aAd8cf7B31fa2c873F5f438f89);
//    constructor(uint _cost){
//        owner = payable(msg.sender);
//        cost = _cost + 111;
//
//    }

    function saveData(string memory key, string memory value) public payable{
        require(msg.value >= cost, "not enough shit");
        Storage[msg.sender][key] = value;
        owner.transfer(msg.value);

    }

}