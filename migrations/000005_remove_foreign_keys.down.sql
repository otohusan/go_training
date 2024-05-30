-- 外部キー制約を削除
ALTER TABLE study_sets DROP CONSTRAINT IF EXISTS fk_user;
ALTER TABLE flashcards DROP CONSTRAINT IF EXISTS fk_study_set;
