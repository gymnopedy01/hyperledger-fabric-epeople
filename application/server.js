const express = require('express')
const app = express()

const path = require('path')

const FabricCAServices = require("fabric-ca-client");
const { Wallets, Gateway } = require("fabric-network");
const fs = require("fs");


const { v4 } = require('uuid');

const PORT = 8080
const HOST = "0.0.0.0"

app.use(express.static(path.join(__dirname, "views")))
app.use(express.urlencoded({extended:false}))
app.use(express.json())

app.get('/', (req,res) => {
    res.sendFile(__dirname+"index.html")
})


app.post("/api/signin.json",async (req,res) => {

    const userId = req.body.userId
    const password = req.body.password

    console.log("POST /singin.json", userId, password);

    try {
        // load the network configuration
        const ccpPath = path.resolve(__dirname, "ccp", "connection-org1.json");
        let ccp = JSON.parse(fs.readFileSync(ccpPath, "utf8"));

        const walletPath = path.join(process.cwd(), "wallet");
        const wallet = await Wallets.newFileSystemWallet(walletPath);
        console.log(`Wallet path: ${walletPath}`);

        const identity = await wallet.get("appUser");
        if (!identity) {
            console.log('An identity for the user "appUser" does not exist in the wallet');
            console.log("Run the registerUser.js application before retrying");
            return;
        }

        // Create a new gateway for connecting to our peer node.
        const gateway = new Gateway();
        await gateway.connect(ccp, {
            wallet,
            identity: "appUser",
            discovery: { enabled: true, asLocalhost: true },
        });

        // Get the network (channel) our contract is deployed to.
        const network = await gateway.getNetwork("mychannel");
        const contract = network.getContract("epeople");
        await contract.submitTransaction("AddUser", userId, password);

        console.log("Transaction has been submitted");
        
        await gateway.disconnect();

        res.status(200).send({"resultCode":"success"});

    } catch (error) {
        console.error(`Failed to submit transaction: ${error}`);
        res.status(500).send({"resultCode":"error", "result": `Failed to evaluate transaction: ${error}`});
    }

})

//민원등록
app.get('/api/complaint-request.json', async (req,res) => {
    const requestId = req.query.requestId
    console.log("GET /complaint-request.json", requestId);

    try {
        // load the network configuration
        const ccpPath = path.resolve(__dirname, "ccp", "connection-org1.json");
        const ccp = JSON.parse(fs.readFileSync(ccpPath, "utf8"));

        // Create a new file system based wallet for managing identities.
        const walletPath = path.join(process.cwd(), "wallet");
        const wallet = await Wallets.newFileSystemWallet(walletPath);
        console.log(`Wallet path: ${walletPath}`);

        // Check to see if we've already enrolled the user.
        const identity = await wallet.get("appUser");
        if (!identity) {
            console.log(
                'An identity for the user "appUser" does not exist in the wallet'
            );
            console.log("Run the registerUser.js application before retrying");
            return;
        }

        // Create a new gateway for connecting to our peer node.
        const gateway = new Gateway();
        await gateway.connect(ccp, {
            wallet,
            identity: "appUser",
            discovery: { enabled: true, asLocalhost: true },
        });

        // Get the network (channel) our contract is deployed to.
        const network = await gateway.getNetwork("mychannel");

        // Get the contract from the network.
        const contract = network.getContract("epeople");

        // Evaluate the specified transaction.
        // queryCar transaction - requires 1 argument, ex: ('queryCar', 'CAR4')
        // queryAllCars transaction - requires no arguments, ex: ('queryAllCars')
        const result = await contract.evaluateTransaction("GetComplaintRequest", requestId);
        console.log(
            `Transaction has been evaluated, result is: ${result.toString()}`
        );

        const res_str = `{"resultCode":"success", "result":${result}}`
        res.set(200).send(JSON.parse(res_str));

        // Disconnect from the gateway.
        await gateway.disconnect();
    } catch (error) {
        console.error(`Failed to evaluate transaction: ${error}`);
        res.status(500).send({"resultCode":"error", "result": `Failed to evaluate transaction: ${error}`});
    } 

})

