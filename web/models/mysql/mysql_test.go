package mysql

import (
	"fmt"
	"testing"
)

func TestGetConnect(t *testing.T){
	db, err := GetConnect()
	if err != nil{
		fmt.Println(db,err)
	}
	//db.CreateTable()
}
