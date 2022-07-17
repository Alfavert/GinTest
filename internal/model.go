package internal

import (
	"fmt"
	"io/ioutil"
	"os"

	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
)

type Data struct {
	Id     int64  `json:"id" binding:"required,gt=0,lte=3"`
	Title  string `json:"title" binding:"required,gt=0,lte=100"`
	Author string `json:"author" binding:"required,gt=0,lte=50"`
	Price  int    `json:"price" binding:"required, gt=0, lte=10"`
}

type DataList struct {
	List []Data `json:"data"`
}

var Dlist DataList

type Cart struct {
	Id int64 `json:"id" binding:"lte3"`
}
type CartList struct {
	List []Cart `json:"list"`
}

var Clist CartList

func Init() {
	bytes, err := ioutil.ReadFile(Config.Filename)
	if err != nil {
		logrus.Panicf("cannot read data file, err %v", err)
	}
	err = jsoniter.Unmarshal(bytes, &Dlist)
	if err != nil {
		logrus.Panicf("cannot unmarshall file, error: %v", err)
	}
	return
}

func UpdateFile(Dlist *DataList) error {
	bytes, err := jsoniter.Marshal(&Dlist)
	if err != nil {
		return fmt.Errorf("cannot marshall dList, error: %v", err)
	}
	file, err := os.OpenFile(Config.Filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		return fmt.Errorf("cannot open file, error: %v", err)
	}
	_, err = file.Write(bytes)
	if err != nil {
		return fmt.Errorf("cannot write into file, error: %v", err)
	}
	file.Close()
	return nil
}
func UpdateCartFile(cList *CartList) error {
	bytes, err := jsoniter.Marshal(&cList)
	if err != nil {
		return fmt.Errorf("cannot marshall dList, error: %v", err)
	}
	file, err := os.OpenFile(Config.Cartname, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		return fmt.Errorf("cannot open file, error: %v", err)
	}
	_, err = file.Write(bytes)
	if err != nil {
		return fmt.Errorf("cannot write into file, error: %v", err)
	}
	file.Close()
	return nil
}
func InitCart() {
	bytes, err := ioutil.ReadFile(Config.Cartname)
	if err != nil {
		logrus.Panicf("cannot read data file, err %v", err)
	}
	err = jsoniter.Unmarshal(bytes, &Clist)
	if err != nil {
		logrus.Panicf("cannot unmarshall file, error: %v", err)
	}
	return
}
