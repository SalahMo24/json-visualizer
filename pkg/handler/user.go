package handler

import (
	"fmt"
	jsonmodule "json-visualizer/pkg/json-module"
	localizationparser "json-visualizer/pkg/localization-parser"
	"json-visualizer/pkg/views/user"
	"path/filepath"
	"time"

	"github.com/labstack/echo/v4"
)

type UserHandler struct{}

func readFile(fileName string) jsonmodule.Input {
	path, err := filepath.Abs(fileName)
	if err != nil {
		panic(err)
	}

	file, err := jsonmodule.NewFile(path)
	if err != nil {
		panic(err)
	}
	start := time.Now()
	entries, err := file.Reader()
	if err != nil {
		panic(err)
	}
	elapsed := time.Since(start)
	fmt.Println("Time elapsed: ", elapsed)
	return entries
}
func (h UserHandler) HandleUser(c echo.Context) error {

	enEntries := readFile("en.json")
	arEntries := readFile("ar.json")
	merged := localizationparser.Merge(enEntries, arEntries, "en", "ar")

	return render(c, user.Show(merged))
}

func (h UserHandler) HandleUpdate(c echo.Context) error {
	data, err := c.MultipartForm()
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
	return c.String(200, "kkk")
}
