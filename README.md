
# XDC transaction signing with Builder Vault


### Prerequisites
  - Go https://go.dev/doc/install
  - IDE recommended
  - Register for a demo Builder Vault tenant: https://www.blockdaemon.com/wallet/builder-vault
  - Download SDK bundle provided in registration email (extract authentication certificates)
  - Place Builder Vault authentication certificates in working directory


### Step 1. Create Wallet with Builder Vault and XDC chain path
```shell
go run 1-create-xdc-wallet/main.go
```
  - note the new XDC Wallet address and fund it through https://faucet.blocksscan.io
  - note the Builder Vault Master Key ID (to be used for future signing)


### Step 2. Craft an unsigned transaction with go-ethereum
  - set the wallet address created in step 1
  - set the destination address to send funds to
```shell
go run 2-craft-transaction/main.go
```
  - check the wallet balance is not 0
  - note the raw unsigned transaction (to be broadcast once signed)
  - note the unsigned transaction hash (to be signed by Builder Vault)


### Step 3. Sign transaction with MPC wallet library
  - set the Builder Vault Master Key ID to be used for signing (created in step 1)
  - set the unsigned transaction hash to be signed by Builder Vault (created in step 2)
```shell
go run 3-sign-transaction/main.go
```
  - note the transaction signature


### Step 4. Broadcast signed raw transaction with XDC client library
  - set the transaction signature hash (created in step 3)
  - set the raw unsigned transaction (created in step 2)
```shell
go run 4-broadcast-signed-transaction/main.go
```
  - check the transaction hash on a block explorer


## XDC Testnet Network details:
  - RPC endpoint: https://rpc.apothem.network
  - Block explorers: https://explorer.apothem.network https://apothem.blocksscan.io
  - Faucets: https://faucet.apothem.network https://faucet.blocksscan.io
  - ChainId: 51
  - Web wallet: https://wallet.apothem.network https://mnrpc.xinfin.network

<!--
ToDo:
- Sequence diagram  
- prize for completion?
-->