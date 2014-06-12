package event

const (
	ACTION_SUCCESS = 1
	ACTION_FAILED  = 2
	ACTION_REJECT  = 3 // 拒绝

	CONNECT_PACKET_TYPE      = 0x01 // 连接指令
	PING_COMMAND             = 0x03 // ping心跳指令
	PONG_COMMAND        byte = 0x04 // pong心跳指令
	LOGOUT_COMMAND      byte = 0x05 // 退出正在玩的游戏

	/**
	 * 玩家玩游戏过程中的type
	 */
	LOG_IN            = 0x01 // 连接游戏，进入一局
	RES_LOCK          = 0x03
	RES_LOCK_COMPLETE = 0x04

	GAME_END   = 0x07
	GAME_RESET = 0x08

	START = 0x1a
	OVER  = 0x1b

	USER_CLOSE_GAME_CODE = 0x05
	USER_OFFLINE_CODE    = 0x06

	LOGIN_GAME  = 0x10 // 登陆游戏
	LOGOUT_GAME = 0x11 // 登出游戏

	GET_FRIENDS     = 0x12 // 获取已经安装该游戏的好友
	ADD_FRIENDS     = 0x13 // 添加好友到chatgame联系列表
	INVITE_PLAYER   = 0x14 // 邀请玩家玩游戏
	REPLY_INVITE    = 0x15 // 响应玩家邀请
	JOIN_WAIT_QUEUE = 0x16 // 进入等待队列，准备随机匹配
	PLAYER_READY    = 0x17 // 玩家就绪
	GAME_START      = 0x18 // 游戏开始

	SAVE_GAME_INFO = 0x30
	SAVE_SEX       = 0x31 // 设置性别
	SCORE_LIST     = 0x32 // 查询游戏排行榜

	SERVER_ERR = 0x50 // 服务内部异常
)

type Event struct {
	code int `json:"type"`
}

type LoginReqEvent struct {
	Event
	SessionId string
	appId     int
}
