package components

import (
	"bytes"
	"encoding/json"
	"net/http"
	"syscall/js"
)

func RenderLoginForm(this js.Value, args []js.Value) interface{} {
	html := `
        <h2 class="mb-4">Login</h2>
        <form id="loginForm">
            <div class="mb-3">
                <label for="username" class="form-label">Username</label>
                <input type="text" class="form-control" id="username" required>
            </div>
            <div class="mb-3">
                <label for="password" class="form-label">Password</label>
                <input type="password" class="form-control" id="password" required>
            </div>
            <button type="submit" class="btn btn-primary">Login</button>
        </form>
    `

	js.Global().Get("document").Call("getElementById", "content").Set("innerHTML", html)

	js.Global().Get("document").Call("getElementById", "loginForm").Call("addEventListener", "submit", js.FuncOf(submitLoginForm))

	return nil
}

func submitLoginForm(this js.Value, args []js.Value) interface{} {
	args[0].Call("preventDefault")

	credentials := map[string]string{
		"username": js.Global().Get("document").Call("getElementById", "username").Get("value").String(),
		"password": js.Global().Get("document").Call("getElementById", "password").Get("value").String(),
	}

	credentialsJSON, _ := json.Marshal(credentials)

	go func() {
		resp, err := http.Post("/api/login", "application/json", bytes.NewBuffer(credentialsJSON))
		if err != nil {
			js.Global().Get("console").Call("error", err.Error())
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			js.Global().Get("alert").Invoke("Login successful")
			js.Global().Call("renderUserList")
		} else {
			js.Global().Get("alert").Invoke("Invalid credentials")
		}
	}()

	return nil
}
