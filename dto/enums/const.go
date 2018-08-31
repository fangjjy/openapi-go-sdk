package enums

const (
	HouseType_Sale  = "Sale"  //出售
	HouseType_Lease = "Lease" //出租
)
const (
	Purpose_House  = "House"  //住宅
	Purpose_Villa  = "Villa"  //别墅
	Purpose_Shop   = "Shop"   //商铺
	Purpose_Office = "Office" //写字楼
)
const (
	PType_SFB    = 2  //搜房帮
	PType_ZFB    = 8  //租房帮
	PType_SHOP   = 16 //商铺帮
	PType_OFFICE = 32 //写字楼
)

const (
	PagePType_PROMOTING = -2 //推广中
	PagePType_ALL       = -1 //全部房源
	PagePType_NONE      = 0  //没有任何推广的房源
	PagePType_SFB       = 2  //搜房帮
	PagePType_ZFB       = 8  //租房帮
	PagePType_SHOP      = 16 //商铺帮
	PagePType_OFFICE    = 32 //写字楼
)
