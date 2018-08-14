package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/itsjamie/go-bindata-templates"
	"github.com/gin-gonic/gin"
	"github.com/nu7hatch/gouuid"
	"github.com/elazarl/go-bindata-assetfs"
	"github.com/spf13/viper"
	"os"
	"fmt"
)

// App struct.
// There is no singleton anti-pattern,
// all variables defined locally inside
// this struct.
type App struct {
	Engine *gin.Engine
	Conf   *viper.Viper
	React  *React
	API    *API
}

// NewApp returns initialized struct
// of main server application.
func NewApp(opts ...AppOptions) *App {
	options := AppOptions{}
	for _, i := range opts {
		options = i
		break
	}

	options.init()


	envVar := os.Getenv("ENV")

	if envVar == ""{
		envVar = "development"
	}

	viper.Set("ENV", envVar)

	viper.SetConfigName("config-" + envVar)
	viper.AddConfigPath("server/config/")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// Make an engine
	engine := gin.New()

	// Logger for each request
	engine.Use(gin.Logger())

	// Recovery for any panics issue in system
	engine.Use(gin.Recovery())

	engine.GET("/favicon.ico", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/static/images/favicon.ico")
	})

	engine.LoadHTMLGlob("server/data/templates/*")

	// Initialize the application
	app := &App{
		Conf:   viper.GetViper(),
		Engine: engine,
		API:    &API{},
		React: NewReact(
			viper.GetString("duktape.path"),
			viper.GetBool("debug"),
			engine,
		),
	}

	// Map app and uuid for every requests
	app.Engine.Use(func(c *gin.Context) {
		c.Set("app", app)
		id, _ := uuid.NewV4()
		c.Set("uuid", id)
		c.Next()
	})

	// Bind api hadling for URL api.prefix
	app.API.Bind(
		app.Engine.Group(
			app.Conf.GetString("api.prefix"),
		),
	)

	// Create file http server from bindata
	fileServerHandler := http.FileServer(&assetfs.AssetFS{
		Asset:     Asset,
		AssetDir:  AssetDir,
		AssetInfo: AssetInfo,
	})

	//Serve static via bindata and handle via react app
	//in case when static file was not found
	engine.NoRoute(func(c *gin.Context){

		if _, err := Asset(c.Request.URL.Path[1:]); err == nil {
			fileServerHandler.ServeHTTP(
				c.Writer,
				c.Request)
			return
		}

		app.React.Handle(c)
	})


	return app
}

// Run runs the app
func (app *App) Run() {
	Must(app.Engine.Run(":" + app.Conf.GetString("port")))
}

// Template is custom renderer for Echo, to render html from bindata
type Template struct {
	templates *template.Template
}

// NewTemplate creates a new template
func NewTemplate() *Template {
	return &Template{
		templates: binhtml.New(Asset, AssetDir).MustLoadDirectory("templates"),
	}
}

// Render renders template
func (t *Template) Render(w io.Writer, name string, data interface{}, c gin.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// AppOptions is options struct
type AppOptions struct{}

func (ao *AppOptions) init() { /* write your own*/ }
