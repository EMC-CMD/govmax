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
