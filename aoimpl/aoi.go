package aoimpl

type AOI interface {
	AddEntity(e Entity)
	Remove(e Entity)
	GetEntity(t EntityType)
}