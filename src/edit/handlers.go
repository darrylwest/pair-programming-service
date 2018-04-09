//
// web socket message handlers
//
// @author darryl.west <darryl.west@ebay.com>
// @created 2018-04-08 09:26:59

package edit

import (
    "fmt"
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
    log.Info("web socket client request: %s", r.URL);
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Error("upgrade error: %s", err)
        return
    }

    defer func() {
        log.Info("client session closed...")
        conn.Close()
    }()

    for {
        _, raw, err := conn.ReadMessage()
        if err != nil {
            log.Warn("client socket down: %s", err);
            break
        }

        log.Info("raw request: %s", raw)
        request := string(raw)
        log.Info("%s", request)
    }
}

// CreateResponseWrapper creates the standard response wrapper for websocket messages
func (h Handlers) CreateResponseWrapper(request string) CommandResponse {
    wrapper := make(map[string]interface{})

    wrapper["request"] = request

    return wrapper
}

// PingHandler return the ping response
func (h Handlers) PingHandler(request string) (CommandResponse, error) {
    wrapper := h.CreateResponseWrapper(request)
    wrapper["response"] = "pong"

    return wrapper, nil
}

// AboutHandler returns the about page
func (h Handlers) AboutHandler(w http.ResponseWriter, r *http.Request) {
    log.Info("show the about page")

    // serve a template file...
    // w.Header().Set("Content-Type", "plain/html")
    fmt.Fprintf(w, "<!doctype html><html><pre>%s</pre><h5>Version %s</h5><html>", logo, Version())
}

// ConfigHandler returns the configuration dialog
func (h Handlers) ConfigHandler (w http.ResponseWriter, r *http.Request) {
    log.Info("show the configuration page")

    // serve a template file...
    // w.Header().Set("Content-Type", "plain/html")
    fmt.Fprintf(w, "<!doctype html><html><pre>%s</pre><h5>Version %s</h5><h4>Configuration</h4><html>", logo, Version())
}
