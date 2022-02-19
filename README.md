# Gigasheet-Hometest

## Usage

1. Clone this Github repo or download the zip file at https://github.com/thaipham98/Gigasheet-Hometest.
2. Open terminal of choice and cd into the Gigasheet-Hometest directory.

## Run the program

```bash
go run server.go
```

## Test
1. Open another terminal

2. Run with curl in terminal

```bash
curl --data '{"ips":["94.142.241.194","192.168.1.1","159.65.180.64"]}' -H "Content-type: application/json" http://localhost:8080/count_ips_in_ipsum
```

2 should be printed out

3. Run the test file

```bash
go test -v
```
