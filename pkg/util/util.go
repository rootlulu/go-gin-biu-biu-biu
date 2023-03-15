package util

import (
	"crypto/md5"
	"encoding/hex"

	// "github.com/rootlulu/go-gin-biu-biu-biu/pkg/logging"
	// todo, the cycle importing.
	"log"

	"github.com/go-ini/ini"
)

func Init() {}

func IniToStruct(cfg *ini.File, section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Parsing the config failed: %v", err)
	}
}

// EncodeMD5 md5 encryption
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}
