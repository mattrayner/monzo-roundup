package spec

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "gopkg.in/jarcoal/httpmock.v1"
  "testing"
  "log"
  "io/ioutil"
)

func TestNetworking(t *testing.T) {
  RegisterFailHandler(Fail)
  log.SetOutput(ioutil.Discard)
  RunSpecs(t, "Network Client Spec Suite")
}

var _ = BeforeSuite(func() {
  // block all HTTP requests
  httpmock.Activate()
})

var _ = BeforeEach(func() {
  // remove any mocks
  httpmock.Reset()
  httpmock.RegisterNoResponder(httpmock.NewStringResponder(500, ""))
})

var _ = AfterSuite(func() {
  httpmock.DeactivateAndReset()
})