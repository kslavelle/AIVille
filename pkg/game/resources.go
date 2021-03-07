package game

// EnergyResource represents a resource that produces energy
type EnergyResource struct {
	Cost     int `json:"cost"`
	CO2      int `json:"co2/hr"`
	MWOutput int `json:"Mw/hr"`
	Workers  int `json:"workers"`
}
