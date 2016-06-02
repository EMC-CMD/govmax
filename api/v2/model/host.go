package model

type CreateHostParam struct {
	HostID      string   `json:"hostId"`
	InitiatorID []string `json:"initiatorId"`
	HostFlags   struct {
		VolumeSetAddressing struct {
			Enabled  bool `json:"enabled"`
			Override bool `json:"override"`
		} `json:"volume_set_addressing"`
		CommonSerialNumber struct {
			Enabled  bool `json:"enabled"`
			Override bool `json:"override"`
		} `json:"common_serial_number"`
		DisableQResetOnUa struct {
			Enabled  bool `json:"enabled"`
			Override bool `json:"override"`
		} `json:"disable_q_reset_on_ua"`
		EnvironSet struct {
			Enabled  bool `json:"enabled"`
			Override bool `json:"override"`
		} `json:"environ_set"`
		AvoidResetBroadcast struct {
			Enabled  bool `json:"enabled"`
			Override bool `json:"override"`
		} `json:"avoid_reset_broadcast"`
		As400 struct {
			Enabled  bool `json:"enabled"`
			Override bool `json:"override"`
		} `json:"as400"`
		Openvms struct {
			Enabled  bool `json:"enabled"`
			Override bool `json:"override"`
		} `json:"openvms"`
		Scsi3 struct {
			Enabled  bool `json:"enabled"`
			Override bool `json:"override"`
		} `json:"scsi_3"`
		Spc2ProtocolVersion struct {
			Enabled  bool `json:"enabled"`
			Override bool `json:"override"`
		} `json:"spc2_protocol_version"`
		ScsiSupport1 struct {
			Enabled  bool `json:"enabled"`
			Override bool `json:"override"`
		} `json:"scsi_support1"`
		ConsistentLun bool `json:"consistent_lun"`
	} `json:"hostFlags"`
}

type ListOfHostsResponse struct {
	NumOfHosts int      `json:"num_of_hosts"`
	HostID     []string `json:"hostId"`
	Success    bool     `json:"success"`
}

type HostResponse struct {
	Host []struct {
		HostID            string   `json:"hostId"`
		NumOfMaskingViews int      `json:"num_of_masking_views"`
		NumOfInitiators   int      `json:"num_of_initiators"`
		NumOfHostGroups   int      `json:"num_of_host_groups"`
		PortFlagsOverride bool     `json:"port_flags_override"`
		ConsistentLun     bool     `json:"consistent_lun"`
		EnabledFlags      string   `json:"enabled_flags"`
		DisabledFlags     string   `json:"disabled_flags"`
		Type              string   `json:"type"`
		Initiator         []string `json:"initiator"`
		Maskingview       []string `json:"maskingview"`
	} `json:"host"`
	Success bool `json:"success"`
}
