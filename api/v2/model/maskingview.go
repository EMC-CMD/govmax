package model

type MaskingViewResponse struct {
	MaskingView []struct {
		MaskingViewID string `json:"maskingViewId"`
		HostID string `json:"hostId"`
		HostGroupID string `json:"hostGroupId"`
		PortGroupID string `json:"portGroupId"`
		StorageGroupID string `json:"storageGroupId"`
	} `json:"maskingView"`
	Success bool `json:"success"`
	Message string `json:"message"`
}
