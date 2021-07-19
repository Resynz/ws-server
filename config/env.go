/**
 * @Author: Resynz
 * @Date: 2021/7/19 14:04
 */
package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type EnvString string

const (
	EnvDefault EnvString = "default"
	EnvDev     EnvString = "dev"
	EnvAudit   EnvString = "audit"
	EnvProd    EnvString = "prod"
)

var (
	Conf      Config
	Env       = EnvDefault
	ClientMap = new(SafeClientMap)
)

func init() {
	env := os.Getenv("ENV")
	if env == "dev" {
		Env = EnvDev
	}
	if env == "audit" {
		Env = EnvAudit
	}
	if env == "prod" {
		Env = EnvProd
	}
	cp := fmt.Sprintf("./configs/%s.json", Env)
	c, err := ioutil.ReadFile(cp)
	if err != nil {
		log.Fatalf("init env config failed! error:%s\n", err.Error())
	}
	if err = json.Unmarshal(c, &Conf); err != nil {
		log.Fatalf("init env config failed! error:%s\n", err.Error())
	}
	ClientMap.clientMap = make(map[string][]*Client)
}
