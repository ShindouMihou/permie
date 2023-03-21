package permie

func setOrDefineList(role string, permissions []string, value bool) {
	if _, exists := roles[role]; !exists {
		perms := Permissions{}
		for _, v := range permissions {
			perms[v] = value
		}
		Define(role, perms)
		return
	}
	for _, v := range permissions {
		roles[role][v] = value
	}
}

func setOrDefineMap(role string, permissions Permissions) {
	if _, exists := roles[role]; !exists {
		Define(role, permissions)
		return
	}
	for k, v := range permissions {
		roles[role][k] = v
	}
}

var roles = make(map[string]Permissions)
