/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-contract-api-go/metadata"
)

func main() {
	vesselContract := new(VesselContract)
	vesselContract.Info.Version = "0.0.1"
	vesselContract.Info.Description = "My Smart Contract"
	vesselContract.Info.License = new(metadata.LicenseMetadata)
	vesselContract.Info.License.Name = "Apache-2.0"
	vesselContract.Info.Contact = new(metadata.ContactMetadata)
	vesselContract.Info.Contact.Name = "John Doe"

	chaincode, err := contractapi.NewChaincode(vesselContract)
	chaincode.Info.Title = "test_contract chaincode"
	chaincode.Info.Version = "0.0.1"

	if err != nil {
		panic("Could not create chaincode from VesselContract." + err.Error())
	}

	err = chaincode.Start()

	if err != nil {
		panic("Failed to start chaincode. " + err.Error())
	}
}
