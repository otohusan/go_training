package studyset

import (
	"database/sql"
	"errors"
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
	query := `SELECT id, user_id, title, description, created_at, updated_at 
			  FROM study_sets 
			  WHERE id = $1`
	row := r.db.QueryRow(query, id)

	studySet := &model.StudySet{}

	err := row.Scan(&studySet.ID, &studySet.UserID, &studySet.Title, &studySet.Description, &studySet.CreatedAt, &studySet.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("study set not found")
		}
		return nil, err
	}
	return studySet, nil
}

// ユーザが作成したすべての学習セットを配列に詰めて返す
func (r *StudySetRepository) GetByUserID(userID string) ([]*model.StudySet, error) {
	query := `SELECT id, user_id, title, description, created_at, updated_at FROM study_sets WHERE user_id = $1`
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var studySets []*model.StudySet

	// クエリをそれぞれstudySetの型にして、studySetsに詰めてる
	for rows.Next() {
		studySet := &model.StudySet{}
		err := rows.Scan(&studySet.ID, &studySet.UserID, &studySet.Title, &studySet.Description, &studySet.CreatedAt, &studySet.UpdatedAt)
		if err != nil {
			return nil, err
		}
		studySets = append(studySets, studySet)
	}

	// エラーなければ返す
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return studySets, nil
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
