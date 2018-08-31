package dto

import (
	"bytes"
	"go-openapi-sdk/dto/enums"
	"go-openapi-sdk/util"
	"strconv"
)

//HouseQueryDTO 房源列表查询
type HouseQueryDTO struct {
	HouseType   string
	PurposeType string
	HouseStatus string
	HouseId     int
	InnerId     string
	ProjName    string
	CurPage     int
	PageSize    int
	QueryPType  int
}

//NewQueryDTO 新建房源列表查询
func NewQueryDTO() *HouseQueryDTO {
	var dto = &HouseQueryDTO{}
	dto.HouseType = enums.HouseType_Sale
	dto.CurPage = 1
	dto.PageSize = 20
	dto.QueryPType = enums.PagePType_ALL
	return dto
}

//FormatQuery 格式化查询条件
func (query *HouseQueryDTO) Format() string {
	bufferstring := bytes.NewBufferString("housetype=" + query.HouseType)
	bufferstring.WriteString("&purposetype=")
	bufferstring.WriteString("&houseStatus=")
	bufferstring.WriteString(query.HouseStatus)
	if query.HouseId > 0 {
		bufferstring.WriteString("&houseid=")
		bufferstring.WriteString(strconv.Itoa(query.HouseId))
	}
	bufferstring.WriteString("&innerid=")
	bufferstring.WriteString(query.InnerId)

	bufferstring.WriteString("&projname=")
	bufferstring.WriteString(query.ProjName)

	bufferstring.WriteString("&curpage=")
	bufferstring.WriteString(strconv.Itoa(query.CurPage))

	bufferstring.WriteString("&pagesize=")
	bufferstring.WriteString(strconv.Itoa(query.PageSize))

	bufferstring.WriteString("&ptype=")
	bufferstring.WriteString(strconv.Itoa(query.QueryPType))
	return bufferstring.String()

}

type HousePageDTO struct {
	/**
	 * 违规房源
	 */
	Violationcount int `json:"violationcount"`
	/**
	 * 过期房源数
	 */
	Overduecount int `json:"overduecount"`
	/**
	 * 查询房源数
	 */
	Housecount int `json:"housecount"`
	/**
	 * 房源列表
	 */
	List []HousePageDetailDTO `json:"list"`
}

type HousePageDetailDTO struct {
	HouseId            int           `json:"houseid"`
	AgentId            int           `json:"agentid"`
	AgentName          string        `json:"agentname"`
	ImageCount         int           `json:"imagecount"`
	IsVideo            string        `json:"isvideo"`
	Address            string        `json:"address"`
	Room               uint8         `json:"room"`
	Hall               uint8         `json:"hall"`
	Toilet             uint8         `json:"toilet"`
	BuildingArea       float32       `json:"buildingArea"`
	Price              float32       `json:"price"`
	PriceType          string        `json:"pricetype"`
	IsbBest            uint8         `json:"isbest"`
	Status             uint8         `json:"status"`
	Source             string        `json:"Source"`
	ProjName           string        `json:"projname"`
	ProjCode           int64         `json:"projCode"`
	LimitDate          int           `json:"limitdate"`
	Tag                uint8         `json:"tag"`
	InsertTime         util.JsonTime `json:"insertTime"`
	IsValid            string        `json:"isvalid"`
	Abnormal           uint8         `json:"abnormal"`
	Purpose            string        `json:"Purpose"`
	Title              string        `json:"boardtitle"`
	PhotoURL           string        `json:"photourl"`
	PromoteProductType int           `json:"promoteproducttype"`
	PromotedTo         []int         `json:"promotedto"`
	Showlevel          int           `json:"showlevel"`
	Tagaddtime         util.JsonTime `json:"tagaddtime"`
	IsRealHouse        uint8         `json:"isRealHouse"`
	IsNewHouse         uint8         `json:"isNewHouse"`
	Deposit            uint8         `json:"Deposit"`
}
