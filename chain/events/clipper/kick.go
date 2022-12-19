package clipper

import (
	"github.com/IR-Digital-Token/x/chain/bindings/clipper"
	"github.com/IR-Digital-Token/x/chain/events"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

type kickHandler struct {
	binding  *clipper.Clipper
	callback events.CallbackFn[clipper.ClipperKick]
}

func NewKickHandler(clipperAddr common.Address, eth *ethclient.Client, callback events.CallbackFn[clipper.ClipperKick]) events.Handler {
	binding, err := clipper.NewClipper(clipperAddr, eth)
	if err != nil {
		panic(err)
	}
	return &kickHandler{binding: binding, callback: callback}
}

func (h *kickHandler) Signature() string {
	return "0x7c5bfdc0a5e8192f6cd4972f382cec69116862fb62e6abff8003874c58e064b8"
}

func (h *kickHandler) DecodeLog(log types.Log) (interface{}, error) {
	return h.binding.ParseKick(log)
}

func (h *kickHandler) HandleEvent(event interface{}) error {
	kick, ok := event.(clipper.ClipperKick)
	if !ok {
		return errors.New("event type is not clipper kick.")
	}
	return h.callback(kick)
}

func (h *kickHandler) DecodeAndHandle(log types.Log) error {
	kick, err := h.binding.ParseKick(log)
	if err != nil {
		return err
	}
	return h.callback(*kick)
}
