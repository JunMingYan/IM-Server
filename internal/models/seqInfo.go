package models

type SeqInfo struct {
	ID       int64  // 主键
	CollName string // 集合名称
	SeqId    int32  // 序列值
}
