//
// service tests
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

func TestService(t *testing.T) {
	g := Goblin(t)

	g.Describe("Service", func() {
		log := edit.CreateLogger()
		log.SetLevel(4)
		cfg := edit.NewDefaultConfig()

		g.It("should create a service struct", func() {
			service, err := edit.NewService(cfg)
			g.Assert(err).Equal(nil)
			g.Assert(fmt.Sprintf("%T", service)).Equal("*edit.Service")
		})

        g.It("should configure the service routes")
	})
}
