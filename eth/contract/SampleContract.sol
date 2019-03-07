pragma solidity ^0.5.0;

contract SampleContract {
    address[] public allowed;
    
    function addAllowed(address _addr) public {
        allowed.push(_addr);
    }
}
