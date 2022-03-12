package middleware

import (
	"fin"
)

// Cors Cross-Origin Resource Share
func Cors(c *fin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	c.Header("Access-Control-Allow-Headers", "X-Requested-With,X_Requested_With,content-type,token,action")
}
