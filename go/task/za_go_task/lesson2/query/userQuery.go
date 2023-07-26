package query

type UserQuery struct {
	Page
	NameLike string `form:"nameLike"`
	AgeStart int    `form:"ageStart"`
	AgeEnd   int    `form:"ageEnd"`
}
