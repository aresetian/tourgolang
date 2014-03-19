package template

import (
    "fmt"
    "net/http"
    "html/template"
)

var page = `<html>
  <body>
    {{template "tags" .Tags}}

    {{template "content" .Content}}

    {{template "comment" .Comment}}
  </body>
</html>`

var tags = `{{define "tags"}}
<div>
    {{.Name}}
<div>
{{end}}`

var content = `{{define "content"}}
<div>
   <p>{{.Title}}</p>
   <p>{{.Content}}</p>
</div>
{{end}}`

var comment = `{{define "comment"}}
<div>
    {{.Note}}
</div>
{{end}}`

type Tags struct {
   Id int
   Name string
}

type Content struct {
   Id int
   Title string
   Content string
}

type Comment struct {
   Id int
   Note string
}

type Page struct {
	Tags *Tags
	Content *Content
	Comment *Comment
}


func Print(res http.ResponseWriter) {
    
    pagedata := &Page{Tags:&Tags{Id:1, Name:"golang"},
			Content: &Content{Id:9, Title:"Hello", Content:"World!"},
    			Comment: &Comment{Id:2, Note:"Good Day!"}}
	tmpl := template.New("page")
	var err error
	if tmpl, err = tmpl.Parse(page); err != nil {
		fmt.Println(err)
	}
	if tmpl, err = tmpl.Parse(tags); err != nil {
		fmt.Println(err)
	}
	if tmpl, err = tmpl.Parse(comment); err != nil {
		fmt.Println(err)
	}
	if tmpl, err = tmpl.Parse(content); err != nil {
		fmt.Println(err)
	}
	tmpl.Execute(res, pagedata)
}