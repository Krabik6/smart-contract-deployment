// SPDX-License-Identifier: GPL-3.0
pragma solidity 0.8.19;
import "second.sol";

contract PublicStorageFuck is Second {
    mapping(address => mapping(string => string)) public Storage;
    uint cost;
    address payable owner;
    constructor(uint _cost){
        owner = payable(msg.sender);
        cost = _cost + 11;

    }

    function saveData(string memory key, string memory value) public payable{
        require(msg.value >= cost, "not enough shit");
        Storage[msg.sender][key] = value;
        owner.transfer(msg.value);

    }

}