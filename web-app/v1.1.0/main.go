package main

import (
   "log"
   "net/http"
   "io/ioutil"
   "text/template"
)

var tmpl = `
<!DOCTYPE html>
<html lang="en">
  <head>
    <link rel="stylesheet" href="css/bootstrap.min.css">
  </head>
  <div class="container">
  <body>
    <h1>Files</h1>
    <h4><small>There are {{.Count}} files.</small></h4>
    <table class="table table-striped">
    <tr>
      <th>Filename</th><th>Last Modified</th>
    </tr>
    {{range $filename, $lastmodified := .File}}<tr>
      <td>{{$filename}}</td><td>{{$lastmodified}}</td>
    </tr>{{end}}
  <table>
  </body>
  </div>
</html>
`

func handler(w http.ResponseWriter, r *http.Request) {
   fileMap := make(map[string]string)
   files, _ := ioutil.ReadDir("files/")
   for _, f := range files {
     fileMap[f.Name()] = f.ModTime().String() 
   }
   type FileData struct {
     File map[string]string
     Count int
   }
   data := FileData{fileMap, len(fileMap)}
   t := template.Must(template.New("env").Parse(tmpl))
   err := t.Execute(w, data)
   if err != nil {
     log.Println("executing template:", err)
   }
}

func main() {
   log.SetFlags(0)
   http.HandleFunc("/", handler)
   http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
   log.Println("Starting server...")
   log.Fatal(http.ListenAndServe(":8081", nil))
}
