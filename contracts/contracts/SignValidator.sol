pragma solidity ^0.5.0;

contract SignValidator {
    mapping(address => uint) public nonce;

    function validateSignature(
        uint8 sigV,
        bytes32 sigR,
        bytes32 sigS,
        bytes memory data, // original message
        address originalSigner
    ) public {
        // EIP191 convertible.
        bytes32 h = keccak256(abi.encodePacked(byte(0x19), byte(0x45), address(this), originalSigner, nonce[originalSigner], data));
        address addressFromSig = ecrecover(h, sigV, sigR, sigS);

        emit Log(addressFromSig);
        emit Log(originalSigner);
//        emit Hash(h);

//        require(addressFromSig == originalSigner);
//        nonce[addressFromSig]++;
    }

    event Log(address);
//    event Hash(bytes32);
}