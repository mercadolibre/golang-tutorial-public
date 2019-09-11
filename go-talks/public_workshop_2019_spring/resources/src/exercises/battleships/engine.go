package main

import (
	"fmt"
	"time"
)

const (
	MAP_HEIGHT      = 8
	MAP_WIDTH       = 16
	NUMBER_OF_SHIPS = 3
	LENGTH_OF_SHIP  = 3
	SPEED           = 2 // 1 -> 100
)

const NOTIFICATION_POSITION = MAP_HEIGHT + 16

var Displacements = map[CardinalDirection]Point{
	EAST:  Point{1, 0},
	SOUTH: Point{0, 1},
	WEST:  Point{-1, 0},
	NORTH: Point{0, -1},
}

func NewEngine(playerRunners []PlayerRunnerFunction) EngineStatus {
	return EngineStatus{
		PlayerAIs:   playerRunners,
		CurrentTurn: 0,
		Players:     make([]PlayerStatus, len(playerRunners)),
		Channel:     make(chan EngineCommand, 1),
	}
}

func (engine EngineStatus) run() {
	renderer := Renderer{}

	for i := range engine.Players {
		channel := make(chan PlayerCommand)
		engine.Players[i] = PlayerStatus{
			Id:      i,
			Channel: channel,
			IsReady: false,
			Ships:   make([]Ship, 0),
			Hits:    make([]Hit, 0),
		}

		go engine.PlayerAIs[i](renderer, channel, engine.Channel)

		renderer.showMapWithShipsAndHitsForPlayer(engine, "Please enter your name...")
		channel <- PlayerCommand{
			Id:   i,
			Type: REQUEST_NAME,
		}

		playerRegistrationReady := false

		for !playerRegistrationReady {
			time.Sleep(500 / SPEED * time.Millisecond)
			playerResponse := <-engine.Channel
			switch playerResponse.Type {
			case PLAYER_RESPONSE_NAME:

				playerNameResponse, ok := playerResponse.Payload.(PlayerResponseName)
				if ok {
					engine.Players[i].Name = string(playerNameResponse)
					renderer.showMapWithShipsAndHitsForPlayer(engine, "Position the first ship!")

					channel <- PlayerCommand{
						Type: REQUEST_SHIP_POSTION,
					}
				}

			case PLAYER_RESPONSE_SHIP_POSITION:
				playerPositionShipResponse, ok := playerResponse.Payload.(PlayerResponseShipPositionPayload)

				if ok {
					ship, err := engine.validateAndCreateShip(engine.Players[i], playerPositionShipResponse)

					if err != nil {
						engine.Players[i].Channel <- PlayerCommand{
							Type: REQUEST_SHIP_POSTION,
						}
					} else {
						engine.Players[i].Ships = append(engine.Players[i].Ships, *ship)
						if len(engine.Players[i].Ships) < NUMBER_OF_SHIPS {
							renderer.showMapWithShipsAndHitsForPlayer(engine, "Position the next ship!")
							channel <- PlayerCommand{
								Type: REQUEST_SHIP_POSTION,
							}
						} else {
							renderer.showMapWithShipsAndHitsForPlayer(engine, "Positioned all the ships!")
							engine.Players[i].IsReady = true
							engine.CurrentTurn = (engine.CurrentTurn + 1) % uint(len(engine.Players))

							playerRegistrationReady = true
						}
					}
				}
			}
		}
	}

	engine.Channel <- EngineCommand{
		Type: ENGINE_START_GAME,
	}

	for {
		time.Sleep(1000 / SPEED * time.Millisecond)
		select {
		case engineCommand := <-engine.Channel:
			switch engineCommand.Type {
			case ENGINE_START_GAME:
				renderer.showStartGameScreen()

				engine.Channel <- EngineCommand{
					Type: ENGINE_TAKE_TURN,
				}
			case ENGINE_TAKE_TURN:
				playerOnTurn := engine.Players[engine.CurrentTurn]

				renderer.showMapWithShipsAndHitsForPlayer(engine, "Please input your shot!")

				lastHit := new(Hit)
				if len(playerOnTurn.Hits) != 0 {
					lastHit = &playerOnTurn.Hits[len(playerOnTurn.Hits)-1]
				}

				playerOnTurn.Channel <- PlayerCommand{
					Type:    REQUEST_SHOT_POSITION,
					Payload: lastHit,
				}
			case PLAYER_RESPONSE_SHOT_POSITION:
				playerOnTurn := engine.Players[engine.CurrentTurn]

				playerShotResponse, ok := engineCommand.Payload.(Point)
				if ok {
					if playerOnTurn.Id == engineCommand.Id {
						hasHitted := false

						for _, player := range engine.Players {
							if player.Id != playerOnTurn.Id {
								hasHitted = engine.checkHit(player, Hit{
									Point: playerShotResponse,
								})
							}
						}

						hit := Hit{
							Point: playerShotResponse,
							IsHit: hasHitted,
						}

						engine.Players[engine.CurrentTurn].Hits = append(playerOnTurn.Hits, hit)

						renderer.showShotSelected(playerShotResponse)

						if hasHitted {
							renderer.showMapWithShipsAndHitsForPlayer(engine, "You hitted the enemy ship!")

							playerOnTurn.Channel <- PlayerCommand{
								Type:    REQUEST_SHOT_POSITION,
								Payload: &hit,
							}
						} else {
							renderer.showMapWithShipsAndHitsForPlayer(engine, "So close!")
							engine.Channel <- EngineCommand{
								Type: ENGINE_GIVE_TURN,
							}
						}
					}
				}
			case ENGINE_GIVE_TURN:
				engine.CurrentTurn = (engine.CurrentTurn + 1) % uint(len(engine.Players))
				if engine.CurrentTurn == 0 {
					if engine.isGameOver() {
						renderer.showEndGameScreen(engine.getResult())

						for _, player := range engine.Players {
							player.Channel <- PlayerCommand{
								Type: GAME_OVER,
							}
						}

						return
					}
				}

				engine.Channel <- EngineCommand{
					Type: ENGINE_TAKE_TURN,
				}
			}
		}
	}
}

