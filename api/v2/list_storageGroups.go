// /provisioning/symmetrix/{symmetrixId}/volume
// GET
// This call queries for a list of volumes ids for the specified symmetrix

// create initiator -> create host -> create hostgroup
// create portgroup
// create volume ->
package api

import (
	"fmt"

	"github.com/emccode/govmax/api/v2/model"
)

func (c Client) ListStorageGroups(symmetrixID string) (model.ListStorageGroupResponse, error) {
	uri := fmt.Sprintf("/restapi/sloprovisioning/symmetrix/%s/storagegroup", symmetrixID)

	listStorageGroupResponse := model.ListStorageGroupResponse{}
	err := c.APICall("GET", uri, model.Empty{}, &listStorageGroupResponse)
	if err != nil {
		return listStorageGroupResponse, fmt.Errorf("Error Getting Storage Groups", err)
	}

	if !listStorageGroupResponse.Success {
		return listStorageGroupResponse, fmt.Errorf("Symmetrix Error Getting Storage Groups")
	}

	return listStorageGroupResponse, nil
}
