package main

import (
	"log"
	"html/template"
	"net/http"
	"io"
)

func main()  {
	http.HandleFunc("/",Index)

	http.ListenAndServe("127.0.0.1:8866",nil)

}

func Index (w http.ResponseWriter , req *http.Request) {
	getTpl(w)
}




func getTpl(w io.Writer)  {
	const Tpl  = `
		<!doctype html>
		<html>
	 	<head>
			<title>{{.Title}}</title>
		</head>
		<body>
			<ul>
				{{ range .Item }}
	 				<li>{{ . }}</li>
				{{ else }}
					<li>Nothing</li>
				{{ end }}
			</ul>
		</body>
		</html>`
	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	// 创建一个新的模板，并且载入内容.
	t ,err := template.New("webpage").Parse(Tpl)

	check(err)

	data := struct{
		Title string
		Item []string
	}{
		Title:"HelloTpl",
		Item:[]string{
			"Hello",
			"heiman",
			"jaja",
		},
	}
	err = t.Execute( w , data)
	check(err)
}