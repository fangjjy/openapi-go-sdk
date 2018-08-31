package service

import (
	"errors"
	"fmt"
	"go-openapi-sdk/dto"
	"go-openapi-sdk/util"
	"strconv"
	"strings"
)

//Save 房源保存
func (req *SFBRequest) Save(authDTO *dto.AtuhDTO, house interface{}, ptype int) (*dto.HouseReturnDTO, error) {
	data := &dto.HouseReturnDTO{}
	header := req.buildHeader(authDTO)
	fmt.Println(header)
	body, err := util.StructToForm(house)
	if err != nil {
		return nil, err
	}
	if ptype != 0 {
		body["ptype"] = ptype
	}
	//
	result, err := req.HttpClient.HTTPPost(authDTO.Domain+"/House/Input", header, body)
	if err != nil {
		return nil, err
	}

	err = buildResult(result, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//Modify 房源修改
func (req *SFBRequest) Modify(authDTO *dto.AtuhDTO, house interface{}, houseid int) (*dto.HouseReturnDTO, error) {
	data := &dto.HouseReturnDTO{}
	header := req.buildHeader(authDTO)
	body, err := util.StructToForm(house)
	if err != nil {
		return nil, err
	}
	body["houseid"] = houseid
	result, err := req.HttpClient.HTTPPost(authDTO.Domain+"/House/update", header, body)
	if err != nil {
		return nil, err
	}

	err = buildResult(result, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//Delete 房源删除
func (req *SFBRequest) Delete(authDTO *dto.AtuhDTO, houseType string, houseId int) (*dto.ReturnDTO, error) {
	header := req.buildHeader(authDTO)
	body := make(map[string]interface{}, 0)
	body["housetype"] = houseType
	body["houseid"] = houseId
	result, err := req.HttpClient.HTTPPost(authDTO.Domain+"/house/delete", header, body)
	if err != nil {
		return nil, err
	}
	data := &dto.ReturnDTO{}
	err = util.JsonDecode(result, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//Deteil 房源详情
func (req *SFBRequest) Deteil(authDTO *dto.AtuhDTO, houseType string, houseId int) (*dto.HouseDetailDTO, error) {
	header := req.buildHeader(authDTO)
	body := make(map[string]interface{}, 0)
	body["housetype"] = houseType
	body["houseid"] = houseId
	result, err := req.HttpClient.HTTPGet(authDTO.Domain+"/house/detail", header, body)
	if err != nil {
		return nil, err
	}
	fmt.Println(result)
	if err != nil {
		return nil, err
	}
	data := &dto.HouseDetailDTO{}
	err = buildResult(result, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//PageList 房源列表
func (req *SFBRequest) PageList(authDTO *dto.AtuhDTO, queryDTO *dto.HouseQueryDTO) (*dto.HousePageDTO, error) {
	header := req.buildHeader(authDTO)
	result, err := req.HttpClient.HTTPGet(authDTO.Domain+"/house/list?"+queryDTO.Format(), header, nil)

	if err != nil {
		return nil, err
	}
	data := &dto.HousePageDTO{}
	err = buildResult(result, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//ReleaseHouse 房源推广
func (req *SFBRequest) ReleaseHouse(authDTO *dto.AtuhDTO, houseType string, houseIds []int, pType int) (*dto.ReturnDTO, error) {
	if houseIds == nil || len(houseIds) == 0 {
		return nil, errors.New("houseIds不可为空")
	}
	header := req.buildHeader(authDTO)
	strhouseids := make([]string, 0)
	for _, v := range houseIds {
		strhouseids = append(strhouseids, strconv.Itoa(v))
	}
	body := make(map[string]interface{})
	body["action"] = "releasehouse"
	body["houseType"] = houseType
	body["houseids"] = strings.Join(strhouseids, ",")
	body["ptype"] = pType

	result, err := req.HttpClient.HTTPPost(authDTO.Domain+"/house/RealeaseHouse", header, body)
	if err != nil {
		return nil, err
	}
	data := &dto.ReturnDTO{}
	err = util.JsonDecode(result, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//UnreleaseHouse 取消房源推广
func (req *SFBRequest) UnreleaseHouse(authDTO *dto.AtuhDTO, houseType string, houseIds []int, pType int) (*dto.ReturnDTO, error) {
	if houseIds == nil || len(houseIds) == 0 {
		return nil, errors.New("houseIds不可为空")
	}
	header := req.buildHeader(authDTO)
	strhouseids := make([]string, 0)
	for _, v := range houseIds {
		strhouseids = append(strhouseids, strconv.Itoa(v))
	}
	body := make(map[string]interface{})
	body["action"] = "unreleasehouse"
	body["houseType"] = houseType
	body["houseids"] = strings.Join(strhouseids, ",")
	body["ptype"] = pType

	result, err := req.HttpClient.HTTPPost(authDTO.Domain+"/house/RealeaseHouse", header, body)
	if err != nil {
		return nil, err
	}
	data := &dto.ReturnDTO{}
	err = util.JsonDecode(result, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// HousePhotoList 房源图片列表
func (req *SFBRequest) HousePhotoList(authDTO *dto.AtuhDTO, houseType string, houseid int) ([]dto.HousePhotoListDTO, error) {
	header := req.buildHeader(authDTO)
	body := make(map[string]interface{}, 0)
	body["housetype"] = houseType
	body["houseid"] = houseid
	result, err := req.HttpClient.HTTPGet(authDTO.Domain+"/House/HousePhotoList", header, body)
	if err != nil {
		return nil, err
	}
	data := make([]dto.HousePhotoListDTO, 0)
	err = buildResult(result, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//HousePhotoAdd 房源图片添加
func (req *SFBRequest) HousePhotoAdd(authDTO *dto.AtuhDTO, photo *dto.HousePhotoAddDTO) error {
	header := req.buildHeader(authDTO)
	body, err := util.StructToForm(photo)
	if err != nil {
		return err
	}
	result, err := req.HttpClient.HTTPGet(authDTO.Domain+"/House/HousePhotoAdd", header, body)
	if err != nil {
		return err
	}
	err = buildComResult(result)
	if err != nil {
		return err
	}
	return nil
}

//HousePhotoModify 修改房源那图片名称
func (req *SFBRequest) HousePhotoModify(authDTO *dto.AtuhDTO, photo *dto.HousePhotoModifyDTO) error {
	if photo == nil || len(photo.MapPhotoIdName) == 0 {
		return errors.New("传入参数有误")
	}
	header := req.buildHeader(authDTO)
	body := make(map[string]interface{}, 0)
	body["houseid"] = photo.HouseId
	body["housetype"] = photo.HouseType
	photonames := make([]string, 0)
	photoIds := make([]string, 0)
	for i, v := range photo.MapPhotoIdName {
		photoIds = append(photoIds, fmt.Sprint(i))
		photonames = append(photonames, v)
	}
	body["photonames"] = strings.Join(photonames, ",")
	body["photoids"] = strings.Join(photoIds, ",")
	result, err := req.HttpClient.HTTPGet(authDTO.Domain+"/House/HousePhotoModify", header, body)
	if err != nil {
		return err
	}
	err = buildComResult(result)
	if err != nil {
		return err
	}
	return nil
}

//HousePhotoDelete 房源图片删除
func (req *SFBRequest) HousePhotoDelete(authDTO *dto.AtuhDTO, houseType string, houseid int, photoids []int) error {
	if photoids == nil || len(photoids) == 0 {
		return errors.New("photoids不可为空")
	}
	header := req.buildHeader(authDTO)
	strphotoids := make([]string, 0)
	for _, v := range photoids {
		strphotoids = append(strphotoids, strconv.Itoa(v))
	}
	body := make(map[string]interface{})
	body["houseType"] = houseType
	body["photoids"] = strings.Join(strphotoids, ",")
	body["houseid"] = houseid
	result, err := req.HttpClient.HTTPPost(authDTO.Domain+"/House/HousePhotoDelete", header, body)
	if err != nil {
		return err
	}
	err = buildComResult(result)
	if err != nil {
		return err
	}
	return nil
}

//HousePhotoMultAdd 房源那图片批量添加
func (req *SFBRequest) HousePhotoMultAdd(authDTO *dto.AtuhDTO, photos *dto.HousePhotoMultAddDTO) (string, error) {
	header := req.buildHeader(authDTO)
	body, err := util.StructToForm(photos)
	if err != nil {
		return "", err
	}
	result, err := req.HttpClient.HTTPPost(authDTO.Domain+"/House/HousePhotoMultAdd", header, body)
	if err != nil {
		return "", err
	}
	var data string
	err = buildResult(result, &data)
	if err != nil {
		return "", err
	}
	return data, nil
}

//HouseTitlePictureModify 修改房源标题图
func (req *SFBRequest) HouseTitlePictureModify(authDTO *dto.AtuhDTO, houseType string, houseid int, photourl string) error {

	header := req.buildHeader(authDTO)
	body := make(map[string]interface{})
	body["houseType"] = houseType
	body["photourl"] = photourl
	body["houseid"] = houseid
	result, err := req.HttpClient.HTTPPost(authDTO.Domain+"/House/HouseTitlePictureModify", header, body)
	if err != nil {
		return err
	}
	err = buildComResult(result)
	if err != nil {
		return err
	}
	return nil

}
