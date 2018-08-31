package dto

//Image 房源录入修改的图像
type Image struct {
	Url  string `json:"url"`
	Name string `json:"name"`
}

// HousePhotoListDTO 图片列表
type HousePhotoListDTO struct {
	HouseId   int    `json:"HouseID"`
	PhotoId   int    `json:"PhotoID"`
	PhotoUrl  string `json:"PHOTOURL"`
	Photoname string `json:"PHOTONAME"`
	Albumid   int    `json:"ALBUMID"`
	Albumname string `json:"ALBUMNAME"`
}

//HousePhotoAddDTO 房源添加接口DTO
type HousePhotoAddDTO struct {
	HouseType string
	HouseId   int
	PhotoName string
	/**
	 * 图片类型（1: 室内图 3:户型图 2:小区相关图 6：内景图 7：外景图 4：交通图 5：平面图）
	 * 注释：
	 * 商铺只能上传  内景图、外景图
	 * 住宅只能上传 户型图、室内图、小区相关图
	 * 写字楼只能上传 交通图、平面图、内景图、外景图
	 * 别墅只能上传 户型图、室内图、小区相关图
	 */
	AlbumId  int
	PhotoUrl string
}

//HousePhotoModifyDTO 房源名称修改
type HousePhotoModifyDTO struct {
	HouseType      string
	HouseId        int
	MapPhotoIdName map[int]string //图片的photoid和名称映射
}

//HousePhotoMultAddDTO  房源图片批量上传DTO
type HousePhotoMultAddDTO struct {
	HouseType string
	HouseId   int
	Image1    []Image
	Image2    []Image
	Image3    []Image
	Image4    []Image
	Image5    []Image
	Image6    []Image
	Image7    []Image
}
