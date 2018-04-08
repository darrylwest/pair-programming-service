//
// config  - application specification and CLI parsing
//
// @author darryl.west <darryl.west@ebay.com>
// @created 2018-04-08 09:26:59

package edit

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

// Config the config structure
type Config struct {
	Port         int
	LogLevel     int
	DbFilename   string
	StaticFolder string
	Timeout      int
}

// NewDefaultConfig default settings
func NewDefaultConfig() *Config {
	cfg := new(Config)

	cfg.Port = 3500
	cfg.LogLevel = 2
	cfg.DbFilename = "data/edit.db"
	cfg.StaticFolder = "public"
	cfg.Timeout = 120 // seconds

	return cfg
}

// ShowHelp dump out the use/command line options
func ShowHelp() {
	fmt.Printf("\n%s USE:\n\n", os.Args[0])
	flag.PrintDefaults()
	fmt.Printf("\n%s Version %s\n", os.Args[0], Version())
}

// ParseArgs parse the command line args
func ParseArgs() *Config {
	dflt := NewDefaultConfig()

	vers := flag.Bool("version", false, "show the version and exit")
	level := flag.Int("loglevel", dflt.LogLevel, "set the server's log level 0..5, default info=2")
	port := flag.Int("port", dflt.Port, "set the server's listening port")
	dbfilename := flag.String("db-filename", dflt.DbFilename, "set the databse file")
	static := flag.String("static", dflt.StaticFolder, "set the static html folder")
	timeout := flag.Int("timeout", dflt.Timeout, "the timeout for both tests and builds in seconds")

	flag.Parse()

	if *vers == true {
		return nil
	}

	fmt.Println(logo)
	fmt.Printf("Version %s\n", Version())

	log.Info("%s Version: %s\n", filepath.Base(os.Args[0]), Version())

	cfg := Config{
		Port:         *port,
		LogLevel:     *level,
		DbFilename:   *dbfilename,
		StaticFolder: *static,
		Timeout:      *timeout,
	}

	log.SetLevel(cfg.LogLevel)

	return &cfg
}
