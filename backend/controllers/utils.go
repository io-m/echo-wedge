package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	client "github.com/io-m/echo-wedge/backend/tcpClient"
)

var (
	tcpClient = client.TCPClient{
		Host: "localhost",
		Port: 8051,
	}
	in = &client.JSONRPC{}
)

func findExact(id string, mapp map[string]string) string {
	var exactID string
	for k, v := range mapp {
		if id == k {
			exactID = v 
			return exactID
		}
	}
	return ""
}

// WedgeCallNetwork is helper function for making tcp call to API gtw server for data related to network
func WedgeCallNetwork(ids map[string]string) (*client.JSONRPCResponseNetwork, error) {
	tcpClient.Start()
	if _, ok := ids["netId"]; !ok {
		in = &client.JSONRPC{
			ID:      uuid.New().String(),
			Jsonrpc: "2.0",
			Method:  http.MethodGet,
			Params: client.Params {
				URL: "/",
			},
		}
		b, err := json.Marshal(in)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}

		if err = tcpClient.Write(b); err != nil {
			return nil, err
		}
		netData := tcpClient.ReadNetwork()
		return netData, nil
	}
	if _, ok := ids["netId"]; ok  {
		netID := findExact("netId", ids)
		in = &client.JSONRPC{
			ID:      uuid.New().String(),
			Jsonrpc: "2.0",
			Method:  http.MethodGet,
			Params: client.Params {
				URL: fmt.Sprintf("/network/%s", netID),
			},
		}
		b, err := json.Marshal(in)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}

		if err = tcpClient.Write(b); err != nil {
			return nil, err
		}
		netData := tcpClient.ReadNetwork()
		return netData, nil
	}
	return nil, errors.New("something went wrong with network api call")
}

// WedgeCallDevice is helper function for making tcp call to API gtw server for data related to device
func WedgeCallDevice(ids map[string]string) (*client.JSONResponseValueDevice, *client.JSONResponseOneDevice , error) {
	tcpClient.Start()
	if _, ok := ids["devId"]; !ok {
		netID := findExact("netId", ids)
		in = &client.JSONRPC{
			ID:      uuid.New().String(),
			Jsonrpc: "2.0",
			Method:  http.MethodGet,
			Params: client.Params {
				URL: fmt.Sprintf("/network/%s/device", netID),
			},
		}
		b, err := json.Marshal(in)
		if err != nil {
			fmt.Println(err.Error())
			return nil, nil, err
		}

		if err = tcpClient.Write(b); err != nil {
			return nil, nil, err
		}
		devData := tcpClient.ReadDeviceValue()
		return devData, nil, nil
	}
	if _, ok := ids["devId"]; ok {
		netID := findExact("netId", ids)
		devID := findExact("devId", ids)
		in = &client.JSONRPC{
			ID:      uuid.New().String(),
			Jsonrpc: "2.0",
			Method:  http.MethodGet,
			Params: client.Params {
				URL: fmt.Sprintf("/network/%s/device/%s", netID, devID),
			},
		}
		b, err := json.Marshal(in)
		if err != nil {
			fmt.Println(err.Error())
			return nil, nil, err
		}

		if err = tcpClient.Write(b); err != nil {
			return nil, nil, err
		}
		devDataOne := tcpClient.ReadDeviceOne()
		return nil, devDataOne, nil
	}
	return nil, nil, errors.New("something went wrong with device api call")
}

// WedgeCallValue is helper function for making tcp call to API gtw server for data related to value
func WedgeCallValue(ids map[string]string) (*client.JSONResponseValueDevice, *client.JSONResponseOneValue , error) {
	tcpClient.Start()
	if _, ok := ids["valId"]; !ok {
		netID := findExact("netId", ids)
		devID := findExact("devId", ids)
		in = &client.JSONRPC{
			ID:      uuid.New().String(),
			Jsonrpc: "2.0",
			Method:  http.MethodGet,
			Params: client.Params {
				URL: fmt.Sprintf("/network/%s/device/%s/value", netID, devID),
			},
		}
		b, err := json.Marshal(in)
		if err != nil {
			fmt.Println(err.Error())
			return nil, nil, err
		}

		if err = tcpClient.Write(b); err != nil {
			return nil, nil, err
		}
		valData := tcpClient.ReadDeviceValue()
		return valData, nil, nil
	}
	if _, ok := ids["valId"]; ok {
		netID := findExact("netId", ids)
		devID := findExact("devId", ids)
		valID := findExact("valId", ids)
		in = &client.JSONRPC{
			ID:      uuid.New().String(),
			Jsonrpc: "2.0",
			Method:  http.MethodGet,
			Params: client.Params {
				URL: fmt.Sprintf("/network/%s/device/%s/value/%s", netID, devID, valID),
			},
		}
		b, err := json.Marshal(in)
		if err != nil {
			fmt.Println(err.Error())
			return nil, nil, err
		}

		if err = tcpClient.Write(b); err != nil {
			return nil, nil, err
		}
		valDataOne := tcpClient.ReadValueOne()
		return nil, valDataOne, nil
	}
	return nil, nil, errors.New("something went wrong with value api call")
}
// WedgeCallState is helper function for making tcp call to API gtw server for data related to state
func WedgeCallState(ids map[string]string) (*client.JSONResponseState, *client.JSONResponseOneState, error) {
	tcpClient.Start()
	if _, ok := ids["stateId"]; !ok {
		netID := findExact("netId", ids)
		devID := findExact("devId", ids)
		valID := findExact("valId", ids)
		in = &client.JSONRPC{
			ID:      uuid.New().String(),
			Jsonrpc: "2.0",
			Method:  http.MethodGet,
			Params: client.Params {
				URL: fmt.Sprintf("/network/%s/device/%s/value/%s/state", netID, devID, valID),
			},
		}
		b, err := json.Marshal(in)
		if err != nil {
			fmt.Println(err.Error())
			return nil, nil, err
		}

		if err = tcpClient.Write(b); err != nil {
			return nil, nil, err
		}
		stateData := tcpClient.ReadState()
		return stateData, nil, nil
	}
	if _, ok := ids["stateId"]; ok {
		netID := findExact("netId", ids)
		devID := findExact("devId", ids)
		valID := findExact("valId", ids)
		stateID := findExact("stateId", ids)
		in = &client.JSONRPC{
			ID:      uuid.New().String(),
			Jsonrpc: "2.0",
			Method:  http.MethodGet,
			Params: client.Params {
				URL: fmt.Sprintf("/network/%s/device/%s/value/%s/state/%s", netID, devID, valID, stateID),
			},
		}
		b, err := json.Marshal(in)
		if err != nil {
			fmt.Println(err.Error())
			return nil, nil, err
		}

		if err = tcpClient.Write(b); err != nil {
			return nil, nil, err
		}
		stateDataOne := tcpClient.ReadStateOne()
		return nil, stateDataOne, nil
	}

	return nil, nil, errors.New("something went wrong with state api call")
}