# nutshttp

A http server for nutsdb

## Example

Run http server and listen ON ":8080".

```bash
go run examples/hello.go
```

**Check example data**

```bash
# Get all members in set
curl http://localhost:8080/set/bucket001/foo


# List all list
curl http://localhost:8080/list/bucket001/key1?start=0&end=10
```


## Auth

1. Enable Auth:
```
nutshttp.EnableAuth = true
```
2. Create Token:
```bash
curl http://127.0.0.1:8080/auth/thisisacert
```
`thisisacert` replace with your pwd or username

3. Use Token:

- Add token to HEADERS:
    > Authorization : Bearer `<token>`
