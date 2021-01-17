package route

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

// Person strcut is part of test.html
type Person struct {
	Name string
}

//File struct is used by HomeDirectory func
type File struct {
	Path    string `json:"path"`
	Name    string `json:"name"`
	Display string `json:"display"`
}

//FileDetail struct is used by HomeDetail func
type FileDetail struct {
	Path       string `json:"path"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Size       int    `json:"size"`
	UploadTime string `json:"uploadtime"`
}

//HomeTemplate struct is used by Home func
type HomeTemplate struct {
	StartPath string
}

//Test is test Request function - e.GET
func Test(c echo.Context) error {
	person := Person{Name: "Tuuna"}
	return c.Render(http.StatusOK, "test.html", person)
}

/*
Home response home.html - 디렉토리가 아닌 기본적인 부분만 넘겨준다.
*/
func Home(c echo.Context) error {
	return c.Render(http.StatusOK, "home.html", nil)
}

/*
HomeDirectory func은 폴더를 fetch로 해당 경로가 요청된다면 요청된 경로를 기반으로 file struct array를 response로 보내준다.
return JSON
*/
func HomeDirectory(c echo.Context) error {
	fmt.Println("QeuryParam : " + c.QueryParam("dir")) //추후 dir의 위치를 기반으로 디렉토리 리스트를 제공한다.
	f := &[2]File{
		{Path: "/root/profile.jpeg",
			Name:    "profile.jpeg",
			Display: "image"},
		{Path: "/root/git-tutorial",
			Name:    "git-tutorial",
			Display: "folder"},
	}
	fmt.Println(f)
	return c.JSON(http.StatusOK, f)
}

//HomeDetail response file detail for example, /home/detail?filepath=/root/profile.jpg
func HomeDetail(c echo.Context) error {
	fmt.Println("QeuryParam : " + c.QueryParam("filepath"))
	fd := &FileDetail{
		Path:       "/root",
		Name:       "profile.jpeg",
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
