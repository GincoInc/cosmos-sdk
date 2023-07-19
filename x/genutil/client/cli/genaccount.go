package cli

const (
	flagVestingStart = "vesting-start-time"
	flagVestingEnd   = "vesting-end-time"
	flagVestingAmt   = "vesting-amount"
	flagAppendMode   = "append"
	flagModuleName   = "module-name"
)

// AddGenesisAccountCmd returns add-genesis-account cobra Command.
// This command is provided as a default, applications are expected to provide their own command if custom genesis accounts are needed.
