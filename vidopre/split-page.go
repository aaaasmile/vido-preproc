package vidopre

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/afero"
)

var (
	appfs = afero.NewOsFs() // Wrapper del file os molto interessante
)

func SplitPages(dirIn string, dirOut string) {
	dir, err := filepath.Abs(dirIn)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Split all pages (*.page) in %s\n", dir)

	items, err := afero.ReadDir(appfs, dirIn)
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range items {
		log.Println("Processing: ", item.Name())
		splitSinglePage(dirIn, item.Name(), dirOut)
	}
}

func splitSinglePage(dirIn string, fname string, dirOut string) {
	afs := &afero.Afero{Fs: appfs}
	bb, err := afs.ReadFile(filepath.Join(dirIn, fname))
	if err != nil {
		log.Fatalf("Cannot process file %s because:%v", fname, err)
	}
	s := string(bb)
	posts := GetSplittedPosts(s)
	fmt.Printf("Posts: \n%v\n", posts)
	os.Exit(1)
}
