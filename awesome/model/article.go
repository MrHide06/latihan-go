package model

// import {
// 	"github.com/MrHide06/awesome/resource"
// }

type Article struct {
	ID,AuthorID int
	Title,Body string 
}

func (a *Article) Hello()string{
	return "Hello" + a.Title
}

func CreateArticle(id int, user User, title, body string)(*Article, error){
	return &Article{
		ID: id,
		AuthorID: user.ID,
		Title: title,
		Body: body,
	},nil
}