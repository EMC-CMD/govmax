package api_test

import (
	// . "github.com/emccode/govmax/api/v2/"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DeleteStorageGroup", func() {
	storageGroupID := "GOVMAXTESTING"
	Describe("Delete an existing storage group", func() {
		BeforeEach(func() {
			err := client.CreateStorageGroup(symmetrixID, "", "", storageGroupID)
			Expect(err).ToNot(HaveOccurred())
		})

		Context("When storage group exist", func() {
			It("deletes successfully", func() {
				err := client.DeleteStorageGroup(symmetrixID, storageGroupID)
				Expect(err).ToNot(HaveOccurred())
			})
		})
	})

	Context("When storage group does not exist", func() {
		It("returns error", func() {
			err := client.DeleteStorageGroup(symmetrixID, "DOESNOTEXIST")
			Expect(err.Error()).To(ContainSubstring("Cannot find Storage Group 'DOESNOTEXIST' for Symmetrix"))
		})
	})
})
