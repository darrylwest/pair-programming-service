//
// handlers tests
//
// @author darryl.west@ebay.com
// @created 2018-02-27 16:22:25
//

package unit

import (
	"edit"
	"fmt"
	"testing"

	. "github.com/franela/goblin"
)

func TestHandlers(t *testing.T) {
	g := Goblin(t)

	g.Describe("Handlers", func() {
		log := edit.CreateLogger()
		log.SetLevel(4)
		cfg := edit.NewDefaultConfig()

		g.It("should create a handlers struct", func() {
			handlers, err := edit.NewHandlers(cfg)
			g.Assert(err).Equal(nil)
			g.Assert(fmt.Sprintf("%T", handlers)).Equal("*edit.Handlers")
		})

        g.It("should initialize the web socket message commands", func() {
			handlers, _ := edit.NewHandlers(cfg)
            commands := handlers.InitCommands()
            g.Assert(len(commands) > 0).IsTrue()
        })

        g.It("should create a response wrapper", func() {
			handlers, _ := edit.NewHandlers(cfg)
            wrapper := handlers.CreateResponseWrapper("/mycommand")
            g.Assert(len(wrapper)).Equal(1)
            g.Assert(wrapper["request"]).Equal("/mycommand")
        })

        g.It("should handle a ping request", func() {
			handlers, _ := edit.NewHandlers(cfg)
            response, err := handlers.PingHandler("/ping")
			g.Assert(err).Equal(nil)
            g.Assert(response["request"]).Equal("/ping")
            g.Assert(response["response"]).Equal("pong")
        })

        g.It("should handle an update request", func() {
			handlers, _ := edit.NewHandlers(cfg)
            response, err := handlers.PingHandler("/ping")
			g.Assert(err).Equal(nil)
            g.Assert(response["request"]).Equal("/ping")
            g.Assert(response["response"]).Equal("pong")
        })
	})
}
