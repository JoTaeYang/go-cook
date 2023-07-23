package protocol

type ChatHeader struct {
	Type int32
	Len  int32
}

type EProto int32

const (
	REQE_ROOMCREATE EProto = 1 //방 생성
	REQE_ROOMIN     EProto = 2 //방 입장
	REQE_ROOMOUT    EProto = 3 //방 나가기
	REQE_ROOMLIST   EProto = 4 //방 목록 조회

	ANSE_ROOMCREATE EProto = 1001 //방 생성 응답
	ANSE_ROOMIN     EProto = 1002 //방 입장 응답
	ANSE_ROOMOUT    EProto = 1003 //방 나가기 응답
	ANSE_ROOMLIST   EProto = 1004 //방 목록 응답
)

type ProtoRoomCreateReq struct {
	RoomName []byte
}

type ProtoRoomCreateAns struct {
	RoomNo   int32
	RoomName []byte
}

type ProtoRoomIn struct {
}
