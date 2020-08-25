package main

import (
	"github.com/JIeeiroSSt/web/controllers/admin"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/go-martini/martini"
)

func main() {
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{Directory: "views"}))

	m.Get("/keikibook/admin", admin.HomeAdmin)

	m.Group("/keikibook/admin/author", func(r martini.Router) {
		r.Get("/", admin.IndexAuthor)
		r.Get("/show", admin.IndexSingleAuthor)
		r.Get("/edit", admin.EditAuthor)
		r.Get("/new", admin.CreateAuthor)
		r.Post("/insert", admin.InsertAuthor)
		r.Post("/update", admin.UpdateAuthor)
		r.Delete("/delete", admin.DeleteAuthor)
	})

	m.Group("/keikibook/admin/book-author", func(r martini.Router) {
		r.Get("/create", admin.CreateBookAuthor)
		r.Post("/insert", admin.InsertAuthBook)
	})

	m.Group("/keikibook/admin/book", func(r martini.Router) {
		r.Get("/", admin.IndexBook)
		r.Get("/show", admin.IndexSingleBook)
		r.Get("/edit", admin.EditBook)
		r.Get("/new", admin.CreateBook)
		r.Post("/insert", admin.InserBook)
		r.Post("/update", admin.UpdateBook)
		r.Delete("/delete", admin.DeleteBook)
	})

	m.Group("/keikibook/admin/role", func(r martini.Router) {
		r.Get("/", admin.ShowChangeRole)
		r.Post("/change", admin.ChangeRole)
	})

	m.Group("/keikibook/admin/new", func(r martini.Router) {
		r.Get("/", admin.IndexNews)
		r.Get("/show", admin.IndexSingleNew)
		r.Get("/edit", admin.EditNew)
		r.Get("/new", admin.CreateNew)
		r.Post("/insert", admin.InserNew)
		r.Post("/update", admin.UpdateNew)
		r.Delete("/delete", admin.DeleteNews)
		r.Get("/check-new", admin.IndexCheckNew)
		r.Post("/check", admin.CheckNew)
	})

	m.Group("/keikibook/admin/user", func(r martini.Router) {
		r.Get("/", admin.IndexUser)
		r.Get("/show", admin.IndexSingleUser)
	})
	m.RunOnAddr(":9000")
}
