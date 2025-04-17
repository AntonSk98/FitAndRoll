package scaler

import "fit_and_roll/backend/config"

// Handler to manage the scale of the application
type ScalerHandler struct {
	dbManager *config.DatabaseManager
}

// Creates a new handler of ScalerHandler
func NewScalerHandler(dbManager *config.DatabaseManager) *ScalerHandler {
	return &ScalerHandler{dbManager: dbManager}
}

// Fetches current root font size
// If the root font size was not yet set a default entry is created
func (handler *ScalerHandler) FetchRootFontSize() (float32, error) {
	var scaler Scaler
	if err := handler.dbManager.DB.FirstOrCreate(&scaler, Scaler{}).Error; err != nil {
		return 0, err
	}
	return scaler.FontSize, nil
}

// Persist root font size
func (handler *ScalerHandler) PersistRootFontSize(fontSize float32) error {
	var scaler Scaler
	result := handler.dbManager.DB.First(&scaler)
	if result.Error != nil {
		return result.Error
	}

	scaler.FontSize = fontSize

	return handler.dbManager.DB.Save(&scaler).Error
}

// Entity representing the scale of the application.
// Signleton entry.
type Scaler struct {
	ID       uint    `gorm:"primaryKey"`
	FontSize float32 `json:"fontSize`
}
