package api

import (
	"fmt"

	"github.com/emccode/govmax/api/v2/model"
)

func (c Client) GetAttachments(symmetrixID, volumeID string) (model.AttachmentsResponse, error) {
  attachedVolumes, err := c.ListVolumes(symmetrixID, "", true)

  if err != nil {
    return model.AttachmentsResponse{}, fmt.Errorf("error getting volume list %s", err)
  }

  initiatorIDs := []string{}
  for _, v := range attachedVolumes.ResultList.Result {
    attachedVolume, err := c.GetVolume(symmetrixID, v.VolumeID)
    if err != nil {
        return model.AttachmentsResponse{}, fmt.Errorf("error getting volume %s", err)
    }

    for _, sID := range attachedVolume.Volume[0].StorageGroupIDs {
      storageGroup, err := c.GetStorageGroup(symmetrixID, sID)
      if err != nil {
        return model.AttachmentsResponse{}, fmt.Errorf("error getting storagegropu %s", err)
      }

      for _, mID := range storageGroup.StorageGroup[0].Maskingview {
        maskingView, err := c.GetMaskingView(symmetrixID, mID)
        if err != nil {
          return model.AttachmentsResponse{}, fmt.Errorf("error getting maskingview %s", err)
        }

        hostGroup, err := c.GetHostGroup(symmetrixID, maskingView.MaskingView[0].HostGroupID)
        if err != nil {
          return model.AttachmentsResponse{}, fmt.Errorf("error getting maskingview %s", err)
        }

        if len(hostGroup.HostGroup) > 0 {
          for _, h := range hostGroup.HostGroup[0].Host {
            for _, in := range h.Initiator {
              fmt.Println("Adding %s", in)
              initiatorIDs = append(initiatorIDs, in)
            }
          }
        }

        host, err  := c.GetHost(symmetrixID, maskingView.MaskingView[0].HostID)
        if err != nil {
          return model.AttachmentsResponse{}, fmt.Errorf("error getting host %s", err)
        }

        if len(host.Host) > 0 {
          for _, in := range host.Host[0].Initiator {
            initiatorIDs = append(initiatorIDs, in)
          }
        }
      }
    }
  }

	return  model.AttachmentsResponse{InitiatorIDs: initiatorIDs}, nil
}
