package api

import "github.com/emccode/govmax/api/v2/model"

//POST /provisioning/symmetrix/{symmetrixId}/host
// This call creates a new Host on a specified Symmetric

func (c Client) CreateHost(p model.CreateHostParam) error {
	// return c.APICall("POST", "/provisioning/symmetrix/{symmetrixId}/host", p)
	return nil
}