func (engine EngineStatus) getResult() (string, string) {
	player1Ships := 0
	player2Ships := 0

	for _, playerShip := range engine.Players[0].Ships {
		if !playerShip.IsSinked {
			player1Ships++
		}
	}

	for _, playerShip := range engine.Players[1].Ships {
		if !playerShip.IsSinked {
			player2Ships++
		}
	}

	if player1Ships == player2Ships {
		return "TIE", ""
	}

	if player1Ships > player2Ships {
		return "WIN", engine.Players[0].Name
	}

	return "WIN", engine.Players[1].Name
}

func (engine EngineStatus) isGameOver() bool {
	for _, playerStatus := range engine.Players {
		isLooser := true
		for _, playerShip := range playerStatus.Ships {
			if !playerShip.IsSinked {
				isLooser = false
			}
		}

		if isLooser {
			return true
		}
	}

	return false
}

func (engine EngineStatus) allPlayersReady() bool {
	isReady := true
	for _, playerStatus := range engine.Players {
		if !playerStatus.IsReady {
			isReady = false
		}
	}

	return isReady
}

func (engine EngineStatus) validateAndCreateShip(
	player PlayerStatus,
	playerShipCommand PlayerResponseShipPositionPayload,
) (*Ship, error) {
	var ship Ship

	for shipModule := 0; shipModule < LENGTH_OF_SHIP; shipModule++ {
		point := Point{
			X: playerShipCommand.Origin.X + (Displacements[playerShipCommand.Orientation].X * shipModule),
			Y: playerShipCommand.Origin.Y + (Displacements[playerShipCommand.Orientation].Y * shipModule),
		}

		if point.X >= MAP_WIDTH || point.X <= 0 || point.Y >= MAP_HEIGHT || point.Y <= 0 {
			return nil, fmt.Errorf("Ship outside map %+v", point)
		}

		for _, playerShip := range player.Ships {
			for _, playerModule := range playerShip.ShipModules {
				if playerModule.X == point.X && playerModule.Y == point.Y {
					return nil, fmt.Errorf("Ship on ship")
				}
			}
		}

		ship.ShipModules = append(ship.ShipModules, Hit{
			Point: point,
			IsHit: false,
		})
	}

	return &ship, nil
}

func (engine EngineStatus) checkHit(player PlayerStatus, hit Hit) bool {
	for i, playerShip := range player.Ships {
		for j, playerModule := range playerShip.ShipModules {
			if playerModule.X == hit.X && playerModule.Y == hit.Y {
				player.Ships[i].ShipModules[j].IsHit = true

				IsSinked := true
				for _, module := range player.Ships[i].ShipModules {
					if !module.IsHit {
						IsSinked = false
					}
				}
				player.Ships[i].IsSinked = IsSinked

				return true
			}
		}
	}

	return false
}
