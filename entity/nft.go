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
