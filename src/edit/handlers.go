//
// web socket message handlers
//
// @author darryl.west <darryl.west@ebay.com>
// @created 2018-04-08 09:26:59

package edit

import (
    "fmt"
    "net/http"
    "strings"
    "time"

    "github.com/gorilla/websocket"
)

// CommandResponse the response from a web socket message
type CommandResponse interface{}

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
    commands["update"] = h.UpdateHandler
    commands["build"] = h.PingHandler
    commands["test"] = h.PingHandler

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

    if len(commands) == 0 {
        h.InitCommands()
    }

    for {
        _, raw, err := conn.ReadMessage()
        if err != nil {
            log.Warn("client socket down: %s", err);
            break
        }

        log.Info("raw request: %s", raw)
        request := string(raw)
        log.Info("%s", request)

        wrapper := h.CreateResponseWrapper(request)
        response, err := h.RequestHandler(request)
        if err != nil {
            wrapper["status"] = "failed"
            wrapper["reason"] = err.Error()
            if response != nil {
                wrapper["response"] = response
            }
        } else {
			wrapper["status"] = "ok"
			wrapper["response"] = response
        }

		wrapper["ts"] = time.Now().Format(time.RFC3339)
		err = conn.WriteJSON(wrapper)
		if err != nil {
			log.Warn("socket error: %s", err)
			break
		}
    }
}

// RequestHandler parses the request and returns a response or error if request is bad
func (h Handlers) RequestHandler(request string) (CommandResponse, error) {
	if len(commands) == 0 {
		h.InitCommands()
	}

	req := strings.Split(request, "/")[1]

	op, ok := commands[req]
	if !ok {
		log.Warn("unrecognized websocket command: %s", request)
		return nil, fmt.Errorf("not a recognized command: %s", request)
	}

	return op(request)
}

// CreateResponseWrapper creates the standard response wrapper for websocket messages
func (h Handlers) CreateResponseWrapper(request string) map[string]interface{} {
    wrapper := make(map[string]interface{})
    wrapper["request"] = request
    wrapper["response"] = ""

    return wrapper
}

// PingHandler return the ping response
func (h Handlers) PingHandler(request string) (CommandResponse, error) {
    return "pong", nil
}

// UpdateHandler broadcast the document delta to all clients
func (h Handlers) UpdateHandler(request string) (CommandResponse, error) {
    return "ok", nil
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
