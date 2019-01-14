# REST API

This is an example of how to serve REST behind the API using go-restful.
Make sure your service name is in the format of `<kubernetes-deployment-name>.<namespace-name>`

## Getting Started

```
consul agent -dev
```

### Run Greeter Service

```shell
make telepresence-local-dev NAMESPACE="irvi-test" DEPLOYMENT="go-micro-greeter" PORT="8080"
go run greeter/srv/main.go
```

### Run Greeter API

```shell
go run api/rest/rest.go --server_address=0.0.0.0:8080
```

### Curl API

```shell
curl http://localhost:8080/greeter
```

Output

```json
{
  "message": "Hello, Stranger!"
}
```

Test a resource

```shell
 curl http://localhost:8080/greeter/John
```

Output
```json
{
  "msg": "Hello, John!"
}
```
