package sdk

import (
	"io/ioutil"

	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

const (
	cryptoPath    = "/home/fabric_release/01_Fabric/hyperledger/crypto-config/peerOrganizations/org1.unifit.com/users/"
	ccpPath       = "/home/fabric_release/01_Fabric/hyperledger/crypto-config/peerOrganizations/org1.unifit.com/connection-org1.yaml"
	channelName   = "unifitchannel"
	chaincodeName = "unifitPublicNFT"
)

func populateWallet(wallet *gateway.Wallet, username string) error {

	certPath := cryptoPath + username + ".org1.unifit.com/msp/signcerts/cert.pem"
	keyPath := cryptoPath + username + ".org1.unifit.com/msp/keystore/"
	// read the certificate pem
	cert, err := ioutil.ReadFile(certPath)
	if err != nil {
		return err
	}

	// there's a single file in this dir containing the private key
	files, err := ioutil.ReadDir(keyPath)
	if err != nil {
		return err
	}
	privateKeyPath := keyPath + files[0].Name()
	key, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		return err
	}

	identity := gateway.NewX509Identity("Org1MSP", string(cert), string(key))
	err = wallet.Put(username, identity)
	if err != nil {
		return err
	}
	return nil
}

func Submit(username string, contractName string, args ...string) (string, error) {

	wallet, err := gateway.NewFileSystemWallet("/home/fabric_release/03_End/NFT-BASE-BACK/sdk/wallet")
	if err != nil {
		return "", err
	}

	if !wallet.Exists(username) {
		err = populateWallet(wallet, username)
		if err != nil {
			return "", err
		}
	}

	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromFile(ccpPath)),
		gateway.WithIdentity(wallet, username),
	)
	if err != nil {
		return "", err
	}
	defer gw.Close()
	network, err := gw.GetNetwork(channelName)
	if err != nil {
		return "", err
	}

	contract := network.GetContract(chaincodeName)
	// result, err := contract.SubmitTransaction("MintWithTokenURI", string(id), "http://example.com")
	result, err := contract.SubmitTransaction(contractName, args...)
	if err != nil {
		return "", err
	}

	return string(result), nil
}

func Evalute(username string, contractName string, args ...string) (string, error) {

	wallet, err := gateway.NewFileSystemWallet("/home/fabric_release/03_End/NFT-BASE-BACK/sdk/wallet")
	if err != nil {
		return "", err
	}

	if !wallet.Exists(username) {
		err = populateWallet(wallet, username)
		if err != nil {
			return "", err
		}
	}

	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromFile(ccpPath)),
		gateway.WithIdentity(wallet, username),
	)
	if err != nil {
		return "", err
	}
	defer gw.Close()
	network, err := gw.GetNetwork(channelName)
	if err != nil {
		return "", err
	}

	contract := network.GetContract(chaincodeName)
	result, err := contract.EvaluateTransaction(contractName, args...)
	if err != nil {
		return "", err
	}

	return string(result), nil
}
