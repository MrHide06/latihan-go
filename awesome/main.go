package main

import (
	"fmt"
	"github.com/MrHide06/awesome/model"
)

func main() {
	user := model.CreateUser(10, "Foobar", "foobar", "foo@bar.com","password")
	article, _ := model.CreateArticle(1,  *user, "Title","Body")

	fmt.Println(article)
	fmt.Println(article.Hello())
}