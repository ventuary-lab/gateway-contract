const converter = require("./helper/converter.js")
const enums = require("./helper/enums.js")
const types = require("./helper/types.js")
const Supersymmetry = artifacts.require("Supersymmetry");

contract("Mint test", async accounts => {
    let susyContract = null;
    let requestHash = "";
    let amount = "10000000000000000000"

    before(async () => {
        susyContract = await Supersymmetry.deployed([accounts[0],accounts[1],accounts[2],accounts[3],accounts[4]]);
    });

    async function mint(){
        let result = await susyContract.createMintRequest("sender", amount, {from: accounts[5] })
        let requestHash = result.logs[0].args.requestHash
        const request = await susyContract.requests.call(requestHash);

        assert.equal(request.status, enums.Supersymmetry.New.slice(3,4));
        assert.equal(request.rType, types.Supersymmetry.Mint);
        assert.equal(request.owner, accounts[5]);
        assert.equal(request.target, "sender");
        assert.equal(request.tokenAmount, amount);
        return requestHash
    }
    
    async function vote(requestHash, status){
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
        assert.equal(request.rType, types.Supersymmetry.Mint);
        assert.equal(request.owner, accounts[5]);
        assert.equal(request.target, "sender");
        assert.equal(request.tokenAmount, amount);

        assert.equal(res.logs[0].event, "StatusChanged");
        assert.equal(res.logs[0].args.status, status.slice(3,4));
        assert.equal(res.logs[0].args.requestHash, requestHash);
    }

    it("Mint", async () => {
        requestHash = await mint();
    });
    
    it("Confirm", async () => {
        await vote(requestHash, enums.Supersymmetry.Success)
    });


    it("Mint #2", async () => {
        requestHash = await mint();
    });

    it("Reject", async () => {
        await vote(requestHash, enums.Supersymmetry.Rejected)
    });
});