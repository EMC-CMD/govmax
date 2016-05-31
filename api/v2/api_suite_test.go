package api_test

import (
	"os"

	"github.com/emccode/govmax/api/v2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

var host, port, username, password, symmetrixID string
var insecure bool
var client api.Client

func TestApi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Api Suite")
}

var _ = BeforeSuite(func() {
	host = os.Getenv("VMAX_HOST")
	port = os.Getenv("VMAX_PORT")
	username = os.Getenv("VMAX_USERNAME")
	password = os.Getenv("VMAX_PASSWORD")
	insecure = true
	symmetrixID = os.Getenv("VMAX_SYMMETRIC_ID")
	client = api.NewClient(host, username, password, port, insecure)
})
