pragma solidity ^0.5.0;

contract SampleContract {
    User[] public allowed;

    // Data parameter is only a sample
    struct User {
    	address Addr;
	bytes32 Data;
    }
    

    function addAllowed(
    	address _addr, bytes32 _data
    ) external {
    	// verification here

        allowed.push(User(_addr, _data));
    }
}
