package use_case

import (
  _interface "wrk-connector/src/product/application/use-case/interface"
  "wrk-connector/src/product/domain/entity"
  "wrk-connector/src/shared/domain/repository"
  "wrk-connector/src/shared/domain/service"
  log "wrk-connector/src/shared/infrastructure/config"
)

type ProductSendToCreateQueue struct {
  productSendToCreate _interface.ProductSendToCreate
  queue               repository.QueueProducer
}

func NewProductSendToCreateQueue(productSendToCreate _interface.ProductSendToCreate, queue repository.QueueProducer) *ProductSendToCreateQueue {
  return &ProductSendToCreateQueue{
    productSendToCreate: productSendToCreate,
    queue:               queue,
  }
}

func (u *ProductSendToCreateQueue) Do() ([]*entity.Product, error) {
  productsFromUseCase, err := u.productSendToCreate.Do()
  if err != nil {
    log.WithError(err).Error("productsFromUseCase not executing")
    return nil, err
  }

  _ = service.EntityToJson(productsFromUseCase)

 /*var (
    maxLen, _ = strconv.Atoi(os.Getenv("KAFKA_PAYLOAD_MESSAGE_MAX_LENGTH"))
    chunks    = int(float64(len(products)/maxLen)) + 1
    counter   = 0
    message   []*entity.Product
  )*/

 // log.Info("maxLen: %d, chunks: %d", maxLen, chunks)

  //for key, value := range productsFromUseCase {
    //message = append(message, value)
    //if math.Abs(float64(counter-key)) == (float64(maxLen)) {
      //m := service.MessageGenerator(
        //constant.EventName["PRODUCTS_CREATED"],
        //constant.EventDataFormat["JSON"],
        //constant.EventType["CREATE"],
        //"1.0",
        //constant.EventOrigin,
      //)
      //m.Payload = map[string]interface{}{
        //"products": service.EntityToJson(message),
      //}

      //log.Debug("queue message: %s", m)

      //err = u.queue.Send(constant.KafkaTopics["TOPIC_PRODUCTS_CREATE"], m)
      //if err != nil {
        //log.WithError(err).Error("error sending message")

        // os.Exit(0)
        //return nil, err
      //}
    //}

  return nil, nil
}
