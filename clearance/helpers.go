package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func (c *ClearanceContract) AssetExists(ctx contractapi.TransactionContextInterface, memberId string) (bool, error) {
	data, err := ctx.GetStub().GetState(memberId)

	if err != nil {
		return false, err
	}

	return data != nil, nil
}

func (c *ClearanceContract) GetCrewMember(ctx contractapi.TransactionContextInterface, memberId string) (*CrewDetails, error) {
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

func (c *ClearanceContract) GetVessel(ctx contractapi.TransactionContextInterface, vesselID string) (*Vessel, error) {
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

func (c *ClearanceContract) GetCertificate(ctx contractapi.TransactionContextInterface, memberId string) (*VesselCertificates, error) {
	exists, err := c.AssetExists(ctx, memberId)
	if err != nil {
		return nil, fmt.Errorf("Could not read from world state. %s", err)
	} else if !exists {
		return nil, fmt.Errorf("The asset %s does not exist", memberId)
	}

	bytes, _ := ctx.GetStub().GetState(memberId)

	vessel := new(VesselCertificates)

	err = json.Unmarshal(bytes, vessel)

	if err != nil {
		return nil, fmt.Errorf("Could not unmarshal world state data to type Vessel")
	}

	return vessel, nil
}

func (c *ClearanceContract) CrewMembersCheck(ctx contractapi.TransactionContextInterface, pass []string) (bool, string) {
	var badCrewMamber []string
	for _, el := range pass {
		var member, _ = c.GetCrewMember(ctx, el)
		var parsed, _ = time.Parse("YYYY-MM-DD", member.PassportExpiry)
		isValid := time.Now().After(parsed)
		if !isValid {
			badCrewMamber = append(badCrewMamber, member.Name)
		}

	}
	var res = len(badCrewMamber) <= 0
	if res {
		return true, ""
	}

	return false, fmt.Sprintln("Expired documents on crew mwmbers:", strings.Join(badCrewMamber, ""))
}

func (c *ClearanceContract) CertCheck(ctx contractapi.TransactionContextInterface, pass []string) (bool, string) {
	var badCrewMamber []string
	for _, el := range pass {
		var member, _ = c.GetCertificate(ctx, el)
		var parsed, _ = time.Parse("YYYY-MM-DD", member.ExpiryDate)
		isValid := time.Now().After(parsed)
		if !isValid {
			badCrewMamber = append(badCrewMamber, member.Name)
		}
	}
	var res = len(badCrewMamber) <= 0
	if res {
		return true, ""
	}

	return false, fmt.Sprintln("Expired certificates:", strings.Join(badCrewMamber, ""))
}
