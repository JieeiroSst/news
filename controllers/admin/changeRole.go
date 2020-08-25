package admin

import (
	"net/http"

	db "github.com/JIeeiroSSt/web/config"
	view "github.com/JIeeiroSSt/web/config"
	model "github.com/JIeeiroSSt/web/models"
	"github.com/foolin/goview"
)

func ChangeRole(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	if r.Method == "POST" {
		nId := r.URL.Query().Get("id")
		roleId := r.FormValue("key")
		insert, err := db.Prepare("update user set roleId=? where id=?")
		if err != nil {
			panic(err.Error())
		}
		insert.Exec(roleId, nId)
	}
	http.Redirect(w, r, "/keikibook/user", 301)
}

func ShowChangeRole(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	nId := r.URL.Query().Get("id")
	selUser, err := db.Query("select id,name from user where id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	user := model.User{}
	for selUser.Next() {
		var id int
		var name string
		if err = selUser.Scan(&id, &name); err != nil {
			panic(err.Error())
		}
		user.Id = id
		user.Username = name
	}

	selDB, err := db.Query("select id,code from role")
	if err != nil {
		panic(err.Error())
	}
	role := model.Role{}
	res := []model.Role{}
	for selDB.Next() {
		var id int
		var code string
		err = selDB.Scan(&id, &code)
		if err != nil {
			panic(err.Error())
		}
		role.Id = id
		role.Code = code
		res = append(res, role)
	}
	var Role = make(map[int]string)
	for _, index := range res {
		Role[index.Id] = index.Name
	}
	err = view.GvAdmin.Render(w, http.StatusOK, "user/edit", goview.M{
		"user": user,
		"Role": Role,
	})
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
}
