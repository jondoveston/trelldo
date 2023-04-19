package main

import (
	"github.com/spf13/viper"
	"github.com/jondoveston/trelldo/cmd"
)

var (
	version = "0.0.1"
	commit  = "HEAD"
	date    = "now"
	builtBy = "jondoveston"
)

func main() {
	cmd.Version = version
	cmd.Commit = commit
  cmd.Date = date
  cmd.BuiltBy = builtBy

	_ = cmd.Execute()
	_ = viper.WriteConfig()
}
