const Util = require('ethereumjs-util')
const axios = require('axios');

module.exports = async function(callback) {
    try {
        let nonceEndpoint = "http://localhost:3000/auth/raw_message"
        let result = await axios.get(nonceEndpoint)

        let senderAddr = '0x5b7Cfe9978CcFf7fb9C306b65F4799a8b88175C7'
        let privateKeyStr = '224c8771e945db1b2ea65cff8bf616788c01e811f0a9f273e42d126805bcd5d5'
        let privateKey = Buffer.from(privateKeyStr, 'hex')


        let message = result.data.message_to_sign
        let nonce = result.data.auth_nonce
        console.log(message)
        console.log(nonce)

        let prefix = '0x19'
        let version = '0x45'

        let hash = web3.utils.soliditySha3(
            {t: 'bytes', v: prefix },
            {t: 'bytes', v: version },
            {t: 'address', v: senderAddr },
            {t: 'uint', v: nonce },
            {t: 'bytes', v: message },
        )

        let signature = Util.ecsign(
            Buffer.from(Util.stripHexPrefix(hash), 'hex'), Buffer.from(Util.stripHexPrefix(privateKey), 'hex')
        )

        console.log("please give the below as input to API")
        console.log(signature.r.toString('hex'))
        console.log(signature.v)
        console.log(signature.s.toString('hex'))
        console.log(`${signature.v},"0x${signature.r.toString('hex')}", "0x${signature.s.toString('hex')}", "0x${message}", ${senderAddr}`)
    } catch (e) {
        console.error(e)
    }
}