# Mini Wallet Microservice

## System Requirement
- Go 1.19 or newer
- GCC (for sqlite)

## Usage
### Without Docker
- Clone this repository 
- Open your terminal or cmd and change to this project directory on your local machine
- Run the following command to initialize the dependencies and run the project on your local machine
```
go run .
```
- the database (sqllite)  will be automatically created and stored at the root of the project as `test.db` file after the project run successfully
- The api by default will be running on http://localhost:80
