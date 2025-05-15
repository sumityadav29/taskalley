

-- Insert sample projects
INSERT INTO projects (name, description, created_by)
VALUES
  ('Website Redesign', 'Redesign the company website to improve UX.', 1),
  ('Marketing Campaign', 'Launch a new email marketing campaign.', 2);

-- Insert sample tasks
INSERT INTO tasks (project_id, title, description, status, created_by, due_date)
VALUES
  (1, 'Create wireframes', 'Design wireframes for the new homepage.', 'TODO', 1, '2025-06-01 10:00:00'),
  (1, 'Update UI library', 'Replace old components with new ones.', 'IN_PROGRESS', 1, '2025-06-03 17:00:00'),
  (2, 'Write email copy', 'Draft email content for the first campaign.', 'TODO', 2, '2025-06-02 12:00:00'),
  (2, 'Design email template', 'Create a visually engaging template.', 'COMPLETED', 2, '2025-05-30 09:00:00');