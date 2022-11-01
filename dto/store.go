package dto

type StoreUpdateDTO struct {
	ID       uint64 `json:"id" form:"id" binding:"required"`
	Name     string `uri:"name" form:"name" binding:"required"`
	Email    string `uri:"email" form:"email" binding:"required"`
	UserID   uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
	location string `uri:"location" form:"location" bindin:"required"`
}

type StoreCreateDTO struct {
	ID       uint64 `json:"id" form:"id" binding:"required"`
	Name     string `uri:"name" form:"name" binding:"required"`
	location string `uri:"location" form:"location" bindin:"required"`
}
