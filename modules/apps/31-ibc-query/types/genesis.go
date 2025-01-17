package types

// NewGenesisState creates a 31-ibc-query GenesisState instance.
func NewGenesisState(queries []*CrossChainQuery, results []*CrossChainQueryResult) *GenesisState {
	return &GenesisState{
		Queries: queries,
		Results: results,
		PortId:  PortID,
	}
}

// DefaultGenesisState returns a default instance of the 31-ibc-query GenesisState.
func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		Queries: []*CrossChainQuery{},
		Results: []*CrossChainQueryResult{},
		PortId:  PortID,
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	return nil
}
