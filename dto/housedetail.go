package dto

import "go-openapi-sdk/util"

//HouseDetailDTO 房源详情
type HouseDetailDTO struct {

	/**
	 * 房源id
	 */
	HouseId int `json:"houseid"`
	/**
	 * 经纪人id
	 */
	AgentId int `json:"agentid"`
	/**
	 * 房源标题
	 */
	BoardTitle string `json:"boardtitle"`
	/**
	 * 房源描述
	 */
	BoardContent string `json:"boardcontent"`
	/**
	 * 周边配套
	 */
	SubWayInfo string `json:"subwayinfo"`
	/**
	 * 交通状况
	 */
	TrafficInfo string `json:"trafficinfo"`
	/**
	 * 业主心态
	 */
	OwnerMentality string
	/**
	 * 社区配套
	 */
	CommunityMatching string
	/**
	 * 服务介绍
	 */
	ServiceIntroduction string
	/**
	 * 税费分析
	 */
	TaxAnalysis string
	/**
	 * 户型介绍 出租住宅、别墅
	 */
	HouseTypeIntroduction string

	/**
	 * 视频标题
	 */
	VideoTilte string `json:"videotilte"`
	/**
	 * 视频URL
	 */
	VideoURL string `json:"videourl"`
	/**
	 * 房管局房源二维码地址
	 */
	Qrimgurl string
	/**
	 * 城市
	 */
	City string
	/**
	 * 是否推广到了新搜房帮
	 */
	IsPubWsfb int
	/**
	 * 是否推广到了商铺帮
	 */
	IsPubShop int
	/**
	 * 是否推广到办公帮
	 */
	IsPubOffice int

	Floor int
	/**
	 * 总楼层数
	 */
	TotalFloor int `json:"totalfloor"`
	/**
	 * 使用面积
	 */
	Livearea float64 `json:"liveArea"`
	/**
	 * 朝向
	 */
	Forward string
	/**
	 * 建筑年代
	 */
	CreateTime string `json:"createtime"`
	/**
	 * 基础设施
	 */
	Baseservice string `json:"baseService"`
	/**
	 * 装修程度
	 */
	Fitment string
	/**
	 * 看房时间
	 */
	LookHouseType string `json:"lookhousetype"`
	/**
	 * 住宅类型
	 */
	PropertySubType string `json:"propertysubtype"`
	Room            uint16
	Hall            uint16
	Toilet          uint16
	Kitchen         uint16
	Balcony         uint16
	/**
	 * 地下室面积
	 */
	WorkshopArea float64 `json:"workshoparea"`
	/**
	 * 花园面积 别墅
	 */
	SpaceArea float64 `json:"spacearea"`
	/**
	 * 车库数量 （别墅）
	 */
	Garage int
	/**
	 * 车位数量（别墅）
	 */
	ParkingPlace int `json:"parkingplace"`
	/**
	 * 铺面类型（商铺）或地下室类型（别墅）
	 */
	ShopType string `json:"shoptype"`
	/**
	 * 物业费
	 */
	PropertyFee float64 `json:"propertyfee"`
	/**
	 * 物业公司(商铺、写字楼)
	 */
	Propertycompany string `json:"propertyCompany"`
	/**
	 * 目标业态(商铺)
	 */
	Aimoperastion string `json:"aimOperastion"`
	/**
	 * 物业级别(写字楼)
	 */
	Propertygrade string `json:"propertyGrade"`
	/**
	 * 建筑类别 独栋 双拼 联排 叠加	别墅、出售住宅
	 */
	BuildingType string `json:"buildingtype"`
	/**
	 * 厅结构：平层 挑高	别墅、住宅、厂房
	 */
	HouseStructure string `json:"housetructure"`
	Tag            int
	Status         int
	IsValid        int
	Source         string

	District string
	Comarea  string
	Address  string
	Purpose  string
	Projcode uint64
	Projname string
	/**
	 * 建筑面积
	 */
	BuildingArea float64 `json:"buildingarea"`
	/**
	 * 价格
	 */
	Price float64
	/**
	 * 价格单位
	 */
	PriceType string `json:"pricetype"`
	/**
	 * 内部房源编号
	 */
	Innerid string `json:"innerId"`
	/**
	 * 房源封面图
	 */
	PhotoURL string `json:"photourl"`
	/**
	 * 特色分类
	 */
	Feature string
	/**
	 * 房源信息编码
	 */
	InfoCode string `json:"infocode"`
	/**
	 * 录入时间
	 */
	Inserttime util.JsonTime `json:"insertTime"`
	/**
	 * 修改时间
	 */
	ModifyDate util.JsonTime `json:"modifydate"`
	/**
	 * 是否有视频
	 */
	IsVideo int `json:"isvideo"`
	/**
	 * 房源图片数量
	 */
	ImageCount int `json:"imagecount"`
	/**
	 * 是否优质房源
	 */
	IsBest int `json:"isbest"`
	/**
	 * 房源推广状态
	 */
	PromoteProductType int `json:"PromoteProductType"`
	/**
	 * 房源点击付费状态
	 */
	ClickListingStatus int `json:"ClickListingStatus"`
}
