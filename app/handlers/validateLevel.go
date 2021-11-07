package handlers

import (
	"encoding/json"
	"fmt"
	"github/miguelapabenedit/youngdevs-api/app/data"
	"github/miguelapabenedit/youngdevs-api/app/enums"
	"github/miguelapabenedit/youngdevs-api/app/repository"
	"io/ioutil"
	"net/http"
)

var (
	levelStateRepo repository.UserLevelState
	levelRepo      repository.GetLevel
)

func NewValidateLevel(lsrepo repository.UserLevelState, lrepo repository.GetLevel) {
	levelStateRepo = lsrepo
	levelRepo = lrepo
}

func ValidateLevel(w http.ResponseWriter, r *http.Request) {
	uls := &data.UserLevelState{}
	bodyContent, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(bodyContent, uls)
	if err != nil {
		fmt.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// hardcodeamos la solucion normal
	uls.UserSolution = `[{"id":5,"display":"WHILE DO","type":1,"condition":{"id":7,"display":"EMPTY CELL","type":2,"cellID":0},"action":{"id":2,"display":"RIGHT","type":0}}]`
	level := levelRepo.GetLevel(int(uls.LevelID))

	commands := []data.Command{}
	err = json.Unmarshal([]byte(uls.UserSolution), &commands)
	if err != nil {
		fmt.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	lvlMap := [][]int{}
	err = json.Unmarshal([]byte([]byte(level.Map)), &lvlMap)
	if err != nil {
		fmt.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if uls.IsSolved = isValidSolution(lvlMap, commands); uls.IsSolved {
		uls.Score = getScore(level, uls.Time, commands)
		levelStateRepo.UpdateLevelState(uls)
	} else {
		fmt.Println("level is invalid")
	}

}

type PlayerPosition struct {
	X int
	Y int
}

func isValidSolution(lvlMap [][]int, solution []data.Command) bool {
	ps := PlayerPosition{0, 0}
	for _, v := range solution {
		m := data.Move{}
		getMovement(v.ID, &m)
		switch v.Type {
		case enums.MOVEMENT:
			move(&ps, m, lvlMap)
		case enums.OPERATION:
			switch v.ID {
			case enums.IF_DO:
				if lvlMap[ps.Y][ps.X] == v.Condition.Cell {
					m := data.Move{}
					getMovement(v.Action.ID, &m)
					move(&ps, m, lvlMap)
				}
			case enums.WHILE_DO:
				m := data.Move{}
				getMovement(v.Action.ID, &m)
				for {
					move(&ps, m, lvlMap)
					if !canMove(ps.X, ps.Y, m, lvlMap) || lvlMap[ps.Y][ps.X] != v.Condition.Cell {
						break
					}
				}
			}
		}

		if lvlMap[ps.Y][ps.X] == 3 {
			return true
		}
	}

	return false
}

func move(ps *PlayerPosition, m data.Move, lvlMap [][]int) {
	if canMove(ps.X, ps.Y, m, lvlMap) && lvlMap[ps.Y+m.Y][ps.X+m.X] != 2 {
		ps.X += m.X
		ps.Y += m.Y
	}
}

func getScore(lvl *data.Level, time int, commands []data.Command) int {
	return calculateTimeScore(lvl, time) + calculateEfficencScore(lvl, commands)
}

func canMove(px int, py int, m data.Move, lvlMap [][]int) bool {
	return px+m.X < len(lvlMap[0]) && px+m.X > 0 && py+m.Y >= 0 && py+m.Y < len(lvlMap)
}

func getMovement(movementID int, m *data.Move) {
	switch movementID {
	case enums.UP:
		m.Y = -1
	case enums.DOWN:
		m.Y = 1
	case enums.RIGHT:
		m.X = 1
	case enums.LEFT:
		m.X = -1
	}
}

func calculateTimeScore(lvl *data.Level, t int) int {
	tMinutes := t / 60
	if tMinutes <= lvl.BestTimeA {
		return 1500
	} else if tMinutes < lvl.BestTimeC {
		return 1000
	}
	return 500
}

func calculateEfficencScore(lvl *data.Level, commands []data.Command) int {
	if len(commands) <= lvl.BestSolution {
		return 2000
	} else if len(commands) <= lvl.BestSolution*2 {
		return 1000
	}
	return 500
}
