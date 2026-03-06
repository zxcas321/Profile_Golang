CREATE TABLE educations (
    id UUID PRIMARY KEY,
    profile_id UUID REFERENCES profiles(id) ON DELETE CASCADE,
    school_name VARCHAR(150),
    degree VARCHAR(150),
    field VARCHAR(150),
    start_year INT,
    end_year INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);