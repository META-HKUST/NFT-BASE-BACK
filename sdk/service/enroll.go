package service

import (
	config_local "NFT-BASE-BACK/config"
	"errors"
	"log"
	"os"
	"os/exec"
)

//  x509::CN=zwang.org1.example.com,OU=client,O=Hyperledger,ST=North /
// Carolina,C=US::CN=ca-org1,OU=Fabric,O=Hyperledger,ST=North Carolina,C=US
func Enroll(username string) error {
	certPath := config_local.CONFIG.CryptoPath + username + config_local.CONFIG.CertPathSuffix
	// certPath := cryptoPath + username + ".org1.unifit.com/msp/signcerts/cert.pem"
	_, err := os.Stat(certPath)
	if err == nil {
		err1 := errors.New(username + " exists!")

		return err1

	} else {
		command := "/home/fabric_release/01_Fabric/NFT-BASE-FABRIC/tools/03_enroll_user.sh "
		command += username
		cmd := exec.Command("/bin/bash", "-c", command)
		_, err = cmd.CombinedOutput()
		if err != nil {
			log.Println("enroll user err: ", err)
			return err
		}
		return nil
	}
}

// // read pem
// data, err := ioutil.ReadFile(certPath)
// if err != nil {
// 	return err
// }
// // pem decode
// certBlock, restBlock := pem.Decode(data)
// if restBlock == nil {
// 	return errors.New("pem decode error")
// }
// // parse cert
// cert, err := x509.ParseCertificate(certBlock.Bytes)
// x509ID := "x509::" + cert.Subject.String() + "::" + cert.Issuer.String()
