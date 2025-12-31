package auth

// private function, not accessible outside this package
func extractSessionID() string {
	return "session_12345"
}

func GetSession() string {
	return extractSessionID()
}
