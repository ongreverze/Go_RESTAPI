package data

type NewData struct {
	Unique_id string `json:"unique_id" binding:"required"`
	Name string `json:"name" binding:"required"`
	Age uint `json:"age" binding:"required"`
}