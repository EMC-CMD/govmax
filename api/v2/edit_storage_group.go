package api

import (
	"fmt"

	"github.com/emccode/govmax/api/v2/model"
)

func (c Client) EditStorageGroup(symmetricID, storageGroupID string, p model.EditStorageGroupParam) error {
	uri := fmt.Sprintf("/restapi/82/sloprovisioning/symmetrix/%s/storagegroup/%s", symmetricID, storageGroupID)

	var editStorageGroupResponse model.EditStorageGroupResponse
	err := c.APICall("PUT", uri, p, &editStorageGroupResponse)
	if err != nil {
		return fmt.Errorf("error creating volume %s", err)
	}

	return nil
}
