package main

import (
	"encoding/json"
	//	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/websocket"
)

/**************************************************************************
 * Function:
 *			开启http服务 监听端口
 * Description: systax --
 *             StartHTTP(addr string) error
 *		output :
 *						error
 *
 * Designer:
 *************************************************************************/
func StartHTTP(addr string) error {
	http.HandleFunc("/", DoAlarm)
	return http.ListenAndServe(addr, nil)
}

/****************************************************************
 *
 *  FUNCTION : This function call used by StartHttp（addr string） .
 *  处理收到的http包，解析json->检测IP合法性->查看报警等级加入相应的队列
 *
 *  DESCRIPTION :
 *         syntax  -- DoAlarm(w http.ResponseWriter, r *http.Request)
 *	input :
 *         http.ResponseWriter  -- w , http.Request -- r.
 * return value  :
 *
 *
 *****************************************************************/
func DoAlarm(w http.ResponseWriter, r *http.Request) {

	var upgrader = websocket.Upgrader{} // use default options
	c, err := upgrader.Upgrade(w, r, nil)

	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}

	body, _ := ioutil.ReadAll(r.Body)
	req := &Req{}
	err := json.Unmarshal(body, req)
	resp := &Resperr{}
	if err != nil {
		resp.ErrMsg = "json error"
		resp.ErrCode = errJson
		b, _ := json.Marshal(resp)
		w.Write(b)
		return
	}

	err = checkSrcIP(req.Hoster.SrcIP)
	if err != nil {
		resp.ErrMsg = err.Error()
		resp.ErrCode = errSrcIP
		b, _ := json.Marshal(resp)
		w.Write(b)
		return
	}

	G := req.Hoster.Grade
	switch {
	case G == 0:
		alarm.sender.SendNow(req)
	case G == 1:
		alarm.io.PacketChan1 <- req
	case G == 2:
		alarm.io.PacketChan2 <- req
	case G == 3:
		alarm.io.PacketChan3 <- req
	default:
		alarm.io.PacketChan3 <- req
	}
}
