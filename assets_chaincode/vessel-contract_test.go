/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const getStateError = "world state get error"

type MockStub struct {
	shim.ChaincodeStubInterface
	mock.Mock
}

func (ms *MockStub) GetState(key string) ([]byte, error) {
	args := ms.Called(key)

	return args.Get(0).([]byte), args.Error(1)
}

func (ms *MockStub) PutState(key string, value []byte) error {
	args := ms.Called(key, value)

	return args.Error(0)
}

func (ms *MockStub) DelState(key string) error {
	args := ms.Called(key)

	return args.Error(0)
}

type MockContext struct {
	contractapi.TransactionContextInterface
	mock.Mock
}

func (mc *MockContext) GetStub() shim.ChaincodeStubInterface {
	args := mc.Called()

	return args.Get(0).(*MockStub)
}

func configureStub() (*MockContext, *MockStub) {
	var nilBytes []byte

	testVessel := new(Vessel)
	testVessel.Name = "set value"
	vesselBytes, _ := json.Marshal(testVessel)

	ms := new(MockStub)
	ms.On("GetState", "statebad").Return(nilBytes, errors.New(getStateError))
	ms.On("GetState", "missingkey").Return(nilBytes, nil)
	ms.On("GetState", "existingkey").Return([]byte("some value"), nil)
	ms.On("GetState", "vesselkey").Return(vesselBytes, nil)
	ms.On("PutState", mock.AnythingOfType("string"), mock.AnythingOfType("[]uint8")).Return(nil)
	ms.On("DelState", mock.AnythingOfType("string")).Return(nil)

	mc := new(MockContext)
	mc.On("GetStub").Return(ms)

	return mc, ms
}

func TestVesselExists(t *testing.T) {
	var exists bool
	var err error

	ctx, _ := configureStub()
	c := new(VesselContract)

	exists, err = c.AssetExists(ctx, "statebad")
	assert.EqualError(t, err, getStateError)
	assert.False(t, exists, "should return false on error")

	exists, err = c.AssetExists(ctx, "missingkey")
	assert.Nil(t, err, "should not return error when can read from world state but no value for key")
	assert.False(t, exists, "should return false when no value for key in world state")

	exists, err = c.AssetExists(ctx, "existingkey")
	assert.Nil(t, err, "should not return error when can read from world state and value exists for key")
	assert.True(t, exists, "should return true when value for key in world state")
}

func TestCreateVessel(t *testing.T) {
	var err error

	ctx, stub := configureStub()
	c := new(VesselContract)

	err = c.CreateVessel(ctx, "statebad", "some value")
	assert.EqualError(t, err, fmt.Sprintf("Could not read from world state. %s", getStateError), "should error when exists errors")

	err = c.CreateVessel(ctx, "existingkey", "some value")
	assert.EqualError(t, err, "The asset existingkey already exists", "should error when exists returns true")

	err = c.CreateVessel(ctx, "missingkey", "some value")
	stub.AssertCalled(t, "PutState", "missingkey", []byte("{\"value\":\"some value\"}"))
}

func TestReadVessel(t *testing.T) {
	var vessel *Vessel
	var err error

	ctx, _ := configureStub()
	c := new(VesselContract)

	vessel, err = c.GetVessel(ctx, "statebad")
	assert.EqualError(t, err, fmt.Sprintf("Could not read from world state. %s", getStateError), "should error when exists errors when reading")
	assert.Nil(t, vessel, "should not return Vessel when exists errors when reading")

	vessel, err = c.GetVessel(ctx, "missingkey")
	assert.EqualError(t, err, "The asset missingkey does not exist", "should error when exists returns true when reading")
	assert.Nil(t, vessel, "should not return Vessel when key does not exist in world state when reading")

	vessel, err = c.GetVessel(ctx, "existingkey")
	assert.EqualError(t, err, "Could not unmarshal world state data to type Vessel", "should error when data in key is not Vessel")
	assert.Nil(t, vessel, "should not return Vessel when data in key is not of type Vessel")

	vessel, err = c.GetVessel(ctx, "vesselkey")
	expectedVessel := new(Vessel)
	expectedVessel.Name = "set value"
	assert.Nil(t, err, "should not return error when Vessel exists in world state when reading")
	assert.Equal(t, expectedVessel, vessel, "should return deserialized Vessel from world state")
}

func TestDeleteVessel(t *testing.T) {
	var err error

	ctx, stub := configureStub()
	c := new(VesselContract)

	err = c.DeleteVessel(ctx, "statebad")
	assert.EqualError(t, err, fmt.Sprintf("Could not read from world state. %s", getStateError), "should error when exists errors")

	err = c.DeleteVessel(ctx, "missingkey")
	assert.EqualError(t, err, "The asset missingkey does not exist", "should error when exists returns true when deleting")

	err = c.DeleteVessel(ctx, "vesselkey")
	assert.Nil(t, err, "should not return error when Vessel exists in world state when deleting")
	stub.AssertCalled(t, "DelState", "vesselkey")
}
