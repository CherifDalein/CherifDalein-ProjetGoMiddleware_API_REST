package repositories

import (
	"API-Creation/internal/helpers"
	"API-Creation/internal/models"
)

func CreateResource(resource *models.Resource) error {
	query := `INSERT INTO resources (id, name, ucaID) VALUES (?, ?, ?)`
	_, err := helpers.DB.Exec(query, resource.ID, resource.Name, resource.UcaID)
	return err
}

func GetAllResources() ([]models.Resource, error) {
	rows, err := helpers.DB.Query("SELECT id, name, ucaID FROM resources")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var resources []models.Resource
	for rows.Next() {
		var resource models.Resource
		if err := rows.Scan(&resource.ID, &resource.Name, &resource.UcaID); err != nil {
			return nil, err
		}
		resources = append(resources, resource)
	}
	return resources, nil
}
