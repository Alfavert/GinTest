package internal

import (
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
)

const base10 = 10

func (Dlist *DataList) GetInfoHandler(ctx *gin.Context) {
	if CheckAuth(ctx) != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	id, err := strconv.ParseInt(ctx.Query("id"), base10, 64)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	for i := range Dlist.List {
		if Dlist.List[i].Id == id {
			ctx.JSON(http.StatusOK, Dlist.List[i])
		}
	}
	ctx.Status(http.StatusNoContent)
}
func (Dlist *DataList) GetAllInfoHandler(ctx *gin.Context) {
	if CheckAuth(ctx) != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	ctx.JSON(http.StatusOK, Dlist)
}
func (Dlist *DataList) AddInfoHandler(ctx *gin.Context) {
	if CheckAuth(ctx) != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	bytes, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		logrus.Warnf("cannot read body, err %v", err)
		return
	}
	if len(bytes) == 0 {
		ctx.AbortWithStatus(http.StatusBadRequest)
		logrus.Warnf("empty body request, err %v", err)
		return
	}
	var Data Data
	err = jsoniter.Unmarshal(bytes, &Data)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		logrus.Warnf("cannot unmarshall body, error:", err)
		return
	}
	Dlist.List = append(Dlist.List, Data)
	err = UpdateFile(Dlist)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		logrus.Warnf("cannot write file, error:", err)
		return
	}
}
func (Clist *CartList) AddInfoCart(ctx *gin.Context) {
	if CheckAuth(ctx) != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	bytes, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		logrus.Warnf("cannot read body, err %v", err)
		return
	}
	if len(bytes) == 0 {
		ctx.AbortWithStatus(http.StatusBadRequest)
		logrus.Warnf("empty body request, err %v", err)
		return
	}
	var cart Cart
	err = jsoniter.Unmarshal(bytes, &cart)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		logrus.Warnf("cannot unmarshall body, error:", err)
		return
	}
	for c := range Dlist.List {
		if cart.Id != Dlist.List[c].Id {
			ctx.AbortWithStatus(http.StatusBadRequest)
			logrus.Warnf("no book with ID")
			return
		}
	}
	Clist.List = append(Clist.List, cart)
	err = UpdateCartFile(Clist)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		logrus.Warnf("cannot write file, error:", err)
		return
	}
}
func (Clist *CartList) Compare(ctx *gin.Context) {
	if CheckAuth(ctx) != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	type temp struct {
		List    []Data
		idCount int64
	}
	tmp := new(temp)
	for i := range Dlist.List {
		for d := range Clist.List {
			if Dlist.List[i].Id == Clist.List[d].Id {
				tmp.List = append(tmp.List, Dlist.List[i])
				tmp.idCount += Dlist.List[i].Id
			}
		}
	}
	if tmp.List != nil {
		ctx.JSON(200, tmp)
	} else {
		ctx.Status(http.StatusNoContent)
	}
}
