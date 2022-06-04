package ComponentDefine

type EnumComponentType int8

type IBaseComponent interface {
	GetType() EnumComponentType
}

type BaseComponent struct {
	Type EnumComponentType
}

func (c BaseComponent) GetType() EnumComponentType {
	return c.Type
}

const (
	Unknown EnumComponentType = iota
	World
	Map
	MapDungeon
	EntityCollect
	Sprite  // 地图精灵
	Role    // 角色实体默认组件
	Dungeon // 副本
	DungeonPool
	Room
	Guild
	UniteState         //抽象出的单个状态
	Movement           //移动任务
	BTAgnet            // 行为树
	Email              //邮件
	Horse              //坐骑
	FightDamageFormula //伤害值计算公式
	Bag
	Pet         // 宠物
	Achievement //成就
	Activity    // 活动
	Designation //称号
	Execution   //执行单元
	Pk
	Wing
	Dress
	Social
	Equip
	EntityStatistics
	EntityGenerator
	Reputation  //声望
	StoragePack //仓库
	RegionLimit //区域限制
	CountGroup
	Buff
	DamageStatistics
	Follow      // 跟随
	Guide       //教学解锁
	Expedition  // 远征
	Talent      //被动技能
	BaseAttr    //基础属性 需要存库的血量、魔法值等
	Skill       // 技能数据类型
	EntitySight // 缓存能看到的周围的Entity信息
	QuestChapter
	GloryLevel       // 荣耀等级
	Store            //商城
	SkillMastery     //技能专精
	Item             //物品
	OnlineReward     //在线奖励
	Chat             //聊天
	Quest            //任务
	FestivalActivity //节日活动
	ActivityStore    //活动商城
	KilledStatistic  // 击杀统计
	Customize        // 定制
	MagicalAccident  //奇遇
)
