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

package routes

import (
	"github.com/kalkayan/gumshoe/controllers"
	"github.com/kalkayan/gumshoe/core"
)

func Register(r *core.Router) {
	// Group routes for v1 api
	v1 := r.Router.Group("v1")
	{
		v1.POST("/register", new(controllers.AuthController).Register)
	}
}
