package codec

import (
	"github.com/streamingfast/bstream"
	"github.com/streamingfast/logging"
)

var zlog, _ = logging.PackageLogger("codec", "github.com/Maar-io/firehose-astar/codec_test")

func init() {
	logging.InstantiateLoggers()

	bstream.GetBlockPayloadSetter = bstream.MemoryBlockPayloadSetter
}
