package service

import (
	"go-openapi-sdk/dto"
	"go-openapi-sdk/dto/enums"
	"testing"
)

func TestSFBRequest_Save(t *testing.T) {
	//	house := getSalehouse()
	house := []interface{}{getSalehouse(), getSaleValla(), getSaleShop(), getSaleOffice(), getLeaseHouse(), getLeaseVilla(), getLeaseShop(), getLeaseOffice()}
	for _, v := range house {
		result, err := Request.Save(Result, v, 2)
		if err != nil {
			t.Fatal(err)
			return
		}
		t.Log(result)
		Request.HttpClient.Log.Println(result, err)
	}
}

func TestSFBRequest_Modify(t *testing.T) {
	house := getSalehouse()
	data, err := Request.Modify(Result, house, 227025582)
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(data)
	Request.HttpClient.Log.Println(data, err)
}
func TestDeteil(t *testing.T) {

	data, err := Request.Deteil(Result, enums.HouseType_Sale, 227025582)
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(data)
	Request.HttpClient.Log.Println(data, err)
}

func TestPageList(t *testing.T) {
	query := dto.NewQueryDTO()
	data, err := Request.PageList(Result, query)
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(data)
	list := data.List
	Request.HttpClient.Log.Println(data, err)
	Request.HttpClient.Log.Println(list, err)
}

func TestReleaseHouse(t *testing.T) {
	houseids := make([]int, 0)
	houseids = append(houseids, 227025582)
	houseids = append(houseids, 216339984)
	houseids = append(houseids, 221072695)
	data, err := Request.ReleaseHouse(Result, enums.HouseType_Sale, houseids, 2)
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(data)
	Request.HttpClient.Log.Println(data, err)
}

func TestUnreleaseHouse(t *testing.T) {
	houseids := make([]int, 0)
	houseids = append(houseids, 227025582)
	houseids = append(houseids, 216339984)
	houseids = append(houseids, 221072695)
	data, err := Request.UnreleaseHouse(Result, enums.HouseType_Sale, houseids, 2)
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(data)
	Request.HttpClient.Log.Println(data, err)
}

func TestHousePhotoList(t *testing.T) {
	data, err := Request.HousePhotoList(Result, enums.HouseType_Sale, 221072695)
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(data)
	Request.HttpClient.Log.Println(data, err)
}

func TestHousePhotoAdd(t *testing.T) {
	photo := &dto.HousePhotoAddDTO{HouseId: 221072695, HouseType: enums.HouseType_Sale, PhotoName: "大图片",
		PhotoUrl: "http://img11.soufunimg.com/viewimage/agents/2018_07/20/M1F/0E/5B/ChCE4ltRa_eIWZpAAAEYJaTpm-gABD_EAGO4IkAARg9345/690x440c.jpg",
		AlbumId:  1,
	}
	err := Request.HousePhotoAdd(Result, photo)
	if err != nil {
		Request.HttpClient.Log.Println(err)
		t.Fatal(err)
		return
	}

}
func TestHousePhotoModify(t *testing.T) {
	//994819
	photo := &dto.HousePhotoModifyDTO{HouseId: 221072695, HouseType: enums.HouseType_Sale, MapPhotoIdName: make(map[int]string)}
	photo.MapPhotoIdName[994813] = "侧视图"
	err := Request.HousePhotoModify(Result, photo)
	if err != nil {
		Request.HttpClient.Log.Println(err)
		t.Fatal(err)
		return
	}
}

func TestHousePhotoDelete(t *testing.T) {
	//994819
	photo := []int{1569734, 1569735, 1569736}
	err := Request.HousePhotoDelete(Result, enums.HouseType_Sale, 221072695, photo)
	if err != nil {
		Request.HttpClient.Log.Println(err)
		t.Fatal(err)
		return
	}
}

