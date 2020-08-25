package admin

import (
	"net/http"

	db "github.com/JIeeiroSSt/web/config"
	view "github.com/JIeeiroSSt/web/config"
	model "github.com/JIeeiroSSt/web/models"
	"github.com/foolin/goview"
)

func CreateBookAuthor(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	selAuth, err := db.Query("select id,name from author")
	if err != nil {
		panic(err.Error())
	}
	author := model.Author{}
	res := []model.Author{}
	for selAuth.Next() {
		var id int
		var name string
		err = selAuth.Scan(&id, &name)
		if err != nil {
			panic(err.Error())
		}
		author.Id = id
		author.Name = name
		res = append(res, author)
	}
	var Author = make(map[int]string)
	for _, index := range res {
		Author[index.Id] = index.Name
	}

	selBook, err := db.Query("select id,name from book")
	if err != nil {
		panic(err.Error())
	}
	book := model.Book{}
	ress := []model.Book{}
	for selBook.Next() {
		var id int
		var name string
		err = selBook.Scan(&id, &name)
		if err != nil {
			panic(err.Error())
		}
		book.Id = id
		book.Name = name
		ress = append(ress, book)
	}
	var Book = make(map[int]string)
	for _, index := range ress {
		Book[index.Id] = index.Name
	}
	err = view.GvAdmin.Render(w, http.StatusOK, "book/changle", goview.M{
		"Author": Author,
		"Book":   Book,
	})
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
}

func InsertAuthBook(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	idAuth := r.FormValue("key")
	idBook := r.FormValue("keys")
	insert, err := db.Prepare("insert into auth_book values(?,?)")
	if err != nil {
		panic(err.Error())
	}
	insert.Exec(idAuth, idBook)
	defer db.Close()
}
