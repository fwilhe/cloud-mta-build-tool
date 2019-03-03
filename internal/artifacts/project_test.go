package artifacts

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/SAP/cloud-mta-build-tool/internal/commands"
)

var _ = Describe("Project", func() {
	var _ = Describe("runBuilder", func() {
		It("Sanity", func() {
			buildersCfg := commands.BuilderTypeConfig
			commands.BuilderTypeConfig = []byte(`
builders:
- name: testbuilder
  info: "installing module dependencies & remove dev dependencies"
  path: "path to config file which override the following default commands"
  type:
  - command: go version
`)
			Ω(runBuilder("testbuilder")).Should(Succeed())
			commands.BuilderTypeConfig = buildersCfg
		})

		It("Fails on command execution", func() {
			buildersCfg := commands.BuilderTypeConfig
			commands.BuilderTypeConfig = []byte(`
builders:
- name: testbuilder
  info: "installing module dependencies & remove dev dependencies"
  path: "path to config file which override the following default commands"
  type:
  - command: go test unknown_test.go
`)
			Ω(runBuilder("testbuilder")).Should(HaveOccurred())
			commands.BuilderTypeConfig = buildersCfg
		})
	})
})
