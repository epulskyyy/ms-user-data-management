package application

import (
	"github.com/subosito/gotenv"
	"ms-user-data-management/router"
)

func init()  {
	gotenv.Load()
}

func StartApp()  {
	router.NewRoute()
}
