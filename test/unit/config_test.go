//
// config tests
//
// @author darryl.west <darwest@ebay.com>
// @created 2018-04-08 09:26:59
//

package unit

import (
	"edit"
	"fmt"
	"testing"

	. "github.com/franela/goblin"
)

func TestConfig(t *testing.T) {
	g := Goblin(t)

	g.Describe("Config", func() {
		edit.CreateLogger()

		g.It("should create a context struct with defaults set", func() {
			cfg := edit.NewDefaultConfig()
			g.Assert(fmt.Sprintf("%T", cfg)).Equal("*edit.Config")
			g.Assert(cfg.Port).Equal(3500)
			g.Assert(cfg.LogLevel > 1).IsTrue()
			g.Assert(cfg.DbFilename).Equal("data/edit.db")
			g.Assert(cfg.StaticFolder).Equal("public")
		})

		g.It("should parse an empty command line and return default config", func() {
			cfg := edit.ParseArgs()
			g.Assert(cfg != nil).IsTrue()
		})
	})
}
