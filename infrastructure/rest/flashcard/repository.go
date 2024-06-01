package flashcard

import (
	"database/sql"
	"errors"
	"go-training/domain/model"
)

type FlashcardRepository struct {
	db *sql.DB
}

func NewFlashcardRepository(db *sql.DB) *FlashcardRepository {
	return &FlashcardRepository{db: db}
}

func (r *FlashcardRepository) Create(authUserID string, flashcard *model.Flashcard) error {
	// フラッシュカードの作成クエリ
	// flashcardが加えられるstudySetにあるuserIDと、
	// リクエストしたuserのidが等しい場合のみflashcardを作成
	query := `
		INSERT INTO flashcards (study_set_id, question, answer)
		SELECT $1, $2, $3
		WHERE EXISTS (
			SELECT 1
			FROM study_sets
			WHERE id = $1 AND user_id = $4
		)
	`

	// クエリの実行
	result, err := r.db.Exec(query, flashcard.StudySetID, flashcard.Question, flashcard.Answer, authUserID)
	if err != nil {
		return err
	}

	// 挿入が成功したか確認
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("not authorized to create flashcard or study set not found")
	}

	return nil
}

func (r *FlashcardRepository) GetByID(id string) (*model.Flashcard, error) {
	query := `SELECT id, study_set_id, question, answer
			  FROM flashcards
			  WHERE id = $1`
	row := r.db.QueryRow(query, id)

	flashcard := &model.Flashcard{}

	err := row.Scan(&flashcard.ID, &flashcard.StudySetID, &flashcard.Question, &flashcard.Answer)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("flashcard not found")
		}
		return nil, err
	}
	return flashcard, nil
}

func (r *FlashcardRepository) GetByStudySetID(studySetID string) ([]*model.Flashcard, error) {
	query := `SELECT id, study_set_id, question, answer
			  FROM flashcards
			  WHERE study_set_id = $1`

	// クエリ実行
	rows, err := r.db.Query(query, studySetID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// 返り値
	var flashcards []*model.Flashcard

	// 結果を詰めていく
	for rows.Next() {
		flashcard := &model.Flashcard{}
		err := rows.Scan(&flashcard.ID, &flashcard.StudySetID, &flashcard.Question, &flashcard.Answer)
		if err != nil {
			return nil, err
		}
		flashcards = append(flashcards, flashcard)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return flashcards, nil
}

func (r *FlashcardRepository) Update(authUserID string, flashcard *model.Flashcard) error {
	return nil
}

func (r *FlashcardRepository) Delete(authUserID, studySetID, flashcardID string) error {
	return nil
}
