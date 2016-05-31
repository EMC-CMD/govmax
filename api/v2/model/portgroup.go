package model

type CreatePortGroupParam struct {
	PortGroupID string `json:"portGroupId"`
	SymmetrixPortKey []struct {
	} `json:"symmetrixPortKey"`
}
