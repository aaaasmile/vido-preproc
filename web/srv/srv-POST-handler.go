package srv

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/aaaasmile/vido-preproc/conf"
)

func getURLForRoute(uri string) string {
	arr := strings.Split(uri, "/")
	//fmt.Println("split: ", arr, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		ss := arr[i]
		if ss != "" {
			if !strings.HasPrefix(ss, "?") {
				//fmt.Printf("Url for route is %s\n", ss)
				return ss
			}
		}
	}
	return uri
}

type CustomBackendHandler struct {
	w   http.ResponseWriter
	req *http.Request
}

func handlePostOperations(w http.ResponseWriter, req *http.Request) error {
	start := time.Now()
	if conf.Current.DebugVerbose {
		log.Println("POST index", req.RequestURI)
	}
	lastPath := getURLForRoute(req.RequestURI)
	log.Println("Check the last path ", lastPath)
	switch lastPath {
	default:
		return fmt.Errorf("%s Not supported", lastPath)
	}

	cusHandler := &CustomBackendHandler{
		w:   w,
		req: req,
	}

	rawbody, err := ioutil.ReadAll(cusHandler.req.Body)
	if err != nil {
		return err
	}
	if conf.Current.DebugVerbose {
		log.Println("Body is: ", string(rawbody))
	}

	switch lastPath {

	default:
		return fmt.Errorf("%s Not supported", lastPath)
	}
	if err != nil {
		return err
	}

	t := time.Now()
	elapsed := t.Sub(start)
	log.Printf("Post handler total call duration: %v\n", elapsed)
	return nil
}
