package model

type SloBasedStorageGroupParam struct {
	SloID                                           string          `json:"sloId"`
	WorkloadSelection                               string          `json:"workloadSelection"`
	NumOfVols                                       int             `json:"num_of_vols"`
	VolumeAttribute                                 VolumeAttribute `json:"volumeAttribute"`
	AllocateCapacityForEachVol                      bool            `json:"allocate_capacity_for_each_vol"`
	PersistPreallocatedCapacityThroughReclaimOrCopy bool            `json:"persist_preallocated_capacity_through_reclaim_or_copy"`
}

type CreateStorageGroupParam struct {
	StorageGroupID            string                      `json:"storageGroupId"`
	CreateEmptyStorageGroup   bool                        `json:"create_empty_storage_group"`
	SrpID                     string                      `json:"srpId"`
	SloBasedStorageGroupParam []SloBasedStorageGroupParam `json:"sloBasedStorageGroupParam"`
}

type ListStorageGroupResponse struct {
	NumOfStorageGroups int      `json:"num_of_storage_groups"`
	StorageGroupID     []string `json:"storageGroupId"`
	Success            bool     `json:"success"`
}

type GetStorageGroupResult struct {
	StorageGroup []struct {
		StorageGroupID     string   `json:"storageGroupId"`
		Slo                string   `json:"slo"`
		Srp                string   `json:"srp"`
		Workload           string   `json:"workload"`
		SloCompliance      string   `json:"slo_compliance"`
		NumOfVols          int      `json:"num_of_vols"`
		NumOfChildSgs      int      `json:"num_of_child_sgs"`
		NumOfParentSgs     int      `json:"num_of_parent_sgs"`
		NumOfMaskingViews  int      `json:"num_of_masking_views"`
		NumOfSnapshots     int      `json:"num_of_snapshots"`
		CapGb              float64  `json:"cap_gb"`
		DeviceEmulation    string   `json:"device_emulation"`
		Type               string   `json:"type"`
		Unprotected        bool     `json:"unprotected"`
		ChildStorageGroup  []string `json:"child_storage_group"`
		ParentStorageGroup []string `json:"parent_storage_group"`
		Maskingview        []string `json:"maskingview"`
		HostIOLimit        struct {
			HostIoLimitMbSec    string `json:"host_io_limit_mb_sec"`
			HostIoLimitIoSec    string `json:"host_io_limit_io_sec"`
			DynamicDistribution string `json:"dynamicDistribution"`
		} `json:"hostIOLimit"`
	} `json:"storageGroup"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type EditStorageGroupParam struct {
	EditStorageGroupActionParam EditStorageGroupActionParam `json:"editStorageGroupActionParam"`
}

type EditStorageGroupActionParam struct {
	ExpandStorageGroupParam *ExpandStorageGroupParam `json:"expandStorageGroupParam,omitempty"`
	RemoveVolumeParam       *RemoveVolumeParam       `json:"removeVolumeParam,omitempty"`
	AddVolumeParam          *AddVolumeParam          `json:"addVolumeParam,omitempty"`
}

type ExpandStorageGroupParam struct {
	NumOfVols        int             `json:"num_of_vols"`
	VolumeAttribute  VolumeAttribute `json:"volumeAttribute"`
	CreateNewVolumes bool            `json:"create_new_volumes"`
	Emulation        string          `json:"emulation"`
}

type VolumeAttribute struct {
	CapacityUnit string `json:"capacityUnit"`
	VolumeSize   string `json:"volume_size"`
}

type EditStorageGroupResponse struct {
	StorageGroupID    string  `json:"storageGroupId"`
	Slo               string  `json:"slo"`
	Srp               string  `json:"srp"`
	Workload          string  `json:"workload"`
	SloCompliance     string  `json:"slo_compliance"`
	NumOfVols         int     `json:"num_of_vols"`
	NumOfChildSgs     int     `json:"num_of_child_sgs"`
	NumOfParentSgs    int     `json:"num_of_parent_sgs"`
	NumOfMaskingViews int     `json:"num_of_masking_views"`
	NumOfSnapshots    int     `json:"num_of_snapshots"`
	CapGb             float64 `json:"cap_gb"`
	DeviceEmulation   string  `json:"device_emulation"`
	Type              string  `json:"type"`
	Unprotected       bool    `json:"unprotected"`
}

type RemoveVolumeParam struct {
	VolumeID []string `json:"volumeId"`
}

type AddVolumeParam struct {
	VolumeID []string `json:"volumeId"`
}
