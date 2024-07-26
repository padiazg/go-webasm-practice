package main

import (
	"syscall/js"

	"github.com/padiazg/wasm-app-test/frontend/components"
)

func main() {
	c := make(chan struct{}, 0)

	js.Global().Set("renderUserList", js.FuncOf(components.RenderUserList))
	js.Global().Set("renderUserForm", js.FuncOf(components.RenderUserForm))
	js.Global().Set("renderLoginForm", js.FuncOf(components.RenderLoginForm))
	js.Global().Set("renderUserUpdateForm", js.FuncOf(components.RenderUserUpdateForm))

	<-c
}
