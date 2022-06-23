package service

import (
	"NFT-BASE-BACK/base"
	"NFT-BASE-BACK/entity"
	"NFT-BASE-BACK/utils"
)

func LoadItemHistory(ItenId string) []string {
	return []string{"Created by mingzheliu-ust-hk"}
}

func CreateItem(UserId string, name string, image string, description string, itemCollection string, category string, label []string) (base.ErrCode, entity.Item) {
	item := entity.Item{
		Name:           name,
		ItemID:         "1012",
		Image:          image,
		CreateTime:     utils.GetTimeNow(),
		Description:    description,
		ItemCollection: itemCollection,
		Category:       category,
		Label:          label,
		CreaterId:      UserId,
		OwnerId:        UserId,
		Favorites:      0,
		History:        LoadItemHistory("1012"),
	}
	return base.Success, item
}

func EditItem(UserId string, name string, image string, description string, itemCollection string, category string, label []string) (base.ErrCode, entity.Item) {
	item := entity.Item{
		Name:           name,
		ItemID:         "1012",
		Image:          image,
		CreateTime:     utils.GetTimeNow(),
		Description:    description,
		ItemCollection: itemCollection,
		Category:       category,
		Label:          label,
		CreaterId:      UserId,
		OwnerId:        UserId,
		Favorites:      0,
		History:        LoadItemHistory("1012"),
	}
	return base.Success, item
}

func GetItem(ItemId string) (base.ErrCode, entity.Item) {
	item := entity.Item{
		Name:           "Pixel Bear With Hammer",
		ItemID:         "1010",
		Image:          "https://img1.baidu.com/it/u=1783064339,1648739044&fm=253&fmt=auto&app=138&f=GIF?w=240&h=240",
		CreateTime:     utils.GetTimeNow(),
		Description:    "A very cute pixel bear with hammer",
		ItemCollection: "Pixel Bear",
		Category:       "image",
		Label:          []string{"pixel", "bear"},
		CreaterId:      "mingzheliu-ust-hk",
		OwnerId:        "mingzheliu-ust-hk",
		Favorites:      1000000,
		History:        []string{"Created by baofuhan-ust-hk", "Transaction: baofuhan-ust-hk to mingzheliu-ust-hk"},
	}
	return base.Success, item
}

// Id可以是collection或者user的id，后边的method是查询方法，collections同理
func GetItems(Id string, pgnumber int, pgsize int, method string) (base.ErrCode, []entity.Item, int) {
	item1 := entity.Item{
		Name:           "Pixel Bear With Hammer",
		ItemID:         "1010",
		Image:          "https://img1.baidu.com/it/u=1783064339,1648739044&fm=253&fmt=auto&app=138&f=GIF?w=240&h=240",
		CreateTime:     utils.GetTimeNow(),
		Description:    "A very cute pixel bear with hammer",
		ItemCollection: "Pixel Bear",
		Category:       "image",
		Label:          []string{"pixel", "bear"},
		CreaterId:      "mingzheliu-ust-hk",
		OwnerId:        "mingzheliu-ust-hk",
		Favorites:      1000000,
		History:        []string{"Created by baofuhan-ust-hk", "Transaction: baofuhan-ust-hk to mingzheliu-ust-hk"},
	}
	item2 := entity.Item{
		Name:           "Pixel Bear With Pet",
		ItemID:         "1011",
		Image:          "https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fimg.zcool.cn%2Fcommunity%2F0178c6577a22b40000018c1b286ea9.gif&refer=http%3A%2F%2Fimg.zcool.cn&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=auto?sec=1656499539&t=27dfbc3a3e02abb7eec805034f4da7e2",
		CreateTime:     utils.GetTimeNow(),
		Description:    "A very cute pixel bear with pet",
		ItemCollection: "Pixel Bear",
		Category:       "image",
		Label:          []string{"pixel", "bear"},
		CreaterId:      "mingzheliu-ust-hk",
		OwnerId:        "mingzheliu-ust-hk",
		Favorites:      1000000,
		History:        []string{"Created by baofuhan-ust-hk", "Transaction: baofuhan-ust-hk to mingzheliu-ust-hk"},
	}
	return base.Success, []entity.Item{item1, item2}, 2
}

func LoadItemIds(CollectionId string) []string {
	return []string{"1000", "1001", "1003", "1004"}
}

