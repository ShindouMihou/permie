package permie

type Permissions map[string]bool

// Define SETS the policies or permissions of the given role. This will overwrite
// any existing permissions if it exists, but in exchange, this operation will
// perform in O(1).
func Define(role string, permissions Permissions) {
	roles[role] = permissions
}

// Change modifies the given role's permissions by overwriting existing policies
// and adding new policies if needed based on the permissions map given.
// Not to be confused with Define which overwrites the entire role's permissions,
// and this operation can take up to O(n).
func Change(role string, permissions Permissions) {
	setOrDefineMap(role, permissions)
}

// Allow permits the given role to the given permissions.
func Allow(role string, permissions ...string) {
	setOrDefineList(role, permissions, true)
}

// Disallow prevents the given role to the given permissions.
func Disallow(role string, permissions ...string) {
	setOrDefineList(role, permissions, false)
}

// IsAllowed checks whether the given role is permitted to the given permission.
func IsAllowed(role string, permission string) bool {
	return roles[role][permission]
}

// Can checks whether the given role has permission to perform all the
// given permissions.
func Can(role string, permissions ...string) bool {
	for _, permission := range permissions {
		if !IsAllowed(role, permission) {
			return false
		}
	}
	return true
}

// CanAtLeast checks whether at least one of the roles is allowed to perform
// the permission listed. This is an O(n) operation since it has to check each
// role individually.
func CanAtLeast(roles []string, permission string) bool {
	for _, role := range roles {
		if IsAllowed(role, permission) {
			return true
		}
	}
	return false
}

// Require checks whether at least one of the roles, or all of the roles combined
// has permission to do all the of given permissions.
func Require(roles []string, permissions ...string) bool {
	var permits = make(map[string]bool)
	for _, role := range roles {
		for _, permission := range permissions {
			if _, exists := permits[permission]; exists {
				if IsAllowed(role, permission) {
					permits[permission] = true
				}
			} else {
				permits[permission] = IsAllowed(role, permission)
			}
		}
	}
	for _, permitted := range permits {
		if !permitted {
			return false
		}
	}
	return true
}
