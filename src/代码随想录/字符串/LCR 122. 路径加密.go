package main

func pathEncryption(path string) string {
	newPath := make([]byte, 0, len(path))

	for i := 0; i <= len(path); i++ {
		if path[i] == '.' {
			newPath = append(newPath, ' ')
			continue
		}
		newPath = append(newPath, path[i])

	}
	return string(newPath)
}
