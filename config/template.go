package config

import (
	"html/template"
	"time"

	"github.com/foolin/goview"
)

//============== admin ============== //
var GvAdmin = goview.New(goview.Config{
	Root:      "views/admin",
	Extension: ".html",
	Master:    "layouts/master",
	Partials:  []string{"partials/ad"},
	Funcs: template.FuncMap{
		"title": "ADMIN",
		"sub": func(a, b int) int {
			return a - b
		},
		"copy": func() string {
			return time.Now().Format("2020")
		},
	},
	DisableCache: true,
})

//============== editor ============== //
var GvEditor = goview.New(goview.Config{
	Root:      "views/editor",
	Extension: ".html",
	Master:    "layouts/master",
	Partials:  []string{"partials/ad"},
	Funcs: template.FuncMap{
		"title": "EDITOR",
		"sub": func(a, b int) int {
			return a - b
		},
		"copy": func() string {
			return time.Now().Format("2020")
		},
	},
	DisableCache: true,
})

//============== subcriber ============== //
var GvSubcriber = goview.New(goview.Config{
	Root:      "views/subcriber",
	Extension: ".html",
	Master:    "layouts/master",
	Partials:  []string{"partials/ad"},
	Funcs: template.FuncMap{
		"title": "SUBCRIBER",
		"sub": func(a, b int) int {
			return a - b
		},
		"copy": func() string {
			return time.Now().Format("2020")
		},
	},
	DisableCache: true,
})

//============== writer ============== //
var GvWriter = goview.New(goview.Config{
	Root:      "views/writer",
	Extension: ".html",
	Master:    "layouts/master",
	Partials:  []string{"partials/ad"},
	Funcs: template.FuncMap{
		"title": "WRITER",
		"sub": func(a, b int) int {
			return a - b
		},
		"copy": func() string {
			return time.Now().Format("2020")
		},
	},
	DisableCache: true,
})

//============== home ============== //
var GvHome = goview.New(goview.Config{
	Root:      "views/home",
	Extension: ".html",
	Master:    "layouts/master",
	Partials:  []string{"partials/ad"},
	Funcs: template.FuncMap{
		"title": "HOME",
		"sub": func(a, b int) int {
			return a - b
		},
		"copy": func() string {
			return time.Now().Format("2020")
		},
	},
	DisableCache: true,
})

//============== home ============== //
var GvMain = goview.New(goview.Config{
	Root:      "views/main",
	Extension: ".html",
	Master:    "layouts/master",
	Partials:  []string{"partials/ad"},
	Funcs: template.FuncMap{
		"title": "keikibook",
		"sub": func(a, b int) int {
			return a - b
		},
		"copy": func() string {
			return time.Now().Format("2020")
		},
	},
	DisableCache: true,
})
