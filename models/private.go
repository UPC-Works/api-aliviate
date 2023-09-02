package models

import "time"

//Model private

type ExamenesAuxiliares struct {
	Examen          bool      `json:"isDeleted"`
	Is_sended_to_delete bool      `json:"isSendedToDelete"`
	Sended_to_delete_at time.Time `json:"sendedToDeleteAt"`
	Deleted_at          time.Time `json:"deletedAt"`
}

