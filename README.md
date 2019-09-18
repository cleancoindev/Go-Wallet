# Go- Wallet

I have created wallet for transfer tokens in between addresses and the trasaction block store in blockchain. Each block contains transferred token value, timestamp, hash of block hash of previous block and sender and recipient addresses. 

Also added wallet module to create new wallet address. In wallet module it first creates Private and Public keys using ecdsa algorithm and on the basis of public and private key it generates unique wallet address which will be used to transfer and recieve token.

# Run
- To run write below command into terminal
   
        cd Blockchain
        go run main.go

# screenshots

I have used gorilla-mux library to create apis
I have attached 5 screenshots with repo to interact with blockchain using postman.
- first screenshot(**addwallet**) : It generates new wallet address.
- second screenshot(**getAllWalletAddresses**) : Using this api you can see all the generated addresses.
- third screenshot(**getAllWalletDetails**) : Using this api you can see all the details of wallet address. Like private key, public key and tokens (In real life we can not show private key or wallet details, it is only for demo purpose).
- fourth screenshot(**transferToken/token/sender/recipient**) : Using this api you can transfer token between addresses and the transaction block will store in ledger. you can see ledger in scrrenshot.
- fifth screenshot(**getAllWalletDetails**) : Again(third screenshot) hit this api to see change in token value in wallet addresses. 

# Webapp

- After Running program write **http://localhost:8000** and you will get template.
- First press **Generate Wallet Address** Button to genrate addresses. Generate more than two addresses.
- Second you can get all generated addresses by pressing on **Get All Wallet Addresses** Button.
- Third you can also get all generated address details by pressing **Get All Wallet Addresses Details** Button. Details like Private key, Public key and Token associated with wallet address.
- Forth you can transfer token in between wallet addresses by giving Token value, sender address, recipient address to form and press send Button.
- To check updated token presse again **Get All Wallet Addresses Details** Button.



