// simple verion
//
// @author darryl.west@ebay.com
// @created 2018-02-27 16:22:25

package edit

import "fmt"

const (
	major = 18
	minor = 4
	patch = 8
	logo  = `
 ___      _       ___                                    _             ___              _        
| _ \__ _(_)_ _  | _ \_ _ ___  __ _ _ _ __ _ _ __  _ __ (_)_ _  __ _  / __| ___ _ ___ _(_)__ ___ 
|  _/ _, | | '_| |  _/ '_/ _ \/ _, | '_/ _, | '  \| '  \| | ' \/ _, | \__ \/ -_) '_\ V / / _/ -_)
|_| \__,_|_|_|   |_| |_| \___/\__, |_| \__,_|_|_|_|_|_|_|_|_||_\__, | |___/\___|_|  \_/|_\__\___|
                              |___/                            |___/                             
`
)

// Version - return the version number as a single string
func Version() string {
	return fmt.Sprintf("%d.%d.%d", major, minor, patch)
}
