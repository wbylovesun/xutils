package xtime

func (cwr *ComparableWeekRange) LastWeekRange() *WeekRange {
	return cwr.NWeeksAgoWeekRange(1)
}

func (cwr *ComparableWeekRange) NWeeksAgoWeekRange(weeks int) *WeekRange {
	var nwr WeekRange
	nwr.from = cwr.from.AddDate(0, 0, -7*weeks)
	nwr.to = cwr.to.AddDate(0, 0, -7*weeks)
	return &nwr
}

func (cwr *ComparableWeekRange) LastNWeeksWeekRange(weeks int) []*WeekRange {
	var nwr []*WeekRange
	for i := 0; i < weeks; i++ {
		nwr = append(nwr, cwr.NWeeksAgoWeekRange(i))
	}
	return nwr
}
