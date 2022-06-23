package service

import (
	"NFT-BASE-BACK/base"
	"NFT-BASE-BACK/entity"
	"NFT-BASE-BACK/model"
	"NFT-BASE-BACK/utils"
)

func CreateCollectionByAccount(UserId string, name string, logoImage string, featureImage string, bannerImage string, description string) (base.ErrCode, entity.Collection) {
	c := entity.Collection{
		CollectionName: name,
		LogoImage:      logoImage,
		FeatureImage:   featureImage,
		BannerImage:    bannerImage,
		ItemsCount:     0,
		Description:    description,
		CreateTime:     utils.GetTimeNow(),
		Owner:          UserId,
	}
	err := model.CreatCollection(c.CollectionName, c.LogoImage, c.FeatureImage, c.BannerImage, c.ItemsCount, c.Description, UserId, UserId, c.CreateTime)
	if err != nil {
		return base.ServerError, entity.Collection{}
	}
	return base.Success, c
}
