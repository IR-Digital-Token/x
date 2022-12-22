// Code generated - DO NOT EDIT.
// This file is a generated event handler and any manual changes will be lost.

package clipper

import (
	"errors"
	"github.com/IR-Digital-Token/x/chain/bindings/clipper"
	"github.com/IR-Digital-Token/x/chain/events"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type FileHandler struct {
	binding  *clipper.Clipper
	callback events.CallbackFn[clipper.ClipperFile]
}

func (h *FileHandler) Signature() string {
	return "0x8fef588b5fc1afbf5b2f06c1a435d513f208da2e6704c3d8f0e0ec91167066ba"
}

func (h *FileHandler) DecodeLog(log types.Log) (interface{}, error) {
	return h.binding.ParseFile(log)
}

func (h *FileHandler) HandleEvent(event interface{}) error {
	e, ok := event.(clipper.ClipperFile)
	if !ok {
		return errors.New("event type is not ClipperFile")
	}
	return h.callback(e)
}

func (h *FileHandler) DecodeAndHandle(log types.Log) error {
	e, err := h.binding.ParseFile(log)
	if err != nil {
		return err
	}
	return h.callback(*e)
}

func NewFileHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[clipper.ClipperFile]) events.Handler {
	b, err := clipper.NewClipper(addr, eth)
	if err != nil {
		panic(err)
	}
	return &FileHandler{
		binding:  b,
		callback: callback,
	}
}
