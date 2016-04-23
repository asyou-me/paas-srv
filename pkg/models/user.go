package models

type user struct {
	Id   string "json:id"
	Auth Auth   "json:auth"
}

type Auth struct {
	Group Group "json:group"
}
