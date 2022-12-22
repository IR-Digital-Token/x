# X

x is a repository that contains common golang packages.

Generate `events.Handler` from abi file.
```shell
go run codegen/events_handler.go --abi=chain/bindings/clipper/abi.json  --output-dir=chain/bindings/clipper --contract=Clipper
go run codegen/events_handler.go --abi=chain/bindings/vat/abi.json  --output-dir=chain/bindings/vat --contract=Vat
```