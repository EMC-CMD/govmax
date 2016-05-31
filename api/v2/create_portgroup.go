// POST /provisioning/symmetrix/{symmetrixId}/portgroup
// This call creates a new Port Group on a specified Symmetrix
package api

import "github.com/emccode/govmax/api/v2/model"

//POST /provisioning/symmetrix/{symmetrixId}/hostgroup

func (c Client) CreatePortGroup(p model.CreatePortGroupParam) error {
	// return c.APICall("POST", "/provisioning/symmetrix/{symmetrixId}/portgroup", p)
	return nil
}
