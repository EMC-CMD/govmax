package api_test

import (
	// . "github.com/emccode/govmax/api/v2"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CreateVolume", func() {
	Context("Given an editing storage group param", func() {
		It("Creates a new volume", func() {
			volumeIDs, err := client.CreateVolume(symmetrixID, "5", "FBA", 2)
			Expect(err).ToNot(HaveOccurred())
			Expect(len(volumeIDs)).To(Equal(2))

			for _, id := range volumeIDs {
				err := client.DeleteVolume(symmetrixID, id)
				Expect(err).ToNot(HaveOccurred())
			}
		})
	})
})
