build-run-html5:
	GOOS=js GOARCH=wasm go build -o export/html5/app.wasm app.go
	cp -r media export/html5/media
	light-server -s export/html5 -p 8001
