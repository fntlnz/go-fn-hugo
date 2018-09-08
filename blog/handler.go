package function

import (
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
	lockFile    = "/tmp/website.lock"
	builtFolder = "/tmp/website-built"
)

// Handle a serverless request
func Handle(req []byte) string {
	// only build the first time
	if _, err := os.Stat(lockFile); os.IsNotExist(err) {
		if _, err := os.Create(lockFile); err != nil {
			log.Panic("could not create lock file", err)
		}
		gitRepo, ok := os.LookupEnv("repository")
		if !ok {
			log.Panic("repository environment variable not configured, needs a repository containing an hugo website")
		}
		if err := buildWebsite(gitRepo); err != nil {
			log.Panic("error building website:", err)
		}
		log.Print("website built")
		return "website not ready yet"
	}

	if _, err := os.Stat(path.Join(builtFolder, "index.html")); os.IsNotExist(err) {
		return "website not ready yet"
	}

	httpPath, ok := os.LookupEnv("Http_Path")
	if !ok {
		log.Panic("Http_Path environment variable not set")
	}

	curPath := path.Join(builtFolder, httpPath)

	stat, err := os.Stat(curPath)
	if err != nil {
		log.Panic(err)
	}
	if stat.IsDir() {
		curPath = path.Join(curPath, "index.html")
	}

	res, err := ioutil.ReadFile(curPath)
	if err != nil {
		log.Panic(err)
	}
	return string(res)
}

func buildWebsite(gitRepo string) error {
	log.Print("website built started for repo", gitRepo)

	_, err := git.PlainClone(cloneFolder, false, &git.CloneOptions{
		URL:      gitRepo,
		Progress: os.Stdout,
	})

	if err != nil {
		return err
	}

	log.Print("repository cloned with success")

	err = os.Chdir(cloneFolder) //todo remove this
	if err != nil {
		return err
	}

	var cfg = &deps.DepsCfg{}

	osFs := afero.NewOsFs()

	csd := hugolib.ConfigSourceDescriptor{
		Fs:         osFs,
		Path:       cloneFolder,
		WorkingDir: cloneFolder,
	}
	config, _, err := hugolib.LoadConfig(csd)
	if err != nil {
		return err
	}
	config.Set("workingDir", cloneFolder)
	config.Set("publishDir", builtFolder)
	config.Set("staticDir", path.Join(cloneFolder, "static"))
	config.Set("themesDir", path.Join(cloneFolder, "themes"))
	if baseurl, ok := os.LookupEnv("baseurl"); ok {
		config.Set("baseurl", baseurl)
	}

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

	log.Print("building website html files")
	sites.Build(hugolib.BuildCfg{})
	return nil
}
