package repositories

import (
	"strings"
	"tucode_v2/internal/interfaces"
	"tucode_v2/internal/models"

	"gorm.io/gorm"
)

// các bộ phận xe, ví dụ: bánh, động cơ...
type RobotRepo struct {
	gormDB *gorm.DB
}

func NewRobotRepo(db *gorm.DB) interfaces.IRobotRepository {
	return &RobotRepo{
		gormDB: db,
	}
}

// tức là một hành động thuộc về RobotRepo.
func (r *RobotRepo) Create(Robot *models.Robots) (int64, error) {
	result := r.gormDB.Create(Robot)
	return int64(Robot.ID), result.Error
}

// Get robot
func (r *RobotRepo) GetList(params models.RobotParams) ([]models.Robots, error) {
	db := r.gormDB.Model(&models.Robots{})
	// Filter
	if params.Color != "" {
		db = db.Where("LOWER(color) = ?", strings.ToLower(params.Color))
	}
	if params.Name != "" {
		db = db.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(params.Name)+"%")
	}
	if params.Model != "" {
		db = db.Where("LOWER(model) = ?", strings.ToLower(params.Model))
	}
	if params.Version != "" {
		db = db.Where("LOWER(version) = ?", strings.ToLower(params.Version))
	}
	if params.ProductionCost > 0 {
		db = db.Where("production_cost = ?", params.ProductionCost)
	}
	if params.Color != "" {
		db = db.Where("LOWER(color) = ?", strings.ToLower(params.Color))
	}
	if params.Camera != "" {
		db = db.Where("LOWER(camera) = ?", strings.ToLower(params.Camera))
	}
	if params.Weight > 0 {
		db = db.Where("weight = ?", params.Weight)
	}
	if params.Height > 0 {
		db = db.Where("height = ?", params.Height)
	}
	if params.Width > 0 {
		db = db.Where("width = ?", params.Width)
	}
	if params.Depth > 0 {
		db = db.Where("depth = ?", params.Depth)
	}
	if params.BatteryLife > 0 {
		db = db.Where("battery_life = ?", params.BatteryLife)
	}
	if params.HasAI != nil {
		db = db.Where("has_ai = ?", *params.HasAI)
	}
	if params.SensorCount > 0 {
		db = db.Where("sensor_count = ?", params.SensorCount)
	}
	if params.MotorType != "" {
		db = db.Where("motor_type = ?", params.MotorType)
	}
	if params.Speed > 0 {
		db = db.Where("speed = ?", params.Speed)
	}
	if params.MaxPayload > 0 {
		db = db.Where("max_payload = ?", params.MaxPayload)
	}
	if params.OperatingTemperatureMin > 0 {
		db = db.Where("operating_temperature_min <= ?", params.OperatingTemperatureMin)
	}
	if params.OperatingTemperatureMax > 0 {
		db = db.Where("operating_temperature_max >= ?", params.OperatingTemperatureMax)
	}
	if params.IPRating != "" {
		db = db.Where("LOWER(ip_rating) >= ?", strings.ToLower(params.IPRating))
	}

	if params.HasWifi != nil {
		db = db.Where("has_ai = ?", *params.HasWifi)
	}
	if params.HasBluetooth != nil {
		db = db.Where("has_ai = ?", *params.HasBluetooth)
	}
	if params.CountryOfOrigin != "" {
		db = db.Where("LOWER(country_of_origin) >= ?", strings.ToLower(params.CountryOfOrigin))
	}
	if params.ManufactureCompany != "" {
		db = db.Where("LOWER(manufacture_company) >= ?", strings.ToLower(params.ManufactureCompany))
	}
	if params.FirmwareVersion != "" {
		db = db.Where("LOWER(firmware_version) >= ?", strings.ToLower(params.FirmwareVersion))
	}
	if params.OperatingHours > 0 {
		db = db.Where("operating_hours = ?", params.OperatingHours)
	}
	if params.CPU != "" {
		db = db.Where("LOWER(cpu) >= ?", strings.ToLower(params.CPU))
	}
	if params.RAM > 0 {
		db = db.Where("ram = ?", params.RAM)
	}
	if params.Storage > 0 {
		db = db.Where("storage = ?", params.Storage)
	}
	if params.HasTouchScreen != nil {
		db = db.Where("has_touch_screen = ?", *params.HasTouchScreen)
	}
	if params.ScreenSize > 0 {
		db = db.Where("screen_size = ?", params.ScreenSize)
	}
	if params.SupportedLanguages != "" {
		db = db.Where("LOWER(supported_languages) >= ?", strings.ToLower(params.SupportedLanguages))
	}
	if params.ConnectivityProtocols != "" {
		db = db.Where("LOWER(connectivity_protocols) >= ?", strings.ToLower(params.ConnectivityProtocols))
	}
	if params.EmergencyShutdownSupport != nil {
		db = db.Where("emergency_shutdown_support = ?", *params.EmergencyShutdownSupport)
	}
	if params.CustomizableModulesAvailable != nil {
		db = db.Where("customizable_modules_available = ?", *params.CustomizableModulesAvailable)
	}
	if params.ProductionDateFrom != "" && params.ProductionDateTo != "" {
		db = db.Where("DATE(production_date) BETWEEN ? AND ?", params.ProductionDateFrom, params.ProductionDateTo)
	} else if params.ProductionDateFrom != "" {
		db = db.Where("DATE(production_date) >= ?", params.ProductionDateFrom)
	} else if params.ProductionDateTo != "" {
		db = db.Where("DATE(production_date) <= ?", params.ProductionDateTo)
	}

	// Sort
	if params.SortBy != "" {
		db = db.Order(params.SortBy)
	}
	offset := (params.Page - 1) * params.Limit
	var robots []models.Robots
	err := db.Offset(offset).Limit(params.Limit).Find(&robots).Error
	return robots, err
}

func (r *RobotRepo) Update(RobotID int, Robot *models.Robots) (int64, error) {
	result := r.gormDB.Model(&models.Robots{}).Where("id = ?", RobotID).Updates(Robot)
	return result.RowsAffected, result.Error
}

// models.Robots{} là kiểu struct đại diện cho bảng bạn muốn thao tác.
// RobotID là ID muốn xóa, GORM sẽ ngầm thực thi:
func (r *RobotRepo) Delete(RobotID int) (int64, error) {
	result := r.gormDB.Delete(&models.Robots{}, RobotID)
	return int64(RobotID), result.Error
}

func (r *RobotRepo) GetByID(RobotID int) (*models.Robots, error) {
	var robot models.Robots
	result := r.gormDB.First(&robot, RobotID)
	return &robot, result.Error
}
