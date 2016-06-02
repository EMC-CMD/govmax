// /82/sloprovisioning/symmetrix
// GET
// This call queries for a list of symmetrixs

package api

import (
	"fmt"

	"github.com/emccode/govmax/api/v2/model"
)

func (c Client) ListSymmetrixs() (model.SymmetrixsResponse, error) {
	uri := fmt.Sprintf("/restapi/82/sloprovisioning/symmetrix")

	var symmetrixsResponse model.SymmetrixsResponse
	err := c.APICall("GET", uri, model.Empty{}, &symmetrixsResponse)
	if err != nil {
		return model.SymmetrixsResponse{}, fmt.Errorf("error getting list of Symmetrixs %s", err)
	}

	return symmetrixsResponse, nil
}
