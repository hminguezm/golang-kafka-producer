package use_case

import (
	"wrk-connector/src/register/domain/repository"
	"wrk-connector/src/register/infrastructure/persistence/postgres/model"
	log "wrk-connector/src/shared/infrastructure/config"
)

type SendToCreate struct {
	registerRepository repository.RegisterPersistence
}

func NewSendToCreate(registerRepository repository.RegisterPersistence) *SendToCreate {
	return &SendToCreate{
		registerRepository: registerRepository,
	}
}

func (u *SendToCreate) Do() error {
	log.Debug("SendToCreate!")

	var register *model.Register
	register, err := u.registerRepository.GetLastRegister()
	if err != nil {
		return err
	}
	log.Debug("create_at: %s", register.CreatedAt)

	// err := u.registerRepository.CreateRegister()
	// if err != nil {
	//   return err
	// }

	return nil
}
