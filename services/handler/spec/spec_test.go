package spec

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "testing"
  "log"
  "io/ioutil"
)

func TestHandlerSpec(t *testing.T) {
  log.SetOutput(ioutil.Discard)
  RegisterFailHandler(Fail)
  RunSpecs(t, "Handler Service Spec Suite")
}
