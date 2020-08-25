package admin

import (
	"net/http"

	db "github.com/JIeeiroSSt/web/config"
	view "github.com/JIeeiroSSt/web/config"
	model "github.com/JIeeiroSSt/web/models"
	session "github.com/JIeeiroSSt/web/utils"
	upload "github.com/JIeeiroSSt/web/utils"
)

func IndexAuthor(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	selDB, err := db.Query("select * from author")
	if err != nil {
		panic(err.Error())
	}
	author := model.Author{}
	res := []model.Author{}

	for selDB.Next() {
		var id int
		var name, image string
		err = selDB.Scan(&id, &name, &image)
		if err != nil {
			panic(err.Error())
		}

		author.Id = id
		author.Name = name
		author.Image = image

		res = append(res, author)
	}
	userName := session.GetUserName(r)
	if userName != "" {
		err = view.GvAdmin.Render(w, http.StatusOK, "author/index", res)
		if err != nil {
			panic(err.Error())
		}
	} else {
		http.Redirect(w, r, "/keikibook/login", 302)
	}
	defer db.Close()
}

func IndexSingleAuthor(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("select * from author where id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	author := model.Author{}

	for selDB.Next() {
		var id int
		var name, image string
		err = selDB.Scan(&id, &name, &image)
		if err != nil {
			panic(err.Error())
		}

		author.Id = id
		author.Name = name
		author.Image = image

	}
	userName := session.GetUserName(r)
	if userName != "" {
		err = view.GvAdmin.Render(w, http.StatusOK, "author/show", author)
		if err != nil {
			panic(err.Error())
		}
	} else {
		http.Redirect(w, r, "/keikibook/login", 302)
	}
	defer db.Close()
}

func EditAuthor(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("select * from author where id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	author := model.Author{}

	for selDB.Next() {
		var id int
		var name, image string
		err = selDB.Scan(&id, &name, &image)
		if err != nil {
			panic(err.Error())
		}

		author.Id = id
		author.Name = name
		author.Image = image

	}
	userName := session.GetUserName(r)
	if userName != "" {
		err = view.GvAdmin.Render(w, http.StatusOK, "author/edit", author)
		if err != nil {
			panic(err.Error())
		}
	} else {
		http.Redirect(w, r, "/keikibook/login", 302)
	}
	defer db.Close()
}

func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	userName := session.GetUserName(r)
	if userName != "" {
		if err := view.GvAdmin.Render(w, http.StatusOK, "author/new", nil); err != nil {
			panic(err.Error())
		}
	} else {
		http.Redirect(w, r, "/keikibook/login", 302)
	}
}

func InsertAuthor(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	name := r.FormValue("name")
	file := upload.UploadFile(w, r, "image")
	insert, err := db.Prepare("insert into author values(?,?)")
	if err != nil {
		panic(err.Error())
	}
	insert.Exec(name, file)
	http.Redirect(w, r, "/keikibook/admin/author", 301)
}

func UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	name := r.FormValue("name")
	file := upload.UploadFile(w, r, "image")
	id := r.FormValue("id")
	update, err := db.Prepare("update author set name=?,file=? where id=?")
	if err != nil {
		panic(err.Error())
	}
	update.Exec(name, file, id)
	http.Redirect(w, r, "/keikibook/admin/author", 301)
}

func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	nID := r.URL.Query().Get("id")
	delete, err := db.Prepare("delete from author where id=?")
	if err != nil {
		panic(err.Error())
	}
	delete.Exec(nID)
	http.Redirect(w, r, "/keikibook/admin/author", 301)
}