func TestHousePhotoMultAdd(t *testing.T) {
	photos := &dto.HousePhotoMultAddDTO{HouseId: 221072695, HouseType: enums.HouseType_Sale}
	photos.Image1 = append(photos.Image1, dto.Image{Url: "http://img11.soufunimg.com/viewimage/agents/2018_07/20/M1F/0E/5B/ChCE4ltRa_eIWZpAAAEYJaTpm-gABD_EAGO4IkAARg9345/690x440c.jpg", Name: "室内图"})
	photos.Image1 = append(photos.Image1, dto.Image{Url: "http://img11.soufunimg.com/viewimage/agents/2018_07/20/M1F/0E/5B/ChCE4ltRa_eIWZpAAAEYJaTpm-gABD_EAGO4IkAARg9345/690x440c.jpg", Name: "室内图"})
	photos.Image1 = append(photos.Image1, dto.Image{Url: "http://img11.soufunimg.com/viewimage/agents/2018_07/20/M1F/0E/5B/ChCE4ltRa_eIWZpAAAEYJaTpm-gABD_EAGO4IkAARg9345/690x440c.jpg", Name: "室内图"})

	str, err := Request.HousePhotoMultAdd(Result, photos)
	if err != nil {

		t.Fatal(err)
		return
	}
	Request.HttpClient.Log.Println(str)
}
func TestHouseTitlePictureModify(t *testing.T) {
	err := Request.HouseTitlePictureModify(Result, enums.HouseType_Sale, 221072695, "http://img11.soufunimg.com/viewimage/agents/2018_07/20/M1F/0E/5B/ChCE4ltRa_eIWZpAAAEYJaTpm-gABD_EAGO4IkAARg9345/690x440c.jpg")
	if err != nil {
		Request.HttpClient.Log.Println(err)
		t.Fatal(err)
		return
	}
}
func getBaseHouse(house *dto.HouseBaseDTO) {
	//17
	house.NewCode = 12132131
	house.ProjName = "前程似锦"
	house.District = "丰台区"
	house.Comarea = "科技园区"
	house.Address = "大望路"
	house.Price = 601.09
	house.BuildingArea = 16.8
	house.InnerId = "951357"
	house.Title = "好房子大房子低价卖了"
	house.InfoCode = "9876543"
	house.AllFloor = 20
	house.PhotoUrl = "http://img11.soufunimg.com/viewimage/agents/2018_07/20/M1F/0E/5B/ChCE4ltRa_eIWZpAAAEYJaTpm-gABD_EAGO4IkAARg9345/690x440c.jpg"
	house.Content = "好房子大房子"
	house.Trafficinfo = "11路公交"
	house.VideoTitle = "房源视频"
	house.VideoUrl = "http://baidu.com"
	house.SubwayInfo = "电梯，天然气"

}

func getSalehouse() *dto.HouseSaleDTO {
	house := dto.NewHouseSaleDTO()
	getBaseHouse(&house.HouseBaseDTO)
	house.CreateTime = "2001-07-06"
	house.LiveArea = 88.8
	house.BaseService = "煤气/天然气,暖气"
	house.Forward = "东北"
	house.PayInfo = "限价房"
	house.LookHouse = "随时看房"
	house.Room = 3
	house.Toilet = 1
	house.Hall = 1
	house.Kitchen = 1
	house.Balcony = 3
	house.Floor = 7
	house.Fitment = "豪华装修"
	house.BuildingType = "跃层"
	house.HouseStructure = "平层"
	house.PropertySubType = "propertysubtype"
	house.Ownermentality = "业主急售"
	house.Serviceintroduction = "帮过户、贷款"
	house.Taxanalysis = "税费超级低，只收1.5"
	house.Communitymatching = "社区配套"
	house.Feature = "满五唯一"
	house.Image1 = append(house.Image1, dto.Image{Url: "http://img11.soufunimg.com/viewimage/agents/2018_07/20/M1F/0E/5B/ChCE4ltRa_eIWZpAAAEYJaTpm-gABD_EAGO4IkAARg9345/690x440c.jpg", Name: "室内图"})
	house.Image1 = append(house.Image1, dto.Image{Url: "http://img11.soufunimg.com/viewimage/agents/2018_07/20/M1F/0E/5B/ChCE4ltRa_eIWZpAAAEYJaTpm-gABD_EAGO4IkAARg9345/690x440c.jpg", Name: "室内图"})
	house.Image1 = append(house.Image1, dto.Image{Url: "http://img11.soufunimg.com/viewimage/agents/2018_07/20/M1F/0E/5B/ChCE4ltRa_eIWZpAAAEYJaTpm-gABD_EAGO4IkAARg9345/690x440c.jpg", Name: "室内图"})
	house.Image2 = append(house.Image2, dto.Image{Url: "http://img11.soufunimg.com/viewimage/agents/2018_07/21/M15/0E/63/ChCE4ltSrmKIWdMlAADULOzwLmEABECMAG440AAANRE828/690x440c.jpg", Name: "户外图"})
	house.Image3 = append(house.Image3, dto.Image{Url: "http://img11.soufunimg.com/viewimage/agents/2018_07/20/M1F/0E/5B/ChCE4ltRa_eIWZpAAAEYJaTpm-gABD_EAGO4IkAARg9345/690x440c.jpg", Name: "室内图"})
	house.Image3 = append(house.Image3, dto.Image{Url: "http://img11.soufunimg.com/viewimage/agents/2018_07/20/M1F/0E/5B/ChCE4ltRa_eIWZpAAAEYJaTpm-gABD_EAGO4IkAARg9345/690x440c.jpg", Name: "室内图"})
	return house
}

