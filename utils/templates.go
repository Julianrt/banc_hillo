package utils

import(
  "log"
  "net/http"
  "html/template"

)

var templates = template.Must(template.New("t").ParseGlob(TemplatesDir()))
var errorTemplate = template.Must(template.ParseFiles(ErrorTemplateDir()))
  
func RenderErrorTemplate(w http.ResponseWriter, status int){
  w.WriteHeader(status)
  errorTemplate.Execute(w, nil)
}

func RenderTemplate(w http.ResponseWriter, name string, data interface{}){
  w.Header().Set("Content-Type", "text/html")
  
  err := templates.ExecuteTemplate(w, name, data)
  
  if err != nil{
    log.Println(err)
    RenderErrorTemplate(w, http.StatusInternalServerError)
  }
}

func TemplatesDir() string{
  //return fmt.Sprintf("%s/templates/**/*.html", server.templateDir)
  return "templates/**/*.html"
}

func ErrorTemplateDir() string{
  //return fmt.Sprintf("%s/templates/error.html", server.templateDir)
  return "templates/error.html"
}
