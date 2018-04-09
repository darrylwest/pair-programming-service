//
// web socket message handlers
//
// @author darryl.west <darryl.west@ebay.com>
// @created 2018-04-08 09:26:59

package edit

import (
    "net/http"
    "github.com/gorilla/websocket"
)

// CommandResponse the response from a web socket message
type CommandResponse map[string]interface{}

// Commander a command type def
type Commander func(string) (CommandResponse, error)

// Handlers the web socket handlers struct
type Handlers struct {
    cfg *Config
}

var (
    upgrader = websocket.Upgrader{}
    commands = make(map[string]Commander)
)

// NewHandlers create a new web socket handlers struct
func NewHandlers(cfg *Config) (*Handlers, error) {
    handlers := Handlers{
        cfg: cfg,
    }

    return &handlers, nil
}

// InitCommands initialize all the web socket message commands
func (h Handlers) InitCommands() map[string]Commander {
    log.Info("initialize web socket commands")
    commands["ping"] = h.PingHandler

    return commands
}

// ClientHandler upgrade the connection for websockets
func (h Handlers) ClientHandler(w http.ResponseWriter, r *http.Request) {

}

// PingHandler return the ping response
func (h Handlers) PingHandler(request string) (CommandResponse, error) {
    response := make(map[string]interface{})
    response["pong"] = true

    return response, nil
}

