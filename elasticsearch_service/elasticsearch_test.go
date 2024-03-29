package elasticsearch_service

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"

	"github.com/cloudfoundry-incubator/cf-test-helpers/cf"
	"github.com/cloudfoundry-incubator/cf-test-helpers/generator"
	"github.com/cloudfoundry-incubator/cf-test-helpers/helpers"
)

var _ = Describe("The Elasticsearch service", func() {

	Describe("adding elasticsearch to an app", func() {
		var (
			appName string
		)

		BeforeEach(func() {
			appName = generator.PrefixedRandomName("CATS-APP-")
			Expect(cf.Cf(
				"push", appName,
				"--no-start",
				"-b", config.GoBuildpackName,
				"-p", "../example_apps/es_test_app",
				"-d", config.AppsDomain,
			).Wait(DEFAULT_TIMEOUT)).To(Exit(0))
		})

		It("can be added to an app", func() {
			instanceName := generator.PrefixedRandomName("ES-SERVICE-")

			Expect(cf.Cf("create-service", "elasticsearch13", "free", instanceName).Wait(DEFAULT_TIMEOUT)).To(Exit(0))
			Expect(cf.Cf("bind-service", appName, instanceName).Wait(DEFAULT_TIMEOUT)).To(Exit(0))
			Expect(cf.Cf("start", appName).Wait(CF_PUSH_TIMEOUT)).To(Exit(0))
			Expect(helpers.CurlAppRoot(appName)).To(ContainSubstring("You Know, for Search"))
		})
	})
})
