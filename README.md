[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fthspinto%2Fisecnet-go.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fthspinto%2Fisecnet-go?ref=badge_shield)

This an go SDK implementation for Intelbras' ISECNet protocol.

# ISEcnet

ISECnet is a proprietary protocol used by Intelbras in their alarm central communication.

# Usage

## Available Commands

```
  help          Help about any command
  partialStatus Get partial central status
  zones         Get Zone status
```

## Example

![Example of the zones and zones -w command running in the terminal](./img/zones.gif)

# Zones description

You can configure the zone description in `.isecnet-go.yaml` to see meaningful names and only zones in use. See `.isecnet-go..yaml.example` for example:

```yaml
zones:
  - id: 1
    name: Front Door
    description: Front dor magnetic sensor
```

```bash
➜  isecnet-go git:(main) ✗ go run . --password 1234 zones
Using config file: ./.isecnet-go.yaml
INFO[0000] Connecting...                                 address="localhost:9009"
+------------+----------+-------+----------+------------+--------+---------------+
|    ZONE    | ANULATED | OPEN  | VIOLATED | LOWBATTERY | TAMPER | SHORT CIRCUIT |
+------------+----------+-------+----------+------------+--------+---------------+
| Front Door | false    | false | false    | false      | false  | false         |
+------------+----------+-------+----------+------------+--------+---------------+
```

# Testing

* `make unit-test`: unit tests

* `make lint`: linter

* `make test`: run all

* `make mock-alarm-central`: starts the mock server

* `go run . server`: starts the gRPC server:

```
// Testing the API
grpcurl -plaintext localhost:8080 ZoneService.GetZones
```

# References

* https://github.com/jrbenito/isec-wireshark
* https://github.com/felipealmeida/amt2018

# License
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fthspinto%2Fisecnet-go.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fthspinto%2Fisecnet-go?ref=badge_large)
