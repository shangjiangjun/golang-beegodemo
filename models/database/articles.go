package database

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// gallery:  ["impro/article/6216e13585dd846393f2a6b14c8ae0e7.jpg", "impro/article/2341f33f4249bdfcae9074052cff5437.jpg"]
type Articles struct {
	Id           int64     `json:"id" orm:"column(id);pk;auto"`
	CategoryId   int8      `json:"category_id" orm:"column(category_id);default(4)" description:"归属分类"`
	Belong       string    `json:"belong" orm:"column(belong);size(50); default(article)" description:"归属：文章,公告"`
	Title        string    `json:"title" orm:"column(title);size(100);" description:"标题"`
	Intro        string    `json:"intro" orm:"column(intro);size(191);null" description:"简介"`
	Image        string    `json:"image" orm:"column(image);size(191)" description:"主图"`
	Gallery      string    `json:"gallery" orm:"column(gallery);null"`
	Tag          string    `json:"tag" orm:"column(tag);size(191);null" description:"标签"`
	IsRecommend  bool      `json:"is_recommend" orm:"column(is_recommend);default(0)" description:"推荐"`
	IsHot        bool      `json:"is_hot" orm:"column(is_hot);default(0)" description:"热门"`
	Content      string    `json:"content" orm:"column(content);null" description:"内容"`
	ReviewCount  uint      `json:"review_count" orm:"column(review_count);default(0)" description:"查阅数量"`
	LikeCount    uint      `json:"like_count" orm:"column(like_count);default(0)" description:"收藏数量"`
	CommentCount uint      `json:"comment_count" orm:"column(comment_count);default(0)" description:"评论数量"`
	ListOrder    uint      `json:"list_order" orm:"column(list_order);default(0)" description:"排序"`
	IsAudit      bool      `json:"is_audit" orm:"column(is_audit);default(1)" description:"发布状态"`
	CreatedAt    time.Time `json:"created_at" orm:"column(created_at);auto_now_add;type(datetime)"` //发布时间
	UpdatedAt    time.Time `json:"updated_at" orm:"column(updated_at);auto_now;type(datetime)"`     //修改时间
}

// 返回表名
func (t *Articles) TableName() string {
	return "articles"
}

func init() {
	orm.RegisterModel(new(Articles))
}
