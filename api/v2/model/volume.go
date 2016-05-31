package model

type CreateVolumeParam struct {
}

type VolumeResponse struct {
	Volume []struct {
		VolumeID           string  `json:"volumeId"`
		Type               string  `json:"type"`
		Emulation          string  `json:"emulation"`
		Ssid               string  `json:"ssid"`
		AllocatedPercent   int     `json:"allocated_percent"`
		CapGb              float64 `json:"cap_gb"`
		CapMb              float64 `json:"cap_mb"`
		CapCyl             int     `json:"cap_cyl"`
		Status             string  `json:"status"`
		Reserved           bool    `json:"reserved"`
		Pinned             bool    `json:"pinned"`
		PhysicalName       string  `json:"physical_name"`
		VolumeIdentifier   string  `json:"volume_identifier"`
		Wwn                string  `json:"wwn"`
		Encapsulated       bool    `json:"encapsulated"`
		NumOfStorageGroups int     `json:"num_of_storage_groups"`
		NumOfFrontEndPaths int     `json:"num_of_front_end_paths"`
		SymmetrixPortKey   []struct {
			DirectorID string `json:"directorId"`
			PortID     string `json:"portId"`
		} `json:"symmetrixPortKey"`
	} `json:"volume"`
	Success bool `json:"success"`
}

type ListVolumeResponse struct {
	ResultList struct {
		Result []struct {
			VolumeID string `json:"volumeId"`
		} `json:"result"`
		From int `json:"from"`
		To   int `json:"to"`
	} `json:"resultList"`
	ID             string `json:"id"`
	Count          int    `json:"count"`
	ExpirationTime int64  `json:"expirationTime"`
	MaxPageSize    int    `json:"maxPageSize"`
}
