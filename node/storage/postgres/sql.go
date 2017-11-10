// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

package postgres

const sqlInsert = "insert into vinodes.nodes(id, parent, lft, rgt) values ($1, $2, $3, $4)"
const sqlPos = "select lft, rgt from vinodes.nodes where id = $1"

const sqlGet = `
select    parent.parent as p_parent,
          child.id as c_id
from      vinodes.nodes parent
left join vinodes.nodes child
on        (parent.id = child.parent)
where parent.id = $1`
