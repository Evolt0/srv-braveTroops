package user

import (
	"encoding/json"
	"github.com/Evolt0/def-braveTroops/consts/status"
	"github.com/Evolt0/def-braveTroops/proto"
	"github.com/Evolt0/def-braveTroops/proto/epkg"
	"github.com/Evolt0/def-braveTroops/proto/prefix"
	"github.com/Evolt0/fabric-sdk-yml/base"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	"log"
)

type Controller struct {
	baseClient *base.Client
}

func NewController(baseClient *base.Client) *Controller {
	return &Controller{baseClient: baseClient}
}

func (c Controller) Create(data *proto.UserReq) (interface{}, error) {
	log.Println(data)
	marshal, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	response, err := c.baseClient.ChannelExecute(
		channel.Request{ChaincodeID: c.baseClient.ChainCodeID, Fcn: prefix.FnCreateUser, Args: [][]byte{marshal}},
		channel.WithRetry(retry.DefaultChannelOpts))
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c Controller) ListLedgerByID(data *proto.BodyData) (interface{}, error) {
	log.Println(data)
	marshal, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	response, err := c.baseClient.ChannelQuery(
		channel.Request{ChaincodeID: c.baseClient.ChainCodeID, Fcn: prefix.FnListLedgerByID, Args: [][]byte{marshal}},
		channel.WithRetry(retry.DefaultChannelOpts))
	if err != nil {
		return nil, err
	}
	result := &proto.Ledger{}
	err = epkg.UnwrapSucc(response.Payload, result)
	if err != nil {
		return nil, epkg.Wrapf(status.InternalServerError, "failed to unmarshal %v", err)
	}
	return result, nil
}

func (c Controller) List() (interface{}, error) {
	response, err := c.baseClient.ChannelQuery(
		channel.Request{ChaincodeID: c.baseClient.ChainCodeID, Fcn: prefix.FnListUser, Args: [][]byte{}},
		channel.WithRetry(retry.DefaultChannelOpts))
	if err != nil {
		return nil, err
	}
	result := &proto.ListUser{}
	err = epkg.UnwrapSucc(response.Payload, result)
	if err != nil {
		return nil, epkg.Wrapf(status.InternalServerError, "failed to unmarshal %v", err)
	}
	return result, nil
}
