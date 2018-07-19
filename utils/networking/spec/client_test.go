package spec


import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "gopkg.in/jarcoal/httpmock.v1"
  "github.com/mattrayner/monzo-roundup/utils/networking"
  "net/http"
  "io/ioutil"
  "bytes"
  "errors"
  "net/http/httptest"
)

type MockHttpClient struct {}

func (m *MockHttpClient) Do(request *http.Request) (*http.Response, error) {
  response := &http.Response{
    Body: request.Body,
    StatusCode: 200,
    Header: request.Header,
  }

  return response, nil
}

type errReader int

func (errReader) Read(p []byte) (n int, err error) {
  return 0, errors.New("test error")
}

var _ = Describe("client", func() {
  Describe("MakeRequest", func() {
    Context("with an error passed", func() {
      It("returns the expected error", func() {
        client := &MockHttpClient{}
        request := &http.Request{}
        err := errors.New("an error");

        _, error := networking.MakeRequest(client, request, err, "https://matt.rayner.io", true, "foo")
        Expect(error).To(HaveOccurred())
        Expect(error.Error()).To(Equal("error creating request object for: https://matt.rayner.io"))
      })
    })

    Context("without an error passed", func() {
      Context( "with the includeAuth flag", func() {
        Context("and an authToken value", func() {
          It("adds the Authorization header", func() {
            httpmock.RegisterResponder("GET", "https://matt.rayner.io",
              func(req *http.Request) (*http.Response, error) {
                Expect(req.Header.Get("Authorization")).To(Equal("Bearer foo"))

                response := &http.Response{
                  Body: ioutil.NopCloser(bytes.NewBuffer([]byte("Test Response"))),
                  StatusCode: 200,
                }

                return response, nil
              },
            )

            client := &http.Client{}
            request, err := http.NewRequest("GET", "https://matt.rayner.io", nil)

            response, error := networking.MakeRequest(client, request, err, "https://matt.rayner.io", true, "foo")
            Expect(response).To(Equal([]byte("Test Response")))
            Expect(error).NotTo(HaveOccurred())
          })
        })

        Context("and no authToken value", func() {
          It("raises an argument error", func() {
            client := &http.Client{}
            request, err := http.NewRequest("GET", "https://matt.rayner.io", nil)

            _, e := networking.MakeRequest(client, request, err, "https://matt.rayner.io", true, "")
            Expect(e).To(HaveOccurred())
            Expect(e.Error()).To(Equal("argument error - authToken is required if includeAuth is true"))
          })
        })
      })

      Context("without the includeAuth flag", func() {
        It("doesn't add the Authorization header", func() {
          httpmock.RegisterResponder("GET", "https://matt.rayner.io",
            func(req *http.Request) (*http.Response, error) {
              Expect(req.Header.Get("Authorization")).To(Equal(""))

              response := &http.Response{
                Body: ioutil.NopCloser(bytes.NewBuffer([]byte("Test Response"))),
                StatusCode: 200,
              }

              return response, nil
            },
          )

          client := &http.Client{}
          request, err := http.NewRequest("GET", "https://matt.rayner.io", nil)

          response, error := networking.MakeRequest(client, request, err, "https://matt.rayner.io", false, "foo")
          Expect(response).To(Equal([]byte("Test Response")))
          Expect(error).NotTo(HaveOccurred())
        })
      })

      Context("with a non-200 HTTP response", func() {
        BeforeEach(func() {
          httpmock.RegisterResponder("GET", "https://matt.rayner.io",
            httpmock.NewStringResponder(500, `{"error": true}`))
        })

        It("returns the expected error", func() {
          client := &http.Client{}
          request, err := http.NewRequest("GET", "https://matt.rayner.io", nil)

          _, e := networking.MakeRequest(client, request, err, "https://matt.rayner.io", false, "foo")
          Expect(e).To(HaveOccurred())
          Expect(e.Error()).To(Equal("Non-200 Status code. Status code: (500), body: {\"error\": true}"))
        })
      })

      Context("with a 200 HTTP response", func() {
        Context("with an error reading the response body", func() {
          It("returns the expected error", func() {
            client := &MockHttpClient{}
            request := httptest.NewRequest(http.MethodGet, "https://matt.rayner.io", errReader(0))

            _, e := networking.MakeRequest(client, request, nil, "https://matt.rayner.io", false, "foo")
            Expect(e).To(HaveOccurred())
            Expect(e.Error()).To(Equal("Error reading body: test error"))
          })
        })

        Context("successfully reading the response body", func() {
          BeforeEach(func() {
            httpmock.RegisterResponder("GET", "https://matt.rayner.io",
              httpmock.NewStringResponder(200, `Test response!`))
          })

          It("returns the expected byte array", func() {
            client := &http.Client{}
            request, err := http.NewRequest("GET", "https://matt.rayner.io", nil)

            expected := make([]byte, 0, 1536)
            expected = append(expected, 'T', 'e', 's', 't', ' ', 'r', 'e', 's', 'p', 'o', 'n', 's', 'e', '!')

            response, e := networking.MakeRequest(client, request, err, "https://matt.rayner.io", false, "foo")
            Expect(response).To(Equal(expected))
            Expect(e).NotTo(HaveOccurred())
          })
        })
      })
    })
  })
})
