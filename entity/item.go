package entity

type Item struct {
	// basic info
	ItemName string `json:"item_name" example:"hahhahah"`
	ItemID   string `json:"item_id" example:"123455"`
	ItemData string `json:"item_data" exmaple:"https://img1.baidu.com/it/u=1783064339,1648739044&fm=253&fmt=auto&app=138&f=GIF?w=240&h=240"`

	CreateTime  string `json:"create-time" example:"2022-06-16 22:04:22"`
	Description string `json:"description" example:"A very cute pixel bear with hammer"`

	CollectionId   string   `json:"collection_id" example:"1"`
	CollectionName string   `json:"collection_name" example:"Pixel Bear"`
	Category       string   `json:"category" example:"Pixel Bear"`
	Label          []string `json:"lable" example:"Music,Comics"`

	CreaterId string `json:"creater_id" example:"mazhengwang-ust-hk"`
	OwnerId   string `json:"owner_id" example:"mazhengwang-ust-hk"`

	FavoriteNum int  `json:"like_count" example:"100"`
	Favorite    bool `json:"favorite" example:"false"`
}
