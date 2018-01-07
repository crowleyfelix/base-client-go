package basetesting

import (
	"fmt"
	"os"

	"github.com/namsral/flag"
)

type config struct {
	URL       string
	AppKey    string
	SecretKey string
	User      string
}

const (
	configurationFailed = "This is a required configuration. Please provide a value."
)

func (c *config) Load() {
	fs := flag.NewFlagSet(os.Args[0], 0)
	fs.StringVar(&c.URL, "url", "", "Url do Cadu")
	fs.StringVar(&c.AppKey, "app_key", "", "Application Key do Cadu")
	fs.StringVar(&c.SecretKey, "secret_key", "", "Secret Key do Cadu")
	fs.StringVar(&c.SecretKey, "user", "", "Secret Key do Cadu")

	err := fs.Parse(os.Args[1:])
	if err != nil {
		panic(fmt.Sprintf("Error on parsing variables: %#v", err))
	}
}

func (c *config) Validate() {
	if c.URL == "" {
		panic(fmt.Sprintf("%s - %s", "url", configurationFailed))
	}
	if c.AppKey == "" {
		panic(fmt.Sprintf("%s - %s", "app_key", configurationFailed))
	}
	if c.SecretKey == "" {
		panic(fmt.Sprintf("%s - %s", "secret_key", configurationFailed))
	}
}
