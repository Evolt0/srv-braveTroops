package base

import (
	"encoding/json"
	"log"

	"github.com/Parker-Yang/fabric-sdk-yml/base"
	"github.com/Parker-Yang/srv-braveTroops/models"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
)

type Controller struct {
	baseClient *base.Client
}

func NewController(baseClient *base.Client) *Controller {
	return &Controller{baseClient: baseClient}
}

func (c *Controller) PutState(data models.PutState) (interface{}, error) {
	log.Println(data)
	marshal, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	response, err := c.baseClient.ChannelExecute(
		channel.Request{ChaincodeID: c.baseClient.ChainCodeID, Fcn: "PutState", Args: [][]byte{marshal}},
		channel.WithRetry(retry.DefaultChannelOpts))
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Controller) GetState(data *models.GetState) (interface{}, error) {
	log.Println(data)
	marshal, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	response, err := c.baseClient.ChannelQuery(
		channel.Request{ChaincodeID: c.baseClient.ChainCodeID, Fcn: "GetState", Args: [][]byte{marshal}},
		channel.WithRetry(retry.DefaultChannelOpts))
	if err != nil {
		return nil, err
	}
	return response, nil
}
