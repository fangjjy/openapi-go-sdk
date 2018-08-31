package dto

import "go-openapi-sdk/dto/enums"

//HouseLeaseDTO 出租住宅录入
type HouseLeaseDTO struct {
	HouseBaseDTO
	/**
	 * 卧室
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
	 * 支付方式(押一付三。年付等)
	 */

	PayDetail string
	/**
	 * 朝向
	 */

	Forward string
	/**
	 * 装修程度
	 */

	Fitment string
	/**
	 * 出租方式：整租；合租
	 */

	LeaseStyle string
	/**
	 * 合租户数
	 */
	RoommateCount uint16
	/**
	 * 合住方式：主卧、次卧、床位、隔断间
	 */
	Leaseroomtype string
	/**
	 * 合租限制 ：性别不限、限女生、限男生、限夫妻
	 */
	Leasedetail string
	/**
	 * 配套设施
	 */

	BaseService string
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
	 * 户型介绍（只有住宅别墅暂时有用）
	 */
	Housetypeintroduction string
	/**
	 * 周边配套
	 */
	Subwayinfo string
	/**
	 * 看房时间
	 */
	LookHouse string
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
	image3 []Image
}

//NewHouseLeaseDTO 新建出租住宅
func NewHouseLeaseDTO() *HouseLeaseDTO {
	house := &HouseLeaseDTO{}
	house.HouseType = enums.HouseType_Lease
	house.PurposeType = enums.Purpose_House
	return house
}

//VillaLeaseDTO 别墅出售录入
type VillaLeaseDTO struct {
	HouseBaseDTO
	/**
	 * 建筑形式:
	 * 独栋
	 * 双拼
	 * 联排
	 * 叠加
	 */
	BuildingType string
	/**
	 * 厅结构 平层\挑高
	 */
	HouseStructure string
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
	 * 出租方式：整租；合租
	 */
	LeaseStyle string
	/**
	 * 支付方式(押一付三。年付等)
	 */
	PayDetail string
	/**
	 * 建筑年代
	 */
	CreateTime string
	/**
	 * 使用面积
	 */
	LiveArea float64

	/**
	 * 进门朝向
	 */
	Forward string
	/**
	 * 地下室类型:     全明     半明     暗
	 */
	ShopType string
	/**
	 * 花园面积（如果有）
	 */
	SpaceArea float64
	/**
	 * 车库数量
	 */
	Garage uint16
	/**
	 * 车位数量（如果有）
	 */
	ParkingPlace uint16
	/**
	 * 装修程度
	 */
	Fitment string
	/**
	 * 配套设施
	 */
	BaseService string
	/**
	 * 室内设施
	 */
	Equitment string
	/**
	 * 看房时间
	 */
	LookHouse string
	/**
	 * 入住时间
	 */
	LiveTime string

	/**
	 * 社（小）区配套（只有住宅别墅有，其他物业类型留备）
	 */
	Communitymatching string
	/**
	 * 服务介绍（只有住宅别墅有，其他物业类型留备）
	 */
	Serviceintroduction string
	/**
	 * 户型介绍（只有住宅别墅暂时有用）
	 */
	Housetypeintroduction string
	/**
	 * 周边配套（只有住宅别墅有，其他物业类型留备）
	 */
	Subwayinfo string
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

//NewVillaLeaseDTO 新建出租别墅
func NewVillaLeaseDTO() *VillaLeaseDTO {
	house := &VillaLeaseDTO{}
	house.HouseType = enums.HouseType_Lease
	house.PurposeType = enums.Purpose_Villa
	return house
}

//ShopLeaseDTO 商铺出租录入
type ShopLeaseDTO struct {
	HouseBaseDTO
	/***
	 * 类别:住宅底商 商业街商铺 临街门面 写字楼配套底商 购物中心/百货 其他
	 */
	SubType string
	/**
	 * 商铺状态: 营业中  闲置中 新铺
	 */
	ShopStatus string
	/**
	 * 是否含物业费/管理费(0为不含；1为含)
	 */
	IsIncludFee int
	/**
	 * 物业费
	 */
	PropFee float64
	/**
	 * 是否转让:1/0
	 */
	IsTransfer int
	/**
	 * 支付方式(押一付三。年付等)
	 */
	PayDetail string
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
	 * 转让费（商铺）
	 */
	TransferFee float64
	/**
	 * 内景图
	 */
	Image6 []Image
	/**
	 * 外景图
	 */
	Image7 []Image
}

//NewShopLeaseDTO 新建出租商铺
func NewShopLeaseDTO() *ShopLeaseDTO {
	house := &ShopLeaseDTO{}
	house.HouseType = enums.HouseType_Lease
	house.PurposeType = enums.Purpose_Shop
	return house
}

//OfficeLeaseDTO 写字楼出租录入
type OfficeLeaseDTO struct {
	HouseBaseDTO
	/***
	 * 类别，纯写字楼 商住楼 商业综合体楼 酒店写字楼
	 */
	SubType string
	/**
	 * 物业费
	 */
	PropFee float64
	/**
	 * 是否含物业费/管理费(0为不含；1为含)
	 */
	IsIncludFee int
	/**
	 * 支付方式(押一付三。年付等)
	 */
	PayDetail string

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

//NewOfficeLeaseDTO 新建出租写字楼
func NewOfficeLeaseDTO() *OfficeLeaseDTO {
	house := &OfficeLeaseDTO{}
	house.HouseType = enums.HouseType_Lease
	house.PurposeType = enums.Purpose_Office
	return house
}
