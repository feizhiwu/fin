package middleware

import (
	"fin"
)

func Init(c *fin.Context) {
	//Cors Cross-Origin Resource Share
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	c.Header("Access-Control-Allow-Headers", "X-Requested-With,X_Requested_With,content-type,token,action")
}
