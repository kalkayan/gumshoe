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
	"github.com/gin-gonic/gin"
)

type Router struct {
	Router *gin.Engine
}

// Register method creates a router's instance for the application and encapsu-
// lates it in the Router.
func (r *Router) Register() {
	// create a gin's default router
	router := gin.Default()

	// encapsulate the above router in Router struct.
	r.Router = router
}
