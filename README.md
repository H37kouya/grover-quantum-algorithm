# Grover quantum search

## 簡易起動

```shell script
$ go run main.go nqubit-csv -n 4

$ go run main.go nqubit-times -n 4

$ go run main.go nqubit-times-all  -s 4 -e 25

$ go run main.go nqubit-random-csv -n 4 -r 0 -i 0

$ go run main.go nqubit-random-times -n 4 -c 1 -r 0 -i 0
```

## build系

```shell script
$ go build -o bin/grover ./

$ ./bin/grover nqubit-csv -n 4

$ ./bin/grover nqubit-times -n 4

# 標準入出力の結果をファイルへ送る
$ ./bin/grover nqubit-times-all -min 4 -max 25 > ./outputs/nqubit-times-all.csv

$ ./bin/grover nqubit-random-csv -n 4 -r 0 -i 0

$ ./bin/grover nqubit-random-times -n 4 -c 1 -r 0 -i 0
```
