package data

import "errors"

var (
	errorDroppingCollection = errors.New("error dropping collection")
	errorDecodingToObjectID = errors.New("error decoding to object id")

	errorInsertingProject          = errors.New("error inserting project")
	errorFindingAllProjects        = errors.New("error finding all projects")
	errorDecodingProjectsIntoSlice = errors.New("error decoding projects into slice")
	errorGettingProjectByTitle     = errors.New("error getting project by title")
	errorUpdatingProject           = errors.New("error updating project")

	errorInsertingCodingSkill          = errors.New("error inserting coding skill")
	errorFindingAllCodingSkills        = errors.New("error finding all coding skills")
	errorDecodingCodingSkillsIntoSlice = errors.New("error decoding coding skills into slice")
	errorGettingCodingSkillByTitle     = errors.New("error getting coding skill by title")
	errorUpdatingCodingSkill           = errors.New("error updating coding skill")

	errorInsertingAchievement          = errors.New("error inserting achievement")
	errorFindingAllAchievements        = errors.New("error finding all achievements")
	errorDecodingAchievementsIntoSlice = errors.New("error decoding achievements into slice")
	errorGettingAchievementByTitle     = errors.New("error getting achievement by title")
	errorUpdatingAchievement           = errors.New("error updating achievement")
)
