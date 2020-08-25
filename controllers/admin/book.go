package admin

import (
	"net/http"

	db "github.com/JIeeiroSSt/web/config"
	view "github.com/JIeeiroSSt/web/config"
	model "github.com/JIeeiroSSt/web/models"
	session "github.com/JIeeiroSSt/web/utils"
	upload "github.com/JIeeiroSSt/web/utils"
)

func IndexBook(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	selDB, err := db.Query("select * from book")
	if err != nil {
		panic(err.Error())
	}
	book := model.Book{}
	res := []model.Book{}
	for selDB.Next() {
		var id int
		var name, file, description string
		err = selDB.Scan(&id, &name, &file, &description)
		if err != nil {
			panic(err.Error())
		}
		book.Id = id
		book.Name = name
		book.File = file
		book.Description = description

		res = append(res, book)
	}
	userName := session.GetUserName(r)
	if userName != "" {
		err = view.GvAdmin.Render(w, http.StatusOK, "book/index", res)
		if err != nil {
			panic(err.Error())
		}
	} else {
		http.Redirect(w, r, "/admin/login", 302)
	}
	defer db.Close()
}

func IndexSingleBook(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("select * from book where id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	book := model.Book{}
	for selDB.Next() {
		var id int
		var name, file, description string
		err = selDB.Scan(&id, &name, &file, &description)
		if err != nil {
			panic(err.Error())
		}
		book.Id = id
		book.Name = name
		book.File = file
		book.Description = description

	}
	userName := session.GetUserName(r)
	if userName != "" {
		err = view.GvAdmin.Render(w, http.StatusOK, "book/show", book)
		if err != nil {
			panic(err.Error())
		}
	} else {
		http.Redirect(w, r, "/admin/login", 302)
	}
	defer db.Close()
}

func EditBook(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("select * from book where id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	book := model.Book{}
	for selDB.Next() {
		var id int
		var name, file, description string
		err = selDB.Scan(&id, &name, &file, &description)
		if err != nil {
			panic(err.Error())
		}
		book.Id = id
		book.Name = name
		book.File = file
		book.Description = description

	}
	userName := session.GetUserName(r)
	if userName != "" {
		err = view.GvAdmin.Render(w, http.StatusOK, "book/edit", book)
		if err != nil {
			panic(err.Error())
		}
	} else {
		http.Redirect(w, r, "/admin/login", 302)
	}
	defer db.Close()
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	userName := session.GetUserName(r)
	if userName != "" {
		if err := view.GvAdmin.Render(w, http.StatusOK, "book/new", nil); err != nil {
			panic(err.Error())
		}
	} else {
		http.Redirect(w, r, "/admin/login", 302)
	}
}

func InserBook(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	name := r.FormValue("name")
	file := upload.UploadFile(w, r, "file")
	description := r.FormValue("description")
	insert, err := db.Prepare("insert into book values(?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	insert.Exec(name, file, description)
	defer db.Close()
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	name := r.FormValue("name")
	file := upload.UploadFile(w, r, "file")
	description := r.FormValue("description")
	id := r.FormValue("id")
	update, err := db.Prepare("update book set name=?,file=?,description=? where id=?")
	if err != nil {
		panic(err.Error())
	}
	update.Exec(name, file, description, id)
	defer db.Close()
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	nId := r.URL.Query().Get("id")
	delete, err := db.Prepare("delete from book where id=?")
	if err != nil {
		panic(err.Error())
	}
	delete.Exec(nId)
	defer db.Close()
}
