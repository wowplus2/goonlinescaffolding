// Code generated by "genprotocol -ver=68269c8bfeecd6e461aa862d64007f60a1aeccd64229d523ae7b99a446255112 -basedir=. -prefix=gos -statstype=int"

package gos_json

import (
	"encoding/json"
	"fmt"

	"github.com/kasworld/goonlinescaffolding/protocol_gos/gos_idcmd"
	"github.com/kasworld/goonlinescaffolding/protocol_gos/gos_idnoti"
	"github.com/kasworld/goonlinescaffolding/protocol_gos/gos_obj"
	"github.com/kasworld/goonlinescaffolding/protocol_gos/gos_packet"
)

// marshal body and append to oldBufferToAppend
// and return newbuffer, body type, error
func MarshalBodyFn(body interface{}, oldBuffToAppend []byte) ([]byte, byte, error) {
	var newBuffer []byte
	mdata, err := json.Marshal(body)
	if err == nil {
		newBuffer = append(oldBuffToAppend, mdata...)
	}
	return newBuffer, 0, err
}

func UnmarshalPacket(h gos_packet.Header, bodyData []byte) (interface{}, error) {
	switch h.FlowType {
	case gos_packet.Request:
		if int(h.Cmd) >= len(ReqUnmarshalMap) {
			return nil, fmt.Errorf("unknown request command: %v %v",
				h.FlowType, gos_idcmd.CommandID(h.Cmd))
		}
		return ReqUnmarshalMap[h.Cmd](h, bodyData)

	case gos_packet.Response:
		if int(h.Cmd) >= len(RspUnmarshalMap) {
			return nil, fmt.Errorf("unknown response command: %v %v",
				h.FlowType, gos_idcmd.CommandID(h.Cmd))
		}
		return RspUnmarshalMap[h.Cmd](h, bodyData)

	case gos_packet.Notification:
		if int(h.Cmd) >= len(NotiUnmarshalMap) {
			return nil, fmt.Errorf("unknown notification command: %v %v",
				h.FlowType, gos_idcmd.CommandID(h.Cmd))
		}
		return NotiUnmarshalMap[h.Cmd](h, bodyData)
	}
	return nil, fmt.Errorf("unknown packet FlowType %v", h.FlowType)
}

var ReqUnmarshalMap = [...]func(h gos_packet.Header, bodyData []byte) (interface{}, error){
	gos_idcmd.Invalid:     unmarshal_ReqInvalid,
	gos_idcmd.Login:       unmarshal_ReqLogin,
	gos_idcmd.Heartbeat:   unmarshal_ReqHeartbeat,
	gos_idcmd.MakeStage:   unmarshal_ReqMakeStage,
	gos_idcmd.EnterStage:  unmarshal_ReqEnterStage,
	gos_idcmd.ChatToStage: unmarshal_ReqChatToStage,
	gos_idcmd.LeaveStage:  unmarshal_ReqLeaveStage,
}

var RspUnmarshalMap = [...]func(h gos_packet.Header, bodyData []byte) (interface{}, error){
	gos_idcmd.Invalid:     unmarshal_RspInvalid,
	gos_idcmd.Login:       unmarshal_RspLogin,
	gos_idcmd.Heartbeat:   unmarshal_RspHeartbeat,
	gos_idcmd.MakeStage:   unmarshal_RspMakeStage,
	gos_idcmd.EnterStage:  unmarshal_RspEnterStage,
	gos_idcmd.ChatToStage: unmarshal_RspChatToStage,
	gos_idcmd.LeaveStage:  unmarshal_RspLeaveStage,
}

var NotiUnmarshalMap = [...]func(h gos_packet.Header, bodyData []byte) (interface{}, error){
	gos_idnoti.Invalid:   unmarshal_NotiInvalid,
	gos_idnoti.StageInfo: unmarshal_NotiStageInfo,
	gos_idnoti.StageChat: unmarshal_NotiStageChat,
}

func unmarshal_ReqInvalid(h gos_packet.Header, bodyData []byte) (interface{}, error) {
	var args gos_obj.ReqInvalid_data
	if err := json.Unmarshal(bodyData, &args); err != nil {
		return nil, err
	}
	return &args, nil
}

