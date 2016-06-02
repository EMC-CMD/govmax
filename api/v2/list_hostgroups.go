// /provisioning/symmetrix/{symmetrixId}/hostgroup
// GET
// This call queries for a list of Host ids for the specified symmetrix

package api

import (
	"fmt"

	"github.com/emccode/govmax/api/v2/model"
)

func (c Client) ListHostGroups(symmetrixID string) (model.ListOfHostGroupsResponse, error) {
	uri := fmt.Sprintf("/restapi/sloprovisioning/symmetrix/%s/hostgroup", symmetrixID)

	listHostGroupsResponse := model.ListOfHostGroupsResponse{}
	err := c.APICall("GET", uri, model.Empty{}, &listHostGroupsResponse)
	if err != nil {
		return listHostGroupsResponse, fmt.Errorf("Error getting hostgroup list%s", err)
	}

	return listHostGroupsResponse, nil
}
