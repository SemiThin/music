package database

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viperErr := viper.ReadInConfig()
	if viperErr != nil {
		fmt.Printf("read database failed: %v", viperErr)
	}

	host := viper.Get("mysql.host")
	port := viper.Get("mysql.port")
	user := viper.Get("mysql.user")
	password := viper.Get("mysql.password")
	database := viper.Get("mysql.database")

	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, database)
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	if Db.Error != nil {
		fmt.Printf("database error %v", Db.Error)
	}
}
