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
	cId, _ := model.GetMaxCollectionId()
	col, _ := model.GetCollection(cId)
	return base.Success, col
}

func EditCollection(c entity.Collection) (base.ErrCode, entity.Collection) {
	if c.CollectionName != "" {
		err := model.EditCollectionName(c.CollectionName, c.CollectionId)
		if err != nil {
			return base.InsertError, entity.Collection{}
		}
	}
	if c.LogoImage != "" {
		err := model.EditLogoImage(c.LogoImage, c.CollectionId)
		if err != nil {
			return base.InsertError, entity.Collection{}
		}
	}
	if c.FeatureImage != "" {
		err := model.EditFeatureImage(c.FeatureImage, c.CollectionId)
		if err != nil {
			return base.InsertError, entity.Collection{}
		}
	}
	if c.BannerImage != "" {
		err := model.EditBannerImage(c.BannerImage, c.CollectionId)
		if err != nil {
			return base.InsertError, entity.Collection{}
		}
	}
	if c.Description != "" {
		err := model.EditDescription(c.Description, c.CollectionId)
		if err != nil {
			return base.InsertError, entity.Collection{}
		}
	}
	g, err := model.GetCollection(c.CollectionId)

	if err != nil {
		return base.ServerError, entity.Collection{}
	}
	return base.Success, g
}
