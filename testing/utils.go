package basetesting

import (
	"encoding/json"
	"io/ioutil"

	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/suite"
)

func File(path string) []byte {
	blob, err := ioutil.ReadFile(path)
	Expect(err).Should(BeNil())
	return blob
}

func LoadJSON(path string, target interface{}) {
	blob := suite.File(path)
	err := json.Unmarshal(blob, &target)
	Expect(err).Should(BeNil())
}

func Config() *config {
	c := new(config)
	c.Load()
	c.Validate()
	return c
}
