package wsLogic

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

func (l *Hub) addUserConn(userId int64, platformEm int64, conn *UserConn) {
	rwLock.Lock()
	defer rwLock.Unlock()
	if oldConnMap, ok := l.wsUserToConn[userId]; ok {
		oldConnMap[platformEm] = conn
		l.wsUserToConn[userId] = oldConnMap
	} else {
		i := make(map[int64]*UserConn)
		i[platformEm] = conn
		l.wsUserToConn[userId] = i
	}
	fmt.Println("addUserConn mid")
	if oldStringMap, ok := l.wsConnToUser[conn]; ok {
		oldStringMap[platformEm] = userId
		l.wsConnToUser[conn] = oldStringMap
	} else {
		i := make(map[int64]int64)
		i[platformEm] = userId
		l.wsConnToUser[conn] = i
	}
	fmt.Println("addUserConn end")
	fmt.Println("userId, platformEM", userId, platformEm)
	count := 0
	for _, v := range l.wsUserToConn {
		count = count + len(v)
	}
	for _, v := range l.wsUserToConn {
		fmt.Println("v: ", v)
	}
}
func (l *Hub) getUserUid(conn *UserConn) (userId int64, platformEm int64) {
	rwLock.RLock()
	defer rwLock.RUnlock()
	if oldStringMap, ok := l.wsConnToUser[conn]; ok {
		for k, v := range oldStringMap {
			platformEm = k
			userId = v
		}
		return userId, platformEm
	}
	return 0, 0
}

func (l *Hub) delUserConn(conn *UserConn) {
	rwLock.Lock()
	defer rwLock.Unlock()
	var platformEm, userId int64
	if oldStringMap, ok := l.wsConnToUser[conn]; ok {
		for k, v := range oldStringMap {
			platformEm = k
			userId = v
		}
		if oldConnMap, ok := l.wsUserToConn[userId]; ok {
			// 因为将来可能有很多个平台
			delete(oldConnMap, platformEm)
			l.wsUserToConn[userId] = oldConnMap
			// 如果所有平台都下线了，就删除这个用户连接
			if len(oldConnMap) == 0 {
				delete(l.wsUserToConn, userId)
			}
		}
		delete(l.wsConnToUser, conn)
	}
	err := conn.Close()
	if err != nil {
		logx.WithContext(l.ctx).Error("close conn err ", "userId: ", userId, "platformEm: ", platformEm)
	}
}

func (l *Hub) GetUserConn(userId int64, platformEm int64) *UserConn {
	rwLock.RLock()
	defer rwLock.RUnlock()
	if oldConnMap, ok := l.wsUserToConn[userId]; ok {
		if conn, flag := oldConnMap[platformEm]; flag {
			return conn
		}
	}
	return nil
}

func (l *Hub) DelUserConn(userId int64, platformEm int64) {
	rwLock.RLock()
	defer rwLock.RUnlock()
	if oldConnMap, ok := l.wsUserToConn[userId]; ok {
		if len(oldConnMap) == 0 {
			delete(l.wsUserToConn, userId)
		} else {
			if _, flag := oldConnMap[platformEm]; flag {
				delete(oldConnMap, platformEm)
			}
		}
	}
}