func getSaleValla() *dto.VillaSaleDTO {
	house := dto.NewVillaSaleDTO()
	getBaseHouse(&house.HouseBaseDTO)
	//22
	house.CreateTime = "2001-07-06"
	house.LiveArea = 88.2
	house.BaseService = "煤气/天然气,暖气"
	house.Forward = "东北"
	house.Room = 2
	house.Toilet = 1
	house.Hall = 2
	house.Kitchen = 1
	house.Balcony = 3
	house.Fitment = "豪华装修"
	house.BuildingType = "跃层"
	house.HouseStructure = "平层"
	house.Ownermentality = "业主急售"
	house.Serviceintroduction = "帮过户、贷款"
	house.Taxanalysis = "税费超级低，只收1.5"
	house.Communitymatching = "社区配套"
	house.WorkshopArea = 15.6
	house.ShopType = "半明"
	house.SpaceArea = 908.0
	house.ParkingPlace = 2
	house.LookHouse = "随时"
	house.Feature = "满五唯一"
	house.Image1 = append(house.Image1, dto.Image{Url: "http://img11.soufunimg.com/viewimage/agents/2018_07/20/M1F/0E/5B/ChCE4ltRa_eIWZpAAAEYJaTpm-gABD_EAGO4IkAARg9345/690x440c.jpg", Name: "室内图"})
	house.Image1 = append(house.Image1, dto.Image{Url: "http://img11.soufunimg.com/viewimage/agents/2018_07/20/M1F/0E/5B/ChCE4ltRa_eIWZpAAAEYJaTpm-gABD_EAGO4IkAARg9345/690x440c.jpg", Name: "室内图"})
	house.Image2 = append(house.Image1, dto.Image{Url: "http://img11.soufunimg.com/viewimage/agents/2018_07/21/M15/0E/63/ChCE4ltSrmKIWdMlAADULOzwLmEABECMAG440AAANRE828/690x440c.jpg", Name: "户外图"})

	return house
}

func getSaleShop() *dto.ShopSaleDTO {
	house := dto.NewShopSaleDTO()
	getBaseHouse(&house.HouseBaseDTO)
	house.SubType = "propertysubtype"
	house.PropFee = 12.30
	house.Floor = 12
	house.IsDivisi = 1
	house.Fitment = "豪华装修"
	house.BaseService = "电梯"
	house.AimOperastion = "饭店"
	house.Feature = "特色"
	return house
}

func getSaleOffice() *dto.OfficeSaleDTO {
	house := dto.NewOfficeSaleDTO()
	getBaseHouse(&house.HouseBaseDTO)
	house.SubType = "subType"
	house.PropFee = 11.2
	house.Floor = 12
	house.IsDivisi = 1
	house.Fitment = "豪华"
	house.PropertyGrade = "甲级"
	return house
}

func getLeaseHouse() *dto.HouseLeaseDTO {
	house := dto.NewHouseLeaseDTO()
	getBaseHouse(&house.HouseBaseDTO)
	house.Room = 2
	house.Toilet = 1
	house.Hall = 2
	house.Kitchen = 1
	house.Balcony = 3
	house.Floor = 10
	house.PayDetail = "押一付三"
	house.Forward = "东北"
	house.LookHouse = "随时看房"
	house.Fitment = "豪华装修"
	house.LeaseStyle = "合租"
	house.RoommateCount = 2
	house.Leaseroomtype = "主卧"
	house.BaseService = "煤气/天然气,暖气"
	house.Feature = "满五唯一"
	house.Serviceintroduction = "帮过户、贷款"
	house.Communitymatching = "社区配套"
	house.Housetypeintroduction = "户型方正"
	house.Subwayinfo = "z周边大商场"
	house.Leasedetail = "限女生"
	house.Image1 = append(house.Image1, dto.Image{Url: "http://img11.soufunimg.com/viewimage/agents/2018_07/20/M1F/0E/5B/ChCE4ltRa_eIWZpAAAEYJaTpm-gABD_EAGO4IkAARg9345/690x440c.jpg", Name: "室内图"})
	house.Image1 = append(house.Image1, dto.Image{Url: "http://img11.soufunimg.com/viewimage/agents/2018_07/20/M1F/0E/5B/ChCE4ltRa_eIWZpAAAEYJaTpm-gABD_EAGO4IkAARg9345/690x440c.jpg", Name: "室内图"})
	house.Image2 = append(house.Image1, dto.Image{Url: "http://img11.soufunimg.com/viewimage/agents/2018_07/21/M15/0E/63/ChCE4ltSrmKIWdMlAADULOzwLmEABECMAG440AAANRE828/690x440c.jpg", Name: "户外图"})
	return house
}

