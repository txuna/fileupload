package libdb

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
)

//FileDB struct is used by virtual_db.json
type FileDB struct {
	ID          int    `json:"id"`
	Filename    string `json:"filename"`
	Display     string `json:"display"`
	Path        string `json:"path"`
	Fullpath    string `json:"fullpath"`
	UploadTime  string `json:"uploadtime"`
	ModfiedTime string `json:"modfiedtime"`
	Type        string `json:"type"`
	Size        int    `json:"size"`
}

//FilesDB func is used by virtual_db.json -> https://stackoverflow.com/questions/21830447/json-cannot-unmarshal-object-into-go-value-of-type
type FilesDB struct {
	Files []FileDB `json:"files"`
}

//LogsDB struct is used by virtual_db.json
type LogsDB struct {
	ID       int    `json:"id"`
	Method   string `json:"method"`
	Filename string `json:"filename"`
	Type     string `json:"type"`
	Path     string `json:"path"`
	Time     string `json:"time"`
}

//Readjson func read json file and return route.FilesDB
func Readjson() (FilesDB, error) {
	var filesDB FilesDB
	r, err := ioutil.ReadFile("testDB/virtual_filedb.json")
	if err != nil {
		log.Fatalln(err)
		return filesDB, errors.New("Failed Read Json File")
	}
	err = json.Unmarshal(r, &filesDB)
	if err != nil {
		log.Fatalln(err)
		return filesDB, errors.New("json load Failed")
	}
	return filesDB, nil
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
