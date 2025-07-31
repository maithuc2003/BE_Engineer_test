package models

import "time"

// RobotModels sang snake_case
type Robots struct {
	ID                           int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name                         string    `json:"name"`
	Description                  string    `json:"description"`
	Model                        string    `json:"model"`
	Version                      string    `json:"version"`
	ProductionCost               float64   `json:"production_cost"`
	ProductionDate               time.Time `json:"production_date"`
	Color                        string    `json:"color"`
	Camera                       string    `json:"camera"`
	Weight                       float64   `json:"weight"`
	Height                       float64   `json:"height"`
	Width                        float64   `json:"width"`
	Depth                        float64   `json:"depth"`
	BatteryLife                  int       `json:"battery_life"`
	HasAI                        bool      `json:"has_ai"`
	SensorCount                  int       `json:"sensor_count"`
	MotorType                    string    `json:"motor_type"`
	Speed                        float64   `json:"speed"`
	MaxPayload                   float64   `json:"max_payload"`
	Material                     string    `json:"material"`
	OperatingTemperatureMin      float64   `json:"operating_temperature_min"`
	OperatingTemperatureMax      float64   `json:"operating_temperature_max"`
	IPRating                     string    `json:"ip_rating"`
	HasWifi                      bool      `json:"has_wifi"`
	HasBluetooth                 bool      `json:"has_bluetooth"`
	CountryOfOrigin              string    `json:"country_of_origin"`
	ManufactureCompany           string    `json:"manufacture_company"`
	FirmwareVersion              string    `json:"firmware_version"`
	LastMaintenance              time.Time `json:"last_maintenance"`
	NextMaintenanceDue           time.Time `json:"next_maintenance_due"`
	OperatingHours               float64   `json:"operating_hours"`
	CPU                          string    `json:"cpu"`
	RAM                          int       `json:"ram"`
	Storage                      int       `json:"storage"`
	HasTouchScreen               bool      `json:"has_touch_screen"`
	ScreenSize                   float64   `json:"screen_size"`
	SupportedLanguages           string    `json:"supported_languages"`
	ConnectivityProtocols        string    `json:"connectivity_protocols"`
	EmergencyShutdownSupport     bool      `json:"emergency_shutdown_support"`
	CustomizableModulesAvailable bool      `json:"customizable_modules_available"`
}

type RobotParams struct {
	ID                           int       `form:"id"`
	Page                         int       `form:"page"`
	Limit                        int       `form:"limit"`
	Name                         string    `form:"name"`
	SortBy                       string    `form:"sort_by"`
	Description                  string    `form:"description"`
	Model                        string    `form:"model"`
	Version                      string    `form:"version"`
	ProductionCost               float64   `form:"production_cost"`
	ProductionDateFrom           string    `form:"production_date_from"`
	ProductionDateTo             string    `form:"production_date_to"`
	Color                        string    `form:"color"`
	Camera                       string    `form:"camera"`
	Weight                       float64   `form:"weight"`
	Height                       float64   `form:"height"`
	Width                        float64   `form:"width"`
	Depth                        float64   `form:"depth"`
	BatteryLife                  int       `form:"battery_life"`
	HasAI                        *bool     `form:"has_ai"`
	SensorCount                  int       `form:"sensor_count"`
	MotorType                    string    `form:"motor_type"`
	Speed                        float64   `form:"speed"`
	MaxPayload                   float64   `form:"max_payload"`
	Material                     string    `form:"material"`
	OperatingTemperatureMin      float64   `form:"operating_temperature_min"`
	OperatingTemperatureMax      float64   `form:"operating_temperature_max"`
	IPRating                     string    `form:"ip_rating"`
	HasWifi                      *bool     `form:"has_wifi"`
	HasBluetooth                 *bool     `form:"has_bluetooth"`
	CountryOfOrigin              string    `form:"country_of_origin"`
	ManufactureCompany           string    `form:"manufacture_company"`
	FirmwareVersion              string    `form:"firmware_version"`
	LastMaintenance              time.Time `form:"last_maintenance"`
	NextMaintenanceDue           time.Time `form:"next_maintenance_due"`
	OperatingHours               float64   `form:"operating_hours"`
	CPU                          string    `form:"cpu"`
	RAM                          int       `form:"ram"`
	Storage                      int       `form:"storage"`
	HasTouchScreen               *bool     `form:"has_touch_screen"`
	ScreenSize                   float64   `form:"screen_size"`
	SupportedLanguages           string    `form:"supported_languages"`
	ConnectivityProtocols        string    `form:"connectivity_protocols"`
	EmergencyShutdownSupport     *bool     `form:"emergency_shutdown_support"`
	CustomizableModulesAvailable *bool     `form:"customizable_modules_available"`
}
