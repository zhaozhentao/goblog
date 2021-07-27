package article

import (
	"github.com/zhaozhentao/goblog/app/models"
	"github.com/zhaozhentao/goblog/pkg/route"
)

// Article 文章模型
type Article struct {
	models.BaseModel
	Title string
	Body  string
}

// Link 方法用来生成文章链接
func (a Article) Link() string {
	return route.Name2URL("articles.show", "id", a.GetStringID())
}
