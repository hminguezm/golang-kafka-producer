package use_case

import (
	repoP "wrk-connector/src/product/domain/repository"

	log "wrk-connector/src/shared/infrastructure/config"
)

type SendToCreate struct {
	// registerRepository repoR.Register
	productRepository repoP.Product
}

func NewSendToCreate(productRepository repoP.Product) *SendToCreate {
	return &SendToCreate{
		// registerRepository: registerRepository,
		productRepository: productRepository,
	}
}

func (u *SendToCreate) Do() error {
	log.Debug("SendToCreate!")

	// var register *model.Register
	// register, err := u.registerRepository.GetLastRegister()
	// if err != nil {
	// return err
	// }
	// log.Debug("create_at: %s", register.CreatedAt)

	_, _ = u.productRepository.FindAll()

	// err := u.registerRepository.CreateRegister()
	// if err != nil {
	//   return err
	// }

	return nil
}
