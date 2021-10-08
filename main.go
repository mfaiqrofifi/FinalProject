package main

import (
	"log"
	"social_media/app/routes"
	_middleware "social_media/app/routes/middleware"
	_userdeteilstatus "social_media/business/users/statusdesc"
	_userUsecaseRegister "social_media/business/users/userRegister"
	_userUsecaseStatus "social_media/business/users/userStatus"
	_usercontrollerdeteil "social_media/controllers/users/http/deteilStatus"
	_userControllerRegister "social_media/controllers/users/http/userRegister"
	_userControllerStatus "social_media/controllers/users/http/userStatus"
	_userdeteilstatusDB "social_media/drivers/databases/detailStatus"
	_registeruserDB "social_media/drivers/databases/userDB"
	_registerstatusDB "social_media/drivers/databases/userStatus"
	_mysqldrver "social_media/drivers/mysql"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile(`app/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}
func dbMigration(db *gorm.DB) {
	db.AutoMigrate(&_registeruserDB.DBRegisters{})
	db.AutoMigrate(&_registerstatusDB.DBStatus{})
	db.AutoMigrate(&_userdeteilstatusDB.DetailStatusDB{})
}

func main() {

	configDB := _mysqldrver.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}
	Conn := configDB.InitialDB()

	dbMigration(Conn)

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	e := echo.New()
	configJWT := _middleware.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}
	userRepositoryRegister := _registeruserDB.NewMysqlUserRepository(Conn)
	userUseCaseRegister := _userUsecaseRegister.NewUserUscase(userRepositoryRegister, timeoutContext, configJWT)
	userControlerRegister := _userControllerRegister.NewController(userUseCaseRegister)

	userRepositoryStatus := _registerstatusDB.NewMysqlUserRepository(Conn)
	userUseCaseStatus := _userUsecaseStatus.NewUserUscase(userRepositoryStatus, timeoutContext, configJWT)
	userControlerStatus := _userControllerStatus.NewController(userUseCaseStatus)

	userRepositoryDeteilStatus := _userdeteilstatusDB.NewMysqlUserRepository(Conn)
	userUseCaseDeteil := _userdeteilstatus.NewUserUscase(userRepositoryDeteilStatus, timeoutContext)
	userControlerDeteil := _usercontrollerdeteil.NewController(userUseCaseDeteil)

	routesInit := routes.ControllerList{
		UserControllerRegister:      *userControlerRegister,
		UserControllerStatus:        *userControlerStatus,
		UserControllerDeteilStattus: *userControlerDeteil,
	}
	routesInit.RouteRegister(e)
	log.Fatal(e.Start(viper.GetString("server.address")))
}
