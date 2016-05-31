package model

type ListOfSLOs struct {
	NumSlos int      `json:"numSlos"`
	SloID   []string `json:"sloId"`
}
