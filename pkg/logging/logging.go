package logging

import "log"

// Init ...
func Init() {
	log.SetFlags(log.Llongfile | log.LstdFlags)
}
