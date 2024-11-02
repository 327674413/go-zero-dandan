package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
)

// 群组元信息结构体
type GroupMetaInfo struct {
	GroupName string           // 群组名称
	Members   map[int64]Member // 群组成员（mapid -> Member）
	MaxID     int64            // 当前群组中最大的 mapid
	QuitBit   []uint64         // 记录哪些成员已退出的位图
}

// 群组成员结构体
type Member struct {
	UserID int64 // 成员的用户ID
	MapID  int64 // 成员的映射ID
}

// 消息详情结构体
type Message struct {
	MessageID int64    // 消息唯一标识符
	ReadBit   []uint64 // 记录每个 mapid 对应的用户是否已读的位图
}

// 初始化群组
func NewGroupMetaInfo(groupName string) *GroupMetaInfo {
	return &GroupMetaInfo{
		GroupName: groupName,
		Members:   make(map[int64]Member),
		QuitBit:   make([]uint64, 1), // 初始化 QuitBit 为 1 个元素
	}
}

// 添加群组成员
func (g *GroupMetaInfo) AddMember(userID int64) {
	mapID := g.MaxID
	member := Member{UserID: userID, MapID: mapID}
	g.Members[mapID] = member
	g.MaxID++
	// 确保 QuitBit 大小足够
	if mapID/64 >= int64(len(g.QuitBit)) {
		g.QuitBit = append(g.QuitBit, 0)
	}
}

// 删除群组成员（逻辑删除）
func (g *GroupMetaInfo) RemoveMember(mapID int64) {
	if _, exists := g.Members[mapID]; exists {
		g.QuitBit[mapID/64] |= (1 << (mapID % 64)) // 在退出位图中标记为已退出
	}
}

// 初始化消息
func NewMessage(messageID int64, maxID int64) *Message {
	return &Message{
		MessageID: messageID,
		ReadBit:   make([]uint64, (maxID+63)/64), // 每64个用户用一个uint64存储
	}
}

// 成员读取消息
func (m *Message) MarkAsRead(mapID int64) {
	m.ReadBit[mapID/64] |= (1 << (mapID % 64)) // 使用位操作记录已读
}

// 检查成员是否已读消息
func (m *Message) IsRead(mapID int64) bool {
	if mapID/64 >= int64(len(m.ReadBit)) {
		return false // 避免越界访问
	}
	return m.ReadBit[mapID/64]&(1<<(mapID%64)) != 0
}

// 将位图转换为字节数组
func BitmapToBytes(bitmap []uint64) []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, bitmap)
	if err != nil {
		log.Fatal("Failed to convert bitmap to bytes:", err)
	}
	return buf.Bytes()
}

// 将字节数组还原为位图
func BytesToBitmap(data []byte) []uint64 {
	buf := bytes.NewReader(data)
	var bitmap []uint64
	err := binary.Read(buf, binary.LittleEndian, &bitmap)
	if err != nil {
		log.Fatal("Failed to convert bytes to bitmap:", err)
	}
	return bitmap
}

// 存储消息的位图数据
func (m *Message) SaveToDatabase() []byte {
	return BitmapToBytes(m.ReadBit)
}

// 从数据库加载消息的位图数据
func (m *Message) LoadFromDatabase(data []byte) {
	m.ReadBit = BytesToBitmap(data)
}

// 打印群组信息
func (g *GroupMetaInfo) PrintGroupInfo() {
	fmt.Printf("Group: %s\n", g.GroupName)
	fmt.Println("Members:")
	for mapID, member := range g.Members {
		status := "Active"
		if mapID/64 < int64(len(g.QuitBit)) && g.QuitBit[mapID/64]&(1<<(mapID%64)) != 0 {
			status = "Quit"
		}
		fmt.Printf("  UserID: %d, MapID: %d, Status: %s\n", member.UserID, member.MapID, status)
	}
}

// 测试代码
func main() {
	// 初始化群组
	group := NewGroupMetaInfo("Golang Developers")

	// 添加群组成员
	group.AddMember(101)
	group.AddMember(102)
	group.AddMember(103)

	// 初始化消息
	message := NewMessage(1, group.MaxID)

	// 成员读取消息
	message.MarkAsRead(2)

	// 检查成员读取状态
	fmt.Printf("User with MapID 2 has read the message: %v\n", message.IsRead(2))

	// 删除成员
	group.RemoveMember(3)

	// 打印群组信息
	group.PrintGroupInfo()

	// 将位图数据保存到数据库（模拟）
	savedData := message.SaveToDatabase()

	// 从数据库加载位图数据（模拟）
	newMessage := NewMessage(2, group.MaxID)
	newMessage.LoadFromDatabase(savedData)

	// 检查加载后的位图是否正确
	fmt.Printf("Loaded Message ReadBit for MapID 2: %v\n", newMessage.IsRead(2))
}
