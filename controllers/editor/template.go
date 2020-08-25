package editor

import "html/template"

var tmpl = template.Must(template.ParseGlob("views/*"))
