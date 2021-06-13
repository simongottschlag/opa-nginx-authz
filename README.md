# opa-nginx-authz
Lab to test opa authz for nginx

## Tests

```shell
curl --data '{"foo": "bar"}' --header "Content-Type: application/json" localhost:8080/echo/test
curl -v --data '{"foo": "bar"}' --header "Content-Type: application/json" localhost:8080/authopa/test
curl -v --data '{"foo": "bar"}' --header "Authorization: Bearer test" --header "Content-Type: application/json" localhost:8080/authopa/test
curl -v --data '{"foo": "bar"}' --header "Authorization: Bearer 2test" --header "Content-Type: application/json" localhost:8080/authopa/test
``` 