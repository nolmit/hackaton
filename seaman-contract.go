/*
 * SPDX-License-Identifier: Apache-2.0
 */

 package main

 import (
	 "encoding/json"
	 "fmt"
 
	 "github.com/hyperledger/fabric-contract-api-go/contractapi"
 )
 
 // SeamanContract contract for managing CRUD for Vessel
type SeamanContract struct {
	contractapi.Contract
}

// VesselExists returns true when asset with given ID exists in world state
func (c *VesselContract) SeamanExists(ctx contractapi.TransactionContextInterface, seamanID string) (bool, error) {
	data, err := ctx.GetStub().GetState(seamanID)

	if err != nil {
		return false, err
	}

	return data != nil, nil
}

// CreateVessel creates a new instance of Vessel
func (c *VesselContract) CreateSeaman(ctx contractapi.TransactionContextInterface, id string, rank string, name string, nationality string, dateofbirth string, placeofbirth string, passportnumber string, passportexpiry string, seamanbooknumber string,
	seamanbookexpiry string) error {
	exists, err := c.VesselExists(ctx, id)
	if err != nil {
		return fmt.Errorf("Could not read from world state. %s", err)
	} else if exists {
		return fmt.Errorf("The asset %s already exists", id)
	}

	seaman := Seaman{
		ID: id,
		Rank: rank,
		Name: name,
		Nationality: nationality,
		DateOfBirth: dateofbirth,
		PlaceofBirth: placeofbirth,
		PassportNumber: passportnumber,
		PassportExpiry: passportexpiry,
		SeamanBookNumber: seamanbooknumber,
		SeamanBookExpiry: seamanbookexpiry,
	}
	bytes, _ := json.Marshal(seaman)

	return ctx.GetStub().PutState(id, bytes)
}