func EditCollection(UserId string, CollectionId string, name string, logoImage string, featureImage string, bannerImage string, description string) (base.ErrCode, entity.Collection) {

	collection := entity.Collection{
		CollectionName: name,
		LogoImage:      logoImage,
		FeatureImage:   featureImage,
		BannerImage:    bannerImage,
		ItemsCount:     12,
		ItemIds:        LoadItemIds(CollectionId),
		Description:    description,
		CreateTime:     utils.GetTimeNow(),
		Owner:          UserId,
	}
	return base.Success, collection
}

func GetCollection(CollectionId string) (base.ErrCode, entity.Collection) {
	collection := entity.Collection{
		CollectionId:   1,
		CollectionName: "Pixel Bear",
		LogoImage:      "https://img2.baidu.com/it/u=3149703497,1232639303&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=888",
		FeatureImage:   "https://img0.baidu.com/it/u=407638665,2322213270&fm=253&fmt=auto&app=120&f=JPEG?w=1422&h=800",
		BannerImage:    "https://img0.baidu.com/it/u=2383738538,1577541218&fm=253&fmt=auto&app=138&f=JPEG?w=1346&h=500",
		ItemsCount:     12,
		ItemIds:        LoadItemIds(CollectionId),
		Description:    "A collection of Pixel Bears",
		CreateTime:     utils.GetTimeNow(),
		Owner:          "mingzheliu-ust-hk",
	}
	return base.Success, collection
}

func GetCollections(Id string, pgnumber int, pgsize int, method string) (base.ErrCode, []entity.Collection, int) {
	collection1 := entity.Collection{
		CollectionId:   1,
		CollectionName: "Pixel Bear",
		LogoImage:      "https://img2.baidu.com/it/u=3149703497,1232639303&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=888",
		FeatureImage:   "https://img1.baidu.com/it/u=2611341694,1602866768&fm=253&fmt=auto&app=138&f=JPEG?w=889&h=500",
		BannerImage:    "https://img0.baidu.com/it/u=2383738538,1577541218&fm=253&fmt=auto&app=138&f=JPEG?w=1346&h=500",
		ItemsCount:     12,
		ItemIds:        LoadItemIds("1"),
		Description:    "A collection of Pixel Bears",
		CreateTime:     utils.GetTimeNow(),
		Owner:          "mingzheliu-ust-hk",
	}
	collection2 := entity.Collection{
		CollectionId:   1,
		CollectionName: "MetaUniverse",
		LogoImage:      "https://img0.baidu.com/it/u=2102284223,1322820129&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=500",
		FeatureImage:   "https://img2.baidu.com/it/u=3756534050,1657480602&fm=253&fmt=auto&app=138&f=JPEG?w=640&h=363",
		BannerImage:    "https://img1.baidu.com/it/u=502242943,1278115791&fm=253&fmt=auto&app=138&f=JPEG?w=499&h=308",
		ItemsCount:     2,
		ItemIds:        LoadItemIds("2"),
		Description:    "A collection of my imagination",
		CreateTime:     utils.GetTimeNow(),
		Owner:          "mingzheliu-ust-hk",
	}
	return base.Success, []entity.Collection{collection1, collection2}, 2
}

func EditProfile(UserId string, BannerImage string, AvatarImage string, Poison string, Campus string) (base.ErrCode, entity.Account) {
	account := entity.Account{
		BannerImage: BannerImage,
		AvatarImage: AvatarImage,
		UserId:      UserId,
		Poison:      Poison,
		Campus:      Campus,
		Token:       GetToken(UserId),
	}
	return base.Success, account
}

func GetUser(id string) (base.ErrCode, entity.Account) {
	user := entity.Account{
		"https://img0.baidu.com/it/u=2549864824,978431990&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=157",
		"https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fc-ssl.duitang.com%2Fuploads%2Fblog%2F202107%2F17%2F20210717232533_2edcf.thumb.1000_0.jpg&refer=http%3A%2F%2Fc-ssl.duitang.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=auto?sec=1656495244&t=e7ca8e69bf79e0903fddc701e225255b",
		id,
		"Student",
		"HKUST-GZ",
		100,
	}
	return base.Success, user
}

func GetToken(UserId string) uint64 {
	return 100
}

func DeleteItem(ItemId string, UserId string) base.ErrCode {
	return base.Success
}

func DeleteCollection(CollectionId string, UserId string) base.ErrCode {
	return base.Success
}
