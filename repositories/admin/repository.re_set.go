package admin

import (
	configs "github.com/UPC-Works/api-aliviate/configs"
)

func Re_SetId(input_id string) error {

	_, err_do := configs.Conn_Re_DB().Get().Do("SET", input_id, input_id, "EX", 7776000)
	if err_do != nil {
		return err_do
	}

	return nil
}
