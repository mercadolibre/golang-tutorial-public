package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func getRandomMapPoint() Point {
	return Point{
		X: rand.Intn(MAP_WIDTH),
		Y: rand.Intn(MAP_HEIGHT),
	}
}

func getRandomMapPointNotInPreviousShots(previousShots []Point) Point {
	found := false
	var newShot Point

	for !found {
		newShot = getRandomMapPoint()
		found = true

		for _, shot := range previousShots {
			if shot.X == newShot.X && shot.Y == newShot.Y {
				found = false
			}
		}
	}

	return newShot
}

func RandomAI(
	renderer Renderer,
	commandChannel chan PlayerCommand,
	engineChannel chan EngineCommand) {

	names := []PlayerResponseName{
		"Popeye",
		"Jack Sparrow",
		"Jacques Costeau",
	}
	nameIndex := rand.Int() % len(names)

	var id int
	name := names[nameIndex]
	previousShots := make([]Point, 0)

	for {
		select {
		case command := <-commandChannel:
			switch command.Type {
			case REQUEST_NAME:
				id = command.Id

				engineChannel <- EngineCommand{
					Id:      id,
					Type:    PLAYER_RESPONSE_NAME,
					Payload: name,
				}

			case REQUEST_SHIP_POSTION:
				engineChannel <- EngineCommand{
					Id:   id,
					Type: PLAYER_RESPONSE_SHIP_POSITION,
					Payload: PlayerResponseShipPositionPayload{
						Origin:      getRandomMapPoint(),
						Orientation: CardinalDirections[rand.Intn(len(CardinalDirections))],
					},
				}

			case REQUEST_SHOT_POSITION:
				newShot := getRandomMapPointNotInPreviousShots(previousShots)
				previousShots = append(previousShots, newShot)

				if command.Payload != nil {
					lastHit := command.Payload

					if lastHit.IsHit {
						renderer.printDebugMessage(fmt.Sprintf("Last shot was successful! %+v", lastHit))
					}
				}

				engineChannel <- EngineCommand{
					Id:      id,
					Type:    PLAYER_RESPONSE_SHOT_POSITION,
					Payload: newShot,
				}

			case GAME_OVER:
				return

			default:
				renderer.printDebugMessage(fmt.Sprintf("Unknown command for Player %s: %s", name, command.Type))
			}
		}
	}
}
