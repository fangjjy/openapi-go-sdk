package service

import (
	"errors"
	"fmt"
	"go-openapi-sdk/dto"
	"go-openapi-sdk/util"
	"net/url"
)

//SFBRequest 搜房帮请求
type SFBRequest struct {
	KeyUnit    *dto.CompanyKeyUnit
	HttpClient *util.HttpClient
	Domain     string
}

//NewSFBRequest 新建一个请求
func NewSFBRequest(keyid int64, secret string) *SFBRequest {
	request := &SFBRequest{
		KeyUnit:    &dto.CompanyKeyUnit{KeyId: keyid, Secret: secret},
		HttpClient: util.NewHttpClient(10),
		Domain:     "http://openapi.agentn.fang.com/",
	}
	return request
}

//ApplyToken 申请密令，以及后续其他操作的域名
//返回数据正确，则dto 的data是AtuhDTO
func (req *SFBRequest) ApplyToken(username, pwd string) (*dto.AtuhDTO, error) {
	header := make(map[string]string, 0)
	body := make(map[string]interface{}, 0)
	header["KeyId"] = fmt.Sprintf("%v", req.KeyUnit.KeyId)
	header["DataType"] = "json"
	header["Accept"] = "text/json"
	var authdto dto.AtuhDTO
	var resultdto dto.ReturnGenericDTO
	resultdto.Data = &authdto
	pwddes, err1 := util.DesEncode(req.KeyUnit.Secret, req.KeyUnit.Secret, pwd)
	if err1 != nil {
		return &authdto, err1
	}
	body["userName"] = username
	body["pwd"] =  url.QueryEscape(pwddes)
	body["keyid"] = req.KeyUnit.KeyId
	body["datatype"] = "json"
	result, err := req.HttpClient.HTTPPost(req.Domain+"/agent/UserLoginAuthenticate/", header, body)
	if err != nil {

		return nil, err
	}
	err = buildResult(result, &authdto)
	if err != nil {
		req.HttpClient.Log.Println(result, err)
		return nil, err
	}
	return &authdto, nil
}

//GetAgentPowerInfo 获取经纪人权限
func (req *SFBRequest) GetAgentPowerInfo(authDTO *dto.AtuhDTO) ([]dto.AgentPowerInfo, error) {
	header := req.buildHeader(authDTO)
	body := make(map[string]interface{}, 0)
	body["param"] = 1231
	result, err := req.HttpClient.HTTPGet(authDTO.Domain+"/Agent/GetAgentPowerInfo", header, body)
	data := make([]dto.AgentPowerInfo, 0)
	if err != nil {
		return nil, err
	}
	err = buildResult(result, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (req *SFBRequest) buildHeader(authdto *dto.AtuhDTO) map[string]string {
	header := make(map[string]string, 0)
	header["KeyId"] = fmt.Sprintf("%v", req.KeyUnit.KeyId)
	header["DataType"] = "json"
	header["Token"] = authdto.Token
	header["RequestID"] = fmt.Sprintf("%v", req.KeyUnit.KeyId)
	return header
}

func buildResult(content string, data interface{}) error {
	var resultdto dto.ReturnGenericDTO
	resultdto.Data = data
	err := util.JsonDecode(content, &resultdto)
	if err != nil {
		return err
	}
	if !resultdto.ReturnDTO.IsSuccess() {
		return errors.New(resultdto.Description)
	}
	return nil
}

func buildComResult(content string) error {
	var resultdto dto.ReturnGenericDTO
	err := util.JsonDecode(content, &resultdto)
	if err != nil {
		return err
	}
	if !resultdto.ReturnDTO.IsSuccess() {
		return errors.New(resultdto.Description)
	}
	return nil
}
