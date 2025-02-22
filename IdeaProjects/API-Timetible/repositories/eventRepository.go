package repositories

import (
	"API-Timetible/helpers"
	"API-Timetible/models"
	"encoding/json"
	"log"
)

func CreateEvent(event *models.Event) error {
	// Convertir resource_id en JSON
	resourceIDs, err := json.Marshal(event.ResourceId)
	if err != nil {
		log.Println("Error marshaling resource_ids:", err)
		return err
	}

	query := `INSERT INTO events (id, summary, description, location, start, end, uid, resource_id) 
              VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	_, err = helpers.DB.Exec(query, event.ID, event.Summary, event.Description, event.Location, event.Start, event.End, event.UID, resourceIDs)
	return err
}

func GetAllEvents() ([]models.Event, error) {
	rows, err := helpers.DB.Query("SELECT id, summary, description, location, start, end, uid, resource_id FROM events")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []models.Event
	for rows.Next() {
		var event models.Event
		var resourceIDs string
		if err := rows.Scan(&event.ID, &event.Summary, &event.Description, &event.Location, &event.Start, &event.End, &event.UID, &resourceIDs); err != nil {
			return nil, err
		}

		// Convertir resource_id de JSON en []uuid.UUID
		if err := json.Unmarshal([]byte(resourceIDs), &event.ResourceId); err != nil {
			return nil, err
		}

		events = append(events, event)
	}
	return events, nil
}
