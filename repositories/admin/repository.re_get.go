package admin

import (
	"github.com/gomodule/redigo/redis"

	configs "github.com/UPC-Works/api-aliviate/configs"
)

func Re_GetId(intput_id string) (string, error) {

	id_admin, err_do := redis.String(configs.Conn_Re_DB().Get().Do("GET", intput_id))
	if err_do != nil {
		return id_admin, err_do
	}

	return id_admin, nil
}
