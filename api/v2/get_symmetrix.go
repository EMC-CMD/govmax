package api

import (
	"fmt"

	"github.com/emccode/govmax/api/v2/model"
)

func (c Client) GetSymmetrix(symmID string) (model.Symmetrix, error) {
	uri := fmt.Sprintf("/restapi/sloprovisioning/symmetrix/%s", symmID)

	var symmetrixResponse model.SymmetrixResponse
	err := c.APICall("GET", uri, model.Empty{}, &symmetrixResponse)
	if err != nil {
		return model.Symmetrix{}, fmt.Errorf("error getting Symmetrix %s", err)
	}

	if len(symmetrixResponse.Symmetrix) == 0 {
		return model.Symmetrix{}, fmt.Errorf("no symmetrix found with the id %s", err)
	}

	symmetrix := model.Symmetrix{
		SymmetrixID: symmetrixResponse.Symmetrix[0].SymmetrixID,
	}

	return symmetrix, nil
}
