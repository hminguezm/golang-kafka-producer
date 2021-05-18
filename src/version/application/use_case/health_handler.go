package use_case

import (
  "encoding/json"
  "fmt"
  "github.com/labstack/echo/v4"
  "io/ioutil"
  "net/http"
  "os"
  log "wrk-connector/src/shared/infrastructure/config"
  "wrk-connector/src/version/domain/entity"
)

type healthHandler struct{}

func NewHealthHandler(e *echo.Echo) {
  basePath := os.Getenv("BASE_PATH")
  if "" == basePath {
    log.Fatal("Base path not config")
    return
  }
  versionPath := fmt.Sprintf("%s/version", basePath)
  h := &healthHandler{}
  e.GET(versionPath, h.HealthCheck)
}

func (h *healthHandler) HealthCheck(c echo.Context) error {
  var healthCheckData entity.HealthCheck
  projectInfo, err := ioutil.ReadFile("./package.json")

  if err != nil {
    log.Fatal("File project information error: %s", err)
  }
  _ = json.Unmarshal(projectInfo, &healthCheckData)

  healthCheck := entity.HealthCheck{
    App:     healthCheckData.App,
    Version: healthCheckData.Version,
    Env:     os.Getenv("ENVIRONMENT_TYPE"),
    Author:  healthCheckData.Author,
  }

  return c.JSON(http.StatusOK, healthCheck)
}
