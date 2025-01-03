1. Make sure have the following installed on your system:
Go (Version: >=1.16)
npm (or yarn)
PostgreSQL (Database)
Git

2. Backend Setup (http://localhost:8080)
navigate to be-project: cd be-project
install dependencies: go mod tidy
Ensure the .env file contains correct configuration values
go run main.go (there is seed data)

3. Frontend Setup (http://localhost:3000)
navigate to frontend-project: cd frontend-project
install dependencies: npm install
npm start

4. Login to the application with the following seeded data:
Username: testuser
Password: testpassword
