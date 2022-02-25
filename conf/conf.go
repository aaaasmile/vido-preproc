package conf

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	PostPerPage          int
	PageSplitterInputDir string
	PostSourceDir        string
	OutDirPage           string
	WebGenWebPageDir     string
	WebGenLocation       string
	WebgenOutIndexFile   string
	VuetifyLibName       string
	VueLibName           string
	RootURLPattern       string
	ServiceURL           string
	DebugVerbose         bool
}

var Current = &Config{
	OutDirPage:  "./data/page-out",
	PostPerPage: 13,
}

func ReadConfig(configfile string) *Config {
	_, err := os.Stat(configfile)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := toml.DecodeFile(configfile, &Current); err != nil {
		log.Fatal(err)
	}
	return Current
}
