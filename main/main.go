package main

import (
	"fmt"
	"os"
	hlfsdk "github.com/bityoga/hlf-driver"
)

type osArgs struct {
	methodToCall string
	user string
	secret  string
	connectionPath string
	channelName string
	chaincodeName string
	fcn string
	args string
}

func readArgs() osArgs {
	args := osArgs{"", "", "", "", "", "", "", ""}
	if len(os.Args) >= 2 {
		if os.Args[1] == "enroll" {
			if len(os.Args) == 5 {
				args.user = os.Args[2]
				args.secret = os.Args[3]
				args.connectionPath = os.Args[4]
			} else {
				fmt.Println("sufficient args not supplied for enroll: methodToCall=enroll user secret connectionPath")
				os.Exit(1)
			}
		} else if os.Args[1] == "query" || os.Args[1] == "invoke" {
			if len(os.Args) == 8 {
				args.user = os.Args[2]
				args.connectionPath = os.Args[3]
				args.channelName = os.Args[4]
				args.chaincodeName = os.Args[5]
				args.fcn = os.Args[6]
				args.args = os.Args[7]
			} else {
				fmt.Println("sufficient args not supplied for query/invoke: methodToCall=invoke/query user secret connectionPath channelName chaincodeName functionToCall args")
				os.Exit(1)
			}
		} else {
			fmt.Println("the method to call (enroll, invoke, query) should be the first arg")
			os.Exit(1)
		}

	} else {
		fmt.Println("the method to call (enroll, invoke, query) should be the first arg")
		os.Exit(1)
	}

	return args
}

func main() {
	args := readArgs()
	var payload string
	var err error
	if os.Args[1] == "enroll" {
		payload, err = hlfsdk.Enroll(args.user, args.secret, args.connectionPath)
	} else if os.Args[1] == "query" {
		payload, err = hlfsdk.Query(args.user, args.connectionPath, args.channelName, args.chaincodeName, args.fcn, args.args)
	} else if os.Args[1] == "invoke" {
		payload, err = hlfsdk.Invoke(args.user, args.connectionPath, args.channelName, args.chaincodeName, args.fcn, args.args)
	}

	if err == nil {
		fmt.Println(payload)
	} else {
		fmt.Println("error: ", err)
	}
}