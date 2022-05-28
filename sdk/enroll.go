package sdk

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"os/exec"
)

func Enroll(username string) (id string, err error) {

	command := "/home/fabric_release/01_Fabric/NFT-BASE-FABRIC/tools/03_enroll_user.sh "
	command += username
	cmd := exec.Command("/bin/bash", "-c", command)

	_, err = cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	// read pem
	certPath := cryptoPath + username + ".org1.example.com/msp/signcerts/cert.pem"
	data, err := ioutil.ReadFile(certPath)
	if err != nil {
		return "", err
	}
	// pem decode
	certBlock, restBlock := pem.Decode(data)
	if restBlock == nil {
		return "", errors.New("pem decode error")
	}

	// parse cert
	cert, err := x509.ParseCertificate(certBlock.Bytes)
	x509ID := "x509::" + cert.Subject.String() + "::" + cert.Issuer.String()

	return x509ID, nil
}
