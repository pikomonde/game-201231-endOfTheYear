package objects

import "github.com/pikomonde/game-201231-endOfTheYear/tools"

// Room is a room
type Room struct {
	LayoutSize tools.Vector2
	Rule       *Rule
	Square     *Square
}
