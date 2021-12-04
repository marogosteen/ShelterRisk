package simulation

import (
	"time"

	"example/OSURisk/person"
)

const (
	groupDurationMinute int = 15
	groupDurationSecond int = groupDurationMinute * 60
	breakFastHour       int = 7
	breakFastSecond     int = breakFastHour * 60 * 60
	lunchHour           int = 12
	lunchSecond         int = lunchHour * 60 * 60
	dinnerHour          int = 18
	dinnerSecond        int = dinnerHour * 60 * 60
)

type mealTime struct {
	mealGroupCount    int
	maxTimeDifference int
	interval          int
}

func newMealTime(people People, interval time.Duration) mealTime {
	mealGroupCount := ((len(people) - 1) / 25) + 1
	m := mealTime{
		mealGroupCount:    mealGroupCount,
		maxTimeDifference: groupDurationSecond * (mealGroupCount - 1),
		interval:          int(interval.Seconds()),
	}
	return m
}

// 時間条件にあったPersonのLifeActionをMealに変える。
func (m *mealTime) setMealTime(people People, currentDate time.Time) People {
	secondsFromMidnight := (currentDate.Hour()*60+currentDate.Minute())*60 + currentDate.Second()

	var nextMealSecond int
	if breakFastSecond <= secondsFromMidnight && secondsFromMidnight <= breakFastSecond+m.maxTimeDifference {
		nextMealSecond = breakFastSecond
	} else if lunchSecond <= secondsFromMidnight && secondsFromMidnight <= lunchSecond+m.maxTimeDifference {
		nextMealSecond = lunchSecond
	} else if dinnerSecond <= secondsFromMidnight && secondsFromMidnight <= dinnerSecond+m.maxTimeDifference {
		nextMealSecond = dinnerSecond
	} else {
		return people
	}

	for i := 0; i < m.mealGroupCount; i++ {
		nextGroupMealSecond := nextMealSecond + groupDurationSecond*i

		if secondsFromMidnight >= nextGroupMealSecond && nextGroupMealSecond+int(m.interval) > secondsFromMidnight {
			for id := i * 25; id < (i+1)*25; id++ {
				if id+1 > len(people) {
					break
				}
				people[id].LifeAction = person.Meal
				people[id].LifeActionElapsedSec = 0
				people[id].PassedCount = 0
			}
			break
		}
	}

	return people
}
