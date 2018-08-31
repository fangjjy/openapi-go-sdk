package dto

//HouseBaseDTO 所有房源的基类
type HouseBaseDTO struct {
	HouseType string

	/**
	 * 楼盘字典名称
	 */

	ProjName string
	/**
	 * 楼盘字典id
	 */
	NewCode uint64
	/**
	 * 楼盘所在区县
	 */

	District string
	/**
	 * 楼盘所在商圈
	 */

	Comarea string
	/**
	 * 物业地址
	 */

	Address string
	/**
	 * 物业类型（住宅，别墅、商铺、写字楼）
	 */
	PurposeType string
	/**
	 * 价格
	 */
	Price float64
	/**
	 * 建筑面积
	 */

	BuildingArea float64
	/**
	 * 房源内部编号
	 */
	InnerId string
	/**
	 * 标 题 图
	 */
	PhotoUrl string
	/**
	 * 房源信息编码
	 */
	InfoCode string
	/**
	 * 总楼层数
	 */
	AllFloor uint16

	/**
	 * 房源描述
	 */
	Content string

	/**
	 * 交通状况
	 */
	Trafficinfo string
	/**
	 * 房源标题
	 */
	Title string
	/**
	 * 周边配套
	 */
	SubwayInfo string
	/**
	 * 视频的地址，如果值是字符串null,表示删除视频，如果不传或是空代表视频不做修改
	 */
	VideoUrl string
	/**
	 * 视频标题
	 */
	VideoTitle string
}
