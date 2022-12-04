# Mini Wallet Microservice
![Github Action](https://github.com/sudtanj/Mini-Wallet-Microservice/actions/workflows/docker-publish.yml/badge.svg)
[![Docker](https://img.shields.io/docker/cloud/build/eaudeweb/scratch?label=Docker&style=flat)](https://hub.docker.com/r/sudtanj/mini-wallet-microservice/builds)

##### Warning: this is not production-ready!

For the exercise, use this documentation to build an API backend service, for managing a simple mini wallet.

Submit a working executable program that can be run in localhost that is based on this documentation as close as you can. The source code should be shared as a publicly accessible repo on either GitHub, GitLab, Bitbucket, etc. Please provide steps on how to "install" and "run" the program.
If working code is not possible, submit detailed technical design specs on the logic and framework that would allow any engineer to implement as with minimal time and supervision.
The idea is that this API is exposed by the wallet service for a wallet feature. Please assume that the customer verification was already done and information/profile was already stored in a separate customer service.

For authentication to this wallet service, pass it as a header called Authorization with the content in the format of Token <my token>.

The API is a HTTP REST API and the responses are returned in JSON with structure based on JSend.

All the best!

## System Requirement
- Go 1.19 or newer
- GCC (for sqlite)
- Docker (if you use the recommended docker or kubernetes usage)
- Kubernetes (if you use the kubernetes usage)

## Usage
### With Docker (Recommended)
- Run the following command will create instance mini wallet container with db store in memory (not persist). Change [YOUR_SELECTED_PORT] to port that's available on your local machine.
```
docker run -p [YOUR_SELECTED_PORT]:80 -e DATABASE_PATH='file::memory:?cache=shared' sudtanj/mini-wallet-microservice:latest
```
### Without Docker
- Clone this repository 
- Open your terminal or cmd and change to this project directory on your local machine
- Run the following command to initialize the dependencies and run the project on your local machine
```
go run .
```
- the database (sqllite)  will be automatically created and stored at the root of the project as `test.db` file after the project run successfully You can customize this behaviour in the env file
- The api by default will be running on http://localhost:80
### With Kubernetes
This project also provided some example if you want to run the mini wallet inside kubernetes cluster with [Kong](https://konghq.com/) as the api gateway. all that needed is to run the following command 
inside the project root folder
```
kubectl create namespace mini-wallet
kubectl apply -f kube.yaml
```
The `kube.yaml` file will use [Kong](https://konghq.com/) as the ingress and will create 2 pods to be served at http://localhost for client outside kubernetes cluster. 
The service will be served at http://mini-wallet-service:5000 inside kubernetes cluster.

## Unit Test
Unit testing can be verified by running command
```
go test ./...
```  
Note: not all function has yet unit tested yet! only utils.
