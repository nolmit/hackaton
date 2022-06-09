/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// VesselContract contract for managing CRUD for Vessel
type VesselContract struct {
	contractapi.Contract
}

func (c *VesselContract) AssetExists(ctx contractapi.TransactionContextInterface, memberId string) (bool, error) {
	data, err := ctx.GetStub().GetState(memberId)

	if err != nil {
		return false, err
	}

	return data != nil, nil
}
