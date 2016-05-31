// This call deletes a storage group on a specified Symmetrix
package api

import (
	"fmt"

	"github.com/EMC-CMD/govmaxapi/model"
)

//DELETE /provisioning/symmetrix/{symmetrixId}/storagegroup/{storageGroupId}

func (c Client) DeleteStorageGroup(symmetricID, storageGroupID string) error {
	uri := fmt.Sprintf("/restapi/sloprovisioning/symmetrix/%s/storagegroup/%s", symmetricID, storageGroupID)

	err := c.APICall("DELETE", uri, model.Empty{}, &model.Empty{})
	if err != nil {
		return fmt.Errorf("error deleting storage group %s", err)
	}

	return nil
}
