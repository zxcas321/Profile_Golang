CREATE TABLE social_media (
    id UUID PRIMARY KEY,
    profile_id UUID REFERENCES profiles(id) ON DELETE CASCADE,
    platform VARCHAR(100),
    url TEXT
);