package main

import (
	"flag"
	"fmt"
	"log"

	"./vidopre"
)

const (
	help_usage       = "Vido site pre-processor (before webgen). Please use --usage to see all options."
	new_title_needed = "New post always needs the title. Please specify it using the --title flag."
)

func main() {
	cmd := flag.String("cmd", "", "Commands: splitpages, newpost, createindex")
	var configfile = flag.String("config", "config.toml", "Configuration file path")
	var ver = flag.Bool("version", false, "Prints current version")
	var title = flag.String("title", "", "Title of the new post")
	var uicmd = flag.String("uicmd", "new", "Edit post with browser ui. Commands: new, last")
	var nobrowser = flag.Bool("nobrowser", false, "Do not open the editor into a new browser page (use it if you have it already open)")
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
		if *title == "" {
			log.Fatalln(new_title_needed)
		}
		vidopre.NewPost(vidopre.Conf.PostSourceDir, *title, "")
		break
	case "createindex":
		vidopre.CreateIndexPostPages(vidopre.Conf.PostSourceDir, vidopre.Conf.OutDirPage, vidopre.Conf.PostPerPage)
		break
	default:
		fmt.Println(help_usage)
	}

	openInBrowser := !*nobrowser
	switch *uicmd {
	case "new":
		if *title == "" {
			log.Fatalln(new_title_needed)
		}
		vidopre.NewPost(vidopre.Conf.PostSourceDir, *title, "")
		vidopre.EditLastPost(vidopre.Conf.PostSourceDir, openInBrowser)
		break
	case "last":
		vidopre.EditLastPost(vidopre.Conf.PostSourceDir, openInBrowser)
		break
	default:
		fmt.Println(help_usage)
	}

}
