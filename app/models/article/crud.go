package article

import (
	"github.com/zhaozhentao/goblog/pkg/model"
	"github.com/zhaozhentao/goblog/pkg/types"
)

// Get 通过 ID 获取文章
func Get(idStr string) (Article, error) {
	var article Article
	id := types.StringToInt(idStr)
	if err := model.DB.First(&article, id).Error; err != nil {
		return article, err
	}

	return article, nil
}
