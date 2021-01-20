package libdb

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"strconv"
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

//LogDB struct is used by virtual_filedb.json
type LogDB struct {
	ID       int    `json:"id"`
	Method   string `json:"method"`
	Filename string `json:"filename"`
	Type     string `json:"type"`
	Path     string `json:"path"`
	Time     string `json:"time"`
}

//LogsDB struct is used by virtual_logdb.json
type LogsDB struct {
	Logs []LogDB
}

var (
	errJSONRead     = errors.New("Failed Read Json File")
	errJSONLoad     = errors.New("json load Failed")
	errAtoiConv     = errors.New("Atoi Convertion Error")
	errIncorretPage = errors.New("Incorrect Page number")
)

//ReadFileDBjson func read json file and return route.FilesDB
func ReadFileDBjson() (FilesDB, error) {
	var filesDB FilesDB
	r, err := ioutil.ReadFile("testDB/virtual_filedb.json")
	if err != nil {
		log.Println(errJSONRead)
		return FilesDB{}, errJSONRead
	}
	err = json.Unmarshal(r, &filesDB)
	if err != nil {
		log.Println(errJSONLoad)
		return FilesDB{}, errJSONLoad
	}
	return filesDB, nil
}

//ReadLogDBjson func read logs json file
func ReadLogDBjson() (LogsDB, error) {
	var logsDB LogsDB
	r, err := ioutil.ReadFile("testDB/virtual_logdb.json")
	if err != nil {
		log.Println(errJSONRead)
		return LogsDB{}, errJSONRead
	}

	err = json.Unmarshal(r, &logsDB)
	if err != nil {
		log.Println(errJSONLoad)
		return LogsDB{}, errJSONLoad
	}
	return logsDB, nil
}

var pageUnit = 2

/*
PagenationLogDB func load logs as much as page number
*/
func PagenationLogDB(pageStr string) (LogsDB, error) {
	logsDB, err := ReadLogDBjson()
	var pageLogDB LogsDB
	if err != nil {
		return LogsDB{}, err
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		log.Println(errAtoiConv)
		return LogsDB{}, errAtoiConv
	}

	var (
		startPage = (page - 1) * pageUnit
		lastPage  = (page-1)*pageUnit + pageUnit
	)

	if startPage < 0 || startPage > len(logsDB.Logs) {
		log.Println(errIncorretPage)
		return LogsDB{}, errIncorretPage
	}

	for i := startPage; i < lastPage; i++ {
		if i > len(logsDB.Logs) {
			return pageLogDB, nil
		}
		pageLogDB.Logs = append(pageLogDB.Logs, logsDB.Logs[i])
	}
	return pageLogDB, nil

}
