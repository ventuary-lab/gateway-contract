const converter = require("./helper/converter.js")
const enums = require("./helper/enums.js")
const types = require("./helper/types.js")
const Supersymmetry = artifacts.require("Supersymmetry");
const ERC20 = artifacts.require("Supersymmetry");

contract("Burn confirm test", async accounts => {
    let susyContract = null;
    
    before(async () => {
        susyContract = await Supersymmetry.deployed([accounts[0],accounts[1],accounts[2],accounts[3],accounts[4]]);
    });

    
    let requestHash = ""
    let amount = "10000000000000000000"

    it("Mint", async () => {
        let result = await susyContract.createBurnRequest("sender", amount, {from: accounts[5] })
        requestHash = result.logs[0].args.requestHash
        const request = await susyContract.requests.call(requestHash);
        assert.equal(request.status, "1");
        assert.equal(request.rType, "1");
        assert.equal(request.owner, accounts[5]);
        assert.equal(request.target, "sender");
        assert.equal(request.tokenAmount, amount);
    });
    it("Confirm", async () => {
        let status = enums.Supersymmetry.Confirm
        let msg = requestHash + status.slice(2,4)
        let v = []
        let r = []
        let s = []
        for (let i = 0; i < 5; i++) {
            let sig = await web3.eth.sign(msg, accounts[i])
            let sigModel = converter.splitSign(sig)
            v.push(sigModel.v)
            r.push(sigModel.r)
            s.push(sigModel.s)
        }
        let res = await susyContract.changeStatus(requestHash, v, r, s, status, {from: accounts[0] })
        const request = await susyContract.requests.call(requestHash);
        assert.equal(request.status, status.slice(3,4));
        assert.equal(request.rType, types.Supersymmetry.Burn);
        assert.equal(request.owner, accounts[5]);
        assert.equal(request.target, "sender");
        assert.equal(request.tokenAmount, amount);

        assert.equal(res.logs[0].event, "StatusChanged");
        assert.equal(res.logs[0].args.status, status.slice(3,4));
        assert.equal(res.logs[0].args.requestHash, requestHash);
    });
});