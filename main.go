package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

func isPicture(fileName string) bool {
	return strings.HasSuffix(fileName, "png") || strings.HasSuffix(fileName, "jpg") || strings.HasSuffix(fileName, "jpeg") || strings.HasSuffix(fileName, "JPG") || strings.HasSuffix(fileName, "PNG") || strings.HasSuffix(fileName, "JPEG")
}

func main() {
	var rootPath string
	if len(os.Args) > 1 {
		rootPath = os.Args[1]
	} else {
		rootPath = "."
	}
	if _, err := os.Stat(rootPath); os.IsNotExist(err) {
		log.Fatal(err)
	}
	files, err := ioutil.ReadDir(rootPath)
	if err != nil {
		log.Fatal(err)
	}
	var names []string
	for _, f := range files {
		name := f.Name()
		if isPicture(name) {
			names = append(names, fmt.Sprintf("<img src=\"%s\" alt=\"%s\"/>", path.Join(rootPath, name), name))
		}
	}
	if len(names) == 0 {
		fmt.Println("no image found")
		return
	}
	htmlBody := fmt.Sprint(strings.Join(names[:], "\n"))
	htmlTemplate := `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<meta http-equiv="X-UA-Compatible" content="ie=edge">
<style type="text/css">
body { display: flex; flex-direction: column; align-items: center; }
img { max-width: 100%s; }
</style>
<title>Picture Gallery</title>
</head>
<body>
%s
</body>
</html>
`
	htmlResult := fmt.Sprintf(htmlTemplate, "%", htmlBody)
	file, err := os.Create("index.html")
	if err != nil {
		log.Fatal(err)
	}
	_, err = file.WriteString(htmlResult)
	if err != nil {
		log.Fatal(err)
		file.Close()
		return
	}
	fmt.Println("create index.html done")
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)
	log.Println("http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
