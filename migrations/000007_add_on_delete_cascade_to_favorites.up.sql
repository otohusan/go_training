-- 既存の外部キー制約を削除
ALTER TABLE favorites DROP CONSTRAINT IF EXISTS favorites_user_id_fkey;
ALTER TABLE favorites DROP CONSTRAINT IF EXISTS favorites_study_set_id_fkey;

-- 新しい外部キー制約を追加
ALTER TABLE favorites
ADD CONSTRAINT favorites_user_id_fkey
FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;

ALTER TABLE favorites
ADD CONSTRAINT favorites_study_set_id_fkey
FOREIGN KEY (study_set_id) REFERENCES study_sets(id) ON DELETE CASCADE;
