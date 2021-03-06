package oak

import (
	"syscall/js"

	"github.com/elliotforbes/oak/http"
	"github.com/elliotforbes/oak/utils"
)

func registerCallbacks() {
	// packages
	utils.RegisterCallbacks()
	http.RegisterCallbacks()
}

func RegisterFunction(funcName string, myfunc func(i []js.Value)) {
	js.Global().Set(funcName, js.NewCallback(myfunc))
}

func Start() {
	println("Oak Framework Initialized")
	registerCallbacks()
}
