package rrogue

type TurnState int

const (
	BeforePlayerAction = iota
	PlayerTurn
	MonsterTurn
	GameOVer
)

func GetNextState(state TurnState) TurnState {
	switch state {
	case BeforePlayerAction:
		return PlayerTurn

	case PlayerTurn:
		return MonsterTurn

	case MonsterTurn:
		return BeforePlayerAction

	case GameOVer:
		return GameOVer

	default:
		return PlayerTurn
	}
}
