package websocket

import (
	"encoding/binary"
	"server/internal/models"
	"time"
)

type IPacker interface {
	Pack(msg models.Message) ([]byte, error)
	Unpack(buffer []byte) (*models.SendMsg, error)
}

type NormalPacker struct {
	ByteOrder binary.ByteOrder
}

const Format = "2006-01-02 15:04:05"

func (p *NormalPacker) Pack(msg models.Message) ([]byte, error) {
	buffer := make([]byte, 8+8+8+8+len(msg.Message))         // 消息长度+消息ID+发送者ID+发送时间+消息内容
	p.ByteOrder.PutUint64(buffer[:8], uint64(len(buffer)))   // 消息长度
	p.ByteOrder.PutUint64(buffer[8:16], msg.MsgID)           // 消息ID
	p.ByteOrder.PutUint64(buffer[16:24], msg.SenderID)       // 发送者ID
	copy(buffer[24:32], []byte(msg.SendTime.Format(Format))) // 发送时间
	copy(buffer[32:], []byte(msg.Message))                   //  消息内容
	//
	return buffer, nil
}

func (p *NormalPacker) Unpack(buffer []byte) (*models.SendMsg, error) {
	msgLen := p.ByteOrder.Uint64(buffer[:8])
	msgID := p.ByteOrder.Uint64(buffer[8:16])
	sendID := p.ByteOrder.Uint64(buffer[16:24])
	timeStr := string(p.ByteOrder.Uint64(buffer[24:32]))
	sendTime, _ := time.Parse(Format, timeStr)
	msg := string(buffer[msgLen-8-8-8-8:])
	//
	message := &models.SendMsg{
		MsgID:    msgID,
		SendID:   sendID,
		SendTime: sendTime,
		Content:  msg,
	}
	//
	return message, nil
}
