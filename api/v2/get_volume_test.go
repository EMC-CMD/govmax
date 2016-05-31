package api_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GetVolume", func() {
	Context("When passing symmetricID and volumeID", func() {
		It("return the corresponding volume response", func() {
			volumeResp, err := client.GetVolume(symmetrixID, "00001")
			Expect(err).ToNot(HaveOccurred())
			Expect(volumeResp.Success).To(BeTrue())
			Expect(volumeResp.Volume[0].Type).To(Equal("TDEV"))
		})
	})
})
