package admin

import (
	"net/http"

	view "github.com/JIeeiroSSt/web/config"
	session "github.com/JIeeiroSSt/web/utils"
)

func HomeAdmin(w http.ResponseWriter, r *http.Request) {
	userName := session.GetUserName(r)
	if userName != "" {
		if err := view.GvAdmin.Render(w, http.StatusOK, "book/index", nil); err != nil {
			panic(err.Error())
		}

	} else {
		http.Redirect(w, r, "/admin/login", 302)
	}
}
