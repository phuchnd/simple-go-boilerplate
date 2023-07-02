package tests

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/phuchnd/simple-go-boilerplate/utils/test"
	"testing"
)

var (
	dockerComposeSetup = test.NewDockerComposeSetup()
	databaseSetup      = test.NewDatabaseSetup(dockerComposeSetup)
)

var _ = BeforeSuite(func() {
	dockerComposeSetup.Setup()
	databaseSetup.Setup()
})

var _ = AfterSuite(func() {
	dockerComposeSetup.Teardown()
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "integration/Integration Tests")
}
