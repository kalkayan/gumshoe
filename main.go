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

package main

import (
	"github.com/kalkayan/gumshoe/config"
	"github.com/kalkayan/gumshoe/core"
	"github.com/kalkayan/gumshoe/models"
	"github.com/kalkayan/gumshoe/routes"
)

func main() {
	// config provides a convient way to parameterize the application with env
	// -ironment variables. The ".env" present at root  is expected to have all
	// the custom environment variables.
	config.Configure()

	// registers a database instance for the application, database connection
	// string is picked from the configurations loaded in the application.
	// RegisterDB binds a new gorm instance to a package level variable "DB".
	core.RegisterDB()

	// automigrate the database if the application is in debug mode. This should
	// be skipped in the production mode as it alters the database tables and
	// schema of the model, which can create unexpected crash of the application.
	if config.Debug {
		core.DB.AutoMigrate(&models.User{})
	}

	// registers a router for the application. Router is an encapsulation over
	// gin framework. Routes are added by passing router's reference.
	router := core.Router{}
	router.Register()

	// routes are registered to the application router. the further addition of
	// routes  should be done in Register method under "routes.go" and not here.
	routes.Register(&router)

	// fire up the already configured application. Process requests with their
	// appropriate response and return back to the client. The application by
	// default runs on port :8080.
	router.Router.Run()
}
