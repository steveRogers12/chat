package redislib

const (
	GetUserIdByToken              = "GetUserIdByToken_" // 根据token获取用户id的
	GetUserTokenById              = "GetUserTokenById_" // 根据id获取用户token的
	GetUserInfoById               = "GetUserInfoById_" // 用户信息
	GetUserIdByName               = "GetUserIdByName_" // 获取用户id的redis key
	RegisterCodeByPhone           = "RegisterCodeByPhone_" // 注册短信验证码key
	LimitNextSendCodeByPhone      = "LimitNextSendCodeByPhone_"// 限制发送短信下次的时间
	CacheUserInfo				  = "CacheUserInfo_" // 缓存用户信息
	CacheFriends				  = "CacheFriends"	// 缓存用户好友
	CacheGroups					  = "CacheGroups"	// 缓存群信息
	CacheMessages				  = "CacheMessages" // 最近消息
)