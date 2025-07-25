package types

import (
	types "github.com/GincoInc/cosmos-sdk/codec/types"
	sdk "github.com/GincoInc/cosmos-sdk/types"
	"github.com/GincoInc/cosmos-sdk/types/msgservice"
)

// RegisterInterfaces registers the interfaces types with the interface registry.
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAuthorizeCircuitBreaker{},
		&MsgResetCircuitBreaker{},
		&MsgTripCircuitBreaker{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
