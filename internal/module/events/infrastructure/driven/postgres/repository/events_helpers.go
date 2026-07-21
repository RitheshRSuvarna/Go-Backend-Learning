package repository

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"common"
)

func uuidStringToPgUUID(id string) (pgtype.UUID, error) {
	u, err := uuid.Parse(id)
	if err != nil {
		return pgtype.UUID{}, err
	}
	return pgtype.UUID{Bytes: u, Valid: true}, nil
}

func pgTypeToString(u pgtype.UUID) string {
	if !u.Valid {
		return ""
	}
	return uuid.UUID(u.Bytes).String()
}

func commonTimeToPG(t common.Time) pgtype.Timestamptz {
    return pgtype.Timestamptz{
        Time:  t.Value(),
        Valid: true,
    }
}