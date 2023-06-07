package config

import {{.authImport}}

type Config struct {
	rest.RestConf
	DB struct {
        DataSource string
    }
    Env string
	{{.auth}}
	{{.jwtTrans}}
}
var Conf Config