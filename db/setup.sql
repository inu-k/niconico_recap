drop table if exists history;

create table history (
  id serial primary key,
  video_url varchar(255) not null,
  title varchar(255) not null,
  watch_date timestamp not null
);