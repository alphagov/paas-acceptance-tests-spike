package elasticsearch_service

import (
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/cloudfoundry-incubator/cf-test-helpers/helpers"
)

var (
	DEFAULT_TIMEOUT = 30 * time.Second
	CF_PUSH_TIMEOUT = 2 * time.Minute

	context helpers.SuiteContext
	config  helpers.Config
)

func TestElasticsearchService(t *testing.T) {
	RegisterFailHandler(Fail)

	config = helpers.LoadConfig()

	if config.DefaultTimeout > 0 {
		DEFAULT_TIMEOUT = config.DefaultTimeout * time.Second
	}
	if config.CfPushTimeout > 0 {
		CF_PUSH_TIMEOUT = config.CfPushTimeout * time.Second
	}

	context = helpers.NewContext(config)
	environment := helpers.NewEnvironment(context)

	BeforeSuite(func() {
		environment.Setup()
	})

	AfterSuite(func() {
		//environment.Teardown()
	})

	RunSpecs(t, "Elasticsearch Service")
}
