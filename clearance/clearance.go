/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

// Clearance stores a value
type Clearance struct {
	Payload ArrivalClearance
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

// Vessel stores a value
type Vessel struct {
	Name      string `json:"value"`
	ImoNumber string `json:"imoNumber"`
}

type ArrivalClearance struct {
	vesselInfo   string   `json:"vesselInfo"`
	crewList     []string `json:"crewList"`
	declarations VesselDeclarations
	certificates []string `json:"certificates"`
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
	ArmsStrongRoom string `json:"armsInStrongRoom"`
}

type BunkerDeclaration struct {
	ArrRobFO string `json:"arrivalRobFO"`
}

type CargoDeclaration struct {
	CargoDescription string `json:"cargoDescription"`
}

type VesselCertificates struct {
	CertificateId  string `json:"id"`
	Name           string `json:"certName"`
	IssueAuthority string `json:"certIssueAuthority"`
	ExpiryDate     string `json:"certExpiry"`
}

type VoyageInformation struct {
	LastPort string `json"lastPort"`
	NextPort string `json"nextort"`
}
