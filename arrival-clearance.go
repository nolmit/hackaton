/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

// ArrivalClearance stores a value
type ArrivalClearance struct {
	vesselInfo VesselParticulars
	crewList []CrewDetails
	declarations VesselDeclarations
	certificates VesselCertificates
	voyageInfo VoyageInformation
}

type VesselParticulars struct {
	ImoNumber      string `json:"imoNumber"`
	Name           string `json:"name"`
	CallSign       string `json:"callSign"`
	FlagState      string `json:"flagState"`
	VesselType     string `json:"vesselType"`
	Grt            string `json:"grt"`
	OwnerName      string `json:"ownerName"`
	OwnerEmail     string `json:"ownerEmail"`
	AgentName      string `json:"agentName"`
	AgentEmail     string `json:"agentEmail"`
	ShipInmarsatNo string `json:"shipInmarsatNo"`
	ShipEmail      string `json:"shipEmail"`
}

type CrewDetails struct {
	SeamanID         string `json:"id"`
	Rank             string `json:"rank"`
	Name             string `json:"name"`
	Nationality      string `json:"nationality"`
	Dob              string `json:"dateOfBirth"`
	Pob              string `json:"placeOfBirth"`
	PassportNumber   string `json:"passportNumber"`
	PassportExpiry   string `json:"passportExpiry"`
	SeamanbookNumber string `json:"seamanBookNumber"`
	SeamanbookExpiry string `json:"seamanBookExpiry"`
}

type VesselDeclarations struct {
	security SecurityDeclaration
	bunker BunkerDeclaration
	cargo CargoDeclaration
}

type SecurityDeclaration struct {
	SecurityLevel string `json:"vesselSecurityLevel"`
	NumOfArms string `json:"armsAndAmmunition"`
	ArmsType string `json:"armsType"`
	ArmsStrongRoom string `json:"armsInStrongRoom"`
}

type BunkerDeclaration struct {
	ArrRobFO string `json:"arrivalRobFO"`
	DepRobFO string `json:"departureRobFO"`
	BunkerFO string `json:"bunkeringFO"`
	ArrRobDO string `json:"arrivalRobDO"`
	DepRobDO string `json:"departureRobDO"`
	BunkerDO string `json:"bunkeringDO"`
}

type CargoDeclaration struct {
	CargoDescription string `json:"cargoDescription"`
    HasDGCargo string `json:"hasDGCargo"`
    DgWeight string `json:"dgCargoWeight"`
    TotalCargo string `json:"totalCargoOnboard"`
    TotalDisch string `json:"cargoToDischarge"`
    TotalLoad string `json:"cargoToLoad"`
}

type VesselCertificates struct {
	CertificateId string `json:"id"`
	Name string `json:"certName"`
	ExpiryDate string `json:"certExpiry"`
	IssueAuthority string `json:"certIssueAuthority"`
}

type VoyageInformation struct {
	LastPort string `json"lastPort"`
	LastFacility string `json"lastPortFacility"`
	Eta string `json"expectedArrival"`
	Etd string `json"expectedDeparture"`
	Purpose string `json"purposeOfCall"`
	ArrivalPort string `json"arrivalPort"`
	NextPort string `json"nextPort"`
}