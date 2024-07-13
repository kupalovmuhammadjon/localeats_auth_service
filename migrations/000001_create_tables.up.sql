CREATE TYPE user_type AS ENUM ('customer', 'chef');

CREATE TABLE users (
    id UUID PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    full_name VARCHAR(100) NOT NULL,
    user_type user_type DEFAULT 'customer' NOT NULL, 
    address TEXT,
    phone_number VARCHAR(20),
    bio TEXT,
    specialties TEXT[], 
    years_of_experience INTEGER CHECK (years_of_experience >= 0),
    is_verified BOOLEAN DEFAULT false,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE user_preferences (
    user_id UUID REFERENCES users(id),
    cuisine_type VARCHAR(50),
    dietary_preferences TEXT[],
    favorite_kitchen_ids UUID[],
    PRIMARY KEY (user_id)
);
