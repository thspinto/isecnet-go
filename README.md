# ISECnet-go

This an go SDK implementation for Intelbras' ISECNet protocol.

## ISEcnet

ISECnet is a proprietary protocol used by Intelbras in their alarm central communication.

## Usage

### Available Commands

```
  help          Help about any command
  partialStatus Get partial central status
```

### Flags

```
      --config string     config file (default is $HOME/.isecnet-go.yaml)
  -h, --help              help for isecnet-go
      --host string       Host or IP of the central (default "localhost")
      --password string   Central password (default "1234")
      --port string       Central port (default "9009")
  -t, --toggle            Help message for toggle
```

### Example

```
go run . --host localhost --password 1234 partialStatus

INFO[0000] Conecting...                                  address="localhost:9009"
+---------+----------+-------+----------+
|  ZONE   | ANULATED | OPEN  | VIOLATED |
+---------+----------+-------+----------+
| Zone 1  | false    | true  | false    |
| Zone 2  | false    | false | false    |
| Zone 3  | false    | false | false    |
| Zone 4  | false    | false | false    |
| Zone 5  | false    | true  | false    |
| Zone 6  | false    | false | false    |
| Zone 7  | false    | false | false    |
| Zone 8  | false    | false | false    |
| Zone 9  | false    | false | false    |
(...)
```

## Testing

* `make unit-test`: unit tests

* `make lint`: linter

* `make test`: run all


## Contributing

## References

* https://github.com/jrbenito/isec-wireshark
* https://github.com/felipealmeida/amt2018



1000 1000
0100 0100
0000 0100
0000 0010
0001 0001
