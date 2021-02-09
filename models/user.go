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

package models

import (
	"github.com/kalkayan/gumshoe/core"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name       string `gorm:"default:null" form:"name"`
	Email      string `gorm:"unique;" form:"email" binding:"required"`
	Password   string `form:"password"`
	Resources  []Resource
	Identities []Identity
}

/*--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+*\
|							Model's Event Hook								  |
|--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--|
|  Models use gorm hooks like BeforeCreate, AfterCreate for performing even   |
|  based operations.														  |
\*--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+*/

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// Hash user's password before saving
	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	u.Password = string(hashed)
	return err
}

func (u *User) Create() {
	// Get the database session
	core.DB.Create(u)
}

// Exist checks and returns if a user with email already exist or not.
func (u *User) Exist() bool {
	var c int64
	core.DB.Model(&User{}).Where("email = ?", u.Email).Count(&c)

	return c > 0
}
