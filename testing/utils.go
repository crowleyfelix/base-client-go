package basetesting

import (
	"encoding/json"
	"io/ioutil"

	. "github.com/onsi/gomega"
)

func File(path string) []byte {
	blob, err := ioutil.ReadFile(path)
	Expect(err).Should(BeNil())
	return blob
}

func LoadJSON(path string, target interface{}) {
	blob := File(path)
	err := json.Unmarshal(blob, &target)
	Expect(err).Should(BeNil())
}

func Config() *config {
	c := new(config)
	c.Load()
	c.Validate()
	return c
}
