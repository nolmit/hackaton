/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

// Vessel stores a value
type Vessel struct {
	Name      string `json:"value"`
	ImoNumber string `json:"imoNumber"`
}

type ArrivalClearance struct {
	vesselInfo   VesselParticulars
	crewList     []CrewDetails
	declarations VesselDeclarations
	certificates VesselCertificates
	voyageInfo   VoyageInformation
}

type VesselParticulars struct {
	ImoNumber string `json:"imoNumber"`
	Name      string `json:"name"`
}

type CrewDetails struct {
	SeamanID       string `json:"id"`
	Name           string `json:"name"`
	PassportNumber string `json:"passportNumber"`
	PassportExpiry string `json:"passportExpiry"`
}

type VesselDeclarations struct {
	security SecurityDeclaration
	bunker   BunkerDeclaration
	cargo    CargoDeclaration
}

type SecurityDeclaration struct {
	SecurityLevel  string `json:"vesselSecurityLevel"`
	NumOfArms      string `json:"armsAndAmmunition"`
	ArmsType       string `json:"armsType"`
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
	HasDGCargo       string `json:"hasDGCargo"`
	DgWeight         string `json:"dgCargoWeight"`
	TotalCargo       string `json:"totalCargoOnboard"`
	TotalDisch       string `json:"cargoToDischarge"`
	TotalLoad        string `json:"cargoToLoad"`
}

type VesselCertificates struct {
	CertificateId  string `json:"id"`
	Name           string `json:"certName"`
	IssueAuthority string `json:"certIssueAuthority"`
	ExpiryDate     string `json:"certExpiry"`
}

type VoyageInformation struct {
	LastPort     string `json"lastPort"`
	LastFacility string `json"lastPortFacility"`
	Eta          string `json"expectedArrival"`
	Etd          string `json"expectedDeparture"`
	Purpose      string `json"purposeOfCall"`
	ArrivalPort  string `json"arrivalPort"`
	NextPort     string `json"nextPort"`
}
