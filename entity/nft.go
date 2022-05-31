package entity

type Item struct {
	// basic info
	Name        string `json:"name"`
	ItemID      string `json:"item-id" example:"123455"`
	Image       string `json:"image" exmaple:"http://www.iamge.com/123455"`
	CreateTime  string `json:"create-time"`
	Description string `json:"description"`

	// characteristic
	ItemCollection string   `json:"collection-id" example:"5"`
	Category       string   `json:"category"`
	Label          []string `json:"lable"`

	// account id info
	CreaterId string `json:"creater-id" example:"mazhengwang-ust-hk"`
	OwnerId   string `json:"owner-id" example:"mazhengwang-ust-hk"`

	// token_uri got from blockchain
	token_uri string `json:"token_Uri"`

	// favorite
	Favorites uint64 `json:"favorites" example:"1"`

	// transaction history
	History []string `json:"history"`
}

type Collection struct {
	// id
	CollectionId   string `json:"collection-id"`
	CollectionName string `json:"collection-name"`

	// images
	LogoImage    string `json:"logo-image" exmaple:"http://www.iamge.com/123455"`
	FeatureImage string `json:"feature-image" exmaple:"http://www.iamge.com/123455"`
	BannerImage  string `json:"banner-image" exmaple:"http://www.iamge.com/123455"`

	// items contained
	ItemsCount uint64   `json:"items"`
	ItemIds    []string `json:"item-ids"`

	// account description and create time
	Description string `json:"description"`
	CreateTime  string `json:"create-time"`
	Owner       string `json:"owner"`
}
