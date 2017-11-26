package entity

type UserAndSkillResponse struct {
	User  UserResponse  `json:"user" xml:"user" form:"user" query:"user"`
	Skill SkillResponse `json:"skill" xml:"skill" form:"skill" query:"skill"`
}
