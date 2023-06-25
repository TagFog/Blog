package Global

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	Config2 "web/Config"
)

var (
	Config *Config2.Config
	DB     *gorm.DB
	Log    *logrus.Logger
)
