package api

import (
	"fmt"

	"github.com/emccode/govmax/api/v2/model"
)

func (c Client) GetVolume(symmetricID, volumeID string) (model.VolumeResponse, error) {
	uri := fmt.Sprintf("/restapi/sloprovisioning/symmetrix/%s/volume/%s", symmetricID, volumeID)

	var volumeResponse model.VolumeResponse
	err := c.APICall("GET", uri, model.Empty{}, &volumeResponse)
	if err != nil {
		return model.VolumeResponse{}, fmt.Errorf("error getting volume %s", err)
	}

	return volumeResponse, nil
}
