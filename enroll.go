package hlfsdk

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"os"
)

func Enroll(user string, secret string, connectionProfilePath string) (string, error) {
	var payload string
	var error error

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
			// If the sdk was created successfully
			ctx := sdk.Context()
			mspClient, err := msp.New(ctx) // Create a content so that we can interact with the  blockchain
			if err != nil { // Not Successful
				error = fmt.Errorf("could not create msp client: %v", err)
			} else { // Successful
				identity, err := mspClient.GetSigningIdentity(user) //We load the identity from local store
				if err == msp.ErrUserNotFound { // If the identity is not present in the local store
					err = mspClient.Enroll(user, msp.WithSecret(secret)) // We enroll and store the identity locally
					if err != nil {
						//If we could not successfully enroll the user
						error = fmt.Errorf("failed to enroll user. either the user is already enrolled. local msp not found or connection profile setting issues: %v", err)
					} else { // We enrolled the user successfully
						identity, err = mspClient.GetSigningIdentity(user) // We load the identity from local store
						if err != nil {
							error = fmt.Errorf("GetSigningIdentity failed: %v", err) // We couldnt load the identity that was enrolled
						}
					}
				} else {
					// The user is already enrolled and we found its MP locally.
					// That's why  we dont override  the global error variable
					fmt.Errorf("user already enrolled. local msp found! %v", err)
				}

				// If we dont have anny errors
				if error == nil {
					// Load the required info from the identity and store them into the Wallet object
					cert := base64.StdEncoding.EncodeToString(identity.PublicVersion().EnrollmentCertificate())
					privKeyName := fmt.Sprintf("%x_sk", identity.PrivateKey().SKI())
					id := identity.Identifier()

					certName := fmt.Sprintf("%s@%s-cert.pem", id.ID, id.MSPID)
					output, err := json.Marshal(&Wallet{cert, certName, privKeyName})
					if err != nil {
						error = fmt.Errorf("GetSigningIdentity failed: %v", err)
					}
					payload = string(output)
				}
			}
		}
	}
	return payload, error
}