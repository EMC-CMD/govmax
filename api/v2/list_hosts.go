// /provisioning/symmetrix/{symmetrixId}/host
// GET
// This call queries for a list of Host ids for the specified symmetrix

package api

import (
	"fmt"

	"github.com/emccode/govmax/api/v2/model"
)

func (c Client) ListHosts(symmetrixID string) (model.ListOfHostsResponse, error) {
	uri := fmt.Sprintf("/restapi/sloprovisioning/symmetrix/%s/host", symmetrixID)

	listHostsResponse := model.ListOfHostsResponse{}
	err := c.APICall("GET", uri, model.Empty{}, &listHostsResponse)
	if err != nil {
		return listHostsResponse, fmt.Errorf("Error getting Hosts %s", err)
	}

	return listHostsResponse, nil
}