//민원조회
app.post('/api/complaint-request.json', async (req,res) => {

//    const requestId = req.body.requestId
    const userId = req.body.userId
    const requesterName =  req.body.requesterName
    const phoneNumber = req.body.phoneNumber
    const address = req.body.address
    const open = req.body.open
    const title = req.body.title
    const content = req.body.content
    const complaintLocation = req.body.complaintLocation
    const requestDate = req.body.requestDate
    const receptionStatus = req.body.receptionStatus
    const receptionDate = req.body.receptionDate


    var uuid = v4();
    const requestId = "REQ:" + uuid;

    console.log("POST /api/complaint-request.json", requestId ,userId ,requesterName ,phoneNumber ,address ,open ,title ,content ,complaintLocation ,requestDate , uuid);

    try {
        // load the network configuration
        const ccpPath = path.resolve(__dirname, "ccp", "connection-org1.json");
        const ccp = JSON.parse(fs.readFileSync(ccpPath, "utf8"));

        const walletPath = path.join(process.cwd(), "wallet");
        const wallet = await Wallets.newFileSystemWallet(walletPath);
        console.log(`Wallet path: ${walletPath}`);

        const identity = await wallet.get("appUser");
        if (!identity) {
            console.log('An identity for the user "appUser" does not exist in the wallet');
            console.log("Run the registerUser.js application before retrying");
            return;
        }

        // Create a new gateway for connecting to our peer node.
        const gateway = new Gateway();
        await gateway.connect(ccp, {
            wallet,
            identity: "appUser",
            discovery: { enabled: true, asLocalhost: true },
        });

        // Get the network (channel) our contract is deployed to.
        const network = await gateway.getNetwork("mychannel");
        const contract = network.getContract("epeople");
        await contract.submitTransaction(
            "AddComplaintRequest"
            , requestId, userId, requesterName, phoneNumber, address, open, title, content, complaintLocation, requestDate
        );

        console.log("Transaction has been submitted");

        await gateway.disconnect();

        res.status(200).send({"resultCode":"success"});

    } catch (error) {
        console.error(`Failed to submit transaction: ${error}`);
        res.status(500).send({"resultCode":"error", "result": `Failed to evaluate transaction: ${error}`});
    }

})

//나의 민원 조회
app.get('/api/my-complaint-requests.json', async (req,res) => {
    console.log('GET /api/my-comlaint-requests.json')
    res.status(200).send({"resultCode":"개발중"});
});

//민원수정
app.put('/api/complaint-request.json', async (req,res) => {

    const requestId = req.body.requestId;
    const userId = req.body.userId
    const requesterName =  req.body.requesterName
    const phoneNumber = req.body.phoneNumber
    const address = req.body.address
    const open = req.body.open
    const title = req.body.title
    const content = req.body.content
    const complaintLocation = req.body.complaintLocation
    const requestDate = req.body.requestDate 

    console.log("PUT /api/complaint-request.json", requestId ,userId ,requesterName ,phoneNumber ,address ,open ,title ,content ,complaintLocation ,requestDate , uuid);

    res.status(200).send({"resultCode":"개발중"});

})


//민원 삭제
app.delete('/api/complaint-request.json', async (req,res) => {
    console.log('DELETE /api/my-comlaint-requests.json')
    res.status(200).send({"resultCode":"개발중"});
});

//나의 민원 리스트
app.get('/api/my-complaint-requests.json', async (req,res) => {
    console.log('GET /api/my-search-comlaint-requests.json')
    res.status(200).send({"resultCode":"개발중"});
});

//유사 민원 검색
app.get('/api/search-complaint-requests.json', async (req,res) => {
    console.log('GET /api/search-comlaint-requests.json')
    res.status(200).send({"resultCode":"개발중"});
});

//민원 결과 
app.get('/api/complaint-result.json', async (req,res) => {
    console.log('GET /api/comlaint-result.json')
    res.status(200).send({"resultCode":"개발중"});
});


//전체민원리스트
app.get('/api/all-complaint-requests.json', async (req,res) => {
    console.log('GET /api/all-comlaint-requests.json')
    res.status(200).send({"resultCode":"개발중"});
});

//민원접수
app.patch('/api/complaint-request-status.json', async (req,res) => {
    console.log('PATCH /api/comlaint-request-status.json')
    res.status(200).send({"resultCode":"개발중"});
});

//민원결과등록
app.post('/api/complaint-result.json', async (req,res) => {
    console.log('POST /api/comlaint-result.json')
    res.status(200).send({"resultCode":"개발중"});
});

//민원결과수정
app.put('/api/complaint-result.json', async (req,res) => {
    console.log('POST /api/comlaint-result.json')
    res.status(200).send({"resultCode":"개발중"});
});

app.listen(PORT, HOST)
console.log(`EPeople Server launched successfully http://${HOST}:${PORT}`)