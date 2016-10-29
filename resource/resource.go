package resource

import (
	"github.com/labstack/echo"
)

type (
	Param struct {
		e    *echo.Echo
		base string
	}
)

func Regist(param Param) {
	router := Param.e.Router()

}

type indexter interface {
	Index() error
}
type creater interface {
	Create() error
}
type storer interface {
	Store() error
}
type shower interface {
	Show(interface{}) error
}
type editer interface {
	Edit(interface{}) error
}
type updater interface {
	Update(interface{}) error
}
type destroier interface {
	Destroy(interface{}) error
}
