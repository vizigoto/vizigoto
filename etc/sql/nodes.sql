begin;

drop schema if exists vinodes cascade;
create schema vinodes
  create table vinodes.nodes (
      id        uuid not null primary key,
      parent    uuid references nodes,
      lft       integer not null default 0,
      rgt       integer not null default 1,
      name      varchar not null,
      kind      varchar not null,
      owner     varchar not null,
      created   timestamp with time zone not null default current_timestamp,
      updated   timestamp with time zone not null default current_timestamp,
      protected boolean not null default 'false'
  );

alter table vinodes.nodes OWNER TO vizi;
alter schema vinodes OWNER TO vizi;

commit;

insert into vinodes.nodes(id, parent_id, name, kind, owner_id, protected) values ('a', null, "x", "folder", "xx", 't')
