-- Copyright 2017. All rights reserved.
-- Use of this source code is governed by a BSD 3-Clause License
-- license that can be found in the LICENSE file.

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
