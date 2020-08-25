package admin

import (
	"net/http"

	db "github.com/JIeeiroSSt/web/config"
	view "github.com/JIeeiroSSt/web/config"
	model "github.com/JIeeiroSSt/web/models"
	session "github.com/JIeeiroSSt/web/utils"
)

func IndexUser(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	selDB, err := db.Query("select * from user")
	if err != nil {
		panic(err.Error())
	}
	user := model.User{}
	res := []model.User{}
	for selDB.Next() {
		var id, roleId, status int
		var username, password string
		err = selDB.Scan(&id, &username, &password, &roleId, &status)
		if err != nil {
			panic(err.Error())
		}
		user.Id = id
		user.Username = username
		user.Password = password
		user.RoleId = roleId
		user.Status = status
		res = append(res, user)
	}
	userName := session.GetUserName(r)
	if userName != "" {
		err = view.GvAdmin.Render(w, http.StatusOK, "user/index", res)
		if err != nil {
			panic(err.Error())
		}
	} else {
		http.Redirect(w, r, "/admin/login", 302)
	}
	defer db.Close()

}

func IndexSingleUser(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("select * from user where id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	user := model.User{}
	for selDB.Next() {
		var id, roleId, status int
		var username, password string
		err = selDB.Scan(&id, &username, &password, &roleId, &status)
		if err != nil {
			panic(err.Error())
		}
		user.Id = id
		user.Username = username
		user.Password = password
		user.RoleId = roleId
		user.Status = status
	}
	userName := session.GetUserName(r)
	if userName != "" {
		err = view.GvAdmin.Render(w, http.StatusOK, "user/show", user)
		if err != nil {
			panic(err.Error())
		}
	} else {
		http.Redirect(w, r, "/admin/login", 302)
	}
	defer db.Close()
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	nId := r.URL.Query().Get("id")
	delete, err := db.Prepare("delete from user where id=?")
	if err != nil {
		panic(err.Error())
	}
	delete.Exec(nId)
	defer db.Close()
	http.Redirect(w, r, "/keikibook/admin/user", 301)
}
