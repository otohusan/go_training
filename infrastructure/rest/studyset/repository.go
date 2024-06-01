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
	query := `SELECT id, user_id, title, description, created_at, updated_at 
			  FROM study_sets 
			  WHERE user_id = $1`
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
	query := `UPDATE study_sets SET title = $1, description = $2 
			  WHERE id = $3 AND user_id = $4`
	result, err := r.db.Exec(query, studySet.Title, studySet.Description, studySetID, authUserID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("not authorized to update study set or study set not found")
	}
	return nil
}

func (r *StudySetRepository) Delete(authUserID, studySetID string) error {
	query := `DELETE FROM study_sets 
			  WHERE id = $1 AND user_id = $2`
	result, err := r.db.Exec(query, studySetID, authUserID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("not authorized to delete study set or study set not found")
	}
	return nil
}

// タイトルが合う学習セットを最大5件送信
func (r *StudySetRepository) SearchByTitle(title string) ([]*model.StudySet, error) {
	query := `SELECT id, user_id, title, description, created_at, updated_at 
			  FROM study_sets 
			  WHERE LOWER(title) LIKE '%' || LOWER($1) || '%' 
			  LIMIT 5`
	rows, err := r.db.Query(query, title)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []*model.StudySet

	// 返信にクエリの結果を詰めてる
	for rows.Next() {
		studySet := &model.StudySet{}
		err := rows.Scan(&studySet.ID, &studySet.UserID, &studySet.Title, &studySet.Description, &studySet.CreatedAt, &studySet.UpdatedAt)
		if err != nil {
			return nil, err
		}
		results = append(results, studySet)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return results, nil
}
