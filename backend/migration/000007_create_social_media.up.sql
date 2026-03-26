CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE social_media (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    profile_id UUID REFERENCES profiles(id) ON DELETE CASCADE,
    platform VARCHAR(100),
    url TEXT
);