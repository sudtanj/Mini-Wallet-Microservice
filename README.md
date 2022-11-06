# Mini Wallet Microservice
![Github Action](https://github.com/sudtanj/Mini-Wallet-Microservice/actions/workflows/docker-publish.yml/badge.svg)
[![Docker](https://img.shields.io/docker/cloud/build/eaudeweb/scratch?label=Docker&style=flat)](https://hub.docker.com/r/sudtanj/mini-wallet-microservice/builds)

Warning: this is not production-ready!

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