package hlfsdk

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"os"
	"strings"
)

func Query(user string, connectionProfilePath string, channelName string, chaincodeName string, fcn string, args string) (string, error) {
	var payload string
	var error error
	fmt.Println(user)
	if _, err := os.Stat(connectionProfilePath); os.IsNotExist(err) {
		error = fmt.Errorf("file does not exist %v", err);
	} else {
		c := config.FromFile(connectionProfilePath)
		// Create a fabric sdk client
		sdk, err := fabsdk.New(c)
		if err != nil { // If the sdk conection with the fabric network was not successful
			error = fmt.Errorf("could not create sdk: %v", err)
			sdk.Close()
		} else {
			defer sdk.Close() // We defer the closing of the connection to later

			// We create a context to communicate with the channel
			clientChannelContext := sdk.ChannelContext(channelName, fabsdk.WithUser(user))
			// We create the channel client
			client, err := channel.New(clientChannelContext)
			if err != nil { // If channel client creation not successful
				error = fmt.Errorf("could not create channel client: %v", err)
			} else {
				// We split the args string into array of strings and convert them to a byte of byte array
				argArr := strings.Split(args, ",")
				for i := range argArr {
					argArr[i] = strings.TrimSpace(argArr[i])

				}
				queryArgs := make([][]byte, len(argArr))
				for i, v := range argArr {
					queryArgs[i] = []byte(v)
				}
				response, err := client.Query(channel.Request{
					ChaincodeID: chaincodeName,
					Fcn:         fcn,
					Args:        queryArgs,
				})
				if err != nil {
					error = fmt.Errorf("query execution failed: %v", err)
				} else {
					payload = string(response.Payload)
				}
			}
		}
	}
	return payload, error
}
