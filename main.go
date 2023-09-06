package main

import (
	"encoding/json"
	"golang.org/x/net/webdav"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func getConfigFileName() (configFileName string, err error) {
	xdgConfigDirName, err := os.UserConfigDir()
	if err != nil {
		return
	}

	twdsConfigDirName := filepath.Join(xdgConfigDirName, "twds")
	configFileName = filepath.Join(twdsConfigDirName, "config.json")

	return
}

func getSpaceDirName(config *Config) (dirName string, err error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return
	}

	dirName = config.Space

	if !filepath.IsAbs(dirName) {
		dirName = filepath.Join(homeDir, dirName)
	}

	return
}

func access_logger(pH *webdav.Handler, pReq *http.Request, err error) {
	log.Print(*pReq, err)
}

type Config struct {
	Listen string
	Prefix string
	Space  string
}

func closeFile(pFile *os.File) {
	if err := pFile.Close(); err != nil {
		log.Fatal(err)
	}
}

func getConfig(pConfig *Config) (err error) {
	configFileName, err := getConfigFileName()
	if err != nil {
		return
	}

	pConfigFile, err := os.Open(configFileName)
	if err != nil {
		return
	}
	defer closeFile(pConfigFile)

	dec := json.NewDecoder(pConfigFile)

	err = dec.Decode(pConfig)
	if err != nil {
		return
	}

	return
}

func main() {
	config := Config{
		Listen: "127.0.0.1:8080",
		Prefix: "/",
		Space:  ".twds/"}

	var err error

	err = getConfig(&config)
	if err != nil {
		log.Fatal(err)
	}

	spaceDirName, err := getSpaceDirName(&config)
	if err != nil {
		log.Fatal(err)
	}

	var fs twdsFS
	fs.init(spaceDirName)

	h := webdav.Handler{
		Prefix:     config.Prefix,
		FileSystem: &fs,
		LockSystem: webdav.NewMemLS()}
	h.Logger = func(pReq *http.Request, err error) { access_logger(&h, pReq, err) }

	http.HandleFunc("/", h.ServeHTTP)
	log.Fatal(http.ListenAndServe(config.Listen, nil))
}
