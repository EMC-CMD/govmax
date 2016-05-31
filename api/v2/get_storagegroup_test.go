package api_test

import (
	// . "github.com/emccode/govmax/api/v2"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GetStoragegroup", func() {
	Context("When storage group ID does not exist", func() {
		It("return an error", func() {
			_, err := client.GetStorageGroup(symmetrixID, "DOESNOTEXISTUNLESSITDOES")
			Expect(err.Error()).To(ContainSubstring("Cannot find Storage Group 'DOESNOTEXISTUNLESSITDOES'"))
		})

	})

	Context("When storage group ID exist", func() {
		It("return information about that storage group", func() {
			storageGroups, err := client.ListStorageGroups(symmetrixID)
			Expect(err).ToNot(HaveOccurred())
			expected := storageGroups.StorageGroupID[0]
			result, err := client.GetStorageGroup(symmetrixID, expected)
			Expect(err).ToNot(HaveOccurred())
			Expect(result.StorageGroup[0].StorageGroupID).To(Equal(expected))
		})
	})
})
