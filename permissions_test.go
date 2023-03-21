package permie

import "testing"

func TestDefine(t *testing.T) {
	Define("test.$1", Permissions{"users.create": true, "users.populate": false})
	if !Can("test.$1", "users.create") {
		t.Fatal("test.$1 which should have users.create permission, does not.")
	}
	if Can("test.$1", "users.populate") {
		t.Fatal("test.$1 which shouldn't have users.populate permission, does.")
	}
}

func TestCan(t *testing.T) {
	Define("test.$2", Permissions{"users.create": false, "users.populate": true})
	if Can("test.$2", "users.create") {
		t.Fatal("test.$2 which shouldn't have users.create permission, does.")
	}
	if !Can("test.$2", "users.populate") {
		t.Fatal("test.$2 which should have users.populate permission, does not.")
	}
}

func TestChange(t *testing.T) {
	Define("test.$3", Permissions{"users.create": false, "users.populate": true})
	Change("test.$3", Permissions{"users.create": true, "users.ban": false})
	if !Can("test.$3", "users.create") {
		t.Fatal("test.$3 which should have users.create permission, does not.")
	}
	if Can("test.$3", "users.ban") {
		t.Fatal("test.$3 which should not have users.ban permission, does.")
	}
}

func TestAllow(t *testing.T) {
	Define("test.$4", Permissions{"users.create": false})
	Allow("test.$4", "users.create")
	if !Can("test.$4", "users.create") {
		t.Fatal("test.$4 which should have users.create permission, does not.")
	}
}

func TestDisallow(t *testing.T) {
	Define("test.$5", Permissions{"users.create": true})
	Disallow("test.$5", "users.create")
	if Can("test.$5", "users.create") {
		t.Fatal("test.$4 which shouldn't have users.create permission, does.")
	}
}

func TestRequire(t *testing.T) {
	Define("test.$6", Permissions{"users.create": true, "users.register": false})
	Define("test.$7", Permissions{"users.create": false, "users.register": false})
	if !Require([]string{"test.$6", "test.$7"}, "users.create") {
		t.Fatal("test.$6 which should have users.create permission, does not.")
	}
	if Require([]string{"test.$6", "test.$7"}, "users.register") {
		t.Fatal("test.$6 and test.$7 which shouldn't have users.register permission, does.")
	}
}

func TestCanAtLeast(t *testing.T) {
	Define("test.$6", Permissions{"users.create": true, "users.register": false})
	Define("test.$7", Permissions{"users.create": false, "users.register": false})
	if !CanAtLeast([]string{"test.$6", "test.$7"}, "users.create") {
		t.Fatal("test.$6 which should have users.create permission, does not.")
	}
	if CanAtLeast([]string{"test.$6", "test.$7"}, "users.register") {
		t.Fatal("test.$6 and test.$7 which shouldn't have users.register permission, does.")
	}
}
