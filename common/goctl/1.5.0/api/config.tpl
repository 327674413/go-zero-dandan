package config

import {{.authImport}}

type Config struct {
	rest.RestConf
	DB struct {
        DataSource string
    }
    I18n struct {
        Default string
        Langs   []string
    }
	{{.auth}}
	{{.jwtTrans}}
}