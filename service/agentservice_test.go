package service

import (
	"fmt"
	"go-openapi-sdk/dto"
	"go-openapi-sdk/util"
	"os"
	"testing"
	"time"
)

var Result *dto.AtuhDTO
var Err error
var keyUnit = &dto.CompanyKeyUnit{KeyId: 123455, Secret: "xxxxxxxxxxxxxxxxxx"}
var client = util.NewHttpClient(10 * time.Second)
var Request = &SFBRequest{KeyUnit: keyUnit, HttpClient: client, Domain: "http://openapi.agents.test.fang.com"}

func init() {
	filename := "./log.txt"
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE, os.ModeAppend)
	if err != nil {
		fmt.Println(err)
	} else {
		Request.HttpClient.Log.SetOutput(file)
	}
	Result, Err = Request.ApplyToken("username", "pwd")
	if Err != nil {
		Request.HttpClient.Log.Fatal(Err)
	}
}
func fileExist(filepath string) bool {
	_, err := os.Stat(filepath)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}
func TestSFBRequest_ApplyToken(t *testing.T) {
	if Err != nil {
		t.Fatal(Err)
		return
	}
	fmt.Println(Result)
	t.Log(*Result)
	Request.HttpClient.Log.Println(*Result)
}

func TestSFBRequest_GetAgentPowerInfo(t *testing.T) {
	dto, err := Request.GetAgentPowerInfo(Result)
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(dto)
}
