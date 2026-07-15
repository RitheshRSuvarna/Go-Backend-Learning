package command

type CreatePlanVersionCommand struct {
	DaysessionID string
	Version      int
	Note         string
}
