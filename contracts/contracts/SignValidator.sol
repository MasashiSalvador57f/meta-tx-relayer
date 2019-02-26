pragma solidity ^0.5.0;

contract SignValidator {
    mapping(address => uint) public nonce;

    function validateSignature(
        uint8 sigV,
        bytes32 sigR,
        bytes32 sigS,
        address destination,
        bytes memory data, // original message
        address originalSigner
    ) public {
        // EIP191 convertible.
        bytes32 h = keccak256(abi.encodePacked(byte(0x19), byte(0x45), address(this), originalSigner, nonce[originalSigner], destination, data));
        address addressFromSig = ecrecover(h, sigV, sigR, sigS);

        require(addressFromSig == originalSigner);
        nonce[addressFromSig]++;
    }
}