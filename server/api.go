package main

import (
	"github.com/gin-gonic/gin"
)

// API is a defined as struct bundle
// for api. Feel free to organize
// your app as you wish.
type API struct{}

// Bind attaches api routes
func (api *API) Bind(group *gin.RouterGroup) {
	group.GET("/v1/conf", api.ConfHandler)
}

// ConfHandler handle the app config, for example
func (api *API) ConfHandler(c *gin.Context) {
	if app, ok := c.Get("app"); ok {
		c.JSON(200, app.(*App).Conf.AllSettings())
	} else {
		c.String(400, "False")
	}
}
