package api

import (
	"fmt"

	"github.com/emccode/govmax/api/v2/model"
)

func (c Client) GetStorageGroup(symmetricID string, storageGroupID string) (model.GetStorageGroupResult, error) {
	uri := fmt.Sprintf("/restapi/sloprovisioning/symmetrix/%s/storagegroup/%s", symmetricID, storageGroupID)

	var getStorageGroupResult model.GetStorageGroupResult
	err := c.APICall("GET", uri, model.Empty{}, &getStorageGroupResult)
	if err != nil {
		return model.GetStorageGroupResult{}, fmt.Errorf("error finding storagegroup %s", err)
	}

	return getStorageGroupResult, nil
}
