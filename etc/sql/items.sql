begin;
drop schema if exists viitems cascade;
create schema viitems
  create table viitems.items (
    id uuid not null primary key,
    kind varchar not null,
    name varchar not null
  )
  create table viitems.folders (
    id uuid not null primary key
  )
  create table viitems.reports (
    id uuid not null primary key,
    content varchar not null default ''
  );
commit;
