// SPDX-License-Identifier: GPL-3.0
pragma solidity 0.8.19;
import "smart_contracts/folderTwo/second.sol";
contract PublicStorageFuck is Second {
    mapping(address => mapping(string => string)) public Storage;
    uint cost = 145890;
    address payable owner = payable(0x204a167f93d8f02758448979E3EA1A8ca41b6c5E);
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