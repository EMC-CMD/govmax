package api

import (
	"fmt"

	"github.com/emccode/govmax/api/v2/model"
)

func (c Client) GetMaskingView(symmetricID string, maskingViewID string) (model.MaskingViewResponse, error) {
	uri := fmt.Sprintf("/restapi/sloprovisioning/symmetrix/%s/maskingview/%s", symmetricID, maskingViewID)

	var maskingViewResponse model.MaskingViewResponse
	err := c.APICall("GET", uri, model.Empty{}, &maskingViewResponse)
	if err != nil {
		return model.MaskingViewResponse{}, fmt.Errorf("error finding maskingview %s", err)
	}

	return maskingViewResponse, nil
}
