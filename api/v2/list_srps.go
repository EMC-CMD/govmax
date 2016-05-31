// /provisioning/symmetrix/{symmetrixId}/srp
// This call queries for a list of SRP ids for the specified symmetrix

package api

import (
	"fmt"

	"github.com/emccode/govmax/api/v2/model"
)

func (c Client) ListSRPs(symmetrixID string) (model.ListOfSRPs, error) {
	uri := fmt.Sprintf("/restapi/82/sloprovisioning/symmetrix/%s/srp", symmetrixID)

	listSRPsResponse := model.ListOfSRPs{}
	err := c.APICall("GET", uri, model.Empty{}, &listSRPsResponse)
	if err != nil {
		return listSRPsResponse, fmt.Errorf("Error getting SRPs %s", err)
	}

	return listSRPsResponse, nil
}
