### Backend service
- [x] POST /register
  - reqBody: email & password
- [x] POST /auth
  - reqBody: email & password
  - return JWT
- [x] POST /create-habit add new habit item
	- reqBody: title, icon, color, daily_goal, weekly_goal, start date
- [ ] GET data (main screen) (all item)
  - [ ] all habits the user has, also
	- [ ] last 7 days tracker for each habit
- [ ] GET detail habit item (all data)
	- [ ] all checklist (calendar & activity)
	- [ ] processed data (graphs)
- [ ] GET basic info item (for add and edit)
	- name, icon, color, frequencies info, start date
- [ ] PUT edit habit item
	- name, icon, color, frequencies info, start date
- [ ] POST checklist (each item)
	- each item -> date (today)
- [ ] PUT checklist (each item)
	- each item -> date (previous date)


```sql
-- Users Table
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP
);

-- Habit Table
CREATE TABLE habits (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    icon VARCHAR(255),
    color VARCHAR(10),  -- You can use HEX color values
    is_active BOOLEAN DEFAULT TRUE,
    start_date DATE,
    daily_goal INTEGER,
    weekly_goal INTEGER,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP
);

-- Tracker Table
CREATE TABLE tracker (
    date DATE,
    habit_id UUID REFERENCES habit(id) ON DELETE CASCADE,
    count INTEGER,
    updated_at TIMESTAMP,
    PRIMARY KEY(date, habit_id)  -- This makes the combination of date and habit_id unique
);
```