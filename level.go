package main

import (
	//	"fmt"
	"runtime"
	"time"
)

type PacketByName map[string][]Req

type PacketIO struct {
	PacketChan1 chan *Req
	PacketChan2 chan *Req
	PacketChan3 chan *Req
}

/*********************************************************************
 * Function: 初始化Io对象 创建三个通信管道
 * Description:
 *	syntax --
 *          RoutineInit() (io *PacketIO)
 * 	input :
 *
 * 	return value : io *PacketIO
 * Designer:
 **************************************************************************/
func RoutineInit() (io *PacketIO) {
	io = &PacketIO{}

	io.PacketChan1 = make(chan *Req, 1024)
	io.PacketChan2 = make(chan *Req, 1024)
	io.PacketChan3 = make(chan *Req, 1024)

	go io.RoutineStart(2)
	go io.RoutineStart(4)
	go io.RoutineStart(6)

	return
}

/*********************************************************************
 * Function: 处理信息线程， 取出管道信息将其打包发送
 * Description:
 *	syntax --
 *          (io *PacketIO)RoutineStart(timeAfter int)
 * 	input :
 *				timeAfter int   需要接受的管道信息和打包发送时间
 * 	return value : io *PacketIO
 * Designer:
 **************************************************************************/
func (io *PacketIO) RoutineStart(timeAfter int) {
	var pChan chan *Req
	switch timeAfter {
	case 2:
		pChan = io.PacketChan1
	case 4:
		pChan = io.PacketChan2
	case 6:
		pChan = io.PacketChan3
	}
	q := QueueInit()
	go SendLoop(q, timeAfter)
	for {
		PtrReq := <-pChan
		qNode := &QueueNode{PtrReq, nil}
		q.PushQueueNode(qNode)
	}

}

/*********************************************************************
 * Function: 间隔循环队列信息 队列不为空则开始计时 到时间发送数据
 * Description:
 *	syntax --  SendLoop(q *Queue, timeAfter int)
 *
 * 	input :
 *				timeAfter int   需要接受的管道信息和打包发送时间
 *				q *Queue: 队列信息
 * 	return value :
 * Designer:
 **************************************************************************/
func SendLoop(q *Queue, timeAfter int) {
	for {
		for q.IsEmpty() {
			runtime.Gosched()
		}
		sendmap := make(PacketByName)
		time.Sleep(time.Second * time.Duration(timeAfter))
		var TailPtr *QueueNode
		TailPtr = q.tail
		for q.head.Next != TailPtr.Next {
			popNode := q.PopQueueNode()
			if popNode != nil {
				sendmap[popNode.PtrReq.Hoster.Email] = append(sendmap[popNode.PtrReq.Hoster.Email], *popNode.PtrReq)
			}
		}
		if len(sendmap) != 0 {
			go alarm.sender.Send(sendmap)
		}
	}
}
