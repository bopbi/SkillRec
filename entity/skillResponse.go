package entity

// SkillResponse package the struct for json response
type SkillResponse struct {
	ID   int    `json:"id" xml:"id" form:"id" query:"id"`
	Name string `json:"name" xml:"name" form:"name" query:"name"`
}
