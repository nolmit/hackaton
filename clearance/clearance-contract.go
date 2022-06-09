/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

//ClearanceContract contract for managing CRUD for Clearance
type ClearanceContract struct {
	contractapi.Contract
}

//ClearanceExists returns true when asset with given ID exists in world state
func (c *ClearanceContract) ClearanceExists(ctx contractapi.TransactionContextInterface, clearanceID string) (bool, error) {
	data, err := ctx.GetStub().GetState(clearanceID)

	if err != nil {
		return false, err
	}

	return data != nil, nil
}

// CreateClearance creates a new instance of Clearance
func (c *ClearanceContract) CreateClearance(ctx contractapi.TransactionContextInterface, clearanceID string, payload ArrivalClearance) error {
	exists, err := c.ClearanceExists(ctx, clearanceID)
	if err != nil {
		return fmt.Errorf("Could not read from world state. %s", err)
	} else if exists {
		return fmt.Errorf("The asset %s already exists", clearanceID)
	}

	exists1, err1 := c.AssetExists(ctx, payload.vesselInfo)
	if err1 != nil {
		return fmt.Errorf("Could not read from world state. %s", err1)
	} else if !exists1 {
		return fmt.Errorf("The vessel %s does not exist", payload.vesselInfo)
	}

	var passportCheck, passportCheckMsg = c.CrewMembersCheck(ctx, payload.crewList)
	var certCheck, certCheckMsg = c.CertCheck(ctx, payload.certificates)
	var clearanceStatus bool
	var clearanceMsg string
	if passportCheck && certCheck {
		clearanceStatus = true
		clearanceMsg = "Cleared"
	} else {
		clearanceMsg = passportCheckMsg + certCheckMsg
	}

	clearance := new(Clearance)
	clearance.Payload = payload
	clearance.Status = clearanceStatus
	clearance.Message = clearanceMsg
	bytes, _ := json.Marshal(clearance)

	return ctx.GetStub().PutState(clearanceID, bytes)
}

// ReadClearance retrieves an instance of Clearance from the world state
func (c *ClearanceContract) ReadClearance(ctx contractapi.TransactionContextInterface, clearanceID string) (*Clearance, error) {
	exists, err := c.ClearanceExists(ctx, clearanceID)
	if err != nil {
		return nil, fmt.Errorf("Could not read from world state. %s", err)
	} else if !exists {
		return nil, fmt.Errorf("The asset %s does not exist", clearanceID)
	}

	bytes, _ := ctx.GetStub().GetState(clearanceID)

	clearance := new(Clearance)

	err = json.Unmarshal(bytes, clearance)

	if err != nil {
		return nil, fmt.Errorf("Could not unmarshal world state data to type Clearance")
	}

	return clearance, nil
}
