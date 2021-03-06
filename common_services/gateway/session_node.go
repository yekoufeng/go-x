package main

import (
	"sync"

	"github.com/fananchong/go-x/common"
	"github.com/fananchong/go-x/common/k8s"
	discovery "github.com/fananchong/go-x/common/k8s/serverlist"
	"github.com/fananchong/go-x/common_services/proto"
	"github.com/fananchong/gotcp"
	proto1 "github.com/gogo/protobuf/proto"
)

type SessionNode struct {
	gotcp.Session
	endpoint *k8s.Endpoint
}

func (this *SessionNode) OnRecv(data []byte, flag byte) {
	if this.IsVerified() == false {
		this.doVerify(data, flag)
		return
	}
	cmd := proto.MsgTypeCmd(gotcp.GetCmd(data))
	switch cmd {
	case proto.MsgTypeCmd_Kick:
		msg := &proto.MsgKick{}
		if gotcp.DecodeCmd(data, flag, msg) != nil {
			if s, loaded := xaccounts.Load(msg.GetUID()); loaded {
				s.(*SessionAccount).Close()
				xaccounts.Delete(msg.GetUID())
			}
		}
	case proto.MsgTypeCmd_ForwardS:
		msg := &proto.MsgForwardS{}
		if gotcp.DecodeCmd(data, flag, msg) == nil {
			xlog.Debugln("decodeMsg fail.")
			return
		}
		if msg.GetType() == 0 {
			ForwardById(msg.GetId(), msg.GetData())
		} else {
			if msg.GetId() == 0 {
				Forward(common.ServerType(msg.GetType()), msg.GetData())
			} else {
				ForwardAll(common.ServerType(msg.GetType()), msg.GetData(), msg.GetId())
			}
		}
	default:
	}
}

func (this *SessionNode) doVerify(data []byte, flag byte) {
	msg := &proto.MsgVerifyS{}
	if gotcp.DecodeCmd(data, flag, msg) == nil {
		xlog.Errorln("decodeMsg fail.")
		this.Close()
		return
	}

	if this.endpoint.Id() != msg.GetId() {
		xlog.Errorln("verify fail. id error")
		this.Close()
		return
	}

	if msg.GetToken() != xargs.Common.IntranetToken {
		xlog.Errorln("verify fail. token error")
		this.Close()
		return
	}

	xnodes.Store(this.endpoint.Id(), this)
	xnodesMutex.Lock()
	defer xnodesMutex.Unlock()
	if _, ok := xnodesByType[this.endpoint.NodeType]; !ok {
		xnodesByType[this.endpoint.NodeType] = make(map[uint32]*SessionNode)
	}
	xnodesByType[this.endpoint.NodeType][this.endpoint.Id()] = this
	this.Verify()
	xlog.Debugln("Id:", msg.GetId(), "verify success.")
}

func (this *SessionNode) OnClose() {
	if this.endpoint != nil {
		discovery.GetNode().OnNodeLeave(this.endpoint)
	}
	if _, loaded := xnodes.Load(this.endpoint.Id()); loaded {
		xnodes.Delete(this.endpoint.Id())
	}
	xnodesMutex.Lock()
	defer xnodesMutex.Unlock()
	if items, ok := xnodesByType[this.endpoint.NodeType]; ok {
		if _, ok2 := items[this.endpoint.Id()]; ok2 {
			delete(items, this.endpoint.Id())
		}
		if len(items) == 0 {
			delete(xnodesByType, this.endpoint.NodeType)
		}
	}
}

func Forward(serverType common.ServerType, data []byte) {
	id, _, _ := xnode.Servers.GetOne(int(serverType))
	if id == 0 {
		xlog.Errorln("no find server. #1")
		return
	}
	ForwardById(id, data)
}

func ForwardById(id uint32, data []byte) {
	if node, loaded := xnodes.Load(id); loaded {
		// 转发的包不可以太大。这里硬编码flag为0
		node.(*SessionNode).Send(data, 0)
	} else {
		xlog.Errorln("no find server. #2")
		return
	}
}

func ForwardAll(serverType common.ServerType, data []byte, excludeId uint32) {
	xnodesMutex.RLock()
	defer xnodesMutex.RUnlock()
	if items, ok := xnodesByType[int(serverType)]; ok {
		for id, node := range items {
			if id != excludeId {
				// 转发的包不可以太大。这里硬编码flag为0
				node.Send(data, 0)
			}
		}
	}
}

func ForwardMsg(serverType common.ServerType, cmd proto.MsgTypeCmd, msg proto1.Message) {
	id, _, _ := xnode.Servers.GetOne(int(serverType))
	if id == 0 {
		xlog.Errorln("no find server. #1")
		return
	}
	if node, loaded := xnodes.Load(id); loaded {
		node.(*SessionNode).SendMsg(uint64(cmd), msg)
	} else {
		xlog.Errorln("no find server. #2")
		return
	}
}

func ForwardMsgAll(serverType common.ServerType, cmd proto.MsgTypeCmd, msg proto1.Message, excludeId uint32) {
	xnodesMutex.RLock()
	defer xnodesMutex.RUnlock()
	if items, ok := xnodesByType[int(serverType)]; ok {
		for id, node := range items {
			if id != excludeId {
				node.SendMsg(uint64(cmd), msg)
			}
		}
	}
}

// 由于Gateway功能相当简单，这里session管理，没有做成单例管理类。
// 请不要模仿这种不好的习惯:)
var xnodes sync.Map
var xnodesByType map[int]map[uint32]*SessionNode = make(map[int]map[uint32]*SessionNode)
var xnodesMutex sync.RWMutex
