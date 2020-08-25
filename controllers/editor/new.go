package editor

import (
	"net/http"
	"strconv"

	db "github.com/JIeeiroSSt/web/config"
	view "github.com/JIeeiroSSt/web/config"
	model "github.com/JIeeiroSSt/web/models"
	session "github.com/JIeeiroSSt/web/utils"
	upload "github.com/JIeeiroSSt/web/utils"
)

func IndexNews(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	selDB, err := db.Query("select id,name,description,checks from new ")
	if err != nil {
		panic(err.Error())
	}

	new := model.New{}
	res := []model.New{}

	for selDB.Next() {
		var id, checks int
		var name, description string
		if err = selDB.Scan(&id, &name, &description, &checks); err != nil {
			panic(err.Error())
		}
		new.Id = id
		new.Name = name
		new.Description = description
		new.Checks = checks
		res = append(res, new)
	}
	userName := session.GetUserName(r)
	if userName != "" {
		if err := view.GvAdmin.Render(w, http.StatusOK, "subcriber/index", nil); err != nil {
			panic(err.Error())
		}
	} else {
		http.Redirect(w, r, "/keikibook/login", 302)
	}
	defer db.Close()
}

func IndexCheckNew(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	selDB, err := db.Query("select id,name,description,checks from new where checks = 0")
	if err != nil {
		panic(err.Error())
	}

	new := model.New{}
	res := []model.New{}

	for selDB.Next() {
		var id, checks int
		var name, description string
		if err = selDB.Scan(&id, &name, &description, &checks); err != nil {
			panic(err.Error())
		}
		new.Id = id
		new.Name = name
		new.Description = description
		new.Checks = checks
		res = append(res, new)
	}
	userName := session.GetUserName(r)
	if userName != "" {
		if err := view.GvAdmin.Render(w, http.StatusOK, "new/check", nil); err != nil {
			panic(err.Error())
		}
	} else {
		http.Redirect(w, r, "/keikibook/login", 302)
	}
	defer db.Close()
}

func IndexSingleNew(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("select id,name,description,checks from new where id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	new := model.New{}
	for selDB.Next() {
		var id, checks int
		var name, description string
		if err = selDB.Scan(&id, &name, &description, &checks); err != nil {
			panic(err.Error())
		}
		new.Id = id
		new.Name = name
		new.Description = description
		new.Checks = checks
	}
	userName := session.GetUserName(r)
	if userName != "" {
		if err := view.GvAdmin.Render(w, http.StatusOK, "subcriber/show", nil); err != nil {
			panic(err.Error())
		}
	} else {
		http.Redirect(w, r, "/keikibook/login", 302)
	}
	defer db.Close()
}

func EditNew(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("select id,name,description,checks from new where id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	new := model.New{}
	for selDB.Next() {
		var id, checks int
		var name, description string
		if err = selDB.Scan(&id, &name, &description, &checks); err != nil {
			panic(err.Error())
		}
		new.Id = id
		new.Name = name
		new.Description = description
		new.Checks = checks
	}
	userName := session.GetUserName(r)
	if userName != "" {
		if err := view.GvAdmin.Render(w, http.StatusOK, "subcriber/edit", nil); err != nil {
			panic(err.Error())
		}
	} else {
		http.Redirect(w, r, "/keikibook/login", 302)
	}
	defer db.Close()
}

func CreateNew(w http.ResponseWriter, r *http.Request) {
	userName := session.GetUserName(r)
	if userName != "" {
		if err := view.GvAdmin.Render(w, http.StatusOK, "subcriber/new", nil); err != nil {
			panic(err.Error())
		}
	} else {
		http.Redirect(w, r, "/keikibook/login", 302)
	}
}

func InserNew(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	name := r.FormValue("name")
	content := r.FormValue("content")
	description := r.FormValue("description")
	image := upload.UploadFile(w, r, "image")
	insert, err := db.Prepare("insert into new(name,content,description,image) values(?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	insert.Exec(name, content, description, image)
	defer db.Close()
	http.Redirect(w, r, "/keikibook/subcriber/new", 302)
}

func UpdateNew(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	name := r.FormValue("name")
	content := r.FormValue("content")
	description := r.FormValue("description")
	image := upload.UploadFile(w, r, "image")
	id := r.FormValue("id")
	update, err := db.Prepare("update new set name=?,content=?,description=?,image=? where id=?")
	if err != nil {
		panic(err.Error())
	}
	update.Exec(name, content, description, image, id)
	defer db.Close()
	http.Redirect(w, r, "/keikibook/subcriber/new", 302)
}

func DeleteNews(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	id := r.URL.Query().Get("id")
	delete, err := db.Prepare("delete from new where id=?")
	if err != nil {
		panic(err.Error())
	}
	delete.Exec(id)
	defer db.Close()
	http.Redirect(w, r, "/keikibook/subcriber/new", 302)
}

func CheckNew(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	id := r.URL.Query().Get("id")
	check := r.URL.Query().Get("checks")
	checks, _ := strconv.Atoi(check)
	var idNews int
	if checks == 0 {
		idNews = 0
		insert, err := db.Prepare("insert into new set checks=? where id=?")
		if err != nil {
			panic(err.Error())
		}
		insert.Exec(idNews, id)
	} else if checks == 1 {
		idNews = 1
		insert, err := db.Prepare("insert into new set checks=? where id=?")
		if err != nil {
			panic(err.Error())
		}
		insert.Exec(idNews, id)
	}
	defer db.Close()
	http.Redirect(w, r, "/keikibook/subcriber/new/check", 302)
}
