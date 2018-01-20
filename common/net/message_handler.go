package net

import ()

type ServerHandler interface {
	Register(server Server)
}

type ClientHandler interface {
	Register(client Client)
}
