package controllers

import (
	"log"
	"net/http"

	db "github.com/JIeeiroSSt/web/config"
	view "github.com/JIeeiroSSt/web/config"
	model "github.com/JIeeiroSSt/web/models"
	session "github.com/JIeeiroSSt/web/utils"
	"golang.org/x/crypto/bcrypt"
)

func ShowLogin(w http.ResponseWriter, r *http.Request) {
	if err := view.GvMain.Render(w, http.StatusOK, "login", nil); err != nil {
		panic(err.Error())
	}
}

func ShowSignUp(w http.ResponseWriter, r *http.Request) {
	if err := view.GvMain.Render(w, http.StatusOK, "signup", nil); err != nil {
		panic(err.Error())
	}
}

func Signup(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		password := r.FormValue("password")
		roleId := 4
		log.Println(name + "|" + password)
		insert, err := db.Prepare("insert into user(username,password,roleId) values(?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 8)
		log.Println(hashedPassword)
		log.Println(string(hashedPassword))
		insert.Exec(name, string(hashedPassword), roleId)
	}
	defer db.Close()
	http.Redirect(w, r, "/keikibook/login", 301)
}

func Signin(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	redirectTarget := "/keikibook"
	if r.Method == "POST" {
		name := r.FormValue("username")
		password := r.FormValue("password")
		result, err := db.Query("select password,roleId from user where username=?", name)
		if err != nil {
			panic(err.Error())
		}
		user := model.User{}
		//role := model.Role{}
		for result.Next() {
			var passwords string
			var roleId int
			if err = result.Scan(&passwords, &roleId); err != nil {
				panic(err.Error())
			}
			resultDB, err := db.Query("select code from role where id=?", roleId)
			if err != nil {
				panic(err.Error())
			}
			for resultDB.Next() {
				var code string
				if err = resultDB.Scan(&code); err != nil {
					panic(err)
				}
				log.Println(code)
				switch code {
				case "ADMINSTRATOR":
					user.Password = passwords
					if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
						w.WriteHeader(http.StatusUnauthorized)
						redirectTarget = "/keikibook/bug"
					}
					redirectTarget = "/keikibook/admin"
				case "EDITOR":
					user.Password = passwords
					if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
						w.WriteHeader(http.StatusUnauthorized)
						redirectTarget = "/keikibook/bug"
					}
					redirectTarget = "/keikibook/editor"
				case "WRITER":
					user.Password = passwords
					if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
						w.WriteHeader(http.StatusUnauthorized)
						redirectTarget = "/keikibook/bug"
					}
					redirectTarget = "/keikibook/writer"
				case "SUBSCRIBER":
					user.Password = passwords
					if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
						w.WriteHeader(http.StatusUnauthorized)
						redirectTarget = "/keikibook/bug"
					}
					redirectTarget = "/keikibook/subscriber"
				default:
					redirectTarget = "/bug"
				}
			}
		}
		session.SetSession(name, w)
	}
	defer db.Close()
	http.Redirect(w, r, redirectTarget, 301)
}
