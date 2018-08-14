# gin-react-starter-kit [![wercker status](https://app.wercker.com/status/cc4ddd2e2fec29ad988115a8b20b830a/s/master "wercker status")](https://app.wercker.com/project/byKey/cc4ddd2e2fec29ad988115a8b20b830a)

> This project was inspired by olebedev's [golang-starter-kit](https://github.com/olebedev/go-starter-kit) and improve some features which can bootstrap SPA web quickly and efficiently based on **Facebook React** and **Gin Golang Server side framework**.

## Features
### Front end
* Routing via [react-router](https://github.com/reactjs/react-router)
* ES6 & JSX via [babel-loader](https://github.com/babel/babel-loader) with minimal runtime dependency footprint
* [Redux](https://rackt.org/redux/) as state container
* [Redux-devtools](https://github.com/gaearon/redux-devtools)
* [Redux Saga](https://github.com/redux-saga) for asynchronous requests
* Hot reloading via [react-transform](https://github.com/gaearon/babel-plugin-react-transform) & [HMR](http://webpack.github.io/docs/hot-module-replacement.html)
* Css styles without global namespace via PostCSS, [css-loader](https://github.com/webpack/css-loader) & css-modules
* Webpack bundle builder
* Eslint and golint rules for Makefile

### Back end
* Server side render via [goja](https://github.com/dop251/goja)
* Api requests between your react application and server side application directly  via [fetch polyfill](https://github.com/olebedev/gojax/tree/master/fetch)
* Title, Open Graph and other domain-specific meta tags render for each page at the server and at the client
* Server side redirect
* Embedding static files into artefact via bindata
* Popular golang [gin](https://github.com/gin-gonic/gin) framework
* Advanced cli via [cli](https://github.com/codegangsta/cli)
* Makefile based projectd
* Separated config files: development, staging, production via [viper](https://github.com/spf13/viper)


## Workflow dependencies

* [Golang](https://golang.org/)
* [Node.js](https://nodejs.org/) with [yarn](https://yarnpkg.com)
* [GNU make](https://www.gnu.org/software/make/)
* [Go Dep](https://github.com/golang/dep) For golang packages management

Note that probably not works at windows.

## Project structure

##### The server's entry point
```
$ tree server
server
├── config <-- Config file will be loaded via viper
│   └── config-development.json
│   └── config-staging.json
│   └── config-production.json
├── api.go
├── app.go
├── bindata.go <-- this file is gitignored, it will appear at compile time
├── data
│   └── templates
│       └── react.html
├── main.go <-- main function declared here
├── react.go
└── utils.go
```

The `./server/` is flat golang package.

##### The client's entry point

It's simple React application

```
$ tree client
client
├── actions.js
├── components
│   ├── app
│   │   ├── favicon.ico
│   │   ├── index.js
│   │   └── styles.css
│   ├── homepage
│   │   ├── index.js
│   │   └── styles.css
│   ├── not-found
│   │   ├── index.js
│   │   └── styles.css
│   └── usage
│       ├── index.js
│       └── styles.css
├── css
│   ├── funcs.js
│   ├── global.css
│   ├── index.js
│   └── vars.js
├── index.js <-- main function declared here
├── reducers.js
├── sagas.js
├── router
│   ├── index.js
│   ├── routes.js
│   └── toString.js
└── store.js
```

The client app will be compiled into `server/data/static/build/`.  Then it will be embedded into go package via _go-bindata_. After that the package will be compiled into binary.

**Convention**: javascript app should declare [_main_](https://github.com/olebedev/go-starter-kit/blob/master/client/index.js#L4) function right in the global namespace. It will used to render the app at the server side.

## Install

Clone the repo:

```
$ https://github.com/ntquan1704/gin-react-starter-kit.git $GOPATH/src/<project>
$ cd $GOPATH/src/<project>
```

Install dependencies:

```
$ make install
```

## Run development

Start dev server:

```
$ make serve
```

that's it. Open [http://localhost:5001/](http://localhost:5001/)(if you use default port) at your browser. Now you ready to start coding your awesome project.

## Build

Install dependencies and type `ENV=production make build`. This rule is producing webpack build and regular golang build after that. Result you can find at `$GOPATH/bin`. Note that the binary will be named **as the current project directory**.

## RUN IN PRODUCTION
- cd $GOPATH/src/<project>
- Run: ENV=production $GOPATH/bin/<project> run

Open [http://localhost:5000/](http://localhost:5000/) 

## License
MIT
