pragma solidity ^0.5.0;

// this contract is only for testing TxRelay.
contract SimpleStorage {
    uint myVariable;
    constructor() public {
        myVariable = 1;
    }

    function set(uint256 x) public {
        myVariable = x;
        emit Set(msg.sender, x);
    }

    function get() public view returns (uint256) {
        return myVariable;
    }
    // for debugging
    event Set(address sender, uint x);
}