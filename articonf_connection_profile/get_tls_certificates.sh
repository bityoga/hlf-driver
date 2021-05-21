#!/bin/bash
set -x #echo on

IP_ADDRESS="188.166.104.30"
REMOTE_MACHINE_ORDERER_TLS_CERT_FILE="/root/hlft-store/orgca/orderer/msp/tls/ca.crt"
REMOTE_MACHINE_PEER2_TLS_CERT_FILE="/root/hlft-store/orgca/peer2/msp/tls/ca.crt"
LOCAL_ORDER_TLS_CERT_FILE="../mspX/keystore/hlft-store/orgca/orderer/msp/tls/ca.crt"
LOCAL_ORDER_PEER2_CERT_FILE="../mspX/keystore/hlft-store/orgca/peer2/msp/tls/ca.crt"

mkdir -p $LOCAL_ORDER_TLS_CERT_FILE
mkdir -p $LOCAL_ORDER_PEER2_CERT_FILE

scp -r root@$IP_ADDRESS:$REMOTE_MACHINE_ORDERER_TLS_CERT_FILE $LOCAL_ORDER_TLS_CERT_FILE &&
scp -r root@$IP_ADDRESS:$REMOTE_MACHINE_PEER2_TLS_CERT_FILE $LOCAL_ORDER_PEER2_CERT_FILE