package dto

import (
	"go-openapi-sdk/dto/enums"
)

//HouseSaleDTO 出售住宅录入
type HouseSaleDTO struct {
	HouseBaseDTO
	/***
	 * 类别，来自与客服后台的录入
	 */
	PropertySubType string
	/**
	 * 室
	 */
	Room uint16
	/**
	 * 厅
	 */
	Hall uint16
	/**
	 * 卫生间
	 */
	Toilet uint16
	/**
	 * 厨房
	 */
	Kitchen uint16
	/**
	 * 阳台
	 */
	Balcony uint16

	/**
	 * 所在楼层
	 */
	Floor uint16
	/**
	 * 产权性质
	 */
	PayInfo string
	/**
	 * 使用面积
	 */
	LiveArea float64
	/**
	 * 建筑年代
	 */
	CreateTime string
	/**
	 * 朝向
	 */
	Forward string
	/**
	 * 装修程度
	 */
	Fitment string
	/**
	 * 配套设施
	 */
	BaseService string
	/**
	 * 看房时间类别
	 */
	LookHouse string

	/**
	 * 房源标签
	 */
	Feature string
	/**
	 * 建筑结构
	 */
	HouseStructure string
	/**
	 * 建筑形式
	 */
	BuildingType string
	/**
	 * 社（小）区配套（只有住宅别墅有，其他物业类型留备）
	 */
	Communitymatching string
	/**
	 * 服务介绍（只有住宅别墅有，其他物业类型留备）
	 */
	Serviceintroduction string
	/**
	 * 业主心态（只有住宅别墅有，其他物业类型留备）
	 */
	Ownermentality string
	/**
	 * 税费分析（只有住宅别墅有，其他物业类型留备）
	 */

	Taxanalysis string

	/**
	 * 室内图
	 */
	Image1 []Image
	/***
	 * 小区图
	 */
	Image2 []Image
	/***
	 * 户型图
	 */
	Image3 []Image
}

//NewHouseSaleDTO 新建出售住宅
func NewHouseSaleDTO() *HouseSaleDTO {
	house := &HouseSaleDTO{}
	house.HouseType = enums.HouseType_Sale
	house.PurposeType = enums.Purpose_Office
	return house
}

//VillaSaleDTO 别墅出售录入
type VillaSaleDTO struct {
	HouseBaseDTO
	/**
	 * 建筑形式
	 */
	BuildingType string
	/**
	 * PayInfo
	 */
	Room uint16
	/**
	 * 厅
	 */
	Hall uint16
	/**
	 * 卫生间
	 */
	Toilet uint16
	/**
	 * 厨房
	 */
	Kitchen uint16
	/**
	 * 阳台
	 */
	Balcony uint16
	/**
	 * 厅结构
	 */
	HouseStructure string
	/**
	 * 地下室类型
	 */
	ShopType string
	/**
	 * 花园面积（如果有）
	 */
	SpaceArea float64
	/**
	 * 车库数量
	 */
	Garage int
	/**
	 * 车位数量（如果有）
	 */
	ParkingPlace int
	/**
	 * 使用面积
	 */
	LiveArea float64
	/**
	 * 建筑年代
	 */
	CreateTime string
	/**
	 * 进门朝向
	 */
	Forward string
	/**
	 * 装修程度
	 */
	Fitment string
	/**
	 * 配套设施
	 */
	BaseService string
	/**
	 * 看房时间
	 */
	LookHouse string

	/**
	 * 房源标签
	 */
	Feature string
	/**
	 * 社（小）区配套（只有住宅别墅有，其他物业类型留备）
	 */
	Communitymatching string
	/**
	 * 服务介绍（只有住宅别墅有，其他物业类型留备）
	 */
	Serviceintroduction string
	/**
	 * 业主心态（只有住宅别墅有，其他物业类型留备）
	 */
	Ownermentality string
	/**
	 * 税费分析（只有住宅别墅有，其他物业类型留备）
	 */

	Taxanalysis string
	/**
	 * 地下室面积
	 */
	WorkshopArea float64
	/**
	 * 室内图
	 */
	Image1 []Image
	/***
	 * 小区图
	 */
	Image2 []Image
	/***
	 * 户型图
	 */
	Image3 []Image
}

//NewVillaSaleDTO 新建出售别墅
func NewVillaSaleDTO() *VillaSaleDTO {
	house := &VillaSaleDTO{}
	house.HouseType = enums.HouseType_Sale
	house.PurposeType = enums.Purpose_Villa
	return house
}

//ShopSaleDTO 商铺出售录入
type ShopSaleDTO struct {
	HouseBaseDTO
	/***
	 * 类别，必须在存在于搜房帮，可以到搜房帮的页面查看
	 */
	SubType string
	/**
	 * 物业费
	 */
	PropFee float64
	/**
	 * 所在楼层
	 */
	Floor uint16
	/**
	 * 是否可分割
	 */
	IsDivisi int
	/**
	 * 装修程度
	 */
	Fitment string
	/**
	 * 配套设施
	 */
	BaseService string
	/**
	 * 目标业态
	 */
	AimOperastion string
	/**
	 * 房源标签
	 */
	Feature string
	/**
	 * 内景图
	 */
	Image6 []Image
	/**
	 * 外景图
	 */
	Image7 []Image
}

//NewShopSaleDTO 新建出售写商铺
func NewShopSaleDTO() *ShopSaleDTO {
	house := &ShopSaleDTO{}
	house.HouseType = enums.HouseType_Sale
	house.PurposeType = enums.Purpose_Shop
	return house
}

//OfficeSaleDTO 写字楼出售录入
type OfficeSaleDTO struct {
	HouseBaseDTO
	/***
	 * 类别，来自与客服后台的录入
	 */
	SubType string
	/**
	 * 物业费
	 */
	PropFee float64
	/**
	 * 所在楼层
	 */
	Floor uint16
	/**
	 * 是否可分割
	 */
	IsDivisi int
	/**
	 * 装修程度
	 */
	Fitment string
	/**
	 * 写字楼级别
	 */
	PropertyGrade string

	/**
	 * 房源标签
	 */
	Feature string
	/**
	 * 交通图
	 */
	Image4 []Image
	/**
	 * 平面图
	 */
	Image5 []Image
	/**
	 * 内景图
	 */
	Image6 []Image
	/**
	 * 外景图
	 */
	Image7 []Image
}

//NewOfficeSaleDTO 新建出售写字楼
func NewOfficeSaleDTO() *OfficeSaleDTO {
	house := &OfficeSaleDTO{}
	house.HouseType = enums.HouseType_Sale
	house.PurposeType = enums.Purpose_Office
	return house
}
