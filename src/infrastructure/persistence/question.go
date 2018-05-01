package persistence

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/hashmup/QuestionBankAPI/src/domain/entity"
	"github.com/hashmup/QuestionBankAPI/src/domain/repository"
	"github.com/jmoiron/sqlx"
)

type questionRepository struct {
	DBClient *sqlx.DB
}

func NewQuestionRepository(dbc *sqlx.DB) repository.QuestionRepository {
	return &questionRepository{
		DBClient: dbc,
	}
}

func (repo *questionRepository) GetQuestionRationales(ctx context.Context, questionID, classID int64) ([]*entity.QuestionRationale, error) {
	sql, args, _ := sq.Select("student_answers.*").From("student_answers").Join("questions ON student_answers.question_id = questions.id").Join("folders ON questions.folder_id = folders.id").Join("classes ON folders.class_id = classes.id").Where(sq.Eq{"questions.id": questionID, "classes.id": classID}).ToSql()
	studentAnswers := []entity.StudentAnswer{}
	err := repo.DBClient.SelectContext(ctx, &studentAnswers, sql, args...)
	if err != nil {
		return nil, err
	}
	questionRationales := []*entity.QuestionRationale{}
	for i := range studentAnswers {
		questionRationales = append(questionRationales, &entity.QuestionRationale{
			Rationale: studentAnswers[i].Rationale,
			AnswerID:  studentAnswers[i].InitialAnswerID,
		})
	}
	return questionRationales, nil
}

func (repo *questionRepository) GetQuestionAnswer(ctx context.Context, userID, questionID int64) (*entity.StudentAnswerResponse, error) {
	sql, args, _ := sq.Select("*").From("student_answers").Where(sq.Eq{"question_id": questionID, "user_id": userID}).ToSql()
	studentAnswer := entity.StudentAnswer{}
	rating := entity.Rating{}
	err := repo.DBClient.GetContext(ctx, &studentAnswer, sql, args...)
	if err != nil {
		return nil, err
	}
	sql, args, _ = sq.Select("*").From("ratings").Where(sq.Eq{"id": studentAnswer.RatingID}).ToSql()
	err = repo.DBClient.GetContext(ctx, &rating, sql, args...)
	if err != nil {
		return nil, err
	}

	return &entity.StudentAnswerResponse{
		Rationale:       studentAnswer.Rationale,
		QuestionID:      studentAnswer.QuestionID,
		UserID:          studentAnswer.UserID,
		Rating:          &rating,
		InitialAnswerID: studentAnswer.InitialAnswerID,
		FinalAnswerID:   studentAnswer.FinalAnswerID,
	}, nil
}

func (repo *questionRepository) PostQuestions(ctx context.Context, userID int64, questionAnswer *entity.QuestionAnswer) error {
	now := time.Now()
	sql := "INSERT INTO ratings (rating_1, rating_2, rating_3, rating_4, created_at, updated_at) VALUES(:rating_1, :rating_2, :rating_3, :rating_4, :created_at, :updated_at)"
	questionAnswer.Rating.CreatedAt = now
	questionAnswer.Rating.UpdatedAt = now

	res, err := repo.DBClient.NamedExecContext(ctx, sql, questionAnswer.Rating)
	if err != nil {
		return err
	}
	ratingID, err := res.LastInsertId()

	studentAnswer := entity.StudentAnswer{
		Rationale:       questionAnswer.Rationale,
		QuestionID:      questionAnswer.QuestionID,
		UserID:          userID,
		RatingID:        ratingID,
		InitialAnswerID: questionAnswer.InitialAnswerID,
		FinalAnswerID:   questionAnswer.FinalAnswerID,
		CreatedAt:       now,
		UpdatedAt:       now,
	}
	sql = "INSERT INTO student_answers (rationale, question_id, user_id, rating_id, initial_answer_id, final_answer_id, created_at, updated_at) VALUES (:rationale, :question_id, :user_id, :rating_id, :initial_answer_id, :final_answer_id, :created_at, :updated_at)"
	_, err = repo.DBClient.NamedExecContext(ctx, sql, studentAnswer)
	if err != nil {
		return err
	}
	return nil
}

