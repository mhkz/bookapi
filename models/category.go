package models

type Category struct {
	Model

	Name        string `json:"name"`        // 分类名
	Slug        string `json:"slug"`        // 别名
	Description string `json:"description"` // 分类描述
	Pid         int    `json:"pid"`         // 父分类 ID
	CreateAt    string `json:"create_at"`   // 创建时间
	ModifiedAt  string `json:"modified_at"` // 修改时间
}
