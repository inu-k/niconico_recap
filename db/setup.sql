drop table if exists history;
drop table if exists video_basic_info;
drop table if exists video_tag_info;

create table history (
  video_id varchar(255) not null,
  watch_date timestamp not null.
  primary key (video_id, watch_date)
);

create table video_basic_info (
  video_id varchar(255) primary key not null,
  title varchar(255) not null,
  thumbnail_url varchar(255) not null
);

create table video_tag_info (
  video_id varchar(255) not null,
  tag varchar(255) not null,
  primary key (video_id, tag)
);
