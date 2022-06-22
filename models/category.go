// 分类相关
package models

// 从数据库读取
type Category struct {
	Cid      int
	Name     string
	CreateAt string
	UpdateAt string
}

type CategoryResponse struct {
	*HomeRerponse
	CategoryName string
}
