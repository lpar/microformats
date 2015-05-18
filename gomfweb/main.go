package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strings"

	"github.com/andyleap/microformats"
)

var indextemplate = template.Must(template.New("index").Parse(index))

var parser = microformats.New()

func main() {
	http.Handle("/parse", http.HandlerFunc(Parse))
	http.Handle("/", http.HandlerFunc(Index))
	http.ListenAndServe(":4001", nil)
}

func Index(rw http.ResponseWriter, req *http.Request) {
	mf := req.FormValue("html")
	parsed := parser.Parse(strings.NewReader(mf))
	parsedjson, _ := json.MarshalIndent(parsed, "", "    ")

	data := struct {
		MF     string
		Parsed string
	}{
		mf,
		string(parsedjson),
	}

	indextemplate.Execute(rw, data)
}

func Parse(rw http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		data := struct {
			MF     string
			Parsed string
		}{
			"",
			"",
		}
		indextemplate.Execute(rw, data)
		return
	}
	mf := req.FormValue("html")
	parsed := parser.Parse(strings.NewReader(mf))
	parsedjson, _ := json.MarshalIndent(parsed, "", "    ")
	
	rw.Write(parsedjson)
}

var index = `<html>
<head>
</head>
<body>
<form method="POST">
<textarea name="html" style="width: 100%;" rows="15">{{.MF}}</textarea>
<br>
<input type="submit" value="Parse"/>
</form><br>
<code><pre>
{{.Parsed}}
</pre></code>
</body>
</html>`
