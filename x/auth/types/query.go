package types

import (
	codectypes "github.com/GincoInc/cosmos-sdk/codec/types"
	sdk "github.com/GincoInc/cosmos-sdk/types"
)

func (m *QueryAccountResponse) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	var account sdk.AccountI
	return unpacker.UnpackAny(m.Account, &account)
}

var _ codectypes.UnpackInterfacesMessage = &QueryAccountResponse{}
