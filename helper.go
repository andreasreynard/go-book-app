package main

import "strings"

func ValidateUserInput(f string, l string, e string, u uint, n uint) (bool, bool, bool) {
	vn := len(f) >= 2 && len(l) >= 2
	ve := strings.Contains(e, "@")
	vt := u > 0 && u <= n
	return vn, ve, vt
}