func unmarshal_RspInvalid(h gos_packet.Header, bodyData []byte) (interface{}, error) {
	var args gos_obj.RspInvalid_data
	if err := json.Unmarshal(bodyData, &args); err != nil {
		return nil, err
	}
	return &args, nil
}

func unmarshal_ReqLogin(h gos_packet.Header, bodyData []byte) (interface{}, error) {
	var args gos_obj.ReqLogin_data
	if err := json.Unmarshal(bodyData, &args); err != nil {
		return nil, err
	}
	return &args, nil
}

func unmarshal_RspLogin(h gos_packet.Header, bodyData []byte) (interface{}, error) {
	var args gos_obj.RspLogin_data
	if err := json.Unmarshal(bodyData, &args); err != nil {
		return nil, err
	}
	return &args, nil
}

func unmarshal_ReqHeartbeat(h gos_packet.Header, bodyData []byte) (interface{}, error) {
	var args gos_obj.ReqHeartbeat_data
	if err := json.Unmarshal(bodyData, &args); err != nil {
		return nil, err
	}
	return &args, nil
}

func unmarshal_RspHeartbeat(h gos_packet.Header, bodyData []byte) (interface{}, error) {
	var args gos_obj.RspHeartbeat_data
	if err := json.Unmarshal(bodyData, &args); err != nil {
		return nil, err
	}
	return &args, nil
}

func unmarshal_ReqMakeStage(h gos_packet.Header, bodyData []byte) (interface{}, error) {
	var args gos_obj.ReqMakeStage_data
	if err := json.Unmarshal(bodyData, &args); err != nil {
		return nil, err
	}
	return &args, nil
}

func unmarshal_RspMakeStage(h gos_packet.Header, bodyData []byte) (interface{}, error) {
	var args gos_obj.RspMakeStage_data
	if err := json.Unmarshal(bodyData, &args); err != nil {
		return nil, err
	}
	return &args, nil
}

func unmarshal_ReqEnterStage(h gos_packet.Header, bodyData []byte) (interface{}, error) {
	var args gos_obj.ReqEnterStage_data
	if err := json.Unmarshal(bodyData, &args); err != nil {
		return nil, err
	}
	return &args, nil
}

func unmarshal_RspEnterStage(h gos_packet.Header, bodyData []byte) (interface{}, error) {
	var args gos_obj.RspEnterStage_data
	if err := json.Unmarshal(bodyData, &args); err != nil {
		return nil, err
	}
	return &args, nil
}

func unmarshal_ReqChatToStage(h gos_packet.Header, bodyData []byte) (interface{}, error) {
	var args gos_obj.ReqChatToStage_data
	if err := json.Unmarshal(bodyData, &args); err != nil {
		return nil, err
	}
	return &args, nil
}

func unmarshal_RspChatToStage(h gos_packet.Header, bodyData []byte) (interface{}, error) {
	var args gos_obj.RspChatToStage_data
	if err := json.Unmarshal(bodyData, &args); err != nil {
		return nil, err
	}
	return &args, nil
}

func unmarshal_ReqLeaveStage(h gos_packet.Header, bodyData []byte) (interface{}, error) {
	var args gos_obj.ReqLeaveStage_data
	if err := json.Unmarshal(bodyData, &args); err != nil {
		return nil, err
	}
	return &args, nil
}

func unmarshal_RspLeaveStage(h gos_packet.Header, bodyData []byte) (interface{}, error) {
	var args gos_obj.RspLeaveStage_data
	if err := json.Unmarshal(bodyData, &args); err != nil {
		return nil, err
	}
	return &args, nil
}

func unmarshal_NotiInvalid(h gos_packet.Header, bodyData []byte) (interface{}, error) {
	var args gos_obj.NotiInvalid_data
	if err := json.Unmarshal(bodyData, &args); err != nil {
		return nil, err
	}
	return &args, nil
}

func unmarshal_NotiStageInfo(h gos_packet.Header, bodyData []byte) (interface{}, error) {
	var args gos_obj.NotiStageInfo_data
	if err := json.Unmarshal(bodyData, &args); err != nil {
		return nil, err
	}
	return &args, nil
}

func unmarshal_NotiStageChat(h gos_packet.Header, bodyData []byte) (interface{}, error) {
	var args gos_obj.NotiStageChat_data
	if err := json.Unmarshal(bodyData, &args); err != nil {
		return nil, err
	}
	return &args, nil
}
