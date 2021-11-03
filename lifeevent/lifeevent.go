package lifeevent

// TODO ご飯イベントを掘り下げる。全員でいくのか、時間はどうかなどなどしっかり考える。
type LifeEvent struct {
	EventNames       []string
	EventSeconds     []int // map使う
	EventelapsedTime int   // イベント経過時間
}

func NewLifeEvent() *LifeEvent {
	lifeEvent := new(LifeEvent)
	lifeEvent.EventNames = []string{
		"Stay", "CheckBoard", "ChangeClothes", "BathRoom", "Eat",
	}
	// TODO 先生にイベント時間の意見を聞く
	// 移動量ではなく、移動時間にした理由は、Person.Move()では移動しないという選択があるから。
	lifeEvent.EventSeconds = []int{
		0, 0, 0, 0,
	}
	return lifeEvent
}
