package models

type User struct {
	Id   string "json:id"
	Auth Auth   "json:auth"
}

type Auth struct {
	Group Group "json:group"
}

type UserCahce struct {
}

func Token() {

}
