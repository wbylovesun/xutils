# xutils
golang基础扩展工具，例如string, slice, time等

golang>=1.18

## xstrings

### 数字转换
- func MustInt(s string) int; // Panic when s is not numeric
- func ToInt(s string) int; // Ignore errors when s is not numeric
### 摘要
- func MD5(v string) string;
- func SHA1(v string) string;
- func SHA256(v string) string;
- func SHA512(v string) string;
- func SHA512384(v string) string;
### 转切片
- func ToIntSlice(str string, separation string) []int;

### 字符串功能
- func Lcfirst(s string) string;
- func Ucfirst(s string) string;
- func StartsWith(s, p string) bool;
- func EndsWith(s, p string) bool

## xslices

切片的工具，如Contains

- func Contains[T SliceElementType](data []T, ct T) bool;

## xtime

### DateRange
创建一个time.Time的时间范围
```go
type DateRange struct {
	from time.Time
	to   time.Time
}
```
### 快捷方法
- func ISOYearWeek(t time.Time) int;
- func YearMonth(t time.Time) int;
- func YearQuarter(t time.Time) int;
- func ShortDate(t time.Time) string;
- func LongDate(t time.Time) string;
- func YmdDate(t time.Time) string;
- func WithTime(t time.Time, hour, minute, second int) time.Time;
- func MonthFirstDay(t time.Time) time.Time;
- func MonthLastDay(t time.Time) time.Time;
- func IsLeap(year int) bool;
- func DaysIn(t time.Time) int;
- func DaysInYearMonth(y, m int) int;
- func LastSecondOf(t time.Time) time.Time;
- func BuildFormat(template string) string;
- func MustOf(ts string) time.Time;
- func Parse(ts string) (time.Time, error);

- func ThisYearRange(includeToday ...bool) *DateRange;
- func LastYearRange() *DateRange;
- func Today() time.Time;
- func TodayLongDate() string;
- func TodayShortDate() string;
- func TodayYmdDate() string;
- func Yesterday() time.Time;
- func YesterdayLongDate() string;
- func YesterdayShortDate() string;
- func YesterdayYmdDate() string;
- func Tomorrow() time.Time;
- func TomorrowLongDate() string;
- func TomorrowShortDate() string;
- func TomorrowYmdDate() string;
- func ISOLastWeek() int;
- func FirstDayOfThisMonth() time.Time;
- func LastDayOfThisMonth() time.Time;
- func Latest30Days() *DateRange;
- func Passed30Days() *DateRange;
- func LatestNDays(n int) *DateRange;
- func PassedNDays(n int) *DateRange;
- func Lastest1Year() *DateRange;
- func Passed1Year() *DateRange;
- func Latest12Months() *MonthRange;
- func Passed12Months() *MonthRange;
- func LatestNMonths(n int) *MonthRange;
- func PassedNMonths(n int) *MonthRange;


## xvalidator

### 使用方法
```go
// 注册新的v10版本校验器
v := xvalidator.NewValidator()
var _ binding.StructValidator = v
binding.Validator = v
```

### time_span
判定是否时间差超过指定的范围。
以";"分隔检查参数：
- 第一个参数是对比的字段
- 支持gt, lt, lte, gte
- gt, gte表示最少要达到多少时间间隔
- lt, lte表示最多不超过多少时间间隔
```go
type testReq struct {
	S time.Time `form:"s" binding:"required,lte=today" time_format:"2006-01-02"`
	T time.Time `form:"t" binding:"required,lte=today,time_span=S;gte:1day;lte:30day" time_format:"2006-01-02"`
}
```

### int_slice
用于判定字符串以某个特定的字符串分隔的校验
- sep 分隔的字符
- 支持gt, gte, lt, lte，检查每个元素是否符合条件
```go
type inOutRequest struct {
	Nodes string `form:"nodes" binding:"omitempty,int_slice=gt:0"`
}
```