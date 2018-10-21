package vidopre

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"

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
	dirOutAbs, err := filepath.Abs(dirOut)
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
		splitSinglePage(dirIn, item.Name(), dirOutAbs)
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
	log.Printf("%d posts recognized in %s\n", len(posts), fname)
	//fmt.Println(posts[len(posts)-1].Content)
	//t := time.Now()
	for _, item := range posts {
		//outFname := filepath.Join(dirOut, fmt.Sprintf("%s-%s-%s-%d%d%d-%s.txt",
		//	item.Year, item.Month, item.Day, t.Hour(), t.Minute(), t.Second(), item.Title))
		outFname := filepath.Join(dirOut, fmt.Sprintf("%s-%s-%s-%s.txt",
			item.Year, item.Month, item.Day, item.Title))

		if err := afs.WriteFile(outFname, []byte(item.Content), 0777); err != nil {
			log.Fatalln("Unable to write file ", err)
		}
		log.Println("Out file written: ", outFname)
	}
}

type CtxIndexPage struct {
	TotPages     int
	PostPerPages int
	CurrPageNum  int
}

func CreateIndexPostPages(dirIn string, dirOut string, postPerPage int) {
	dir, err := filepath.Abs(dirIn)
	if err != nil {
		log.Fatal(err)
	}
	dirOutAbs, err := filepath.Abs(dirOut)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Create all index pages (*.page) from posts in %s\nOutput is %s\n", dir, dirOutAbs)

	items, err := afero.ReadDir(appfs, dirIn) // sorted by name as default
	if err != nil {
		log.Fatal(err)
	}

	numPages := len(items)/postPerPage + 1
	curPageNum := numPages - 1
	log.Printf("Recognized %d posts, pages %d", len(items), numPages)
	ctx := &CtxIndexPage{
		TotPages:     numPages,
		PostPerPages: postPerPage,
		CurrPageNum:  curPageNum,
	}
	startNewPage(tempIndexOtherPages, ctx)

	for i, item := range items {
		log.Printf("Processing: %s, page ix is %d\n", item.Name(), curPageNum)
		if (i > 0) && (i%postPerPage) == 0 {
			curPageNum--
		}
	}
}

func startNewPage(tempContent string, ctx *CtxIndexPage) {
	var t = template.Must(template.New("Page").Parse(tempContent))
	err := t.Execute(os.Stdout, ctx)
	if err != nil {
		log.Fatal("Template error: ", err)
	}
}
