CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS user_resume (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    file_name TEXT NOT NULL,
    raw_text TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_user_resume_user_id
ON user_resume(user_id);

CREATE TABLE IF NOT EXISTS resume_analysis (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    resume_id UUID NOT NULL REFERENCES user_resume(id) ON DELETE CASCADE,
    job_role TEXT NOT NULL,
    score INT NOT NULL CHECK (score BETWEEN 0 AND 100),
    strengths TEXT NOT NULL,
    weaknesses TEXT NOT NULL,
    suggestions TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_resume_analysis_resume_id
ON resume_analysis(resume_id);

CREATE INDEX IF NOT EXISTS idx_resume_analysis_created_at
ON resume_analysis(created_at);

