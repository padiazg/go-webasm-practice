package components

import (
	"bytes"
	"encoding/json"
	"net/http"
	"syscall/js"
)

func RenderUserForm(this js.Value, args []js.Value) interface{} {
	html := `
        <h2 class="mb-4">Create User</h2>
        <form id="userForm">
            <div class="mb-3">
                <label for="username" class="form-label">Username</label>
                <input type="text" class="form-control" id="username" required>
            </div>
            <div class="mb-3">
                <label for="name" class="form-label">Name</label>
                <input type="text" class="form-control" id="name" required>
            </div>
            <div class="mb-3">
                <label for="email" class="form-label">Email</label>
                <input type="email" class="form-control" id="email" required>
            </div>
            <div class="mb-3">
                <label for="password" class="form-label">Password</label>
                <input type="password" class="form-control" id="password" required>
            </div>
            <button type="submit" class="btn btn-primary">Create User</button>
        </form>
    `

	js.Global().Get("document").Call("getElementById", "content").Set("innerHTML", html)

	js.Global().Get("document").Call("getElementById", "userForm").Call("addEventListener", "submit", js.FuncOf(submitUserForm))

	return nil
}

func submitUserForm(this js.Value, args []js.Value) interface{} {
	args[0].Call("preventDefault")

	user := map[string]interface{}{
		"username": js.Global().Get("document").Call("getElementById", "username").Get("value").String(),
		"name":     js.Global().Get("document").Call("getElementById", "name").Get("value").String(),
		"email":    js.Global().Get("document").Call("getElementById", "email").Get("value").String(),
		"password": js.Global().Get("document").Call("getElementById", "password").Get("value").String(),
		"active":   true,
	}

	userJSON, _ := json.Marshal(user)

	go func() {
		resp, err := http.Post("/api/users", "application/json", bytes.NewBuffer(userJSON))
		if err != nil {
			js.Global().Get("console").Call("error", err.Error())
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusCreated {
			js.Global().Get("alert").Invoke("User created successfully")
			js.Global().Call("renderUserList")
		} else {
			js.Global().Get("alert").Invoke("Error creating user")
		}
	}()

	return nil
}
