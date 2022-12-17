CREATE TABLE IF NOT EXISTS customers (
  id serial primary key,
  first_name varchar(255) not null,
  last_name varchar(255),
  gender varchar(20) not null,
  email varchar(255) not null,
  password varchar(255) not null,
  created_at timestamp with time zone default timezone('utc'::text, now()) not null,
  updated_at timestamp with time zone default timezone('utc'::text, now()) not null
);
