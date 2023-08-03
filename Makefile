# Makefile

# Компилируем код для Linux
build_linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/server-linux main.go

# Компилируем код для macOS
build_macos:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/server-macos main.go

# Компилируем код для Windows
build_windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/server-windows.exe main.go

# Запускаем сервер локально
run:
	go run main.go

# Очистка временных файлов и бинарных исполняемых файлов
clean:
	rm -rf bin
