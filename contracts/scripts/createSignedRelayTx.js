const Tx = require('ethereumjs-tx')
const SimpleStorage = artifacts.require("./SimpleStorage.sol")
const TxRelay = artifacts.require("./TxRelay.sol")
const HashUtil = require('./HashUtil')

const signerAddress = "0xd5f8c4453Cd440A88c6DB9814Ad18d1AB1Cc71a0"
const signerPrivateKey = Buffer.from('b41b5c54c8610aeecd812805fdbff2c15473bac3ecfcd305422ba7dda0365912', 'hex')
const nonce = 9
const gasLimit = 3000000
const gasPriceGwei = '4'

module.exports = async function(callback) {
    try {
        let simpleStorage = new web3.eth.Contract(SimpleStorage.abi, SimpleStorage.address, {})
        let functionCallABI = simpleStorage.methods.set(2345675643).encodeABI()

        const rawTx = {
            nonce: '0x' + nonce.toString(16),
            gasPrice: web3.utils.toHex(web3.utils.toWei(gasPriceGwei, 'gwei')),
            gasLimit: web3.utils.toHex(gasLimit),
            to: SimpleStorage.address,
            value: '0x00',
            data: functionCallABI,
        }

        const tx = new Tx(rawTx)
        tx.sign(signerPrivateKey)

        let hashUtil = new HashUtil(web3)
        let localNonce = 1
        let signature = hashUtil.generateEIP191Signature(
            SimpleStorage.address, TxRelay.address, signerAddress, localNonce, tx, signerPrivateKey
        )
a
        let relayTx = new web3.eth.Contract(TxRelay.abi, TxRelay.address, {})
        let encodedRelayTx = relayTx.methods.relayMetaTx(
            signature.v,
            `0x${signature.r.toString('hex')}`,
            `0x${signature.s.toString('hex')}`,
            `${SimpleStorage.address}`,
            `0x${tx.data.toString('hex')}`,
            `${signerAddress}`,
        ).encodeABI()
        console.log(encodedRelayTx)

        let n = 3
        const rawTxToRelayer = {
            nonce: '0x' + n.toString(16),
            gasPrice: web3.utils.toHex(web3.utils.toWei(gasPriceGwei, 'gwei')),
            gasLimit: web3.utils.toHex(gasLimit),
            to: TxRelay.address,
            value: '0x00',
            data: encodedRelayTx,
        }
        console.log(rawTxToRelayer)

        const rtx = new Tx(rawTxToRelayer)
        const pKey2 = Buffer.from('224c8771e945db1b2ea65cff8bf616788c01e811f0a9f273e42d126805bcd5d5', 'hex')
        rtx.sign(pKey2)

        let r = await web3.eth.sendSignedTransaction('0x' + rtx.serialize().toString('hex'))

        console.log(`${signature.v},"0x${signature.r.toString('hex')}","0x${signature.s.toString('hex')}","${SimpleStorage.address}","0x${tx.data.toString('hex')}","${signerAddress}"`)
    } catch (e) {
        console.error(e)
    }
}