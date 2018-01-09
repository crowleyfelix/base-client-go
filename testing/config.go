package basetesting

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/namsral/flag"
)

type config struct {
	URL string
}

const (
	configurationFailed = "This is a required configuration. Please provide a value."
)

func (c *config) Load() {

	if err := godotenv.Load(); err != nil {
		println("Error loading .env file")
	}

	fs := flag.NewFlagSet(os.Args[0], 0)
	fs.StringVar(&c.URL, "url", "", "Url do Cadu")

	err := fs.Parse(os.Args[1:])
	if err != nil {
		panic(fmt.Sprintf("Error on parsing variables: %#v", err))
	}
}

func (c *config) Validate() {
	if c.URL == "" {
		panic(fmt.Sprintf("%s - %s", "url", configurationFailed))
	}
}
