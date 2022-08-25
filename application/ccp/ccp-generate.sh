#!/bin/bash

function one_line_pem {
    echo "`awk 'NF {sub(/\\n/, ""); printf "%s\\\\\\\n",$0;}' $1`"
}

function json_ccp {
    local PP=$(one_line_pem $5)
    local CP=$(one_line_pem $6)
    sed -e "s/\${ORG}/$1/" \
        -e "s/\${HOST}/$2/" \
        -e "s/\${P0PORT}/$3/" \
        -e "s/\${CAPORT}/$4/" \
        -e "s#\${PEERPEM}#$PP#" \
        -e "s#\${CAPEM}#$CP#" \
        ${DIR}/ccp-template.json
}

function yaml_ccp {
    local PP=$(one_line_pem $5)
    local CP=$(one_line_pem $6)
    sed -e "s/\${ORG}/$1/" \
        -e "s/\${HOST}/$2/" \
        -e "s/\${P0PORT}/$3/" \
        -e "s/\${CAPORT}/$4/" \
        -e "s#\${PEERPEM}#$PP#" \
        -e "s#\${CAPEM}#$CP#" \
        ${DIR}/ccp-template.yaml | sed -e $'s/\\\\n/\\\n          /g'
}

# Where am I?
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

NET_DIR_PATH="${DIR}/../../network"

ORGNAME=Org1
HOSTNAME=org1
P0PORT=7051
CAPORT=7054
PEERPEM=${NET_DIR_PATH}/organizations/peerOrganizations/org1.example.com/tlsca/tlsca.org1.example.com-cert.pem
CAPEM=${NET_DIR_PATH}/organizations/peerOrganizations/org1.example.com/ca/ca.org1.example.com-cert.pem

echo "$(json_ccp $ORGNAME $HOSTNAME $P0PORT $CAPORT $PEERPEM $CAPEM)" > ${DIR}/connection-org1.json
echo "$(yaml_ccp $ORGNAME $HOSTNAME $P0PORT $CAPORT $PEERPEM $CAPEM)" > ${DIR}/connection-org1.yaml

ORGNAME=Org2
HOSTNAME=org2
P0PORT=9051
CAPORT=8054
PEERPEM=${NET_DIR_PATH}/organizations/peerOrganizations/org2.example.com/tlsca/tlsca.org2.example.com-cert.pem
CAPEM=${NET_DIR_PATH}/organizations/peerOrganizations/org2.example.com/ca/ca.org2.example.com-cert.pem

echo "$(json_ccp $ORGNAME $HOSTNAME $P0PORT $CAPORT $PEERPEM $CAPEM)" > ${DIR}/connection-org2.json
echo "$(yaml_ccp $ORGNAME $HOSTNAME $P0PORT $CAPORT $PEERPEM $CAPEM)" > ${DIR}/connection-org2.yaml

ORGNAME=Uni1
HOSTNAME=uni1
P0PORT=11051
CAPORT=11054
PEERPEM=${NET_DIR_PATH}/organizations/peerOrganizations/uni1.example.com/tlsca/tlsca.uni1.example.com-cert.pem
CAPEM=${NET_DIR_PATH}/organizations/peerOrganizations/uni1.example.com/ca/ca.uni1.example.com-cert.pem

echo "$(json_ccp $ORGNAME $HOSTNAME $P0PORT $CAPORT $PEERPEM $CAPEM)" > ${DIR}/connection-uni1.json
echo "$(yaml_ccp $ORGNAME $HOSTNAME $P0PORT $CAPORT $PEERPEM $CAPEM)" > ${DIR}/connection-uni1.yaml

ORGNAME=Uni2
HOSTNAME=uni2
P0PORT=13051
CAPORT=13054
PEERPEM=${NET_DIR_PATH}/organizations/peerOrganizations/uni2.example.com/tlsca/tlsca.uni2.example.com-cert.pem
CAPEM=${NET_DIR_PATH}/organizations/peerOrganizations/uni2.example.com/ca/ca.uni2.example.com-cert.pem

echo "$(json_ccp $ORGNAME $HOSTNAME $P0PORT $CAPORT $PEERPEM $CAPEM)" > ${DIR}/connection-uni2.json
echo "$(yaml_ccp $ORGNAME $HOSTNAME $P0PORT $CAPORT $PEERPEM $CAPEM)" > ${DIR}/connection-uni2.yaml