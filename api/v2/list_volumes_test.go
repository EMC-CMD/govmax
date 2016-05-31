package api_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ListVolumes", func() {
	Context("when given a symmetrix id", func() {
		It("returns a listvolumeresponse", func() {
			_, err := client.ListVolumes(symmetrixID, "")
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
