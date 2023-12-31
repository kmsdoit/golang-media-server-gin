package util

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DB_INIT(user string, password string, host string, name string) *gorm.DB {

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       user + ":" + password + "@" + "tcp(" + host + ":3306" + ")" + "/" + name + "?charset=utf8&parseTime=True&loc=Local",
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	return db
}
