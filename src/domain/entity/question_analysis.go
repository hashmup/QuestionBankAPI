package entity

type QuestionAnalysis struct {
	QuestionStudentAnswer       *QuestionStudentAnswer       `json:"question_student_answer"`
	QuestionStudentAnswerRating *QuestionStudentAnswerRating `json:"question_student_answer_rating"`
	QuestionStudentSwitchAnswer *QuestionStudentSwitchAnswer `json:"question_student_switch_answer"`
}

type QuestionStudentAnswerRating struct {
	AverageRating1 float64 `json:"average_rating_1"`
	AverageRating2 float64 `json:"average_rating_2"`
	AverageRating3 float64 `json:"average_rating_3"`
	AverageRating4 float64 `json:"average_rating_4"`
}

type QuestionStudentAnswer struct {
	RateAnswer1 float64 `json:"rate_answer_1"`
	RateAnswer2 float64 `json:"rate_answer_2"`
	RateAnswer3 float64 `json:"rate_answer_3"`
	RateAnswer4 float64 `json:"rate_answer_4"`
}

type QuestionStudentSwitchAnswer struct {
	CorrectToWrong float64 `json:"correct_to_wrong"`
	WrongToCorrect float64 `json:"wrong_to_correct"`
}
