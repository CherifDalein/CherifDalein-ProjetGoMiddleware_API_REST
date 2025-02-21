package repositories

import (
	"API-Creation/internal/helpers"
	"API-Creation/internal/models"
	"github.com/google/uuid"
)

// Créer une alerte
func CreateAlert(alert *models.Alert) error {
	alert.ID = uuid.New() // Génère un nouvel UUID
	query := `INSERT INTO alerts (id, resource_id, email, IsALL) VALUES (?, ?, ?, ?)`
	_, err := helpers.DB.Exec(query, alert.ID, alert.ResourceID, alert.Email, alert.IsAll)
	return err
}

// Récupérer toutes les alertes
func GetAllAlerts() ([]models.Alert, error) {
	rows, err := helpers.DB.Query("SELECT id, resource_id, email, IsAll FROM alerts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var alerts []models.Alert
	for rows.Next() {
		var alert models.Alert
		if err := rows.Scan(&alert.ID, &alert.ResourceID, &alert.Email, &alert.IsAll); err != nil {
			return nil, err
		}
		alerts = append(alerts, alert)
	}
	return alerts, nil
}
