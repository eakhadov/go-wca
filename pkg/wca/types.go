package wca

type ERole uint64

type EDataFlow uint64

type REFERENCE_TIME int64

type DataFlow uint64

const (
	In DataFlow = iota
	Out
)

type ConnectorType uint64

const (
	Unknown_Connector ConnectorType = iota
	Physical_Internal
	Physical_External
	Software_IO
	Software_Fixed
	Network
)

type PartType uint64

const (
	Connector PartType = iota
	Subunit
)
