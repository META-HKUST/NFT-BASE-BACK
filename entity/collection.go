package entity

type Collection struct {
	// id
	CollectionId   int    `json:"collection_id" db:"collection_id"`
	CollectionName string `json:"collection_name" db:"collection_name"`

	// images
	LogoImage    string `json:"logo_image"  db:"logo_image"`
	FeatureImage string `json:"feature_image" db:"feature_image"`
	BannerImage  string `json:"banner_image"  db:"banner_image"`

	// items contained
	ItemsCount int      `json:"items_count" db:"items_count"`
	ItemIds    []string `json:"item-ids" db:"item-ids"`

	// account description and create time
	Description string `json:"description" db:"description"`
	Owner       string `json:"owner" db:"owner"`
	OwnerName   string `json:"owner_name" db:"owner_name"`
	CreateTime  string `json:"created_at" db:"created_at"`
}
