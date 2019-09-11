package main

type EngineCommandType string

const (
	PLAYER_RESPONSE_NAME          EngineCommandType = "RESPONSE_NAME"
	ENGINE_PLAYER_READY           EngineCommandType = "PLAYER_READY"
	ENGINE_START_GAME             EngineCommandType = "ENGINE_START_GAME"
	ENGINE_TAKE_TURN              EngineCommandType = "TAKE_TURN"
	PLAYER_RESPONSE_SHIP_POSITION EngineCommandType = "RESPONSE_SHIP_POSITION"
	PLAYER_RESPONSE_SHOT_POSITION EngineCommandType = "PLAYER_RESPONSE_SHOT_POSITION"
	ENGINE_GIVE_TURN              EngineCommandType = "ENGINE_GIVE_TURN"
)

type EngineCommand struct {
	Id      int
	Type    EngineCommandType
	Payload interface{}
}

type PlayerResponseName string

type PlayerResponseShotPayload Point

type PlayerResponseShipPositionPayload struct {
	Origin      Point
	Orientation CardinalDirection
}

type EngineStatus struct {
	PlayerAIs   []PlayerRunnerFunction
	CurrentTurn uint
	Players     []PlayerStatus
	Channel     chan EngineCommand
}

type PlayerStatus struct {
	Id      int
	Name    string
	IsReady bool
	Ships   []Ship
	Hits    []Hit
	Channel chan PlayerCommand
}

type Ship struct {
	ShipModules []Hit
	IsSinked    bool
}

type Hit struct {
	Point
	IsHit bool
}

type Point struct {
	X int
	Y int
}

type PlayerRunnerFunction func(
	Renderer,
	chan PlayerCommand,
	chan EngineCommand,
)

type PlayerCommandType string

const (
	REQUEST_NAME          PlayerCommandType = "REQUEST_NAME"
	REQUEST_SHIP_POSTION  PlayerCommandType = "REQUEST_SHIP_POSTION"
	REQUEST_SHOT_POSITION PlayerCommandType = "REQUEST_SHOT_POSITION"
	GAME_OVER             PlayerCommandType = "GAME_OVER"
)

type PlayerCommand struct {
	Id      int
	Type    PlayerCommandType
	Payload *Hit
}

type CardinalDirection string

const (
	NORTH CardinalDirection = "NORTH"
	EAST  CardinalDirection = "EAST"
	SOUTH CardinalDirection = "SOUTH"
	WEST  CardinalDirection = "WEST"
)

var CardinalDirections = []CardinalDirection{
	NORTH, EAST, SOUTH, WEST,
}
