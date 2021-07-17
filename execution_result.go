package flow

type ExecutionResult struct {
	PreviousResultID Identifier // commit of the previous ER
	BlockID          Identifier // commit of the current block
	Chunks           []*Chunk
	ServiceEvents    []*ServiceEvent
}

type Chunk struct {
	CollectionIndex      uint
	StartState           StateCommitment // start state when starting executing this chunk
	EventCollection      Identifier      // Events generated by executing results
	BlockID              Identifier      // Block id of the execution result this chunk belongs to
	TotalComputationUsed uint64          // total amount of computation used by running all txs in this chunk
	NumberOfTransactions uint16          // number of transactions inside the collection
	Index                uint64          // chunk index inside the ER (starts from zero)
	EndState             StateCommitment // EndState inferred from next chunk or from the ER
}

type ServiceEvent struct {
	Type    string
	Payload []byte
}