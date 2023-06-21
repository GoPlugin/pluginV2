======================================
STEP 1:	SETTING UP THE NODE
======================================
1) Install Go1.20 using the link => https://go.dev/doc/install. Set GOROOT, GOPATH & PATH env variable using, export PATH=$GOPATH/bin:$PATH & export GOPATH=/Users/$USER/go
2) Install Postgresql using the link => https://wiki.postgresql.org/wiki/Detailed_installation_guides(version >=11.x and <=15.x)
3) Ensure you have python3
4) Download codebase: 'git clone https://github.com/raghu0891/pluginV2.git && cd pluginV2'
5) Install NodeJs (compatible version in MacOs to install pnpm: v16.14.0) using the link => https://nodejs.org/en/download/package-manager
6) Install pnpm using 'node install -g pnpm'
7) Now set the env variables present in '.env' file
8) Now run 'make install'.
   a) While intstallation it will ask for 'Keystore Password' and then 'confirm password'. 
   b) After setting up the keystore password, you need to provide a valid email and password which will be used to loginto http://localhost:6688
9) Once the build is successfull start the node using the command 'chainlink node start'
10) Login to the UI(http://localhost:6688) and go to 'Key Management' in settings and then fund the Address under 'EVM Chain Accounts' with 1 XDC.


======================================
STEP 2:	SETTING UP ORACLE CONTRACT
======================================
1) Copy the Operator.sol from the link => https://github.com/smartcontractkit/chainlink/blob/develop/contracts/src/v0.7/Operator.sol
2) Compile and Deploy it in remix, while deploying you need to give the below mentioned two parameters.
	LINK: 0xff7412ea7c8445c46a8254dfb557ac1e48094391 for Mainnet
	OWNER:<Your_Wallet_Address>
3) Once the contract is deployed copy the Oracle Contract Address
4) In the deployed contract, go to 'setAuthorizedSenders' and provide the node address in array Ex:["0xdEad....."] and then click on 'transact'. This will perform the fulfillment.


======================================
STEP 3:	SETTING UP JOB 
======================================
1) Copy the Job spec from the page => https://docs.chain.link/chainlink-nodes/job-specs/direct-request-get-uint256
2) Replace 'YOUR_ORACLE_CONTRACT_ADDRESS' with the Oracle Contract address in the job spec(there will be 2 instance to be chaged).
3) Now go to the logged in UI(http://localhost:6688) nad click on the 'Jobs', in the resultant page submit the job spec.
4) Once the Job spec is submitted successfully, copy the job spec id from 'Description'.


======================================
STEP 4:	SETTING UP INTERNAL CONTRACT
======================================
1) Copy the  contract from the link => https://docs.chain.link/any-api/get-request/examples/single-word-response
2) In the constructor part of the contract override the existing values with the values we got in STEP 2, STEP 3.
	setChainlinkToken => 0xff7412ea7c8445c46a8254dfb557ac1e48094391 for Mainnet
	setChainlinkOracle => <Oracle address from STEP 2>
	jobId => <Job id from STEP 3>
3) Now Compile and Deploy this contract using remix.  
4) Once the contract is deployed successfully, click on 'requestVolumeData' and check the job run in UI.
5) After the job run is 'Completed', click on 'volume' in deployed contract to get the value.
