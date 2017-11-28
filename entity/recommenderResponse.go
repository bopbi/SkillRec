package entity

// RecommenderResponse package the struct for json response
type RecommenderResponse struct {
	ID      int    `json:"id" xml:"id" form:"id" query:"id"`
	Name    string `json:"name" xml:"name" form:"name" query:"name"`
	Email   string `json:"email" xml:"email" form:"email" query:"email"`
	SkillID int    `json:"skill_id" xml:"skill_id" form:"skill_id" query:"skill_id"`
}
