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

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kalkayan/gumshoe/models"
)

type AuthController struct{}

func (c AuthController) Register(ctx *gin.Context) {
	var f models.User
	err := ctx.ShouldBindJSON(&f)

	// Validate the request
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	// Check if the user is unique or not
	if f.Exist() {
		ctx.JSON(http.StatusConflict, gin.H{"error": "User already exists."})
		return
	}

	// Register and return a new user
	f.Create()
	ctx.JSON(http.StatusOK, gin.H{"message": f})
}
