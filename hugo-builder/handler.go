package function

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/gohugoio/hugo/deps"
	"github.com/gohugoio/hugo/helpers"
	"github.com/gohugoio/hugo/hugofs"
	"github.com/gohugoio/hugo/hugolib"
	"github.com/spf13/afero"
	jww "github.com/spf13/jwalterweatherman"
	"gopkg.in/src-d/go-git.v4"
)

const (
	cloneFolder = "/tmp/website"
	builtFolder = "/tmp/website-built"
)

// Handle a serverless request
func Handle(req []byte) string {
	gitRepo := string(req)
	if len(gitRepo) > 0 {
		if err := buildWebsite(gitRepo); err != nil {
			// todo(fntlnz): how to log errors?
			return "website build failed"
		}
	}

	// todo(fntlnz): how to serve the entire website?
	res, err := ioutil.ReadFile(path.Join(builtFolder, "index.html"))
	if err != nil {
		return "website not built yet, come back later"
	}
	return string(res)
}

func buildWebsite(gitRepo string) error {
	if len(gitRepo) == 0 {
		return fmt.Errorf("git repository not provided")
	}

	os.RemoveAll(cloneFolder)

	_, err := git.PlainClone(cloneFolder, false, &git.CloneOptions{
		URL:      gitRepo,
		Progress: os.Stdout,
	})

	if err != nil {
		return err
	}

	err = os.Chdir(cloneFolder)
	if err != nil {
		return err
	}

	var cfg = &deps.DepsCfg{}
	osFs := afero.NewOsFs()

	config, err := hugolib.LoadConfig(osFs, "", "")
	if err != nil {
		return err
	}
	config.Set("workingDir", cloneFolder)
	config.Set("publishDir", builtFolder)

	cfg.Cfg = config
	cfg.Logger = jww.NewNotepad(jww.LevelInfo, jww.LevelTrace, os.Stdout, ioutil.Discard, "", log.Ldate|log.Ltime)
	cfg.Logger.SetLogThreshold(jww.LevelTrace)
	cfg.Logger.SetStdoutThreshold(jww.LevelTrace)

	fs := hugofs.NewFrom(osFs, config)

	config.Set("cacheDir", helpers.GetTempDir("hugo_cache", fs.Source))

	cfg.Fs = fs

	sites, err := hugolib.NewHugoSites(*cfg)
	if err != nil {
		return err
	}

	sites.Build(hugolib.BuildCfg{})
	return nil
}
