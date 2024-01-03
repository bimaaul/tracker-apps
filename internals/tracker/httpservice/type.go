package httpservice

import (
	trackerService "github.com/bimaaul/tracker-apps/internals/tracker/service"
	"gorm.io/gorm"
)

type Handler struct {
	DB         *gorm.DB
	trackerSrv trackerService.TrackerServiceProvider
}
