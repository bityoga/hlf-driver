# Hyperledger Fabric Golang sdk driver for [react-native-hlf-wrapper](https://github.com/bityoga/react-native-hlf-wrapper)
A project to generate the required native modules for hyperledger fabric go sdk to be used as a react native module.
# Dependencies
Make sure ```golang`` installed and ```GO_PATH``` is set

Run the following in the terminal to get the required dependencies
```
go get github.com/hyperledger/fabric-sdk-go
go get golang.org/x/mobile/cmd/gomobile
```

Also, navigate to this cloned folder in your terminal and run the following commands in your terminal
```
go mod init
go mod tidy
```

# Binding the SDK package
Run the following command to bind generate the java code:
## Android
```
gomobile bind -o ./bin/hlfsdk.aar -target=android github.com/bityoga/hlf-driver
```
This should generate two files, sdk.aar and sdk-sources.jar. 

These files should be moved to a new folder  called "libs" inside the react native project's android directory.
Make sure to add these libs to your android project

## IOS
```
gomobile bind -ldflags "-w" -o ./bin/hlfsdk.framework -target ios github.com/bityoga/hlf-driver
```
This command removes the debug information, reducing the resulting file size under the GitHub size limit.

Place the generated hlfsdk.framework directory into a new folder  called "libs" inside the react native project's ios directory.
Make sure to add these libs to your ios project

## Functionalities
Only the enroll, query and invoke functionalities are provided. This sdk wrapper is created to be used as native module for react-native. There the idea is that a client mobile application is able to only perform these operations and not any admin operations. Admin operations like channel creation, chaincode installation should be performed by a server side sdk and not a mobile client.

Enroll User: Enrolls a user and stores their crypto materials
```
go run main/main.go enroll username userpw connection_profile.json
```
Query Chaincode: Queries a chaincode
```
go run main/main.go query username connection_profile.json channel_name chaincode_name function_name args,seperated,by,comma
```
Invoke Chaincode: Invokes a function in the chaincode
```
go run main/main.go invoke username connection_profile.json channel_name chaincode_name function_name args,seperated,by,comma
```
