begin;
drop schema if exists vinodes cascade;
create schema vinodes
  create table vinodes.nodes (
    id uuid not null primary key,
    parent uuid references nodes,
    lft integer not null default 0,
    rgt integer not null default 1
  );
commit;
