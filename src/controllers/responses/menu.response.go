package responses

import (
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/consts"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/domains/models"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/utils"
)

/*
メニューリスト取得
*/
type GetMenuListResponse struct {
	PagingResponse
	Items []GetMenuListItem `json:"items"`
}

type GetMenuListItem struct {
	Id           uint   `json:"id"`
	CreatedAt    string `json:"createdAt"`
	Name         string `json:"name"`
	Price        int64  `json:"price"`
	Category     string `json:"category"`
	MealTimeType string `json:"mealTimeType"`
}

func NewGetMenuListResponse(menuList []models.Menu, totalCount int) GetMenuListResponse {
	res := GetMenuListResponse{
		Items: []GetMenuListItem{},
	}

	for _, menu := range menuList {
		res.TotalCount = totalCount
		res.PerPageCount = consts.PerPageCount
		res.Items = append(res.Items, GetMenuListItem{
			Id:           menu.Id,
			CreatedAt:    utils.TimeFormat(utils.ConvertTimeUtcToJst(menu.CreatedAt), "yyyyMMddHHmmss"),
			Name:         menu.Name,
			Price:        menu.Price,
			Category:     menu.Category,
			MealTimeType: menu.MealTimeType,
		})
	}
	return res
}

/*
メニューガチャ取得
*/
type GetMenuGachaResponse struct {
	Budget     int                `json:"budget"`
	TotalPrice int                `json:"totalPrice"`
	Items      []GetMenuGachaItem `json:"items"`
}

type GetMenuGachaItem struct {
	Id           uint   `json:"id"`
	CreatedAt    string `json:"createdAt"`
	Name         string `json:"name"`
	Price        int64  `json:"price"`
	Category     string `json:"category"`
	MealTimeType string `json:"mealTimeType"`
}

func NewGetMenuGachaResponse(menuList []models.Menu, budget int) GetMenuGachaResponse {
	res := GetMenuGachaResponse{
		Items: []GetMenuGachaItem{},
	}

	res.Budget = budget
	for _, menu := range menuList {
		res.TotalPrice = res.TotalPrice + int(menu.Price)
		res.Items = append(res.Items, GetMenuGachaItem{
			Id:           menu.Id,
			CreatedAt:    utils.TimeFormat(utils.ConvertTimeUtcToJst(menu.CreatedAt), "yyyyMMddHHmmss"),
			Name:         menu.Name,
			Price:        menu.Price,
			Category:     menu.Category,
			MealTimeType: menu.MealTimeType,
		})
	}
	return res
}
