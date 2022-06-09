package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func (s *VesselContract) InitVessels(ctx contractapi.TransactionContextInterface, crewList []VesselParticulars) error {

	for _, asset := range crewList {
		assetJSON, err := json.Marshal(asset)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(asset.ImoNumber, assetJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}

	return nil
}

func (c *VesselContract) GetVessel(ctx contractapi.TransactionContextInterface, vesselID string) (*Vessel, error) {
	exists, err := c.AssetExists(ctx, vesselID)
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

// CreateVessel creates a new instance of Vessel
func (c *VesselContract) CreateVessel(ctx contractapi.TransactionContextInterface, vesselID string, name string) error {
	exists, err := c.AssetExists(ctx, vesselID)
	if err != nil {
		return fmt.Errorf("Could not read from world state. %s", err)
	} else if exists {
		return fmt.Errorf("The asset %s already exists", vesselID)
	}

	vessel := new(Vessel)
	vessel.Name = name
	vessel.ImoNumber = vesselID

	bytes, _ := json.Marshal(vessel)

	return ctx.GetStub().PutState(vesselID, bytes)
}

func (c *VesselContract) DeleteVessel(ctx contractapi.TransactionContextInterface, vesselID string) error {
	exists, err := c.AssetExists(ctx, vesselID)
	if err != nil {
		return fmt.Errorf("Could not read from world state. %s", err)
	} else if !exists {
		return fmt.Errorf("The asset %s does not exist", vesselID)
	}

	return ctx.GetStub().DelState(vesselID)
}
