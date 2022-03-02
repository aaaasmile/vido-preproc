package srv

import (
	"log"
	"path/filepath"

	"github.com/aaaasmile/vido-preproc/conf"
	"github.com/aaaasmile/vido-preproc/web/srv/helper"
	"github.com/spf13/afero"
)

var (
	appfs = afero.NewOsFs() // Wrapper del file os molto interessante
)

func (bl *CustomBackendHandler) getPostlistFromFS(rawbody []byte) error {
	list, err := getPostList(conf.Current.PostSourceDir)
	if err != nil {
		return err
	}
	rspdata := struct {
		List   []string
		Status string
	}{
		List:   list,
		Status: "OK",
	}

	return helper.WriteJsonResp(bl.w, rspdata)

}

func getPostList(dirIn string) ([]string, error) {
	dir, err := filepath.Abs(dirIn)
	if err != nil {
		return nil, err
	}
	log.Println("Get post list from :", dir)

	items, err := afero.ReadDir(appfs, dirIn) // sorted by name as default
	if err != nil {
		return nil, err
	}

	res := []string{}

	for _, item := range items {
		sourceFname := filepath.Join(dir, item.Name())
		res = append(res, sourceFname)
	}
	log.Println("Recognized posts: ", len(res))
	return res, nil
}
