package main

import (
	"flag"
	"fmt"

	"./vidopre"
)

func main() {
	cmd := flag.String("cmd", "", "Comnads: splitpages, newpost, createindex")
	var configfile = flag.String("config", "config.toml", "Configuration file path")
	var ver = flag.Bool("version", false, "Prints current version")
	var title = flag.String("title", "", "Title of the new post")
	flag.Parse()

	if *ver {
		fmt.Println("vido-preproc version ", vidopre.BuildNr)
		return
	}

	vidopre.ReadConfig(*configfile)

	switch *cmd {
	case "splitpages":
		vidopre.SplitPages(vidopre.Conf.PageSplitterInputDir, vidopre.Conf.PostSourceDir)
		break
	case "newpost":
		vidopre.NewPost(vidopre.Conf.PostSourceDir, *title, "lol")
		break
	case "createindex":
		vidopre.CreateIndexPostPages(vidopre.Conf.PostSourceDir, vidopre.Conf.OutDirPage, vidopre.Conf.PostPerPage)
		break
	default:
		fmt.Println("Vido site pre-processor (before webgen). Please use --usage to see all options.")
	}

}
