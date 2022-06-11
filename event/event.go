package event

import (
	"encoding/hex"
	"fmt"
	fabAPI "github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
	"log"
)

const (
	cryptoPath    = "/home/fabric_release/01_Fabric/hyperledger/crypto-config/peerOrganizations/org1.unifit.com/users/"
	ccpPath       = "/home/fabric_release/01_Fabric/hyperledger/crypto-config/peerOrganizations/org1.unifit.com/connection-org1.yaml"
	channelName   = "unifitchannel"
	chaincodeName = "unifitPublicNFT"
)

func regitserEvent(client *gateway.Contract, chaincodeID string) (fabAPI.Registration, <-chan *fabAPI.CCEvent) {
	eventName := "Transfer"
	reg, notifier, err := client.RegisterEvent(eventName)
	if err != nil {
		log.Printf("注册链码事件失败: %s", err)
	}
	log.Println("Registered chaincode event\n")
	return reg, notifier
}

func ChainCodeEventListener(client *gateway.Contract, ccID string) fabAPI.Registration {
	fmt.Println(client.Name())
	reg, notifier := regitserEvent(client, ccID)
	fmt.Print(notifier)
	// consume event

	go func() {
		for e := range notifier {
			log.Printf("Receive cc event,ccid: %s\n eventName:%s",ccID,e.EventName)
		}
	}()

	return reg
}


func BlockListener(ec *gateway.Network) fabAPI.Registration {
	// Register monitor block event
	beReg, beCh, err := ec.RegisterBlockEvent()
	if err != nil {
		log.Printf("Register block event error: %v", err)
	}
	log.Println("Registered block event")

	// Receive block event
	go func() {
		for e := range beCh {
			log.Printf("Receive block event:\nSourceURL: %v\nNumber: %v\nHash"+
				": %v\nPreviousHash: %v\n\n",
				e.SourceURL,
				e.Block.Header.Number,
				hex.EncodeToString(e.Block.Header.DataHash),
				hex.EncodeToString(e.Block.Header.PreviousHash))
		}
	}()

	return beReg
}








