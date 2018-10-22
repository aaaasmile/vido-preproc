package vidopre

import (
	"fmt"
	"log"
	"path/filepath"
	"text/template"

	"github.com/spf13/afero"
)

type CtxIndexPage struct {
	TotPages     int
	PostPerPages int
	CurrPageNum  int
	ZeroPage     bool
	FirstPage    bool
	LastPage     bool
	NavDet       []NavPageDetail
}

type NavPageDetail struct {
	PageIx     string
	IsSelected bool
	IsLast     bool
}

func (p *CtxIndexPage) setCurrPage(pageNum int) {
	p.CurrPageNum = pageNum
	if pageNum == 0 {
		p.ZeroPage = true
		p.FirstPage = false
	} else if pageNum == 1 {
		p.ZeroPage = false
		p.FirstPage = true
	}
	if pageNum == (p.TotPages - 1) {
		p.LastPage = true
	} else {
		p.LastPage = false
	}
	p.NavDet = nil
	for i := 1; i < p.TotPages; i++ {
		s := fmt.Sprintf("%02d", i)
		p.NavDet = append(p.NavDet, NavPageDetail{PageIx: s, IsSelected: (i == pageNum), IsLast: i == p.TotPages-1})
	}
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
			writeReversePostInFile(f, reversed, ctx)
			curPageNum--
			f.Close()
			ctx.CurrPageNum = curPageNum
			reversed = nil
			f = startNewPage(getTemplateByPage(curPageNum), ctx, dirOutAbs)
		}
	}
	writeReversePostInFile(f, reversed, ctx)
	f.Close()
}

func writeReversePostInFile(f afero.File, arr []string, ctx *CtxIndexPage) {
	for i := len(arr) - 1; i >= 0; i-- {
		src := arr[i]
		if i != len(arr)-1 {
			src = "\r\n\r\n" + arr[i]
		}
		if _, err := f.WriteString(src); err != nil {
			log.Fatal("Error on merge file", err)
		}
	}
	if _, err := f.WriteString("\r\n\r\n"); err != nil {
		log.Fatal("Error on merge file", err)
	}
	appendNav(ctx, f)
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

	appendNav(ctx, f)

	return f
}

func appendNav(ctx *CtxIndexPage, f afero.File) {
	ctx.setCurrPage(ctx.CurrPageNum)
	var t2 = template.Must(template.New("PageNav").Parse(tempNavInPage))
	err := t2.Execute(f, ctx) // Nota come l'interfaccia File sia anche io.Writer
	if err != nil {
		log.Fatal("Template error: ", err)
	}
}

func getOutPageFileName(curPageNum int, dirOutAbs string) string {
	if curPageNum == 0 {
		return filepath.Join(dirOutAbs, "index.page")
	}
	return filepath.Join(dirOutAbs, fmt.Sprintf("index_%02d.page", curPageNum))
}

func getTemplateByPage(curPageNum int) string {
	if curPageNum == 0 {
		return tempIndex0Page
	}
	return tempIndexOtherPages
}
