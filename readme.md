# WebAsm with Golang practice
I started this project in my journey to learn how to use WebAsm for Web UI with Golang

## Build
First step is to build the frontend binary which is compiled to wasm
```bash
$ cd frontend
$ GOOS=js GOARCH=wasm go build -o ../static/main.wasm
```

Then we can build the backend to serve the app
```bash
$ cd backend
$ go build -o ../bin/server
```

## Run
Now we can run app
```bash
./bin/server
```

## Features
The app will allow you to Add/Update/List users, and also let you "login" with them. There is no database, all data is stored in memory and will vanish when the app stops.

We use Bootstrap to style a bit the pages