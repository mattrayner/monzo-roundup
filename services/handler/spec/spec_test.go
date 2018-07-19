package spec

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "testing"
)

func TestHandlerSpec(t *testing.T) {
  RegisterFailHandler(Fail)
  RunSpecs(t, "Handler Service Spec Suite")
}
