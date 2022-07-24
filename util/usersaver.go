package util

// Mp Save login user
var Mp map[string]string

func init() {
	Mp = make(map[string]string, 20)
}

// Set Save or Update token
func Set(name string, token string) {
	Mp[name] = token
}

// Get Get the login information
func Get(name string) (string, bool) {
	v, ok := Mp[name]
	return v, ok
}
