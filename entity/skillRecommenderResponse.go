package entity

// SkillRecommenderResponse package the struct for json response
type SkillRecommenderResponse struct {
	Skill SkillResponse  `json:"skill" xml:"skill" form:"skill" query:"skill"`
	User  []UserResponse `json:"recommenders" xml:"recommenders" form:"recommenders" query:"recommenders"`
}
