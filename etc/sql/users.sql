-- Copyright 2017. All rights reserved.
-- Use of this source code is governed by a BSD 3-Clause License
-- license that can be found in the LICENSE file.

begin;
drop schema if exists viusers cascade;
create schema viusers
  create table viusers.users (
    id uuid not null primary key,
    username varchar not null
  );
commit;
