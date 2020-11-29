package service

import (
	"fmt"
	"github.com/kallydev/privacy/config"
	"github.com/kallydev/privacy/database"
	"github.com/kallydev/privacy/database/table"
	"github.com/kallydev/privacy/ent"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net"
	"strconv"
)

type Service struct {
	client    *ent.Client
	databases []database.Database
	config    *config.Config
	instance  *echo.Echo
}

func NewService(configPath string) *Service {
	conf, err := config.NewConfig(configPath)
	if err != nil {
		log.Panicln("failed to load config file")
	}
	instance := echo.New()
	instance.HidePort = true
	instance.HideBanner = true
	return &Service{
		config:   conf,
		instance: instance,
	}
}

func (svc *Service) loadRouter() {
	instance := svc.instance
	instance.HTTPErrorHandler = func(err error, ctx echo.Context) {
		_ = NewResponse(ctx, err, nil)
	}
	instance.Static("/", "../website/build")
	apiGroup := instance.Group("/api")
	{
		apiGroup.GET("/query", svc.queryHandlerFunc)
	}
}

func (svc *Service) LoadDatabase() (err error) {
	svc.client, err = ent.Open("sqlite3", fmt.Sprintf("file:%s", svc.config.Database.Path))
	if err != nil {
		return err
	}
	tablesConfig := svc.config.Database.Tables
	if tablesConfig.QQ {
		svc.databases = append(svc.databases, &table.QQDatabase{
			Client: svc.client,
		})
	}
	if tablesConfig.JD {
		svc.databases = append(svc.databases, &table.JDDatabase{
			Client: svc.client,
		})
	}
	if tablesConfig.SF {
		svc.databases = append(svc.databases, &table.SFDatabase{
			Client: svc.client,
		})
	}
	return nil
}

func (svc *Service) Start() (err error) {
	if err := svc.LoadDatabase(); err != nil {
		return err
	}
	defer func() {
		_ = svc.client.Close()
	}()
	svc.loadRouter()
	httpConfig := svc.config.HttpConfig
	address := net.JoinHostPort(httpConfig.Host, strconv.Itoa(int(httpConfig.Port)))
	if httpConfig.TLS != nil {
		return svc.instance.StartTLS(address, httpConfig.TLS.CertPath, httpConfig.TLS.KeyPath)
	}
	return svc.instance.Start(address)
}
