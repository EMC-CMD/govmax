package api

import (
	"fmt"

	"github.com/emccode/govmax/api/v2/model"
)

func (c Client) CreateVolume(symmetricID, size, emulation string, numOfVols int) ([]string, error) {
	tmpStorageGroup := "GOVMAXAPITEMP"
	err := c.CreateStorageGroup(symmetricID, "", "", tmpStorageGroup)
	if err != nil {
		return []string{}, fmt.Errorf("error creating temporary storage group %s", err)
	}
	defer c.DeleteStorageGroup(symmetricID, tmpStorageGroup)

	// add volume to storagegroup (EditStorageGroup)
	p := model.EditStorageGroupParam{
		EditStorageGroupActionParam: model.EditStorageGroupActionParam{
			ExpandStorageGroupParam: &model.ExpandStorageGroupParam{
				NumOfVols: numOfVols,
				VolumeAttribute: model.VolumeAttribute{
					CapacityUnit: "GB",
					VolumeSize:   size,
				},
				CreateNewVolumes: true,
				Emulation:        emulation,
			},
		},
	}

	err = c.EditStorageGroup(symmetricID, tmpStorageGroup, p)
	if err != nil {
		return []string{}, fmt.Errorf("error adding volume into storage group %s", err)
	}
	createdVolumes, err := c.ListVolumes(symmetricID, tmpStorageGroup, false)
	if err != nil {
		return []string{}, fmt.Errorf("error listing volumes in temp storage group %s", tmpStorageGroup)
	}

	volumeIDs := make([]string, len(createdVolumes.ResultList.Result))
	for i, vol := range createdVolumes.ResultList.Result {
		err = c.RemoveVolumeFromStorageGroup(symmetricID, tmpStorageGroup, vol.VolumeID)
		if err != nil {
			return []string{}, fmt.Errorf("error deleting volume %s from temporary storage group %s", vol.VolumeID)
		}
		volumeIDs[i] = vol.VolumeID
	}

	return volumeIDs, nil
}
