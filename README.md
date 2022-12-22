# X

x is a repository that contains common golang packages.

Generate `events.Handler` from abi file.
```shell
go run codegen/events_handler.go --abi=chain/bindings/clipper/abi.json --binding-package=github.com/IR-Digital-Token/x/chain/bindings/clipper --output-dir=chain/events/clipper --contract=Clipper
go run codegen/events_handler.go --abi=chain/bindings/vat/abi.json  --binding-package=github.com/IR-Digital-Token/x/chain/bindings/vat --output-dir=chain/events/vat --contract=Vat
```