package repository

import (
	"database/sql"

	"github.com/lib/pq"

	"github.com/bopbi/SkillRec/checker"
	"github.com/bopbi/SkillRec/entity"
)

// GetSkillRecommendersByUserID return array of UserResponse
func GetSkillRecommendersByUserID(db *sql.DB, id int) []*entity.SkillRecommenderResponse {

	skillListResponse := GetSkillsByUserID(db, id)
	var skillsIDs []int

	for _, skillResponseForID := range skillListResponse {
		skillsIDs = append(skillsIDs, skillResponseForID.ID)
	}

	queryString := `
	SELECT u.id, u.name, u.email, us.skill_id FROM user_skill_recommenders urs
    	INNER JOIN user_skills us ON urs.user_skill_id = us.id
		INNER JOIN users u ON urs.recommender_id = u.id
		INNER JOIN skills s ON us.skill_id = s.id
	WHERE urs.user_skill_id = ANY($1)
	ORDER BY us.skill_id ASC
	`
	var recommenderListResponse []*entity.RecommenderResponse
	rows, err := db.Query(queryString, pq.Array(skillsIDs))
	checker.CheckErr(err)
	for rows.Next() {
		recommenderResponse := new(entity.RecommenderResponse)
		err = rows.Scan(&recommenderResponse.ID, &recommenderResponse.Name, &recommenderResponse.Email, &recommenderResponse.SkillID)
		recommenderListResponse = append(recommenderListResponse, recommenderResponse)
	}

	var startIndex = 0
	var skillRecommenders []*entity.SkillRecommenderResponse
	for _, skillResponse := range skillListResponse {
		skillRecommender := new(entity.SkillRecommenderResponse)
		skillRecommender.Skill = *skillResponse
		skillRecommender.User = make([]entity.UserResponse, 0)
		// grouping recommender of each skill
		for index := startIndex; index < len(recommenderListResponse); index++ {
			recommenderResponse := recommenderListResponse[index]
			if recommenderResponse.SkillID == skillResponse.ID {
				userResponse := new(entity.UserResponse)
				userResponse.ID = recommenderResponse.ID
				userResponse.Name = recommenderResponse.Name
				userResponse.Email = recommenderResponse.Email
				skillRecommender.User = append(skillRecommender.User, *userResponse)
				startIndex = index + 1
			} else {
				break
			}
		}
		// end grouping recommender
		skillRecommenders = append(skillRecommenders, skillRecommender)
	}

	return skillRecommenders
}
