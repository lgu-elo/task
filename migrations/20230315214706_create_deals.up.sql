CREATE TABLE tasks (
  id serial primary key,
  description text,
  project_id int not null,
  user_id int not null,
  status varchar(50)
);