package serverLib

import (
	"encoding/json"
)

type CmdRoomInfo struct {
	RoomInfo []*RoomInfo `json:"roomInfo"`
}

func (cri *CmdRoomInfo) String() string {
	s, err := json.Marshal(cri)
	if err != nil {
		return "roomInfo {}"
	}
	return "roomInfo " + string(s)
}

type CmdOutJoinRoom struct {
	RoomId string `json:"roomId"`
}

func (c *CmdOutJoinRoom) String() string {
	s, err := json.Marshal(c)
	if err != nil {
		return "joinRoom {}"
	}
	return "joinRoom " + string(s)
}
