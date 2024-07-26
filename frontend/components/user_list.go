package components

import (
	"encoding/json"
	"fmt"
	h "html"
	"net/http"
	"syscall/js"
)

func RenderUserList(this js.Value, args []js.Value) interface{} {
	go func() {
		resp, err := http.Get("/api/users")
		if err != nil {
			js.Global().Get("console").Call("error", err.Error())
			return
		}
		defer resp.Body.Close()

		var users []map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&users)

		html := `
        <h2 class="mb-4">User List</h2>
        <table class="table table-striped table-hover">
            <thead>
                <tr>
                    <th>Username</th>
                    <th>Name</th>
                    <th>Email</th>
                    <th>Active</th>
                </tr>
            </thead>
            <tbody>
        `
		for _, user := range users {
			html += fmt.Sprintf(`
                <tr class="user-row" data-id="%s" style="cursor: pointer;">
                    <td>%s</td>
                    <td>%s</td>
                    <td>%s</td>
                    <td>%s</td>
                </tr>
            `,
				user["id"],
				h.EscapeString(user["username"].(string)),
				h.EscapeString(user["name"].(string)),
				h.EscapeString(user["email"].(string)),
				(map[bool]string{true: "Yes", false: "No"})[user["active"].(bool)])
		}
		html += "</tbody></table>"

		js.Global().Get("document").Call("getElementById", "content").Set("innerHTML", html)

		// Add click event listeners to user rows
		userRows := js.Global().Get("document").Call("querySelectorAll", ".user-row")
		for i := 0; i < userRows.Length(); i++ {
			userRows.Index(i).Call("addEventListener", "click", js.FuncOf(handleUserRowClick))
		}
	}()

	return nil
}

func handleUserRowClick(this js.Value, args []js.Value) interface{} {
	userId := this.Get("dataset").Get("id").String()
	js.Global().Call("renderUserUpdateForm", userId)
	return nil
}
