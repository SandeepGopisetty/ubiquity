package scbe_test

import (
	"log"
	"net/http"
	"os"

	httpmock "gopkg.in/jarcoal/httpmock.v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/IBM/ubiquity/fakes"
	"github.com/IBM/ubiquity/resources"

	"github.com/IBM/ubiquity/local/scbe"
)

var _ = Describe("scbeLocalClient", func() {
	var (
		client            resources.StorageClient
		logger            *log.Logger
		fakeScbeDataModel *fakes.FakeScbeDataModel
		fakeConfig        resources.ScbeConfig
		err               error
	)
	BeforeEach(func() {
		logger = log.New(os.Stdout, "ubiquity scbe: ", log.Lshortfile|log.LstdFlags)
		fakeScbeDataModel = new(fakes.FakeScbeDataModel)
		fakeConfig = resources.ScbeConfig{ConfigPath: "/tmp", ScbeURL: "http://scbe.com"}
		client, err = scbe.NewScbeLocalClientWithHTTPClientAndDataModel(logger, fakeConfig, fakeScbeDataModel, &http.Client{})
		Expect(err).ToNot(HaveOccurred())

	})

	Context(".Activate", func() {
		It("should succeed when httpClient returns statusAccepted", func() {
			httpmock.RegisterResponder("POST", "http://scbe.com/activate",
				httpmock.NewStringResponder(http.StatusAccepted, `[]`))

			err = client.Activate()
			Expect(err).ToNot(HaveOccurred())
		})

		It("should fail when httpClient returns http.StatusNotAcceptable", func() {
			httpmock.RegisterResponder("POST", "http://scbe.com/activate",
				httpmock.NewStringResponder(http.StatusNotAcceptable, `[]`))

			err = client.Activate()
			Expect(err).To(HaveOccurred())
		})
	})

})
