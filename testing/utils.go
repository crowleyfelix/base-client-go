package basetesting

import (
	"encoding/json"
	"io/ioutil"

	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
}

func (suite *Suite) File(path string) []byte {
	blob, err := ioutil.ReadFile(path)
	Expect(err).Should(BeNil())
	return blob
}

func (suite *Suite) LoadJSON(path string, target interface{}) {
	blob := suite.File(path)
	err := json.Unmarshal(blob, &target)
	Expect(err).Should(BeNil())
}

func (suite *Suite) Config() *config {
	c := new(config)
	c.Load()
	c.Validate()
	return c
}
