package model

type Toilet struct {
	Id    int64  `xorm:"pk autoincr int(64)" form:"id" json:"id"`
	Title string `xorm:"varchar(40)" json:"title" form:"title"`
	Kind  int    `xorm:"integer" json:"kind" form:"kind"`
}
