/*
Copyright 2021 Manish Sahani (kalkayan)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package core

import (
	"log"

	"github.com/go-gorm/mysql"
	"gorm.io/gorm"
)

type Database struct {
	database *gorm.DB
}

var DB *gorm.DB

func RegisterDB() {
	dsn := "root:password@tcp(localhost:3306)/gumshoe?charset=utf8&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database.")
	}
	DB = database
	log.Println("Database connected")
}
