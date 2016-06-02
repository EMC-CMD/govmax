package api

import (
	"fmt"

	"github.com/emccode/govmax/api/v2/model"
)

func (c Client) GetHostGroup(symmetricID, hostGroupID string) (model.HostGroupResponse, error) {
	uri := fmt.Sprintf("/restapi/sloprovisioning/symmetrix/%s/hostgroup/%s", symmetricID, hostGroupID)

	var hostGroupResponse model.HostGroupResponse
	err := c.APICall("GET", uri, model.Empty{}, &hostGroupResponse)
	if err != nil {
		return model.HostGroupResponse{}, fmt.Errorf("error getting hostgroup %s", err)
	}

	return hostGroupResponse, nil
}
