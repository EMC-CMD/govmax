// POST /provisioning/symmetrix/{symmetrixId}/portgroup
// This call creates a new Port Group on a specified Symmetrix
package api

import (
	"fmt"

	"github.com/emccode/govmax/api/v2/model"
)

//POST /provisioning/symmetrix/{symmetrixId}/hostgroup

func (c Client) CreateStorageGroup(symmetricID, SRPID, SLOID, storageGroupID string) error {
	var err error
	if SRPID == "" {
		SRPID, err = c.getDefaultSRP(symmetricID)
		if err != nil {
			return err
		}
	}
	if SLOID == "" {
		SLOID, err = c.getDefaultSLO(symmetricID)
		if err != nil {
			return err
		}
	}

	uri := fmt.Sprintf("/restapi/sloprovisioning/symmetrix/%s/storagegroup", symmetricID)

	requestedSG := model.CreateStorageGroupParam{
		StorageGroupID:          storageGroupID,
		CreateEmptyStorageGroup: true,
		SrpID: SRPID,
		SloBasedStorageGroupParam: []model.SloBasedStorageGroupParam{
			model.SloBasedStorageGroupParam{
				SloID:             SLOID,
				WorkloadSelection: "none",
				NumOfVols:         0,
				VolumeAttribute: model.VolumeAttribute{
					CapacityUnit: "CYL",
					VolumeSize:   "0",
				},
				AllocateCapacityForEachVol:                      false,
				PersistPreallocatedCapacityThroughReclaimOrCopy: false,
			},
		},
	}

	var result model.GenericResult
	err = c.APICall("POST", uri, &requestedSG, &result)
	if err != nil {
		return fmt.Errorf("error creating storage group %s", err)
	}

	if !result.Success {
		return fmt.Errorf("symmetrix error creating storage group%s", result.Message)
	}

	return nil
}

func (c Client) getDefaultSRP(symmetrixID string) (string, error) {
	SRPs, err := c.ListSRPs(symmetrixID)
	if err != nil {
		return "", fmt.Errorf("Error getting first srp from srp list %s", err)
	}

	return SRPs.SrpID[0], nil
}

func (c Client) getDefaultSLO(symmetrixID string) (string, error) {
	SLOs, err := c.ListSLOs(symmetrixID)
	if err != nil {
		return "", fmt.Errorf("Error getting first slo from slo list %s", err)
	}

	return SLOs.SloID[0], nil

}
