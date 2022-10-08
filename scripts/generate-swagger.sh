cd ..

if $(go env GOPATH)/bin/swag init; then
    echo "Swagger files were generated"
else
    echo "Error! Make sure swagger module is installed."
    echo "Run 'go install github.com/swaggo/swag/cmd/swag@latest' to install it."
    exit 1
fi