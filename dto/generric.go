package dto

//ReturnDTO 一般返回類型
type ReturnDTO struct {
	Code        int32
	Description string
}

//IsSuccess 返回结果是否成功
func (dto ReturnDTO) IsSuccess() bool {
	if dto.Code == 1 {
		return true
	}
	return false
}

//ReturnGenericDTO 接口返回带结果数据
type ReturnGenericDTO struct {
	ReturnDTO
	Data interface{} `json:"returnmsgs"`
}

//AtuhDTO 授权返回数据
type AtuhDTO struct {
	Token   string `json:"Token"` //经纪人后续操作的密令
	Domain  string `json:"URL"`   //经纪人后续操作的域名
	BUserId int64  //用户的B端id
}

//AgentPowerInfo 用户权限
type AgentPowerInfo struct {
	/**
	 * 最大可推广房源量
	 */
	HouseMax int `json:"housemax"`
	/**
	 * 已推广房源量
	 */
	HouseUseCount int `json:"houseusecount"`
	/**
	 * 套餐版本
	 */
	ProductVersion int `json:"productversion"`
	/**
	 * 产品类型（说明：2 搜房帮； 4 品牌公寓； 8 租房帮； 16 商铺帮； 32 办公帮）
	 */
	ProductType int `json:"producttype"`
}

//CompanyKeyUnit 从搜房帮申请的key
type CompanyKeyUnit struct {
	KeyId  int64
	Secret string
}

//HouseReturnDTO 房源录入修改返回实体
type HouseReturnDTO struct {
	/**
	 * 房源id
	 */
	HouseId int32
	/**
	 * 房源的内部编号
	 */
	Innerid string
	/**
	 * 房源
	 */
	HouseUrl string
	/**
	 * 房源推广成功是1，否则是2
	 */
	Flag int
	/**
	 * 推广的端，表示房源推广到了那些端
	 */
	Promotedto []int
}
