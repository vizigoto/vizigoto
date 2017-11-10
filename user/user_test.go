package user_test

import (
	"testing"

	"github.com/vizigoto/vizigoto/user"
)

func TestNew(t *testing.T) {
	name := "Maria"
	u := user.New(name)

	if u.Name != name {
		t.Fatal("name error")
	}
}