func (repo *questionRepository) SearchQuestions(ctx context.Context, name, tag string) ([]*entity.QuestionRequest, error) {
	questions := []entity.Question{}
	var query sq.SelectBuilder
	if tag == "" {
		query = sq.Select("*").From("questions")
	} else {
		query = sq.Select("questions.*").From("questions")
	}

	if tag != "" {
		query = query.Join("question_tag_assoc ON question_tag_assoc.question_id = questions.id").Join("tags ON question_tag_assoc.tag_id = tags.id").Where(sq.Expr("tags.name LIKE ?", "%"+tag+"%"))
		if name != "" {
			query = query.Where(sq.Expr("questions.question LIKE ?", "%"+name+"%"))
		}
	} else {
		if name != "" {
			query = query.Where(sq.Expr("question LIKE ?", "%"+name+"%"))
		}
	}

	sql, args, _ := query.ToSql()
	err := repo.DBClient.SelectContext(ctx, &questions, sql, args...)
	if err != nil {
		return nil, err
	}
	questionRequests := []*entity.QuestionRequest{}
	for i := range questions {
		answers := []entity.Answer{}
		sql, args, _ = sq.Select("*").From("answers").Where(sq.Eq{"id": []int64{questions[i].Answer1, questions[i].Answer2, questions[i].Answer3, questions[i].Answer4}}).ToSql()
		err := repo.DBClient.SelectContext(ctx, &answers, sql, args...)
		if err != nil {
			return nil, err
		}
		tags := []entity.Tag{}
		sql, args, _ = sq.Select("tags.*").From("tags").Join("question_tag_assoc ON question_tag_assoc.tag_id = tags.id").Where(sq.Eq{"question_tag_assoc.question_id": questions[i].ID}).ToSql()
		err = repo.DBClient.SelectContext(ctx, &tags, sql, args...)
		if err != nil {
			return nil, err
		}
		_answers := []*entity.Answer{}
		_tags := []string{}
		for j := range answers {
			_answers = append(_answers, &answers[j])
		}
		for j := range tags {
			_tags = append(_tags, tags[j].Name)
		}

		questionRequests = append(questionRequests, &entity.QuestionRequest{
			QuestionID:      questions[i].ID,
			Question:        questions[i].Question,
			Solved:          false,
			Answers:         _answers,
			Tags:            _tags,
			CorrectAnswerID: questions[i].CorrectAnswerID,
		})
	}
	return questionRequests, nil
}

func (repo *questionRepository) AnalyzeQuestion(ctx context.Context, questionID int64) (*entity.QuestionAnalysis, error) {
	question := entity.Question{}
	rating := entity.Rating{}
	studentAnswers := []entity.StudentAnswer{}
	sql, args, _ := sq.Select("*").From("questions").Where(sq.Eq{"id": questionID}).ToSql()
	err := repo.DBClient.GetContext(ctx, &question, sql, args...)
	if err != nil {
		return nil, err
	}
	sql, args, _ = sq.Select("*").From("student_answers").Where(sq.Eq{"question_id": questionID}).ToSql()
	err = repo.DBClient.SelectContext(ctx, &studentAnswers, sql, args...)
	if err != nil {
		return nil, err
	}
	// answer ratio + switch answer
	answers := []int64{0, 0, 0, 0}
	answerRating := entity.QuestionStudentAnswerRating{
		AverageRating1: 0.0,
		AverageRating2: 0.0,
		AverageRating3: 0.0,
		AverageRating4: 0.0,
	}
	answerSwitch := entity.QuestionStudentSwitchAnswer{
		CorrectToWrong: 0.0,
		WrongToCorrect: 0.0,
	}
	for i := range studentAnswers {
		answers[studentAnswers[i].FinalAnswerID]++
		if question.CorrectAnswerID == studentAnswers[i].InitialAnswerID && question.CorrectAnswerID == studentAnswers[i].FinalAnswerID {
			answerSwitch.CorrectToWrong += 1.0
		}
		if question.CorrectAnswerID == studentAnswers[i].FinalAnswerID && question.CorrectAnswerID == studentAnswers[i].InitialAnswerID {
			answerSwitch.WrongToCorrect += 1.0
		}
		sql, args, _ := sq.Select("*").From("ratings").Where(sq.Eq{"id": studentAnswers[i].RatingID}).ToSql()
		repo.DBClient.GetContext(ctx, &rating, sql, args...)
		answerRating.AverageRating1 += float64(rating.Rating1)
		answerRating.AverageRating2 += float64(rating.Rating2)
		answerRating.AverageRating3 += float64(rating.Rating3)
		answerRating.AverageRating4 += float64(rating.Rating4)
	}
	answerSwitch.CorrectToWrong /= float64(len(studentAnswers))
	answerSwitch.WrongToCorrect /= float64(len(studentAnswers))
	answerRating.AverageRating1 /= float64(len(studentAnswers))
	answerRating.AverageRating2 /= float64(len(studentAnswers))
	answerRating.AverageRating3 /= float64(len(studentAnswers))
	answerRating.AverageRating4 /= float64(len(studentAnswers))
	return &entity.QuestionAnalysis{
		QuestionStudentAnswer: &entity.QuestionStudentAnswer{
			AnswerNum1: answers[0],
			AnswerNum2: answers[1],
			AnswerNum3: answers[2],
			AnswerNum4: answers[3],
		},
		QuestionStudentAnswerRating: &answerRating,
		QuestionStudentSwitchAnswer: &answerSwitch,
	}, nil
}
