package vcago

func SliceContains(list []string, element string) bool {
	for _, s := range list {
		if element == s {
			return true
		}
	}
	return false
}
