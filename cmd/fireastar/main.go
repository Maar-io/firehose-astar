package main

import (
	"github.com/Maar-io/firehose-astar/codec"
	pbastar "github.com/Maar-io/firehose-astar/pb/sf/astar/type/v1"
	firecore "github.com/streamingfast/firehose-core"
)

func main() {
	firecore.Main(&firecore.Chain[*pbastar.Block]{
		ShortName:            "astar",
		LongName:             "Astar",
		ExecutableName:       "dummy-blockchain",
		FullyQualifiedModule: "github.com/Maar-io/firehose-astar",
		Version:              version,

		Protocol:        "ACM",
		ProtocolVersion: 1,

		FirstStreamableBlock: 1,

		BlockFactory:         func() firecore.Block { return new(pbastar.Block) },
		ConsoleReaderFactory: codec.NewConsoleReader,

		Tools: &firecore.ToolsConfig[*pbastar.Block]{
			BlockPrinter: printBlock,
		},
	})
}

// Version value, injected via go build `ldflags` at build time, **must** not be removed or inlined
var version = "dev"
