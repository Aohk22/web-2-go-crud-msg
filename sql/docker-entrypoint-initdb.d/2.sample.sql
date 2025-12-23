-- ai generated sample data
-- Users
INSERT INTO users (name, passhash) VALUES
('Alice' , '$2a$10$kVde2ggYcmJTe/HUy/EZoObZWFzo3e94mBBv5bmtaFXY4/ph/AgI2'),
('Bob' , '$2a$10$r6E3Xw/8Er1TBs2SUW7u3eYyK9WG0peK1E6Bjf0Bu54eGM.W8P5eO'),
('Charlie' , '$2a$10$Ab.SFh0FxEbWfNeseZRjiOgum8nFasnQPQKOMsL.yvwyok6u9zPhu'),
('Diana' , '$2a$10$AH6qAOCdJPs.2ZyEdlH8Iu2Ci8IiSb0L7EQIUMYoXwm0zQ8YcQqfa'),
('Eve' , '$2a$10$zpMA6OkDEKlF24ZlKKlsS.gwgfdnnC5SJx22CA6IvYxUeaXr5NySu');
-- password: alice123
-- password: bobsecure
-- password: charlie!
-- password: diana2025
-- assword: eve_hacker

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
('2025-11-25 09:12:22+00', 'Can someone help me with a PostgreSQL query? I keep getting ''relation does not exist', 2, 3),
('2025-11-25 09:14:08+00', 'Did you forget to set the search_path or quote the table name?', 3, 3),
('2025-11-25 09:16:55+00', 'It might be a case-sensitivity issue. Try ''MyTable'' with quotes', 5, 3),
('2025-12-01 11:26:40+00', 'Alright team, sprint starts today. Let''s crush it! 🚀', 1, 5),
('2025-12-01 11:28:15+00', 'Ready when you are. Backend API is deployed.', 2, 5),
('2025-12-01 11:29:02+00', 'Frontend build is passing, just fixing some styling bugs', 4, 5),
('2025-12-02 08:15:33+00', 'Good morning! Who''s doing standup first?', 1, 5);
