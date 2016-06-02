package model

type SymmetrixsResponse struct {
	SymmetrixID          []string `json:"symmetrixId"`
	NumOfSymmetrixArrays int      `json:"num_of_symmetrix_arrays"`
	Success              bool     `json:"success"`
}

type SymmetrixResponse struct {
	Symmetrix []struct {
		SymmetrixID   string `json:"symmetrixId"`
		DeviceCount   int    `json:"device_count"`
		Ucode         string `json:"ucode"`
		Model         string `json:"model"`
		Local         bool   `json:"local"`
		SloCompliance struct {
			SloStable   int `json:"slo_stable"`
			SloMarginal int `json:"slo_marginal"`
			SloCritical int `json:"slo_critical"`
		} `json:"sloCompliance"`
		VirtualCapacity struct {
			UsedCapacityGb  float64 `json:"used_capacity_gb"`
			TotalCapacityGb float64 `json:"total_capacity_gb"`
		} `json:"virtualCapacity"`
	} `json:"symmetrix"`
	Success bool `json:"success"`
}

type Symmetrix struct {
	SymmetrixID string
}
