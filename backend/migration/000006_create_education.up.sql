CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE educations (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    profile_id UUID REFERENCES profiles(id) ON DELETE CASCADE,
    school_name VARCHAR(150),
    degree VARCHAR(150),
    field VARCHAR(150),
    start_year INT,
    end_year INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);