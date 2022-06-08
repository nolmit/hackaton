/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// VesselContract contract for managing CRUD for Vessel
type VesselContract struct {
	contractapi.Contract
}

// VesselExists returns true when asset with given ID exists in world state
func (c *VesselContract) VesselExists(ctx contractapi.TransactionContextInterface, vesselID string) (bool, error) {
	data, err := ctx.GetStub().GetState(vesselID)

	if err != nil {
		return false, err
	}

	return data != nil, nil
}

// CreateVessel creates a new instance of Vessel
/*
func (c *VesselContract) CreateVessel(ctx contractapi.TransactionContextInterface, vesselID string, 
	imonumber string, name string, callsign string, flagstate string) error {
	exists, err := c.VesselExists(ctx, vesselID)
	if err != nil {
		return fmt.Errorf("Could not read from world state. %s", err)
	} else if exists {
		return fmt.Errorf("The asset %s already exists", vesselID)
	}

	vessel := Vessel{
		IMONumber: imonumber,
		Name: name,
		CallSign: callsign,
		FlagState: flagstate,
	}
	bytes, _ := json.Marshal(vessel)

	return ctx.GetStub().PutState(vesselID, bytes)
}
*/

// ReadVessel retrieves an instance of Vessel from the world state
func (c *VesselContract) ReadVessel(ctx contractapi.TransactionContextInterface, vesselID string) (*Vessel, error) {
	exists, err := c.VesselExists(ctx, vesselID)
	if err != nil {
		return nil, fmt.Errorf("Could not read from world state. %s", err)
	} else if !exists {
		return nil, fmt.Errorf("The asset %s does not exist", vesselID)
	}

	bytes, _ := ctx.GetStub().GetState(vesselID)

	vessel := new(Vessel)

	err = json.Unmarshal(bytes, vessel)

	if err != nil {
		return nil, fmt.Errorf("Could not unmarshal world state data to type Vessel")
	}

	return vessel, nil
}

// UpdateVessel retrieves an instance of Vessel from the world state and updates its value
func (c *VesselContract) UpdateVessel(ctx contractapi.TransactionContextInterface, vesselID string, newValue string) error {
	exists, err := c.VesselExists(ctx, vesselID)
	if err != nil {
		return fmt.Errorf("Could not read from world state. %s", err)
	} else if !exists {
		return fmt.Errorf("The asset %s does not exist", vesselID)
	}

	vessel := new(Vessel)
	//vessel.Value = newValue

	bytes, _ := json.Marshal(vessel)

	return ctx.GetStub().PutState(vesselID, bytes)
}

// DeleteVessel deletes an instance of Vessel from the world state
func (c *VesselContract) DeleteVessel(ctx contractapi.TransactionContextInterface, vesselID string) error {
	exists, err := c.VesselExists(ctx, vesselID)
	if err != nil {
		return fmt.Errorf("Could not read from world state. %s", err)
	} else if !exists {
		return fmt.Errorf("The asset %s does not exist", vesselID)
	}

	return ctx.GetStub().DelState(vesselID)
}

// InitVessels adds a base set of cars to the ledger
func (s *VesselContract) InitVessels(ctx contractapi.TransactionContextInterface) error {
	vessels := []Vessel{
		{IMONumber: "9876541", Name: "MV Good Hope", CallSign: "XXDF", FlagState: "Norway", VesselType: "Container", GRT: "12569", Owner: "Own Ship Holding", OwnerEmail: "owner@osh.no", Agent: "Smoke Lad", AgentEmail: "smoke@agentx.com", InMarSatNo: "987-65646321", ShipEmail: "master@goodhope.osh.no" },
		{IMONumber: "9876542", Name: "MV Good Hope", CallSign: "XXDF", FlagState: "Norway", VesselType: "Container", GRT: "12569", Owner: "Own Ship Holding", OwnerEmail: "owner@osh.no", Agent: "Smoke Lad", AgentEmail: "smoke@agentx.com", InMarSatNo: "987-65646321", ShipEmail: "master@goodhope.osh.no" },
	}

	for i, vessel := range vessels {
		vesselAsBytes, _ := json.Marshal(vessel)
		err := ctx.GetStub().PutState("VESSEL"+strconv.Itoa(i), vesselAsBytes)

		if err != nil {
			return fmt.Errorf("Failed to put to world state. %s", err.Error())
		}
	}

	return nil
}
