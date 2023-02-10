package routes

import (
	"database/sql"
	"projek_fisioterapi/configs"
	"projek_fisioterapi/repositories"
	userGormRepo "projek_fisioterapi/repositories/gorm/users"
	uService "projek_fisioterapi/services/users"

	"gorm.io/gorm"
)

type Payload struct {
	Config      *configs.Config
	DBGorm      *gorm.DB
	DBSql       *sql.DB
	userRepo    repositories.IUserRepository
	userService uService.IUserServices
}

// Init Repo -----------------------------------------------------------------------------------------------------------------
func (p *Payload) InitRepo() {
	p.InitUserRepo()
}

func (p *Payload) InitUserRepo() {
	p.userRepo = userGormRepo.NewGorm(p.DBGorm)
}

// User -----------------------------------------------------------------------------------------------------------------

func (p *Payload) GetUserServices() uService.IUserServices {
	if p.userService == nil {
		p.InitUserService()
	}
	return p.userService
}
func (p *Payload) InitUserService() {
	if p.userRepo == nil {
		p.InitUserRepo()
	}

	p.userService = uService.NewUserServices(p.userRepo)
}
