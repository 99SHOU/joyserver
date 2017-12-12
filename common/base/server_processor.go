package base

import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/name5566/leaf/log"
	"reflect"
)

// -------------------------
// | id | protobuf message |
// -------------------------
type Processor struct {
	littleEndian bool
	msgInfo      map[uint16]*MsgInfo
	msgIDs       map[reflect.Type]uint16
}

type MsgInfo struct {
	msgType    reflect.Type
	msgHandler MsgHandler
}

type MsgHandler func(interface{}, interface{})

func NewProcessor() *Processor {
	p := new(Processor)
	p.littleEndian = false
	p.msgInfo = make(map[uint16]*MsgInfo)
	p.msgIDs = make(map[reflect.Type]uint16)
	return p
}

// It's dangerous to call the method on routing or marshaling (unmarshaling)
func (p *Processor) SetByteOrder(littleEndian bool) {
	p.littleEndian = littleEndian
}

// It's dangerous to call the method on routing or marshaling (unmarshaling)
func (p *Processor) Register(id uint16, msg interface{}) {
	msgType := reflect.TypeOf(msg)
	if msgType == nil || msgType.Kind() != reflect.Ptr {
		log.Fatal("protobuf message pointer required")
	}
	if _, ok := p.msgInfo[id]; ok {
		log.Fatal("id %v is already registered as %v ", id, p.msgInfo[id].msgType)
	}

	if _, ok := p.msgIDs[msgType]; ok {
		log.Fatal("msgType %v is already registered as %v ", msgType, p.msgIDs[msgType])
	}

	mi := new(MsgInfo)
	mi.msgType = msgType
	mi.msgHandler = nil

	p.msgInfo[id] = mi
	p.msgIDs[msgType] = id
}

func (p *Processor) SetHandler(id uint16, msghandler MsgHandler) {
	mi, ok := p.msgInfo[id]
	if !ok {
		log.Fatal("id %v is not registered", id)
		return
	}

	if mi.msgHandler != nil {
		log.Fatal("msg %v handler is already registered", mi.msgType)
		return
	}

	mi.msgHandler = msghandler
}

func (p *Processor) Dispatch(msg interface{}, userdata interface{}) {
	id, ok := p.msgIDs[reflect.TypeOf(msg)]
	if !ok {
		log.Error("msg %v is not registered", reflect.TypeOf(msg))
		return
	}

	mi, ok := p.msgInfo[id]
	if !ok {
		log.Error("id %v is not registered", id)
		return
	}

	if mi.msgHandler == nil {
		log.Error("msg %v handler is not registered", mi.msgType)
		return
	}

	mi.msgHandler(msg, userdata)
}

// goroutine safe
func (p *Processor) Unmarshal(data []byte) (interface{}, error) {
	if len(data) < 2 {
		return nil, errors.New("protobuf data too short")
	}

	// id
	var id uint16
	if p.littleEndian {
		id = binary.LittleEndian.Uint16(data)
	} else {
		id = binary.BigEndian.Uint16(data)
	}

	// msg
	mi, ok := p.msgInfo[id]
	if !ok {
		return nil, fmt.Errorf("message id %v not registered", id)
	}

	msg := reflect.New(mi.msgType.Elem()).Interface()
	return msg, proto.UnmarshalMerge(data[2:], msg.(proto.Message))
}

// goroutine safe
func (p *Processor) Marshal(msg interface{}) ([][]byte, error) {
	id, ok := p.msgIDs[reflect.TypeOf(msg)]
	if !ok {
		return nil, fmt.Errorf("msg %v is not registered", reflect.TypeOf(msg))
	}

	_id := make([]byte, 2)
	if p.littleEndian {
		binary.LittleEndian.PutUint16(_id, id)
	} else {
		binary.BigEndian.PutUint16(_id, id)
	}

	// data
	data, err := proto.Marshal(msg.(proto.Message))
	return [][]byte{_id, data}, err
}
