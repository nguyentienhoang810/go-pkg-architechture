package root

import (
	"github.com/gorilla/mux"
)

type AppService interface {
	//SetRoutes ...
	SetRouters()
}


type HTTPService interface {
	// Router ..
	MuxRouter() *mux.Router
}

type MySQLService interface {
	
}

type UserService interface {
	All() ([]*User, error)
	Get(id int) (*User, error)
	Create(u *User) error
	Delete(id int) error
}

type ProductService interface {
	All() ([]*Product, error)
	Get(id int) (*Product, error)
}