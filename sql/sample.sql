-- ai generated sample data
-- Users
INSERT INTO users (name, passhash) VALUES
('Alice', '$2b$12$Kixh5eX5z3fZ8jQ9vN7m5O8pL9kJ7hG5fD3sA1bC9xE7vT5rW2qU6'),  -- password: alice123
('Bob', '$2b$12$LmN7pQ8rT2vX5yZ9aB3cD5fH7jK9mP1qS3tU5wY8zA2cE4gI6kM8o'),    -- password: bobsecure
('Charlie', '$2b$12$XyZ9aB3cD5fH7jK9mP1qS3tU5wY8zA2cE4gI6kM8oP0qR2sT4uV6wX8'), -- password: charlie!
('Diana', '$2b$12$AbC3dE5fG7hI9jK1mN3oP5qR7sT9uV1wX3yZ5aB7cD9eF1gH3iJ5k'),   -- password: diana2025
('Eve', '$2b$12$QwE4rT6yU8iO0pA2sD4fG6hJ8kL0mN2oP4qS6tU8vW0xY2zA4bC6dE8f'); -- password: eve_hacker

-- Rooms
INSERT INTO rooms (time, name) VALUES
('2025-11-15 10:30:00+00', 'general'),
('2025-11-20 14:15:00+00', 'random'),
('2025-11-25 09:00:00+00', 'programming-help'),
('2025-11-28 18:45:00+00', 'off-topic'),
('2025-12-01 11:20:00+00', 'project-xyz');

-- Users joining rooms (user_room_join)
INSERT INTO user_room_join (time, user_id, room_id) VALUES
('2025-11-15 10:35:00+00', 1, 1),  -- Alice joins general
('2025-11-15 10:40:00+00', 2, 1),  -- Bob joins general
('2025-11-15 11:00:00+00', 3, 1),  -- Charlie joins general
('2025-11-20 14:20:00+00', 1, 2),  -- Alice joins random
('2025-11-20 14:25:00+00', 4, 2),  -- Diana joins random
('2025-11-25 09:05:00+00', 2, 3),  -- Bob joins programming-help
('2025-11-25 09:10:00+00', 3, 3),  -- Charlie joins programming-help
('2025-11-25 09:15:00+00', 5, 3),  -- Eve joins programming-help
('2025-12-01 11:25:00+00', 1, 5),  -- Alice joins project-xyz
('2025-12-01 11:30:00+00', 2, 5),  -- Bob joins project-xyz
('2025-12-01 11:35:00+00', 4, 5);  -- Diana joins project-xyz

-- Sample messages
INSERT INTO messages (time, content, user_id, room_id) VALUES
('2025-11-15 10:36:12+00', 'Hey everyone! Just joined the server 👋', 1, 1),
('2025-11-15 10:37:05+00', 'Welcome Alice!', 2, 1),
('2025-11-15 10:38:30+00', 'Hi guys, anyone up for some coding tonight?', 3, 1),
('2025-11-20 14:22:10+00', 'https://i.imgur.com/5zN0vXj.jpeg', 1, 2),
('2025-11-20 14:23:45+00', 'lmao what is this 🤣', 4, 2),
('2025-11-25 09:12:22+00', 'Can someone help me with a PostgreSQL query? I keep getting "relation does not exist"', 2, 3),
('2025-11-25 09:14:08+00', 'Did you forget to set the search_path or quote the table name?', 3, 3),
('2025-11-25 09:16:55+00', 'It might be a case-sensitivity issue. Try "MyTable" with quotes', 5, 3),
('2025-12-01 11:26:40+00', 'Alright team, sprint starts today. Let''s crush it! 🚀', 1, 5),
('2025-12-01 11:28:15+00', 'Ready when you are. Backend API is deployed.', 2, 5),
('2025-12-01 11:29:02+00', 'Frontend build is passing, just fixing some styling bugs', 4, 5),
('2025-12-02 08:15:33+00', 'Good morning! Who''s doing standup first?', 1, 5);
