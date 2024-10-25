# WebAsm with Golang Practice

A simple project demonstrating WebAssembly (WASM) for Web UI using Golang.

## Features

- Add/Update/List users
- Basic user authentication
- In-memory data storage (non-persistent)
- Bootstrap-styled UI

## Prerequisites

- Go 1.16+
- Make (optional, for using Makefile commands)

## Building

1. Build the frontend WASM binary:

```bash
cd frontend && GOOS=js GOARCH=wasm go build -o ../static/main.wasm
```

2. Build the backend server:

```bash
cd backend && go build -o ../bin/server
```

## Running

Start the application:

```bash
./bin/server
```

The app will be available at `http://localhost:8080` by default.

## Usage

1. Open the application in your browser
2. Create a new user using the "Create User" form
3. Log in with the created user credentials
4. Explore the user management features

## Development

To make changes:

1. Modify frontend code in `frontend/`
2. Rebuild the WASM binary
3. Modify backend code in `backend/`
4. Rebuild the server
5. Restart the application

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Buy me a coffee
[!["Buy Me A Coffee"](https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png)](https://buymeacoffee.com/padiazgy)


## License

This project is licensed under the GNU General Public License v3.0 - see the [LICENSE](LICENSE) file for details.