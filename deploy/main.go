package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"time"

	"./depl"
)

var (
	defOutDir = "..\\..\\tmp\\vido-preproc-deployed"
)

func main() {
	const (
		invidosite = "InvidositeHtmlgit"
	)
	var outdir = flag.String("outdir", "",
		fmt.Sprintf("Output zip directory. If empty use the hardcoded one: %s\n", defOutDir))

	var target = flag.String("target", invidosite,
		fmt.Sprintf("Target of deployment: %s", invidosite))

	flag.Parse()

	rootDirRel := ".."
	pathItems := []string{"vido-preproc.exe", "static", "templates"}
	switch *target {
	case invidosite:
		pathItems = append(pathItems, "deploy/config_files/target_config.toml")
	default:
		log.Fatalf("Target is %s not recognized", *target)
	}
	log.Printf("Create the zip package for target %s", *target)

	outFn := getOutFileName(*outdir, *target)
	depl.CreateDeployZip(rootDirRel, pathItems, outFn, func(pathItem string) string {
		if strings.HasPrefix(pathItem, "deploy/config_files") {
			return "config.toml"
		}
		return pathItem
	})
}

func getOutFileName(outdir string, tgt string) string {
	if outdir == "" {
		outdir = defOutDir
	}
	vn := depl.GetVersionNrFromFile("../vidopre/conf.go", "")
	log.Println("Version is ", vn)

	currentTime := time.Now()
	s := fmt.Sprintf("VidoPreProc_%s_%s_%s.zip", strings.Replace(vn, ".", "-", -1), currentTime.Format("02012006-150405"), tgt) // current date-time stamp using 2006 date time format template
	s = filepath.Join(outdir, s)
	return s
}
