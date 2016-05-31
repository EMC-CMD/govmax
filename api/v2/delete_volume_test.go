package api_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = XDescribe("DeleteVolume", func() {
	Context("When passing an incorrect volumeID", func() {
		It("Return an error", func() {
			volumeID := "999999"
			err := client.DeleteVolume(symmetrixID, volumeID)
			Expect(err).To(MatchError("Error removing volume error reading json response invalid character '<' looking for beginning of value. Response Body <?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\"?>\n<exception xmlns=\"http://www.emc.com/em/2012/07/univmax/restapi/common\" xmlns:ns2=\"http://www.emc.com/em/2012/07/univmax/restapi/82/performance\" xmlns:ns3=\"http://www.emc.com/em/2012/07/univmax/restapi/81/performance\" xmlns:ns4=\"http://www.emc.com/em/2012/07/univmax/restapi/performance\" xmlns:ns5=\"http://www.emc.com/em/2012/07/univmax/restapi/wlpd\" xmlns:ns6=\"http://www.emc.com/em/2012/07/univmax/restapi/wlp\" xmlns:ns7=\"http://www.emc.com/em/2012/07/univmax/restapi/82/internal/performance\" xmlns:ns8=\"http://www.emc.com/em/2012/07/univmax/restapi/mobile\" xmlns:ns9=\"http://www.emc.com/em/2012/07/univmax/restapi/management\" xmlns:ns10=\"http://www.emc.com/em/2012/07/univmax/restapi/82/wlp\" xmlns:ns11=\"http://www.emc.com/em/2012/07/univmax/restapi/82/sloprovisioning\" xmlns:ns12=\"http://www.emc.com/em/2012/07/univmax/restapi/81/vvol\" xmlns:ns13=\"http://www.emc.com/em/2012/07/univmax/restapi/sloprovisioning\" xmlns:ns14=\"http://www.emc.com/em/2012/07/univmax/restapi/provisioning\" xmlns:ns15=\"http://www.emc.com/em/2012/07/univmax/restapi/82/migration\" xmlns:ns16=\"http://www.emc.com/em/2012/07/univmax/restapi/system\" xmlns:ns17=\"http://www.emc.com/em/2012/07/univmax/restapi/82/vvol\" xmlns:ns18=\"http://www.emc.com/em/2012/07/univmax/restapi/82/replication\" xmlns:ns19=\"http://www.emc.com/em/2012/07/univmax/restapi/83/replication\" xmlns:ns20=\"http://www.emc.com/em/2012/07/univmax/restapi/82/provisioning\" xmlns:ns21=\"http://www.emc.com/em/2012/07/univmax/restapi/hinting\" xmlns:ns22=\"http://www.emc.com/em/2012/07/univmax/restapi/replication\" xmlns:ns23=\"http://www.emc.com/em/2012/07/univmax/restapi/81/replication\" xmlns:ns24=\"http://www.emc.com/em/2012/07/univmax/restapi/82/internal/cirrus\" xmlns:ns25=\"http://www.emc.com/em/2012/07/univmax/restapi/82/system\" message=\"Cannot find Volume '999999' for Symmetrix 000196801181\"/>\n\n"))
		})
	})

	Context("When passing in a correct volumeID", func() {
		XIt("Return nothing", func() {
			volumeID := "000E4"
			err := client.DeleteVolume(symmetrixID, volumeID)
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
