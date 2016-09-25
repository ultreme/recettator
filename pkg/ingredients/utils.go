package ingredients

func beginsWithVoyel(word string) bool {
	switch word[0] {
	case 'a', 'e', 'i', 'o', 'u', 'y':
		return true
	default:
		return false
	}
}
