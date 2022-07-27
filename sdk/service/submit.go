package service

import (
	"io/ioutil"
	"log"

	config_local "NFT-BASE-BACK/config"

	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

const (
// cryptoPath    = "/home/fabric_release/01_Fabric/hyperledger/crypto-config/peerOrganizations/org1.unifit.com/users/"
// ccpPath       = "/home/fabric_release/01_Fabric/hyperledger/crypto-config/peerOrganizations/org1.unifit.com/connection-org1.yaml"
// channelName   = "unifitchannel"
// chaincodeName = "unifitPublicNFT"
)

func populateWallet(wallet *gateway.Wallet, username string) error {

	certPath := config_local.CONFIG.CryptoPath + username + config_local.CONFIG.CertPathSuffix
	// certPath := cryptoPath + username + ".org1.unifit.com/msp/signcerts/cert.pem"
	keyPath := config_local.CONFIG.CryptoPath + username + config_local.CONFIG.KeyPathSuffix
	// keyPath := cryptoPath + username + ".org1.unifit.com/msp/keystore/"
	// read the certificate pem
	cert, err := ioutil.ReadFile(certPath)
	if err != nil {
		log.Println(err)
		return err
	}

	// there's a single file in this dir containing the private key
	files, err := ioutil.ReadDir(keyPath)
	if err != nil {
		log.Println(err)
		return err
	}
	privateKeyPath := keyPath + files[0].Name()
	key, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		log.Println(err)
		return err
	}

	identity := gateway.NewX509Identity(config_local.CONFIG.MspId, string(cert), string(key))
	err = wallet.Put(username, identity)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func Submit(username string, contractName string, args ...string) (string, error) {

	wallet, err := gateway.NewFileSystemWallet(config_local.CONFIG.WalletPath)
	if err != nil {
		log.Println(err)
		return "", err
	}

	if !wallet.Exists(username) {
		err = populateWallet(wallet, username)
		if err != nil {
			log.Println(err)
			return "", err
		}
	}

	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromFile(config_local.CONFIG.CcpPath)),
		gateway.WithIdentity(wallet, username),
	)

	if err != nil {
		log.Println(err)
		return "", err
	}
	defer gw.Close()

	network, err := gw.GetNetwork(config_local.CONFIG.ChannelName)
	if err != nil {
		log.Println(err)
		return "", err
	}

	contract := network.GetContract(config_local.CONFIG.ChaincodeName)

	log.Println("args: ", args)

	result, err := contract.SubmitTransaction(contractName, args...)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return string(result), nil
}

func Evaluate(username string, contractName string, args ...string) (string, error) {

	wallet, err := gateway.NewFileSystemWallet(config_local.CONFIG.WalletPath)
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
		gateway.WithConfig(config.FromFile(config_local.CONFIG.CcpPath)),
		gateway.WithIdentity(wallet, username),
	)
	if err != nil {
		return "", err
	}
	defer gw.Close()
	network, err := gw.GetNetwork(config_local.CONFIG.ChannelName)
	if err != nil {
		return "", err
	}

	contract := network.GetContract(config_local.CONFIG.ChaincodeName)
	result, err := contract.EvaluateTransaction(contractName, args...)
	if err != nil {
		return "", err
	}

	return string(result), nil
}
