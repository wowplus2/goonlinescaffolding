// Code generated by "genprotocol -ver=68269c8bfeecd6e461aa862d64007f60a1aeccd64229d523ae7b99a446255112 -basedir=. -prefix=gos -statstype=int"

package gos_handlenoti

/* obj base demux fn map template

var DemuxNoti2ObjFnMap = [...]func(me interface{}, hd gos_packet.Header, body interface{}) error {
gos_idnoti.Invalid : objRecvNotiFn_Invalid,
gos_idnoti.StageInfo : objRecvNotiFn_StageInfo,
gos_idnoti.StageChat : objRecvNotiFn_StageChat,

}

	func objRecvNotiFn_Invalid(me interface{}, hd gos_packet.Header, body interface{}) error {
		robj , ok := body.(*gos_obj.NotiInvalid_data)
		if !ok {
			return fmt.Errorf("packet mismatch %v", body )
		}
		return fmt.Errorf("Not implemented %v", robj)
	}

	func objRecvNotiFn_StageInfo(me interface{}, hd gos_packet.Header, body interface{}) error {
		robj , ok := body.(*gos_obj.NotiStageInfo_data)
		if !ok {
			return fmt.Errorf("packet mismatch %v", body )
		}
		return fmt.Errorf("Not implemented %v", robj)
	}

	func objRecvNotiFn_StageChat(me interface{}, hd gos_packet.Header, body interface{}) error {
		robj , ok := body.(*gos_obj.NotiStageChat_data)
		if !ok {
			return fmt.Errorf("packet mismatch %v", body )
		}
		return fmt.Errorf("Not implemented %v", robj)
	}

*/
