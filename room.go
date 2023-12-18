package parasut

import (
	"github.com/WEG-Technology/room"
)

type IParasutRoom interface {
	room.IRoom
	ContactIndex(params ContactIndexParams) room.IResponse
}

type ParasutRoom struct {
	*room.Room
}

func (r ParasutRoom) ContactIndex(params ContactIndexParams) room.IResponse {
	return r.Send(ContactIndex(params))
}

func (r ParasutRoom) CreateBasicPurchaseBill(params ContactIndexParams) room.IResponse {
	return r.Send(ContactIndex(params))
}

func NewRoom() IParasutRoom {
	return ParasutRoom{room.NewRoom(
		NewConnection(),
		room.WithAuth(room.NewTokenAuthStrategy(NewAuthRequest())),
	)}
}
