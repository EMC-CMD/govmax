package api_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GetVolume", func() {
	Context("When passing symmetricID and volumeID", func() {
		It("return the corresponding volume response", func() {
			hosts, err := client.ListHosts(symmetrixID)
			Expect(err).ToNot(HaveOccurred())
			Expect(len(hosts.HostID) > 0).To(BeTrue())

			hostID := hosts.HostID[0]
			host, err := client.GetHost(symmetrixID, hostID)
			Expect(err).ToNot(HaveOccurred())
			Expect(host.Host[0].HostID).To(Equal(hostID))
		})
	})
})
