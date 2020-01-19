package main


import (
	"banhmi/app"
	"banhmi/http"
	"banhmi/mysql"
)

var (
	appsv = &app.Service{}
	httpsv = &http.Service{}
	myqslsv = &mysql.Service{}
)

type service struct {
	App *app.Service
	HTTP *http.Service
	MySQL *mysql.Service
}

func initServices() *service {

	appsv.Httpsv = httpsv
	appsv.Mysqlsv = myqslsv
	httpsv.Root = appsv
	httpsv.NewRouter()
	httpsv.SetRouters()
	
	return &service {
		App: appsv,
		HTTP: httpsv,
		MySQL: myqslsv,
	}
}