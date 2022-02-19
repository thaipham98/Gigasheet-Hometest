# Gigasheet-Hometest

## Usage

### 1. Clone this Github repo or download the zip file at https://github.com/thaipham98/Gigasheet-Hometest.
### 2. Open terminal of choice and cd into the Gigasheet-Hometest directory.

## Run the program

```bash
go run server.go
```

## Test
### 1. Open another terminal and cd into Gigasheet-Hometest directory again.

### 2. Run with curl 
```bash
curl --data '{"ips":["94.142.241.194","192.168.1.1","159.65.180.64"]}' -H "Content-type: application/json" http://localhost:8080/count_ips_in_ipsum
```
#### The output printed out should be:
    
```bash
2
```
```bash
curl --data '{"ips":["192.168.1.1","159.65.180.64"]}' -H "Content-type: application/json" http://localhost:8080/count_ips_in_ipsum
```
#### The output printed out should be:
    
```bash
1
```

### 3. Run the test file
```bash
go test -v
```
