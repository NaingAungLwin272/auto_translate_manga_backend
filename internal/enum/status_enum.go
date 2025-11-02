package enum

type MangaStatus string

const (
	StatusOngoing   MangaStatus = "ongoing"
	StatusCompleted MangaStatus = "completed"
	StatusHiatus    MangaStatus = "hiatus"
)

func (s MangaStatus) IsValid() bool {
	switch s {
	case StatusOngoing, StatusCompleted, StatusHiatus:
		return true
	default:
		return false
	}
}
