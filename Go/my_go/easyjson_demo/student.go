package proto

import "time"

//easyjson:easyjson
type School struct {
	Name string `easyjson:"name"`
	Addr string `easyjson:"addr"`
}

//easyjson:easyjson
type Student struct {
	Id       int       `easyjson:"id"`
	Name     string    `easyjson:"s_name"`
	School   School    `easyjson:"s_chool"`
	Birthday time.Time `easyjson:"birthday"`
}
