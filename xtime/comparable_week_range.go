package xtime

import "fmt"

func (cwr *ComparableWeekRange) SameDurationLastWeekRange() *WeekRange {
	return cwr.SameDurationNWeeksAgoWeekRange(1)
}

func (cwr *ComparableWeekRange) SameDurationNWeeksAgoWeekRange(weeks int) *WeekRange {
	var nwr WeekRange
	nwr.from = cwr.from.AddDate(0, 0, -7*weeks)
	nwr.to = cwr.to.AddDate(0, 0, -7*weeks)
	return &nwr
}

func (cwr *ComparableWeekRange) SameDurationLastNWeeksWeekRange(weeks int) []*WeekRange {
	var nwr []*WeekRange
	for i := weeks; i > 0; i-- {
		nwr = append(nwr, cwr.SameDurationNWeeksAgoWeekRange(i))
	}
	return nwr
}

func (cwr *ComparableWeekRange) LastWeekRange() *WeekRange {
	var nwr WeekRange
	nwr.to = cwr.from.AddDate(0, 0, -1)
	nwr.from = nwr.to.AddDate(0, 0, -6)
	return &nwr
}

func (cwr *ComparableWeekRange) NWeeksAgoWeekRange(weeks int) *WeekRange {
	var nwr WeekRange
	fmt.Println(cwr.FromAsShortDate(), cwr.ToAsShortDate())
	nwr.to = cwr.from.AddDate(0, 0, -7*(weeks-1)-1)
	nwr.from = nwr.to.AddDate(0, 0, -6)
	return &nwr
}

func (cwr *ComparableWeekRange) LastNWeeksWeekRange(weeks int) []*WeekRange {
	var nwr []*WeekRange
	for i := weeks; i > 0; i-- {
		nwr = append(nwr, cwr.NWeeksAgoWeekRange(i))
	}
	return nwr
}
