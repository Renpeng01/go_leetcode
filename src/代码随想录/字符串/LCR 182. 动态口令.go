package main

func dynamicPassword(password string, target int) string {
	return password[target:] + password[:target]
}
