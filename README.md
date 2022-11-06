# Mini Wallet Microservice
![Github Action](https://github.com/sudtanj/Mini-Wallet-Microservice/actions/workflows/docker-publish.yml/badge.svg)
[![Docker](https://img.shields.io/docker/cloud/build/eaudeweb/scratch?label=Docker&style=flat)](https://hub.docker.com/r/sudtanj/mini-wallet-microservice/builds)

## System Requirement
- Go 1.19 or newer
- GCC (for sqlite)

## Usage
### With Docker (Recommended)
### Without Docker
- Clone this repository 
- Open your terminal or cmd and change to this project directory on your local machine
- Run the following command to initialize the dependencies and run the project on your local machine
```
go run .
```
- the database (sqllite)  will be automatically created and stored at the root of the project as `test.db` file after the project run successfully
- The api by default will be running on http://localhost:80
