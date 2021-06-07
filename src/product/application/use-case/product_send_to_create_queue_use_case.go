package use_case

import (
  "math"
  "os"
  "strconv"
  _interface "wrk-connector/src/product/application/use-case/interface"
  "wrk-connector/src/product/domain/entity"
  "wrk-connector/src/shared/domain/constant"
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


var (
  maxLen, _ = strconv.Atoi(os.Getenv("KAFKA_PAYLOAD_MESSAGE_MAX_LENGTH"))
  message   []*entity.Product
  counter   = 0
)


func (u *ProductSendToCreateQueue) Do() ([]*entity.Product, error) {
  productsFromUseCase, err := u.productSendToCreate.Do()
  if err != nil {
    log.WithError(err).Error("productsFromUseCase not executing")
    return nil, err
  }

  products := service.EntityToJson(productsFromUseCase)

  var chunks = int(float64(len(products)/maxLen)) + 1
  var chunkPos = 1
  log.Info("maxLen: %d, chunks: %d", maxLen, chunks)

  for key, value := range productsFromUseCase {
    message = append(message, value)
    if math.Abs(float64(counter-key)) == (float64(maxLen)) {
      m := service.MessageGenerator(
        constant.EventName["PRODUCTS_CREATED"],
        constant.EventDataFormat["JSON"],
        constant.EventType["CREATE"],
        "1.0",
        constant.EventOrigin,
      )
      m.Payload = map[string]interface{}{
        "products": message,
      }

      err = u.queue.Send(constant.KafkaTopics["TOPIC_PRODUCTS_CREATE"], m)
      if err != nil {
        log.WithError(err).Error("error sending message")

        os.Exit(0)
        return nil, err
      }
      chunkPos++
      counter = key
      message = nil
    }
    if chunkPos == chunks && key+1 == len(productsFromUseCase) {
      m := service.MessageGenerator(
        constant.EventName["PRODUCTS_CREATED"],
        constant.EventDataFormat["JSON"],
        constant.EventType["CREATE"],
        "1.0",
        constant.EventOrigin,
      )
      m.Payload = map[string]interface{}{
        "products": message,
      }

      err = u.queue.Send(constant.KafkaTopics["TOPIC_PRODUCTS_CREATE"], m)
      if err != nil {
        log.WithError(err).Error("error sending message")

        os.Exit(0)
        return nil, err
      }
      chunkPos++
      counter = key
      message = nil
    }
  }

    return nil, nil
  }
