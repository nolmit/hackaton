/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

// Vessel stores a value
type Vessel struct {
	IMONumber string `json:"imoNumber"`
	Name string `json:"name"`
	CallSign string `json:"callSign"`
    FlagState string `json:"flagState"`
	VesselType string `json:"vesselType"`
	GRT string `json:"grt"`
	Owner string `json:"ownerName"`
	OwnerEmail string `json:"ownerEmail"`
	Agent string `json:"agentName"`
	AgentEmail string `json:"agentEmail"`
	InMarSatNo string `json:"shipInmarsatNo"`
	ShipEmail string `json:"shipEmail"`
}