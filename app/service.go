package app

import (
	"banhmi/root"
)

type Service struct {
	Httpsv root.HTTPService
	MySQLsv root.MySQLService
}