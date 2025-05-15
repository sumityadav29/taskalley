


-- Enum type for task status
CREATE TYPE task_status AS ENUM ('TODO', 'IN_PROGRESS', 'COMPLETED');

-- Projects table
CREATE TABLE projects (
                          id SERIAL PRIMARY KEY,
                          name TEXT NOT NULL,
                          description TEXT,
                          created_by INTEGER NOT NULL,
                          created_at TIMESTAMP NOT NULL DEFAULT NOW(),
                          updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Tasks table
CREATE TABLE tasks (
                       id SERIAL PRIMARY KEY,
                       project_id INTEGER NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
                       title TEXT NOT NULL,
                       description TEXT,
                       status task_status NOT NULL DEFAULT 'TODO',
                       created_by INTEGER NOT NULL,
                       due_date TIMESTAMP,
                       created_at TIMESTAMP NOT NULL DEFAULT NOW(),
                       updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
