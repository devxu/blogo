package models

import (
	"time"
)

type Comment struct {
	Id      int64
	PostId  int64 `xorm:"post_id bigint"`
	Name    string
	Email   string
	Website string
	Message string `xorm:"text"`
	Created time.Time
}

//func (comment *Comment) Validate(v *revel.Validation) {
//	v.Required(comment.Name).Message("请填写昵称")
//	v.Required(comment.Message).Message("请填写留言内容")
//	if utf8.RuneCountInString(comment.Name) > 10 {
//		err := &revel.ValidationError{Message: "昵称过长", Key: "comment.Name"}
//		v.Errors = append(v.Errors, err)
//	}
//	if utf8.RuneCountInString(comment.Message) > 200 {
//		err := &revel.ValidationError{Message: "留言内容不能超过200个字", Key: "comment.Message"}
//		v.Errors = append(v.Errors, err)
//	}
//}

func (comment *Comment) GetCreated() time.Time {
	return comment.Created.Add(-time.Hour * 8)
}
