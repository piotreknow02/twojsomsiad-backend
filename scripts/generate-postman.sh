cd ..

# load sample env configuration
export $(cat .env.example | xargs)

go run main.go --postman