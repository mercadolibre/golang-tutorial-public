package main

import (
	"time"

	tm "github.com/buger/goterm"
)

type Renderer struct{}

func (r Renderer) showStartGameScreen() {
	tm.Clear()
	tm.MoveCursor(30, 10)
	tm.Print(tm.Bold("STARTING NEW GAME!"))
	tm.Flush()
}

func (r Renderer) showEndGameScreen(result string, winnerName string) {
	tm.Clear()
	tm.MoveCursor(30, 9)
	if result == "WIN" {
		tm.Print(tm.Bold("WINNER IS " + winnerName))
	} else {
		tm.Print(tm.Bold("IT A TIE!!"))
	}
	tm.MoveCursor(1, 20)
	tm.Flush()
}

func (r Renderer) showShotSelected(playerShot Point) {
	tm.MoveCursor(playerShot.X+1, playerShot.Y+2)
	tm.Print(tm.Background(tm.Color(tm.Bold("X"), tm.RED), tm.YELLOW))
	tm.Flush()
	time.Sleep(500 / SPEED * time.Millisecond)
}

func (r Renderer) printDebugMessage(message string) {
	tm.MoveCursor(0, NOTIFICATION_POSITION)
	tm.Print(tm.Bold(message))
	tm.Flush()
}

func (renderer Renderer) showMapWithShipsAndHitsForPlayer(engineStatus EngineStatus, message string) {
	tm.Clear()
	activePlayer := engineStatus.Players[engineStatus.CurrentTurn]
	renderer.drawMap(engineStatus, "Hits", 0, message)
	renderer.drawMap(engineStatus, "Ships", MAP_WIDTH+20, "")

	var shipColor int
	for _, ship := range activePlayer.Ships {
		for _, shipModule := range ship.ShipModules {
			if shipModule.IsHit {
				shipColor = tm.RED
			} else {
				shipColor = tm.GREEN
			}
			tm.MoveCursor(MAP_WIDTH+shipModule.X+20, shipModule.Y+2)
			tm.Print(tm.Background(tm.Color(string("O"), shipColor), tm.BLUE))
		}
	}

	for _, hit := range activePlayer.Hits {
		if hit.IsHit {
			shipColor = tm.RED
		} else {
			shipColor = tm.GREEN
		}
		tm.MoveCursor(hit.X+1, hit.Y+2)
		tm.Print(tm.Background(tm.Color(string("X"), shipColor), tm.BLUE))
	}

	tm.Flush()
}

func (r Renderer) drawMap(engineStatus EngineStatus, label string, xDisplacement int, message string) {
	tm.MoveCursor(xDisplacement+1, 1)
	tm.Print(tm.Bold(label))
	for row := 1; row < MAP_HEIGHT+1; row++ {
		for column := xDisplacement; column < xDisplacement+MAP_WIDTH; column++ {
			tm.MoveCursor(column+1, row+1)
			tm.Print(tm.Background(tm.Color(string(" "), tm.MAGENTA), tm.CYAN))
		}
	}

	if xDisplacement == 0 {
		tm.MoveCursor(1, MAP_HEIGHT+3)
		tm.Print(tm.Bold("Current Turn: " + engineStatus.Players[engineStatus.CurrentTurn].Name))
	}

	tm.MoveCursor(xDisplacement, MAP_HEIGHT+5)
	tm.Print(tm.Bold(message))

	if xDisplacement == 0 {
		tm.MoveCursor(1, MAP_HEIGHT+7)
		tm.Printf("Player 1 Ships")
		for i, ship := range engineStatus.Players[0].Ships {
			tm.MoveCursor(14, MAP_HEIGHT+8+i)
			tm.Printf("%+v", ship)
		}
		tm.MoveCursor(0, MAP_HEIGHT+11)
		tm.Printf("Player 2 Ships")
		for i, ship := range engineStatus.Players[1].Ships {
			tm.MoveCursor(14, MAP_HEIGHT+12+i)
			tm.Printf("%+v", ship)
		}
	}
}
