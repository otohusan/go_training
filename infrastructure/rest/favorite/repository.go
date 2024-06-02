package favorite

import (
	"database/sql"
	"errors"
	"go-training/domain/model"
)

type FavoriteRepository struct {
	db *sql.DB
}

func NewFavoriteRepository(db *sql.DB) *FavoriteRepository {
	return &FavoriteRepository{db: db}
}

func (r *FavoriteRepository) AddFavorite(userID, studySetID string) error {
	query := `INSERT INTO favorites (user_id, study_set_id) VALUES ($1, $2)`
	_, err := r.db.Exec(query, userID, studySetID)
	if err != nil {
		return err
	}

	return nil
}

func (r *FavoriteRepository) RemoveFavorite(userID, studySetID string) error {
	query := `DELETE FROM favorites WHERE user_id = $1 AND study_set_id = $2`
	result, err := r.db.Exec(query, userID, studySetID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("favorite not found")
	}

	return nil
}

func (r *FavoriteRepository) GetFavoritesByUserID(userID string) ([]*model.Favorite, error) {
	// 使い所が見当たらないから、現段階では実装しない
	// GetFavoriteStudySetsByUserIDによって、studysetを取得して
	// frontでローカルストレージにでもそのstudysetIDを保存しとけば代用できると思う
	return nil, nil
}

func (r *FavoriteRepository) IsFavorite(userID, studySetID string) (bool, error) {
	// 使い所が見当たらないから、現段階では実装しない
	// GetFavoriteStudySetsByUserIDによって、studysetを取得して
	// frontでローカルストレージにでもそのstudysetIDを保存しとけば代用できると思う
	return false, nil
}

func (r *FavoriteRepository) GetFavoriteStudySetsByUserID(userID string) ([]*model.StudySet, error) {
	query := `
		SELECT s.id, s.user_id, s.title, s.description, s.created_at, s.updated_at
		FROM study_sets s
		JOIN favorites f ON s.id = f.study_set_id
		WHERE f.user_id = $1
	`
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var studySets []*model.StudySet
	for rows.Next() {
		studySet := &model.StudySet{}
		err := rows.Scan(&studySet.ID, &studySet.UserID, &studySet.Title, &studySet.Description, &studySet.CreatedAt, &studySet.UpdatedAt)
		if err != nil {
			return nil, err
		}
		studySets = append(studySets, studySet)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return studySets, nil
}
