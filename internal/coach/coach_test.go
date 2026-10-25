package coach

import (
	"context"
	"testing"
)

func TestRulesCoachReturnsMissionForWeek(t *testing.T) {
	mission, err := NewRulesCoach().NextMission(context.Background(), Goal{Stack: "Go"}, Progress{Week: 2})
	if err != nil {
		t.Fatal(err)
	}
	if mission.Title != "Construí la primera versión" || mission.XP == 0 {
		t.Fatalf("misión inesperada: %+v", mission)
	}
}

func TestRulesCoachReviewsEmptyCheckIn(t *testing.T) {
	feedback, err := NewRulesCoach().ReviewCheckIn(context.Background(), Goal{}, " ")
	if err != nil {
		t.Fatal(err)
	}
	if feedback.Next == "" {
		t.Fatal("se esperaba una siguiente acción")
	}
}
