package domain

// User 领域对象，是 DDD 中的 entity
// BO(business object)
type User struct {
	Id          int64  `json:"id"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	NickName    string `json:"nick_name"`
	Description string `json:"description"`
	Birth       string `json:"birth"`
}

//type Address struct {
//}