func getLeaseVilla() *dto.VillaLeaseDTO {
	house := dto.NewVillaLeaseDTO()
	getBaseHouse(&house.HouseBaseDTO)
	house.BuildingType = "独栋"
	house.HouseStructure = "平层"
	house.Room = 2
	house.Toilet = 1
	house.Hall = 2
	house.Kitchen = 1
	house.Balcony = 3
	house.PayDetail = "押一付三"
	house.CreateTime = "2011-01-01"
	house.LiveArea = 12.3
	house.Forward = "东北"
	house.ShopType = "全明"
	house.SpaceArea = 19.2
	house.Garage = 3
	house.Fitment = "豪华装修"
	house.LeaseStyle = "合租"
	house.BaseService = "煤气/天然气,暖气"
	house.Equitment = "电话"
	house.LookHouse = "随时看房"
	house.LiveTime = "2019-02-03"
	house.Serviceintroduction = "帮过户、贷款"
	house.Communitymatching = "社区配套"
	house.Housetypeintroduction = "户型方正"
	house.Subwayinfo = "z周边大商场"
	house.Image1 = append(house.Image1, dto.Image{Url: "http://img11.soufunimg.com/viewimage/agents/2018_07/20/M1F/0E/5B/ChCE4ltRa_eIWZpAAAEYJaTpm-gABD_EAGO4IkAARg9345/690x440c.jpg", Name: "室内图"})
	house.Image1 = append(house.Image1, dto.Image{Url: "http://img11.soufunimg.com/viewimage/agents/2018_07/20/M1F/0E/5B/ChCE4ltRa_eIWZpAAAEYJaTpm-gABD_EAGO4IkAARg9345/690x440c.jpg", Name: "室内图"})
	house.Image2 = append(house.Image1, dto.Image{Url: "http://img11.soufunimg.com/viewimage/agents/2018_07/21/M15/0E/63/ChCE4ltSrmKIWdMlAADULOzwLmEABECMAG440AAANRE828/690x440c.jpg", Name: "户外图"})

	return house
}

func getLeaseShop() *dto.ShopLeaseDTO {
	house := dto.NewShopLeaseDTO()
	getBaseHouse(&house.HouseBaseDTO)
	house.SubType = "住宅底商"
	house.ShopStatus = "营业中"
	house.IsIncludFee = 1
	house.PropFee = 2.0
	house.IsTransfer = 1
	house.PayDetail = "押一付三"
	house.Floor = 2
	house.IsDivisi = 0
	house.Fitment = "豪华装修"
	house.BaseService = "煤气/天然气,暖气"
	house.AimOperastion = "赌场"
	house.Feature = "有钱就行,满五唯一"
	house.TransferFee = 0.2

	house.Image6 = append(house.Image6, dto.Image{Url: "http://img11.soufunimg.com/viewimage/agents/2018_07/20/M1F/0E/5B/ChCE4ltRa_eIWZpAAAEYJaTpm-gABD_EAGO4IkAARg9345/690x440c.jpg", Name: "室内图"})
	house.Image7 = append(house.Image7, dto.Image{Url: "http://img11.soufunimg.com/viewimage/agents/2018_07/20/M1F/0E/5B/ChCE4ltRa_eIWZpAAAEYJaTpm-gABD_EAGO4IkAARg9345/690x440c.jpg", Name: "室内图"})
	house.Image7 = append(house.Image7, dto.Image{Url: "http://img11.soufunimg.com/viewimage/agents/2018_07/21/M15/0E/63/ChCE4ltSrmKIWdMlAADULOzwLmEABECMAG440AAANRE828/690x440c.jpg", Name: "户外图"})

	return house
}

func getLeaseOffice() *dto.OfficeLeaseDTO {
	house := dto.NewOfficeLeaseDTO()
	getBaseHouse(&house.HouseBaseDTO)
	house.SubType = "纯写字楼"
	house.PropFee = 1.2
	house.IsIncludFee = 1
	house.PayDetail = "押一付三"
	house.Floor = 2
	house.IsDivisi = 0
	house.Fitment = "豪华装修"
	house.PropertyGrade = "甲级"
	house.Feature = "有钱就行,满五唯一"
	house.Image6 = append(house.Image6, dto.Image{Url: "http://img11.soufunimg.com/viewimage/agents/2018_07/20/M1F/0E/5B/ChCE4ltRa_eIWZpAAAEYJaTpm-gABD_EAGO4IkAARg9345/690x440c.jpg", Name: "室内图"})
	house.Image7 = append(house.Image7, dto.Image{Url: "http://img11.soufunimg.com/viewimage/agents/2018_07/20/M1F/0E/5B/ChCE4ltRa_eIWZpAAAEYJaTpm-gABD_EAGO4IkAARg9345/690x440c.jpg", Name: "室内图"})
	house.Image4 = append(house.Image4, dto.Image{Url: "http://img11.soufunimg.com/viewimage/agents/2018_07/21/M15/0E/63/ChCE4ltSrmKIWdMlAADULOzwLmEABECMAG440AAANRE828/690x440c.jpg", Name: "户外图"})
	return house
}
