# Secret-Keeper

Author: Charlie Gutierrez  [charliegut14@gmail.com](mailto:charliegut14@gmail.com)

## About Project
Secret Keeper is a CLI based tool written in Go used to encrypt/store decrypt/get key value pairs


## Usage
1. Clone the repository
2. Compile the code with go build -o ./secret-keeper cmd/cli.go   
3. Specify the key to use for encoding as well as the key/value to encrypt
./secret-keeper set secret_api_key "some-value" -k=someencodingkey
4. Get your keys/values
./secret-keeper get secret_api_key -k=someencodingkey


## Contact
Charlie Gutierrez
[charliegut14@gmail.com](mailto:charliegut14@gmail.com)
