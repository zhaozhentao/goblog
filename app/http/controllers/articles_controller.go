package controllers

import (
	"database/sql"
	"fmt"
	"github.com/zhaozhentao/goblog/pkg/database"
	"github.com/zhaozhentao/goblog/pkg/logger"
	"github.com/zhaozhentao/goblog/pkg/route"
	"github.com/zhaozhentao/goblog/pkg/types"
	"html/template"
	"net/http"
)

type ArticleController struct {
}

type Article struct {
	Title, Body string
	ID          int64
}

func getArticleByID(id string) (Article, error) {
	article := Article{}
	query := "SELECT * FROM articles WHERE id = ?"
	err := database.DB.QueryRow(query, id).Scan(&article.ID, &article.Title, &article.Body)
	return article, err
}

func (*ArticleController) Show(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouteVariable("id", r)

	article, err := getArticleByID(id)

	// 3. 如果出现错误
	if err != nil {
		if err == sql.ErrNoRows {
			// 3.1 数据未找到
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 文章未找到")
		} else {
			// 3.2 数据库错误
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		}
	} else {
		// 4. 读取成功，显示文章
		tmpl, err := template.New("show.gohtml").
			Funcs(template.FuncMap{
				"RouteName2URL": route.Name2URL,
				"Int64ToString": types.Int64ToString,
			}).
			ParseFiles("resources/views/articles/show.gohtml")
		logger.LogError(err)
		tmpl.Execute(w, article)
	}
}
