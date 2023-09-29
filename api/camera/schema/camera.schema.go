package schema

import "gorm.io/gorm"

type CameraSchema struct {
	gorm.Model
	RtspAddress string
}
