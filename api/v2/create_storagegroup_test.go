package api_test

import (
	// . "github.com/emccode/govmax/api/v2/"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CreateStoragegroup", func() {
	storageGroupID := "GOVMAXTESTING"
	AfterEach(func() {
		err := client.DeleteStorageGroup(symmetrixID, storageGroupID)
		Expect(err).ToNot(HaveOccurred())
	})

	Context("When storage group exists", func() {
		It("returns error", func() {
			err := client.CreateStorageGroup(symmetrixID, "", "", storageGroupID)
			Expect(err).ToNot(HaveOccurred())
			err = client.CreateStorageGroup(symmetrixID, "", "", storageGroupID)
			Expect(err.Error()).To(ContainSubstring("Cannot use the specified name because it's already in use"))
		})
	})

	Context("When valid parameters passed", func() {
		It("has no err, and we see storage group using get storage group", func() {
			err := client.CreateStorageGroup(symmetrixID, "", "", storageGroupID)
			Expect(err).ToNot(HaveOccurred())
		})
	})

})
