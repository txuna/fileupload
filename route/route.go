package route

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/labstack/echo"
	"github.com/txuna/fileupload/libdb"
)

//File struct is used by HomeDirectory func
type File struct {
	Path     string `json:"path"`
	Filename string `json:"filename"`
	Display  string `json:"display"`
}

//Dir struct is used by HomeDirectory func
type Dir struct {
	Fs   []File `json:"fs"`
	Path string `json:"path"`
}

//FileDetail struct is used by HomeDetail func , maybe add display member
type FileDetail struct {
	Path       string `json:"path"`
	Filename   string `json:"filename"`
	Type       string `json:"type"`
	Size       int    `json:"size"`
	UploadTime string `json:"uploadtime"`
}

/*
Home response home.html - 디렉토리가 아닌 기본적인 부분만 넘겨준다.
*/
func Home(c echo.Context) error {
	return c.Render(http.StatusOK, "home.html", nil)
}

/*
HomeDirectory func은 폴더를 fetch로 해당 경로가 요청된다면 요청된 경로를 기반으로 file struct array를 response로 보내준다.
return JSON - 보낼 때 현재 디렉토리 위치도 보내야 하는데 + 뒤로가기 버튼도 만들어야 함.
요청된 경로가  파일인지 폴더인지 서버단에서도 구별이 필요함.
*/
func HomeDirectory(c echo.Context) error {

	filesDB, err := libdb.ReadFileDBjson()
	if err != nil {
		return err
	}
	requestPath := c.Param("path") //추후 dir의 위치를 기반으로 디렉토리 리스트를 제공한다.
	var file []File

	for _, fileDB := range filesDB.Files {
		if fileDB.Path == requestPath {
			file = append(file, File{
				Path:     fileDB.Path + "->" + fileDB.Filename,
				Filename: fileDB.Filename,
				Display:  fileDB.Display,
			})
		}
	}

	dir := &Dir{
		Fs:   file,
		Path: requestPath,
	}
	return c.JSON(http.StatusOK, &dir)
}

//HomeDetail response file detail for example, /home/detail?filepath=/root/profile.jpg
func HomeDetail(c echo.Context) error {
	requestPath := c.Param("path")
	filesDB, err := libdb.ReadFileDBjson()
	var filedetail FileDetail
	if err != nil {
		return err
	}
	for _, fileDB := range filesDB.Files {
		if fileDB.Fullpath == requestPath {
			filedetail = FileDetail{
				Path:       fileDB.Path,
				Filename:   fileDB.Filename,
				Type:       fileDB.Type,
				Size:       fileDB.Size,
				UploadTime: fileDB.UploadTime,
			}
		}
	}

	return c.JSON(http.StatusOK, filedetail)
}

//Stats response stats.html
func Stats(c echo.Context) error {
	return c.File("template/stats.html")
}

/*
StatsLog func response logs => page넘버가 넘어오는데 한 페이지당 10개씩 보여준다.
*/
func StatsLog(c echo.Context) error {
	pageStr := c.Param("page")
	fmt.Println(reflect.TypeOf(pageStr))
	logsDB, err := libdb.PagenationLogDB(pageStr)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, logsDB.Logs)
}
