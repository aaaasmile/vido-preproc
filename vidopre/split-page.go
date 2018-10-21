package vidopre

import (
	"fmt"
	"log"
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

	f := startNewPage(getTemplateByPage(curPageNum), ctx, dirOutAbs)
	reversed := []string{}

	for i, item := range items {
		log.Printf("Processing: %s, page ix is %d\n", item.Name(), curPageNum)
		afs := &afero.Afero{Fs: appfs}
		sourceFname := filepath.Join(dir, item.Name())
		src, err := afs.ReadFile(sourceFname)
		if err != nil {
			log.Fatal("Error reading file ", err)
		}
		reversed = append(reversed, string(src))

		if (i > 0) && ((i+1)%postPerPage) == 0 {
			writeReversePostInFile(f, reversed)
			curPageNum--
			f.Close()
			ctx.CurrPageNum = curPageNum
			reversed = nil
			f = startNewPage(getTemplateByPage(curPageNum), ctx, dirOutAbs)
		}
	}
	writeReversePostInFile(f, reversed)
	f.Close()
}

func writeReversePostInFile(f afero.File, arr []string) {
	for i := len(arr) - 1; i >= 0; i-- {
		src := arr[i]
		if i != len(arr)-1 {
			src = "\r\n\r\n" + arr[i]
		}
		if _, err := f.WriteString(src); err != nil {
			log.Fatal("Error on merge file", err)
		}
	}
}

func startNewPage(tempContent string, ctx *CtxIndexPage, dirOutAbs string) afero.File {
	afs := &afero.Afero{Fs: appfs}
	pageFileName := getOutPageFileName(ctx.CurrPageNum, dirOutAbs)
	f, err := afs.Create(pageFileName)
	if err != nil {
		log.Fatalln("Error create file", err)
	}
	log.Println("Initialize file ", pageFileName)

	var t = template.Must(template.New("Page").Parse(tempContent))
	err = t.Execute(f, ctx) // Nota come l'interfaccia File sia anche io.Writer
	if err != nil {
		log.Fatal("Template error: ", err)
	}
	return f
}

func getOutPageFileName(curPageNum int, dirOutAbs string) string {
	if curPageNum == 0 {
		return filepath.Join(dirOutAbs, "index.page")
	}
	return filepath.Join(dirOutAbs, fmt.Sprintf("index_%d.page", curPageNum))
}

func getTemplateByPage(curPageNum int) string {
	if curPageNum == 0 {
		return tempIndex0Page
	}
	return tempIndexOtherPages
}
