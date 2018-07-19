package spec

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  uut "github.com/mattrayner/monzo-roundup/services/dynamodb"
  "github.com/mattrayner/monzo-roundup/types/services/dynamodb"
)

var _ = Describe("Contracts", func() {

  var (
    service uut.Service
  )

  BeforeEach(func() {
    service = uut.NewService()
    service.Start()
  })

  It("should support 'GetItem' contract method", func() {
    out, err := service.GetUser(&dynamodb.GetUserInput{AccountID: "acc_1234"})
    Expect(err).ToNot(HaveOccurred())

    Expect(out.User.AccountID).To(Equal("acc_1234"))
    Expect(out.User.RefreshKey).To(Equal("refresh_1234"))
    Expect(out.User.AuthKey).To(Equal("rauth_1234"))
  })
})