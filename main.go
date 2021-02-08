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
	"github.com/kalkayan/gumshoe/routes"
)

func main() {
	// config provides a convient way to parameterize application with environ-
	// ment variables. A ".env" is expected in the root directory to create any
	// custom environment variables.
	config.Configure()

	// registers a database for the application, database connection string is
	// picked from the configurations loaded in the application. DB is a pkg
	// level variable and is available for all the models.
	core.DB.Register()

	// registers a router for the application, router is created in this scope
	// and routes are registered by passing router's reference. Router is an
	// encapsulation over gin framework.
	router := core.Router{}
	router.Register()

	// routes are registered to the application router. for further addition of
	// routes in the router, register them in Register method under "routes.go"
	// and not here.
	routes.Register(&router)

	// fire up the already booted application, and process request with their
	// specific response.
	router.Router.Run()
}
