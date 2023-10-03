package trading

import "sort"

// unsafeDateIsTradingDay date是否交易日?默认是今天
func unsafeDateIsTradingDay(date ...string) bool {
	theDay := Today()
	if len(date) > 0 {
		theDay = FixTradeDate(date[0])
	}
	lastDay := unsafeLastTradeDate()
	if lastDay == theDay {
		return true
	}
	return false
}

// unsafeLastTradeDate 获得最后一个交易日
func unsafeLastTradeDate() string {
	today := IndexToday()
	tradeDates := unsafeDates()
	end := sort.SearchStrings(tradeDates, today)
	lastDay := tradeDates[end]
	if lastDay > today {
		end = end - 1
	}
	return tradeDates[end]
}