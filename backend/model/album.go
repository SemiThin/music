package model

type Album struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Singer   string `json:"singer"`
	SingerId int8   `json:"singer_id"`
	Cover    string `json:"cover"`
	Pubtime  int64  `json:"pubtime"`
	Pubdate  string `json:"pubdate,omitempty"`
}

func (Album) TableName() string {
	return "album"
}
