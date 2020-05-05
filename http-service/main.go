package main

import (
	"net/http"
	"strconv"
	"github.com/labstack/echo"
	"github.com/MrHide06/http-service/model"
)

func main(){
	store := model.NewArticleStoreInMemory()

	e := echo.New()

	e.GET("/articles", func(c echo.Context)error{
		articles := store.ArticleMap

		return c.JSON(http.StatusOK, articles)
	})

	e.GET("/articles/:id", func(c echo.Context)error{
		articles := store.ArticleMap
		id, _ := strconv.Atoi(c.Param("id"))
		return c.JSON(http.StatusOK, articles[id - 1])
	})

	e.POST("/articles", func(c echo.Context)error{
		title := c.FormValue("Title")
		body := c.FormValue("Body")

		article, _ := model.CreateArticle(title, body)

		store.Save(article)

		return c.JSON(http.StatusOK, article)
	})

	e.PUT("/articles/:id", func(c echo.Context)error{
		title := c.FormValue("Title")
		body := c.FormValue("Body")
		id, _ := strconv.Atoi(c.Param("id"))
		articles := store.ArticleMap

		articles[id - 1] = model.Article{ID: id, Title:title, Body: body}

		return c.JSON(http.StatusOK, articles)
	})

	e.DELETE("/articles/:id", func(c echo.Context)error{
		id, _ := strconv.Atoi(c.Param("id"))
		articles := store.ArticleMap

		articles = append(articles[:id - 1], articles[id:]...)
		
		return c.JSON(http.StatusOK, articles)
	})
	e.Logger.Fatal(e.Start(":8080"))

	// e.GET("/", func(c echo.Context) error{
	// 	nama := c.QueryParam("nama")
	// 	age := c.QueryParam("age")
	// 	return c.String(http.StatusOK, nama+age)
	// })
	// e.Logger.Fatal(e.Start(":8080"))


	// e.POST("/", func(c echo.Context) error{
	// 	nama := c.FormValue("nama")
	// 	age := c.FormValue("age")
	// 	return c.String(http.StatusOK, nama+," ",+age)
	// })
	// e.Logger.Fatal(e.Start(":8080"))

	// e.GET("articles", func(c echo.Context) error{
	// 	return c.String(http.StatusOK, "untuk mendapatkan list")
	// })

	// e.POST("articles", func(c echo.Context) error{
	// 	return c.String(http.StatusOK, "untuk create articles")
	// })

	// e.PUT("articles/:id", func(c echo.Context) error{
	// 	return c.String(http.StatusOK, "untuk update Article")
	// })

	// e.DELETE("articles/:id", func(c echo.Context) error{
	// 	return c.String(http.StatusOK, "untuk delete article")
	// })

	// e.GET("articles/:id", func(c echo.Context) error{
	// 	return c.String(http.StatusOK, "untuk get article by id")
	// })

	// e.Logger.Fatal(e.Start(":8080"))

}