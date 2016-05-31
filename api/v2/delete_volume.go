package api

import (
	"fmt"

	"github.com/emccode/govmax/api/v2/model"
)

func (c Client) DeleteVolume(symmetricID, volumeID string) error {
	uri := fmt.Sprintf("/restapi/sloprovisioning/symmetrix/%s/volume/%s", symmetricID, volumeID)
	err := c.APICall("DELETE", uri, model.Empty{}, &model.Empty{})
	if err != nil {
		return fmt.Errorf("Error removing volume %s", err)
	}

	return nil
}

func (c Client) RemoveVolumeFromStorageGroup(symmetricID, storageGroupID, volumeID string) error {
	removeVolumeParam := model.EditStorageGroupParam{
		EditStorageGroupActionParam: model.EditStorageGroupActionParam{
			RemoveVolumeParam: &model.RemoveVolumeParam{
				VolumeID: []string{
					volumeID,
				},
			},
		},
	}
	return c.EditStorageGroup(symmetricID, storageGroupID, removeVolumeParam)
}
