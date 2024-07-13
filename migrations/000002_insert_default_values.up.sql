INSERT INTO users (id, username, email, password_hash, full_name, user_type, address, phone_number, bio, specialties, years_of_experience, is_verified, created_at, updated_at)
VALUES
    (uuid_generate_v4(), 'johndoe', 'john.doe@example.com', 'hashed_password_1', 'John Doe', 'customer', '123 Main St', '123-456-7890', 'Bio of John Doe', '{"Italian", "French"}', 5, true, NOW(), NOW()),
    (uuid_generate_v4(), 'janedoe', 'jane.doe@example.com', 'hashed_password_2', 'Jane Doe', 'chef', '456 Elm St', '234-567-8901', 'Bio of Jane Doe', '{"Japanese", "Mexican"}', 10, true, NOW(), NOW()),
    (uuid_generate_v4(), 'bobjones', 'bob.jones@example.com', 'hashed_password_3', 'Bob Jones', 'customer', '789 Oak St', '345-678-9012', 'Bio of Bob Jones', '{}', 0, false, NOW(), NOW()),
    (uuid_generate_v4(), 'alicewilliams', 'alice.williams@example.com', 'hashed_password_4', 'Alice Williams', 'chef', '321 Pine St', '456-789-0123', 'Bio of Alice Williams', '{"Chinese"}', 8, true, NOW(), NOW()),
    (uuid_generate_v4(), 'michaelbrown', 'michael.brown@example.com', 'hashed_password_5', 'Michael Brown', 'customer', '654 Cedar St', '567-890-1234', 'Bio of Michael Brown', '{}', 2, false, NOW(), NOW()),
    (uuid_generate_v4(), 'emilydavis', 'emily.davis@example.com', 'hashed_password_6', 'Emily Davis', 'chef', '987 Spruce St', '678-901-2345', 'Bio of Emily Davis', '{"American"}', 15, true, NOW(), NOW()),
    (uuid_generate_v4(), 'chriswilson', 'chris.wilson@example.com', 'hashed_password_7', 'Chris Wilson', 'customer', '159 Birch St', '789-012-3456', 'Bio of Chris Wilson', '{}', 1, false, NOW(), NOW()),
    (uuid_generate_v4(), 'sarahmartinez', 'sarah.martinez@example.com', 'hashed_password_8', 'Sarah Martinez', 'chef', '753 Walnut St', '890-123-4567', 'Bio of Sarah Martinez', '{"Italian", "American"}', 12, true, NOW(), NOW()),
    (uuid_generate_v4(), 'davidsmith', 'david.smith@example.com', 'hashed_password_9', 'David Smith', 'customer', '852 Maple St', '901-234-5678', 'Bio of David Smith', '{}', 3, false, NOW(), NOW()),
    (uuid_generate_v4(), 'lindajohnson', 'linda.johnson@example.com', 'hashed_password_10', 'Linda Johnson', 'chef', '951 Cherry St', '012-345-6789', 'Bio of Linda Johnson', '{"Mexican", "French"}', 7, true, NOW(), NOW());
INSERT INTO user_preferences (user_id, cuisine_type, dietary_preferences, favorite_kitchen_ids, created_at, updated_at)
VALUES
    ((SELECT id FROM users WHERE username = 'johndoe'), 'Italian', '{"Vegetarian"}', '{"123e4567-e89b-12d3-a456-426614174000"}', NOW(), NOW()),
    ((SELECT id FROM users WHERE username = 'janedoe'), 'Japanese', '{}', '{"123e4567-e89b-12d3-a456-426614174001"}', NOW(), NOW()),
    ((SELECT id FROM users WHERE username = 'bobjones'), 'Mexican', '{"Vegan"}', '{}', NOW(), NOW()),
    ((SELECT id FROM users WHERE username = 'alicewilliams'), 'Chinese', '{}', '{"123e4567-e89b-12d3-a456-426614174002"}', NOW(), NOW()),
    ((SELECT id FROM users WHERE username = 'michaelbrown'), 'French', '{}', '{}', NOW(), NOW()),
    ((SELECT id FROM users WHERE username = 'emilydavis'), 'American', '{"Gluten-Free"}', '{"123e4567-e89b-12d3-a456-426614174003"}', NOW(), NOW()),
    ((SELECT id FROM users WHERE username = 'chriswilson'), 'Italian', '{}', '{}', NOW(), NOW()),
    ((SELECT id FROM users WHERE username = 'sarahmartinez'), 'Mexican', '{"Dairy-Free"}', '{"123e4567-e89b-12d3-a456-426614174004"}', NOW(), NOW()),
    ((SELECT id FROM users WHERE username = 'davidsmith'), 'Japanese', '{}', '{}', NOW(), NOW()),
    ((SELECT id FROM users WHERE username = 'lindajohnson'), 'French', '{"Vegetarian"}', '{"123e4567-e89b-12d3-a456-426614174005"}', NOW(), NOW());
