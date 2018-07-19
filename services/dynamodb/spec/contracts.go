package spec

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  . "github.com/mattrayner/monzo-roundup/utils/testing"
  uut "github.com/mattrayner/monzo-roundup/services/handler"
  _dynamodb "github.com/mattrayner/monzo-roundup/services/dynamodb"
  _monzo "github.com/mattrayner/monzo-roundup/services/monzoapi"
  "github.com/mattrayner/monzo-roundup/types/services/handler"
  "github.com/mattrayner/monzo-roundup/types/services/dynamodb"
  "github.com/mattrayner/monzo-roundup/types/services/monzoapi"
)

var _ = Describe("Contracts", func() {

}