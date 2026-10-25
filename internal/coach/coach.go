// Package coach contiene el coach de progreso de DevSquads.
package coach

import (
	"context"
	"strings"
)

type Goal struct {
	Role  string `json:"role"`
	Stack string `json:"stack"`
	Why   string `json:"why"`
}

type Progress struct {
	Week           int  `json:"week"`
	CompletedSteps int  `json:"completed_steps"`
	CheckedIn      bool `json:"checked_in"`
	Blocked        bool `json:"blocked"`
}

type Mission struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Minutes     int    `json:"minutes"`
	XP          int    `json:"xp"`
}

type Feedback struct {
	Message string `json:"message"`
	Next    string `json:"next"`
}

// Coach es la frontera estable entre la aplicación y el agente.
// Hoy usa reglas deterministas; mañana puede delegar en un modelo de IA.
type Coach interface {
	CreateSprint(context.Context, Goal) ([]Mission, error)
	NextMission(context.Context, Goal, Progress) (Mission, error)
	ReviewCheckIn(context.Context, Goal, string) (Feedback, error)
}

type RulesCoach struct{}

func NewRulesCoach() RulesCoach { return RulesCoach{} }

func (RulesCoach) CreateSprint(_ context.Context, goal Goal) ([]Mission, error) {
	stack := strings.TrimSpace(goal.Stack)
	if stack == "" {
		stack = "tu stack"
	}
	return []Mission{
		{Title: "Define una idea pequeña", Description: "Escribí quién la usaría y qué problema resuelve con " + stack + ".", Minutes: 25, XP: 80},
		{Title: "Construí la primera versión", Description: "Elegí una sola acción y hacela funcionar de punta a punta.", Minutes: 40, XP: 120},
		{Title: "Pedí feedback", Description: "Compartí una demo y anotá tres cosas para mejorar.", Minutes: 30, XP: 100},
		{Title: "Publicá tu demo", Description: "Prepará un README, una captura y una explicación de lo que aprendiste.", Minutes: 30, XP: 150},
	}, nil
}

func (c RulesCoach) NextMission(ctx context.Context, goal Goal, progress Progress) (Mission, error) {
	missions, err := c.CreateSprint(ctx, goal)
	if err != nil {
		return Mission{}, err
	}
	index := progress.Week - 1
	if index < 0 {
		index = 0
	}
	if index >= len(missions) {
		index = len(missions) - 1
	}
	return missions[index], nil
}

func (RulesCoach) ReviewCheckIn(_ context.Context, _ Goal, checkIn string) (Feedback, error) {
	if strings.TrimSpace(checkIn) == "" {
		return Feedback{Message: "Tu check-in necesita una acción concreta.", Next: "Escribí qué vas a terminar en los próximos 30 minutos."}, nil
	}
	return Feedback{Message: "Buen check-in. Ya convertiste la intención en un compromiso visible.", Next: "Trabajá 30 minutos y compartí qué resultado obtuviste, aunque sea pequeño."}, nil
}
