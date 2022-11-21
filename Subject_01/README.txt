# Run the Go code without creating a binary executale
go run hello-world.go

# Compile the Go code and create a binary executable file
go build hello-world.go

# Compile the Go code and create a binary executable file for given OS
env GOOS=linux GOARCH=amd64 go build -a -o hello hello-world.go