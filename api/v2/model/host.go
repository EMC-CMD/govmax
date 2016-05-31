package model

type CreateHostParam struct {
	HostID string `json:"hostId"`
	InitiatorID []string `json:"initiatorId"`
	HostFlags struct {
		VolumeSetAddressing struct {
			Enabled bool `json:"enabled"`
			Override bool `json:"override"`
		} `json:"volume_set_addressing"`
		CommonSerialNumber struct {
			Enabled bool `json:"enabled"`
			Override bool `json:"override"`
		} `json:"common_serial_number"`
		DisableQResetOnUa struct {
			Enabled bool `json:"enabled"`
			Override bool `json:"override"`
		} `json:"disable_q_reset_on_ua"`
		EnvironSet struct {
			Enabled bool `json:"enabled"`
			Override bool `json:"override"`
		} `json:"environ_set"`
		AvoidResetBroadcast struct {
			Enabled bool `json:"enabled"`
			Override bool `json:"override"`
		} `json:"avoid_reset_broadcast"`
		As400 struct {
			Enabled bool `json:"enabled"`
			Override bool `json:"override"`
		} `json:"as400"`
		Openvms struct {
			Enabled bool `json:"enabled"`
			Override bool `json:"override"`
		} `json:"openvms"`
		Scsi3 struct {
			Enabled bool `json:"enabled"`
			Override bool `json:"override"`
		} `json:"scsi_3"`
		Spc2ProtocolVersion struct {
			Enabled bool `json:"enabled"`
			Override bool `json:"override"`
		} `json:"spc2_protocol_version"`
		ScsiSupport1 struct {
			Enabled bool `json:"enabled"`
			Override bool `json:"override"`
		} `json:"scsi_support1"`
		ConsistentLun bool `json:"consistent_lun"`
	} `json:"hostFlags"`
}
