package model

type ListOfSRPs struct {
	NumSrps int      `json:"numSrps"`
	SrpID   []string `json:"srpId"`
}
