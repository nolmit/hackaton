/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-contract-api-go/metadata"
)

func main() {
	clearanceContract := new(ClearanceContract)
	clearanceContract.Info.Version = "0.0.1"
	clearanceContract.Info.Description = "My Smart Contract"
	clearanceContract.Info.License = new(metadata.LicenseMetadata)
	clearanceContract.Info.License.Name = "Apache-2.0"
	clearanceContract.Info.Contact = new(metadata.ContactMetadata)
	clearanceContract.Info.Contact.Name = "John Doe"

	chaincode, err := contractapi.NewChaincode(clearanceContract)
	chaincode.Info.Title = "clearance chaincode"
	chaincode.Info.Version = "0.0.1"

	if err != nil {
		panic("Could not create chaincode from ClearanceContract." + err.Error())
	}

	err = chaincode.Start()

	if err != nil {
		panic("Failed to start chaincode. " + err.Error())
	}
}
