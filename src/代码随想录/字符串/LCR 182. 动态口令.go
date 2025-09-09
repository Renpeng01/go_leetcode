package main

func dynamicPassword(password string, target int) string {

	newPassword := ""
	return password[target:] + password[:target]
}
