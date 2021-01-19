package route

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

//File struct is used by HomeDirectory func
type file struct {
	Path     string `json:"path"`
	Filename string `json:"filename"`
	Display  string `json:"display"`
}

//Dir struct is used by HomeDirectory func
type Dir struct {
	Fs   []file `json:"fs"`
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

//FilesDB struct is used by virtual_db.json
type FilesDB struct {
	ID          int    `json:"id"`
	Filename    string `json:"filename"`
	Display     string `json:"display"`
	Path        string `json:"path"`
	UploadTime  string `json:"uploadtime"`
	ModfiedTime string `json:"modfiedtime"`
	Type        string `json:"type"`
	Size        int    `json:"size"`
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

	requestPath := c.Param("path") //추후 dir의 위치를 기반으로 디렉토리 리스트를 제공한다.
	fmt.Println("Request Directory Path : ", requestPath)
	var f []file
	if requestPath == "root" {
		f = append(f, file{Path: "root->profile.jpeg",
			Filename: "profile.jpeg",
			Display:  "image"})
		f = append(f, file{Path: "root->git-tutorial",
			Filename: "git-tutorial",
			Display:  "folder"})
	} else if requestPath == "root->git-tutorial" {
		f = append(f, file{Path: "root",
			Filename: "..",
			Display:  "folder"})
		f = append(f, file{Path: "root->git-tutorial->helloworld",
			Filename: "helloworld",
			Display:  "folder"})
	}
	dir := &Dir{
		Fs:   f,
		Path: requestPath,
	}

	return c.JSON(http.StatusOK, dir)
}

//HomeDetail response file detail for example, /home/detail?filepath=/root/profile.jpg
func HomeDetail(c echo.Context) error {
	fmt.Println("Param : " + c.Param("path"))
	fd := &FileDetail{
		Path:       "root",
		Filename:   "profile.jpeg",
		Type:       "JPEG",
		Size:       32321,
		UploadTime: "2021-01-13 13:32:17",
	}
	return c.JSON(http.StatusOK, fd)
}

/*
TODO
1. 폴더나 파일을 더블클릭했을 떄 해당 오프젝트가 폴더라면 폴더의 경로를 서버로 보내고 그 경로에 맞는 디렉토리 목록을 받는다. -> HomeDirectory 재사용
*/

//Stats response stats.html
func Stats(c echo.Context) error {
	return c.File("template/stats.html")
}
