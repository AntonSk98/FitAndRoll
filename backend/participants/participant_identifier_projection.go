package participants

// ParticipantIdentifierProjection is a projection of the Participant entity
// that contains only the ID field.
// It is used to fetch the ID of a participant based on their name and surname.
// This is useful for operations where we only need the ID and not the full participant details.
type ParticipantIdentifierProjection struct {
	ID *uint
}
