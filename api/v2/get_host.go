package api

import (
	"fmt"

	"github.com/emccode/govmax/api/v2/model"
)

func (c Client) GetHost(symmetricID, hostID string) (model.HostResponse, error) {
	uri := fmt.Sprintf("/restapi/sloprovisioning/symmetrix/%s/host/%s", symmetricID, hostID)

	var hostResponse model.HostResponse
	err := c.APICall("GET", uri, model.Empty{}, &hostResponse)
	if err != nil {
		return model.HostResponse{}, fmt.Errorf("error getting host %s", err)
	}

	return hostResponse, nil
}
