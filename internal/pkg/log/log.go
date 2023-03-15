package log

// this package target to init the log path, but not should use it directly.

import (
	"github.com/rootlulu/go-gin-biu-biu-biu/internal/config"
	"github.com/rootlulu/go-gin-biu-biu-biu/pkg/logging"
)

func Init() {
	logging.LogFile(config.App.LogPath)
}
