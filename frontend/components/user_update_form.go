package components

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"syscall/js"
)

func RenderUserUpdateForm(this js.Value, args []js.Value) interface{} {
	if len(args) < 1 {
		js.Global().Get("console").Call("error", "User ID is required")
		return nil
	}

	userId := args[0].String()

	go func() {
		resp, err := http.Get("/api/users/" + userId)
		if err != nil {
			js.Global().Get("console").Call("error", err.Error())
			return
		}
		defer resp.Body.Close()

		var user map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&user)

		html := fmt.Sprintf(`
            <h2 class="mb-4">Update User</h2>
            <form id="userUpdateForm">
                <input type="hidden" id="userId" value="%s">
                <div class="mb-3">
                    <label for="username" class="form-label">Username</label>
                    <input type="text" class="form-control" id="username" value="%s" required>
                </div>
                <div class="mb-3">
                    <label for="name" class="form-label">Name</label>
                    <input type="text" class="form-control" id="name" value="%s" required>
                </div>
                <div class="mb-3">
                    <label for="email" class="form-label">Email</label>
                    <input type="email" class="form-control" id="email" value="%s" required>
                </div>
                <div class="mb-3">
                    <div class="form-check">
                        <input class="form-check-input" type="checkbox" id="active" %s>
                        <label class="form-check-label" for="active">Active</label>
                    </div>
                </div>
                <button type="submit" class="btn btn-primary">Update User</button>
            </form>
        `,
			user["id"],
			user["username"],
			user["name"],
			user["email"],
			map[bool]string{true: "checked", false: ""}[user["active"].(bool)])

		js.Global().Get("document").Call("getElementById", "content").Set("innerHTML", html)

		js.Global().Get("document").Call("getElementById", "userUpdateForm").Call("addEventListener", "submit", js.FuncOf(submitUserUpdateForm))
	}()

	return nil
}

func submitUserUpdateForm(this js.Value, args []js.Value) interface{} {
	args[0].Call("preventDefault")

	userId := js.Global().Get("document").Call("getElementById", "userId").Get("value").String()
	user := map[string]interface{}{
		"username": js.Global().Get("document").Call("getElementById", "username").Get("value").String(),
		"name":     js.Global().Get("document").Call("getElementById", "name").Get("value").String(),
		"email":    js.Global().Get("document").Call("getElementById", "email").Get("value").String(),
		"active":   js.Global().Get("document").Call("getElementById", "active").Get("checked").Bool(),
	}

	userJSON, _ := json.Marshal(user)

	go func() {
		client := &http.Client{}
		req, err := http.NewRequest("PUT", "/api/users/"+userId, bytes.NewBuffer(userJSON))
		if err != nil {
			js.Global().Get("console").Call("error", err.Error())
			return
		}
		req.Header.Set("Content-Type", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			js.Global().Get("console").Call("error", err.Error())
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			js.Global().Get("alert").Invoke("User updated successfully")
			js.Global().Call("renderUserList")
		} else {
			js.Global().Get("alert").Invoke("Error updating user")
		}
	}()

	return nil
}
