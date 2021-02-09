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

package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Name stores the application's name. The value is used when the application
// needs to place its name in an event, UI or any other location.
var Name string

// Debug stores the flag
var Debug bool

func Configure() {
	// Load the .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file, please check for .env file in the root directory.")
	}

	// Update the application's configurations
	Name = os.Getenv("APP_NAME")
	Debug, _ = strconv.ParseBool(os.Getenv("APP_DEBUG"))
}
