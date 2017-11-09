begin;
drop schema if exists viusers cascade;
create schema viusers
  create table viusers.users (
    id uuid not null primary key,
    username varchar not null
  );
commit;
