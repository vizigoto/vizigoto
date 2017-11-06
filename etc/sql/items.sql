begin;

drop schema if exists viitems cascade;
create schema viitems
  create table viitems.folders (
      id        uuid not null primary key
  )
  create table viitems.reports (
      id        uuid not null primary key,
      content   varchar not null default ''
  );

commit;
