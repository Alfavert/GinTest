package internal

import (
	"io/ioutil"

	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
)

type Configs struct {
	Filename      string `json:"filename" binding:"required"`
	BasicLogin    string `json:"basiclogin" binding:"required"`
	BasicPassword string `json:"basicpassword" binding:"required"`
	Cartname      string `json:"cartname"`
}

var Config Configs

func (conf *Configs) ParseConfigs() {
	bytes, err := ioutil.ReadFile("configs/config.json")
	if err != nil {
		logrus.Panicf("cannot read config file, err %v", err)
	}
	err = jsoniter.Unmarshal(bytes, conf)
	if err != nil {
		logrus.Panicf("cannot unmarshal config file, err %v", err)
	}

}
