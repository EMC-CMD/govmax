// /provisioning/symmetrix/{symmetrixId}/slo
// GET
// This call queries for a list of SLO ids for the specified symmetrix

package api

import (
	"fmt"

	"github.com/emccode/govmax/api/v2/model"
)

func (c Client) ListSLOs(symmetrixID string) (model.ListOfSLOs, error) {
	uri := fmt.Sprintf("/restapi/sloprovisioning/symmetrix/%s/slo", symmetrixID)

	listSLOsResponse := model.ListOfSLOs{}
	err := c.APICall("GET", uri, model.Empty{}, &listSLOsResponse)
	if err != nil {
		return listSLOsResponse, fmt.Errorf("Error getting SLOs %s", err)
	}

	return listSLOsResponse, nil
}
