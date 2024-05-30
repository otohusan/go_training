-- 外部キー制約を再追加
ALTER TABLE study_sets ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id);
ALTER TABLE flashcards ADD CONSTRAINT fk_study_set FOREIGN KEY (study_set_id) REFERENCES study_sets(id);
