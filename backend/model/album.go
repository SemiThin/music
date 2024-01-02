package model

type Album struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Singer   string `json:"singer"`
	SingerId int8   `json:"singer_id"`
	Cover    string `json:"cover"`
	Pubdate  int    `json:"pubdate"`
}

func (Album) TableName() string {
	return "album"
}
