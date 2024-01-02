package model

type Song struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	AlbumId  int    `json:"album_id"`
	Source   string `json:"source"`
	Lyric    string `json:"lyric"`
	Duration int16  `json:"duration"`
	Album    Album  `json:"album";gorm:"foreignKey:AlbumId;references:Id"`
}

func (Song) TableName() string {
	return "song"
}
