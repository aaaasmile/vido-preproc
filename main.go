package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/aaaasmile/vido-preproc/conf"
	"github.com/aaaasmile/vido-preproc/vidopre"
	"github.com/aaaasmile/vido-preproc/web"
	"github.com/aaaasmile/vido-preproc/web/idl"
)

const (
	help_usage       = "Vido site pre-processor (before webgen). Please use --usage to see all options."
	new_title_needed = "New post always needs the title. Please specify it using the --title flag."
)

func main() {
	cmd := flag.String("cmd", "", "Commands: splitpages, newpost, createindex")
	var configfile = flag.String("config", "config.toml", "Configuration file path")
	var ver = flag.Bool("ver", false, "Prints current version")
	var title = flag.String("title", "", "Title of the new post")
	var uicmd = flag.String("uicmd", "", "Edit post with browser ui. Commands: new, last")
	var nobrowser = flag.Bool("nobrowser", false, "Do not open the editor into a new browser page (use it if you have it already open)")
	flag.Parse()

	if *ver {
		fmt.Println("vido-preproc version ", idl.Buildnr)
		fmt.Println(help_usage)
		return
	}

	if *uicmd != "" || *cmd != "" {
		conf.ReadConfig(*configfile)
	}

	switch *cmd {
	case "splitpages":
		vidopre.SplitPages(conf.Current.PageSplitterInputDir, conf.Current.PostSourceDir)
	case "newpost":
		if *title == "" {
			log.Fatalln(new_title_needed)
		}
		vidopre.NewPost(conf.Current.PostSourceDir, *title, "")
	case "createindex":
		vidopre.CreateIndexPostPages(conf.Current.PostSourceDir, conf.Current.OutDirPage, conf.Current.PostPerPage)
	}

	openInBrowser := !*nobrowser
	switch *uicmd {
	case "new":
		if *title == "" {
			log.Fatalln(new_title_needed)
		}
		vidopre.NewPost(conf.Current.PostSourceDir, *title, "")
		vidopre.EditLastPost(conf.Current.PostSourceDir, openInBrowser)
	case "last":
		vidopre.EditLastPost(conf.Current.PostSourceDir, openInBrowser)
	}
	log.Println("Start the service as console process")
	if err := web.RunService(nil, nil, *configfile); err != nil {
		log.Fatal("Error: ", err)
	}
}
