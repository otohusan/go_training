-- 既存の外部キー制約を削除
ALTER TABLE flashcards DROP CONSTRAINT IF EXISTS flashcards_study_set_id_fkey;

-- 元の外部キー制約を追加
ALTER TABLE flashcards
ADD CONSTRAINT flashcards_study_set_id_fkey
FOREIGN KEY (study_set_id) REFERENCES study_sets(id);
