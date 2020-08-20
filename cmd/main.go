/*
 * Author: shikanon (shikanon@tensorbytes.com)
 * File Created Time: 2020-07-21 5:52:17
 *
 * Project: editor
 * File: main.go
 * Description:
 *
 */
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"filesystem"
	_ "filesystem/cmd/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// API version
const Version = "v1"

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /filesystem/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

func main() {
	router := gin.Default()

	// Simple group: v1
	v1 := router.Group(fmt.Sprintf("/filesystem/%v/", Version))
	{
		v1.POST("/upload", filesystem.Upload)
		v1.POST("/delete", filesystem.Delete)
		v1.POST("/copy", filesystem.Copy)
		v1.POST("/upload_policy", filesystem.UoloadPolicy)
	}
	// 接口文档
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8080")
}
