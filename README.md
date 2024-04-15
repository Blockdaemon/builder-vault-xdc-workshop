
# XDC transaction signing with Builder Vault

```mermaid
sequenceDiagram
    autonumber
    actor client1 as XDC client
    participant Blockchain as XDC Blockchain<br> RPC API
    actor client2 as BuilderVault client
    box Builder Vault
      participant TSM1 as MPC Wallet <br>(private key share 1)
      participant TSM2 as MPC Wallet <br>(private key share 2)
      participant TSM3 as MPC Wallet <br>(private key share 3)
    end
    #note over client,TSM2: Create wallet
    opt
      client2 ->> TSM1: create master key
      client2 ->> TSM1: create XDC wallet 
    end

    #note over client,TSM2: Create transations
    client1 ->> Blockchain: get blockchain inputs (gas, nonce, balance) for new tx<br>(sender wallet)
    client1 ->> client1: construct unsigned tx
    client1 ->> client2: request signature (unsigned tx hash)

    client2 ->> TSM1: request share1 signature (unsigned tx hash)
    TSM1 -->> client2: return partial signature
    client2 ->> TSM2: request share2 signature (unsigned tx hash)
    TSM2 -->> client2: return partial signature
    client2 ->> TSM3: request share3 signature (unsigned tx hash)
    TSM3 -->> client2: return partial signature
    client2 ->> client2: combine partial signatures
    client2 -->> client1: return signature
    client1 ->> Blockchain: broadcast signed tx<br>(signed tx)
```

### Prerequisites
  - Go https://go.dev/doc/install
  - Register a demo Builder Vault tenant: https://www.blockdaemon.com/wallet/builder-vault
  - Download SDK bundle provided in registration email (extract auth certificates)
  - Place Builder Vault `client.crt` & `client.key` in working directory

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
  - RPC endpoint: https://erpc.apothem.network
  - Block explorers: https://explorer.apothem.network https://apothem.blocksscan.io
  - Faucets: https://faucet.apothem.network https://faucet.blocksscan.io
  - ChainId: 51
  - Web wallet: https://wallet.apothem.network https://mnrpc.xinfin.network
  - TX decoder: https://rawtxdecode.in/
