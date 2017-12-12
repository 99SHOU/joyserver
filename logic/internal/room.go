package internal

import ()

type ROOM_STATU int

const (
	ROOM_STATU_INVALID    ROOM_STATU = 1 + iota //非法状态
	ROOM_STATU_WAIT_READY                       //等待所有人准备
	ROOM_STATU_RUN                              //运行状态
	ROOM_STATU_DESTROY                          //准备销毁
)

func NewRoom(roomId uint32) *Room {
	room := &Room{roomId: roomId, roomStatu: ROOM_STATU_INVALID}
	room.players = make(map[uint32]Player)

	return room
}

type Room struct {
	roomId    uint32
	players   map[uint32]Player
	roomStatu ROOM_STATU
}

func (r *Room) AddPlayer(playerId uint32) bool {
	if _, ok := r.players[playerId]; ok {
		return false
	}

	return true
}

func (r *Room) RemovePlayer(playerId uint32) bool {
	if _, ok := r.players[playerId]; !ok {
		return false
	}

	delete(r.players, playerId)

	return true
}

func (r *Room) Init() {

	r.roomStatu = ROOM_STATU_WAIT_READY
}

func (r *Room) Destroy() {

}

func (r *Room) Run() {

}
