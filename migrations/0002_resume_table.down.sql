-- Drop indexes
DROP INDEX IF EXISTS idx_resume_analysis_created_at;
DROP INDEX IF EXISTS idx_resume_analysis_resume_id;
DROP INDEX IF EXISTS idx_user_resume_user_id;

DROP TABLE IF EXISTS resume_analysis;
DROP TABLE IF EXISTS user_resume;

DROP EXTENSION IF EXISTS "uuid-ossp";
