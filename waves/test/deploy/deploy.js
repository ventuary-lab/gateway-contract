const wvs = 10 ** 8;

describe('Deploy script', async function () {
    before(async function () {
        await setupAccounts({
            contract: 100000 * wvs
        })

    });
    it('Deploy contract', async function () {
        const setScriptTx = setScript({ script: compile(file("../script/gateway.ride")), fee: 1400000,}, accounts.contract); 

        const signedIssueTx = issue({
            name: 'TToken',
            description: 'Test token',
            quantity: "10000000000000000000000000",
            decimals: 8
        }, accounts.contract)

        await broadcast(signedIssueTx)
        await broadcast(setScriptTx)

        const constructorData = data({
            data: [
                { key: "asset_id", value: signedIssueTx.id },
                { key: "admins", value: "AXbaBkJNocyrVpwqTzD4TpUY8fQ6eeRto9k1m2bNCzXV,AXbaBkJNocyrVpwqTzD4TpUY8fQ6eeRto9k1m2bNCzXV" +
                ",AXbaBkJNocyrVpwqTzD4TpUY8fQ6eeRto9k1m2bNCzXV,AXbaBkJNocyrVpwqTzD4TpUY8fQ6eeRto9k1m2bNCzXV,AXbaBkJNocyrVpwqTzD4TpUY8fQ6eeRto9k1m2bNCzXV"},
                
                { key: "bftCoefficient", value: 1}
            ],
            fee: 500000
        }, accounts.contract);
        await broadcast(constructorData)

        console.log(accounts.contract)
    })
})