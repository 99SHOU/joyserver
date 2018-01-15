package internal

type Player struct {
	playerId   uint32
	account    string
	playerName string
	roomId     uint32
	isOnline   bool
}

func (p *Player) GetPlayerId() uint32 {
	return p.playerId
}

func (p *Player) GetPlayerName() string {
	return p.playerName
}
