-- Copyright 2017. All rights reserved.
-- Use of this source code is governed by a BSD 3-Clause License
-- license that can be found in the LICENSE file.

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
