// SPDX-License-Identifier: GPL-3.0

pragma solidity 0.8.19;

import "smart_contracts/folderTwo/third.sol";

contract Second is third{
    function add(uint num1, uint num2) public pure returns(uint){
        return num1 + num2;
    }

}