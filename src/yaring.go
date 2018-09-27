package main

import (
	"flag"
	"fmt"
	"jwt"
)

var (
	tomlConf = flag.String("config", "/tmp/musickit.toml", "TOML format config file.")
	logFile  = flag.String("log", "/tmp/log_musickit.log", "The log file.")
)

func main() {

	flag.Parse()

	if *tomlConf == "" {
		fmt.Println("The header and playload configure file must be specified")
		return
	}

	jwtToken, err := jwt.Token(*tomlConf)
	if nil != err {
		fmt.Println("Create Apple Music Developer Token Failure")
	}

	fmt.Println(jwtToken)

}
