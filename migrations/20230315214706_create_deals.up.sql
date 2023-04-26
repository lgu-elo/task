CREATE TABLE tasks (
  id serial primary key,
  description text,
  project_id int not null,
  amount int not null,
  client_name varchar(255) not null,
  user_id int not null
);