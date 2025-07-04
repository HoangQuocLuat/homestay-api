package main

import (
	"back-end/cmd/config"
	"back-end/cmd/database/postgresql"
	"back-end/cmd/handler"
	"back-end/cmd/svc"
	config2 "back-end/core/config"
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	configType          = flag.String("config-type", "file", "Config type: file or remote")
	configFilePath      = flag.String("config-file-path", "configs", "Config file path: path to config dir")
	configRemoteAddress = flag.String("config-remote-address", "127.0.0.1:8500", "Address of remote config: ip:port")
	configRemoteKeys    = flag.String("config-remote-keys", "", "Keys on remote config: separate by ,")
)

func initConfig() {
	configSource := &config2.Viper{
		ConfigType:    *configType,
		FilePath:      *configFilePath,
		RemoteAddress: *configRemoteAddress,
		RemoteKeys:    *configRemoteKeys,
	}
	err := configSource.InitConfig()
	if err != nil {
		panic(err)
	}
}

func getBootstrapConfig() *config.Config {
	bc := &config.Config{}
	if err := viper.UnmarshalKey("http", &bc.Http); err != nil {
		panic(err)
	}
	if err := viper.UnmarshalKey("database", &bc.PostgreSQL); err != nil {
		panic(err)
	}
	return bc
}

func main() {
	flag.Parse()
	initConfig()
	c := getBootstrapConfig()

	homestayDB, err := postgresql.NewHomestayDB()
	if err != nil {
		panic(err)
	}

	// roomDB, err := postgresql.NewRoomDB()
	// if err != nil {
	// 	panic(err)
	// }

	// serviceDB, err := postgresql.NewServiceDB()
	// if err != nil {
	// 	panic(err)
	// }

	serverCtx := svc.NewServiceContext(*c, homestayDB, nil, nil)
	router := gin.Default()

	handler.RegisterHandlers(router, serverCtx)

	fmt.Printf("Starting server at %s:%s...\n", c.Http.Path, c.Http.Port)
	err = router.Run(fmt.Sprintf("%s:%s", c.Http.Path, c.Http.Port))
	if err != nil {
		panic(err)
	}
}
