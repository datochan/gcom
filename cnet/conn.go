package cnet

import (
	"errors"
	"net"
	"sync/atomic"
	"sync"
	"fmt"
	"time"
	"github.com/datochan/gcom/utils"
)

var (
	ErrSessionClosed = errors.New("net: 会话已被关闭")
	ErrSendChanBlocking = errors.New("net: 发送通道已经塞满")
)

// 事件处理句柄,用于解析相应的封包
type PacketHandler func(session *Session, packet interface{})

type Dispatcher struct {
	rwlock     sync.RWMutex    			// 写互斥避免并发状态下相互干扰
	handlerMap map[uint32]PacketHandler // 时间过程回调句柄, key是事件ID
}

/**
 * 事件分发器
 */
func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		handlerMap: make(map[uint32]PacketHandler),
	}
}

/**
 * 添加新的事件处理器
 * uint32 id: 事件ID
 */
func (p *Dispatcher) AddHandler(id uint32, handler PacketHandler) {
	p.rwlock.Lock()
	defer p.rwlock.Unlock()
	p.handlerMap[id] = handler
}

/**
 * 卸载新的事件处理器
 * uint32 id: 事件ID
 */
func (p *Dispatcher) DelHandler(id uint32) {
	p.rwlock.Lock()
	defer p.rwlock.Unlock()
	delete(p.handlerMap, id)
}

func (p *Dispatcher) GetHandler(id uint32) PacketHandler{
	p.rwlock.Lock()
	defer p.rwlock.Unlock()

	handler, ok := p.handlerMap[id]
	if ok {
		return handler
	}

	return nil
}

/**
 * 事件处理过程
 */
func (p *Dispatcher) HandleProc(session *Session, packet interface{}) {
	p.rwlock.RLock()
	defer p.rwlock.RUnlock()

	// 子类中实现事件分发功能
	//buffer := packet
	//h, ok := p.handlerMap[packet.GetType()]
	//if ok {
	//	h(session, packet)
	//} else {
	//	logger.Error("NOT FOUND")
	//}

}

// 此接口定义了如何组包,解包,发送
// todo notice: 如果要正常发送需要实现此接口
type IPacketProtocol interface {
	// 读取封包
	ReadPacket(s *Session) (interface{}, error)
	// 组装封包
	BuildPacket(pkgNode interface{}) []byte
	// 发送数据
	SendPacket(conn net.Conn, buff []byte) error
}

type Session struct {
	closed         int32
	conn           net.Conn
	sendChan       chan interface{}
	stopedChan     chan interface{}
	syncChan       chan interface{}
	closeCallback  func(*Session)
	sendCallback   func(*Session, interface{})
	packetHandler  PacketHandler
	packetProtocol IPacketProtocol
}

func NewSession(conn net.Conn, protocol IPacketProtocol, handler PacketHandler, sendChanSize int) *Session {
	return &Session{
		closed:         -1,
		conn:           conn,
		sendChan:       make(chan interface{}, sendChanSize),
		stopedChan:     make(chan interface{}),
		syncChan:       make(chan interface{}),
		packetHandler:  handler,
		packetProtocol: protocol,
	}
}

// 发起一个TCP请求
func Dial(network, address string, protocol IPacketProtocol, handler PacketHandler, sendChanSize int) (*Session, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return NewSession(conn, protocol, handler, sendChanSize), nil
}

// return net.Conn
func (s *Session) RawConn() net.Conn {
	return s.conn
}

// 关闭连接并释放相关资源.
func (s *Session) Close() error {
	if atomic.CompareAndSwapInt32(&s.closed, 0, 1) {
		s.conn.Close()
		close(s.stopedChan)
		if s.closeCallback != nil {
			s.closeCallback(s)
		}
	}
	return nil
}

func (s *Session) SetCloseCallback(callback func(*Session)) {
	s.closeCallback = callback
}

func (s *Session) SetSendCallback(callback func(*Session, interface{})) {
	s.sendCallback = callback
}

// 设置封包处理器, 接收到任何封包都会回调此处理器
func (s *Session) SetPacketHandler(handler PacketHandler) {
	s.packetHandler = handler
}

// SetProtocol can set a new IPacketProtocol.
func (s *Session) SetProtocol(protocol IPacketProtocol) {
	s.packetProtocol = protocol
}

// 设置发送队列缓冲区的大小
func (s *Session) SetSendChanSize(chanSize int) {
	s.sendChan = make(chan interface{}, chanSize)
}

// 获取发送队列缓冲区的大小
func (s *Session) GetSendChanSize() int {
	return cap(s.sendChan)
}

// 开始会话，循环监听发送与接收
func (s *Session) Start() {
	if atomic.CompareAndSwapInt32(&s.closed, -1, 0) {
		go s.sendLoop()
		go s.recvLoop()
	}
}

func (s *Session) recvLoop() {
	defer s.Close()

	for {
		recvBuff, err := s.packetProtocol.ReadPacket(s)
		if nil != err {
			time.Sleep(time.Millisecond*100)   // 等待0.1秒
			continue
		}

		if nil != recvBuff && utils.SizeStruct(recvBuff) > 0 {
			<- s.syncChan
			s.packetHandler(s, recvBuff) // 任务封包分发
		}
	}
}

func (s *Session) sendLoop() {
	defer s.Close()

	var err error

	for {
		select {
		case packet, ok := <-s.sendChan: {
				if !ok { return }

				pkgcnt := s.packetProtocol.BuildPacket(packet)
				err = s.packetProtocol.SendPacket(s.conn, pkgcnt)

				if err != nil { return }

				if s.sendCallback != nil { s.sendCallback(s, packet) }
			}
		case <-s.stopedChan: {
				fmt.Println("exit,exit,exit")
				return
			}
		}
	}
}

// AsyncSend queue the packet to the chan of send,
// if the send channel is full, return ErrSendChanBlocking.
// if the session had been closed, return ErrSessionClosed
//func (s *Session) AsyncSend(packet interface{}) error {
func (s *Session) Send(packet interface{}) error {
	select {
	case s.sendChan <- packet:
		s.syncChan <- packet
	case <-s.stopedChan:
		return ErrSessionClosed
	default:
		return ErrSendChanBlocking
	}
	return nil
}