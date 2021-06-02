package use_case

import (
  "os"
  repoP "wrk-connector/src/product/domain/repository"
  repoR "wrk-connector/src/register/domain/repository"
  "wrk-connector/src/shared/domain/repository"
  log "wrk-connector/src/shared/infrastructure/config"
)

type SendToCreate struct {
  registerRepository repoR.Register
  productRepository  repoP.Product
  producerMessage    repository.QueueProducer
}

func NewSendToCreate(
  registerRepository repoR.Register,
  productRepository repoP.Product,
  producerMessage    repository.QueueProducer,
) *SendToCreate {
  return &SendToCreate{
    registerRepository: registerRepository,
    productRepository:  productRepository,
    producerMessage: producerMessage,
  }
}

func (u *SendToCreate) Do() error {
  log.Debug("SendToCreate!")
  //err := u.registerRepository.CreateRegister()
  //if err != nil {
  //  return err
  //}

  //var register *dto.RegisterGetLastedDTO
  //register, err = u.registerRepository.GetLastRegister()
  //if err != nil {
    //return err
   //}
  //log.Debug("create_at: %s", register.CreatedAt)


  //dateExec := "2021-05-26T12:14:06"

  //productsDB, err := u.productRepository.FindToCreate(register.CreatedAt)
  //if nil != err {
    //log.WithError(err).Fatal("error could not get product from db")
    //os.Exit(0)
  //}

  //if len(productsDB) == 0 {
    //log.Info("Could not found products")
   //os.Exit(0)
  //}

  // log.Error("productsDB: ",service.EntityToJson(productsDB))

  //message := service.MessageGenerator(
    //constant.EventName["PRODUCTS_CREATED"],
    //constant.EventDataFormat["JSON"],
    //constant.EventType["CREATE"],
    //"1.0",
    //constant.EventOrigin,
  //)
  //message.Payload = map[string]interface{}{
    //"products": "service.EntityToJson(productsDB)",
  //}

  //log.Debug("message", message)

  //err := u.producerMessage.Send(constant.KafkaTopics["TOPIC_PRODUCTS_CREATE"], message)
  //if err != nil {
    //log.WithError(err).Error("error sending message")

    //os.Exit(0)
    //return err
  //}

  os.Exit(0)
  return nil
}
