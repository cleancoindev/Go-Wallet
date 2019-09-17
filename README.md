# Go- Wallet

I have created wallet for transfer tokens in between addresses and the trasaction block store in blockchain. Each block contains transferred token value, timestamp, hash of block hash of previous block and sender and recipient addresses. 

Also added wallet module to create new wallet address. In wallet module it first creates Private and Public keys using ecdsa algorithm and on the basis of public and private key it generates unique wallet address which will be used to transfer and recieve token.

# Run
- To run write below command into terminal
   
        cd Blockchain
        go run main.go

# screenshots

I have used gorrilla-mux library to create apis
I have attached 5 screenshots with repo to interact with blockchain using postman.
- first screenshot(addwallet) : It generates new wallet address.
- second screenshot(getAllWalletAddresses) : Using this api you can see all the generated addresses.
- third screenshot(getAllWalletDetails) : Using this api you can see all the details of wallet address. Like private key, public key and tokens (In real life we can not show private key or wallet details, it is only for demo purpose).
- fourth screenshot(transferToken/token/sender/recipient) : Using this api you can transfer token between addresses and the transaction block will store in ledger. you can see ledger in scrrenshot.
- fifth screenshot(getAllWalletDetails) : Again(third screenshot) hit this api to see change in token value in wallet addresses. 

