package global

import (
	"gorm-demo/config"
	"sync"
	"time"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

// 通用表数据
type GVA_MODEL struct {
	ID        uint           `gorm:"primarykey"` // 主键ID
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}

var (
	SERVE_ADDR string
	GVA_VP     *viper.Viper
	GVA_CONFIG config.Server
	MY_SQL     *gorm.DB
	RW_LOCK    sync.RWMutex
	LOCK       sync.Mutex
)

const (
	ConfigEnv         = "GVA_CONFIG"
	ConfigDefaultFile = "config.yaml"
	ConfigTestFile    = "config.test.yaml"
	ConfigDebugFile   = "config.debug.yaml"
	ConfigReleaseFile = "config.release.yaml"
)