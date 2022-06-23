package entity

type Collection struct {
	// id
	CollectionId   int    `json:"collection-id"`
	CollectionName string `json:"collection-name"`

	// images
	LogoImage    string `json:"logo-image" exmaple:"http://www.iamge.com/123455"`
	FeatureImage string `json:"feature-image" exmaple:"http://www.iamge.com/123455"`
	BannerImage  string `json:"banner-image" exmaple:"http://www.iamge.com/123455"`

	// items contained
	ItemsCount int      `json:"items"`
	ItemIds    []string `json:"item-ids"`

	// account description and create time
	Description string `json:"description"`
	CreateTime  string `json:"create-time"`
	Owner       string `json:"owner"`
}
