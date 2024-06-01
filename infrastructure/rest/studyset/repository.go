package studyset

import (
	"database/sql"
	"go-training/domain/model"
)

type StudySetRepository struct {
	db *sql.DB
}

func NewStudySetRepository(db *sql.DB) *StudySetRepository {
	return &StudySetRepository{db: db}
}

func (r *StudySetRepository) Create(studySet *model.StudySet) error {
	query := `INSERT INTO study_sets (user_id, title, description) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(query, studySet.UserID, studySet.Title, studySet.Description)
	if err != nil {
		return err
	}
	return nil
}

func (r *StudySetRepository) GetByID(id string) (*model.StudySet, error) {
	return nil, nil
}

func (r *StudySetRepository) GetByUserID(userID string) ([]*model.StudySet, error) {
	return nil, nil
}

func (r *StudySetRepository) Update(authUserID, studySetID string, studySet *model.StudySet) error {
	return nil
}

func (r *StudySetRepository) Delete(authUserID, studySetID string) error {
	return nil
}
func (r *StudySetRepository) SearchByTitle(title string) ([]*model.StudySet, error) {
	return nil, nil
}
