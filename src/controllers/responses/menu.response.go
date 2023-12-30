package responses

import "github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/models"

type GetMenuListResponse struct {
	Items []GetMenuListItem `json:"items"`
}

type GetMenuListItem struct {
	CreatedAt    string `json:"createdAt"`
	Name         string `json:"name"`
	Price        int64  `json:"price"`
	Category     string `json:"category"`
	MealTimeType string `json:"mealTimeType"`
}

func NewGetMenuListResponse(menuList []models.Menu) GetMenuListResponse {
	res := GetMenuListResponse{
		Items: []GetMenuListItem{},
	}

	for _, menu := range menuList {
		res.Items = append(res.Items, GetMenuListItem{
			CreatedAt:    menu.CreatedAt.Format("2006-01-02 15:04:05"),
			Name:         menu.Name,
			Price:        menu.Price,
			Category:     menu.Category,
			MealTimeType: menu.MealTimeType,
		})
	}
	return res
}
