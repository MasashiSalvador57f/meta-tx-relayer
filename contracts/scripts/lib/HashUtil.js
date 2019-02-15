const Util = require('ethereumjs-util')
const leftPad = require('left-pad')

class HashUtil {
    constructor(web3Instance) {
        this.web3 = web3Instance
    }

    generateEIP191Signature(destinationAddr, relayTxAddr, signerAddr, relayNonce, signedRawTx, signerPrivateKey) {
        let prefix = '0x19'
        let version = '0x45'

        let hash = this.web3.utils.soliditySha3(
            { t: "bytes", v: prefix },
            { t: "bytes", v: version },
            { t: "address", v: relayTxAddr },
            { t: "address", v: signerAddr },
            { t: "uint", v: relayNonce },
            { t: "address", v: destinationAddr },
            { t: "bytes", v: signedRawTx.data.toString('hex') },
        )

        let signature = HashUtil._sign(hash, signerPrivateKey)

        return new EIP191Signature(signature, hash)

    }
    static _sign(msg, privateKey) {
        return Util.ecsign(
            Buffer.from(Util.stripHexPrefix(msg), 'hex'), Buffer.from(Util.stripHexPrefix(privateKey), 'hex')
        )
    }

}

class EIP191Signature {
    constructor(sig, hash) {
        this.sig = sig
        this.hash = hash
        this.v = sig.v
        this.r = sig.r
        this.s = sig.s
    }
}

module.exports = HashUtil