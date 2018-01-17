package net

import ()

type MessageHandler interface {
	Register(server Server)
}
