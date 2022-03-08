package srv

import (
	"html/template"
	"log"
	"net/http"
	"net/url"

	"github.com/aaaasmile/vido-preproc/conf"
	"github.com/aaaasmile/vido-preproc/util"
	"github.com/aaaasmile/vido-preproc/web/idl"
)

type PageCtx struct {
	Buildnr        string
	RootUrl        string
	VuetifyLibName string
	VueLibName     string
	Env            string
}

func handleIndexGet(w http.ResponseWriter, req *http.Request) error {
	u, _ := url.Parse(req.RequestURI)

	log.Println("GET requested ", u)

	lastPath := getURLForRoute(req.RequestURI)
	log.Println("Check the last path ", lastPath)

	switch lastPath {
	default:
		return handleGetApp(w, req)
	}
}

func handleGetApp(w http.ResponseWriter, req *http.Request) error {
	pagectx := PageCtx{
		RootUrl:        conf.Current.RootURLPattern,
		Buildnr:        idl.Buildnr,
		VuetifyLibName: conf.Current.VuetifyLibName,
		VueLibName:     conf.Current.VueLibName,
		Env:            "Dev",
	}
	templName := "templates/vue/index.html"

	tmplIndex := template.Must(template.New("AppIndex").ParseFiles(util.GetFullPath(templName)))

	err := tmplIndex.ExecuteTemplate(w, "base", pagectx)
	if err != nil {
		return err
	}
	return nil
}
