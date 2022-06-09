package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func (s *VesselContract) InitCrew(ctx contractapi.TransactionContextInterface, crewList []CrewDetails) error {

	for _, asset := range crewList {
		assetJSON, err := json.Marshal(asset)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(asset.PassportNumber, assetJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}

	return nil
}

func (c *VesselContract) GetCrewMember(ctx contractapi.TransactionContextInterface, memberId string) (*CrewDetails, error) {
	exists, err := c.AssetExists(ctx, memberId)
	if err != nil {
		return nil, fmt.Errorf("Could not read from world state. %s", err)
	} else if !exists {
		return nil, fmt.Errorf("The asset %s does not exist", memberId)
	}

	bytes, _ := ctx.GetStub().GetState(memberId)

	vessel := new(CrewDetails)

	err = json.Unmarshal(bytes, vessel)

	if err != nil {
		return nil, fmt.Errorf("Could not unmarshal world state data to type Vessel")
	}

	return vessel, nil
}

// GetAllAssets returns all assets found in world state
func (s *VesselContract) GetAllCrew(ctx contractapi.TransactionContextInterface) ([]*CrewDetails, error) {
	// range query with empty string for startKey and endKey does an
	// open-ended query of all assets in the chaincode namespace.
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var assets []*CrewDetails
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var asset CrewDetails
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return nil, err
		}
		assets = append(assets, &asset)
	}

	return assets, nil
}
