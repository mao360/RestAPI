package RestAPI

type User struct {
	Id       int    `json:"id"` // возможно не тот json тег
	Name     string `json:"name"`
	UserName string `json:"username"`
	Password string `json:"password"`
}
