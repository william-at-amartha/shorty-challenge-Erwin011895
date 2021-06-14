# shorty-challenge

### Config
```
cp conf/app.ini.example conf/app.ini
``` 


## Build
### Go Module
```
go mod tidy
go build
```

## Run

```
./shorty-challenge
```


## Testing
### Create Mock

see `script/generate_mock.sh`

### Unit Tests

#### generate output file

```
go test ./... -coverprofile cp.out
```

#### generate coverage.html
```
go tool cover -html=cp.out -o coverage.html
```
