CREATE TYPE user_type AS ENUM ('customer', 'chef');

CREATE TABLE users (
    id UUID PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    full_name VARCHAR(100) NOT NULL,
    user_type user_type DEFAULT 'customer' NOT NULL, 
    address TEXT,
    phone_number VARCHAR(20) not null,
    bio TEXT,
    specialties TEXT[], 
    years_of_experience INTEGER CHECK (years_of_experience >= 0),
    is_verified BOOLEAN DEFAULT false,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE user_preferences (
    user_id UUID REFERENCES users(id),
    cuisine_type VARCHAR(50),
    dietary_preferences TEXT[],
    favorite_kitchen_ids UUID[],
    PRIMARY KEY (user_id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE refresh_tokens (
    id uuid DEFAULT gen_random_uuid() not null,
    user_id uuid REFERENCES users(id) not null,
    token text UNIQUE not null,
    expires_at bigint not null,
    revoked boolean DEFAULT false,
    created_at TIMESTAMP default CURRENT_TIMESTAMP not null,
    deleted_at TIMESTAMP
);

CREATE INDEX idx_refresh_tokens_user_id ON refresh_tokens(user_id);
CREATE INDEX idx_refresh_tokens_token ON refresh_tokens(token);
