package command

import (
	"testing"
	"time"

	"github.com/afreakk/i3statusbear/internal/config"
	"github.com/afreakk/i3statusbear/internal/protocol"
)

func TestCommand(t *testing.T) {
	cfg := config.Config{}
	cfg.MinimumRenderInterval = "1us"
	output := protocol.Output{}
	output.Init(cfg)
	module := config.Module{}
	module.CommandName = "echo"
	module.CommandArgs = []string{"hi"}
	module.Sprintf = "%s"
	commandFunc := Command(&output, module)
	commandFunc()
	time.Sleep(time.Second)
	result := output.Messages[0].FullText
	if result != "hi" {
		t.Errorf("expecting hi, got %s", result)
	}

}

func BenchmarkCommand(b *testing.B) {
	cfg := config.Config{}
	cfg.MinimumRenderInterval = "1us"
	output := protocol.Output{}
	output.Init(cfg)
	module := config.Module{}
	module.CommandName = "echo"
	module.CommandArgs = []string{"hi"}
	module.Sprintf = "%s"
	commandFunc := Command(&output, module)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		commandFunc()
	}
}
