package vidopre

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
	UiServerUrl          string
	WebGenWebPageDir     string
	WebGenLocation       string
	WebgenOutIndexFile   string
}

const (
	BuildNr = "0.2.20181228-00"
)

var Conf = &Config{
	OutDirPage:  "./data/page-out",
	PostPerPage: 13,
}

func ReadConfig(configfile string) *Config {
	_, err := os.Stat(configfile)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := toml.DecodeFile(configfile, &Conf); err != nil {
		log.Fatal(err)
	}
	return Conf
}
