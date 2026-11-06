# xutils

golang 基础扩展工具，例如 string, slice, time, ptr, calc 等。

## 环境要求

Go >= 1.21（使用泛型与标准库 `cmp.Ordered`）

## 安装

```bash
go get github.com/wbylovesun/xutils
```

## 包总览

| 包 | 说明 |
|---|---|
| `types` | 泛型数值约束（`Number` / `Integer` / `SignedNumber` / `UnsignedNumber`） |
| `xutils`（根包） | 任意值的深拷贝 `Copy` |
| `xstrings` | 字符串：默认值、摘要、大小写、数值解析与转切片 |
| `xslice` | 切片工具：排序、查找、集合运算、数值转换、指针切片互转 |
| `xptr` | 指针工具：取指针、解引用、数值指针互转、nil-safe 操作 |
| `xcalc` | 数值计算：百分比、增长率、金额格式化、数值比较、字节转换 |
| `xtime` | 时间：范围、快捷方法、时间段切片 `TimeSpan` |
| `xmath` | 数学辅助：四舍五入 |
| `xrand` | 随机数、随机字符串、ULID |
| `xserializer` | 序列化辅助类型 |
| `xvalidator` | gin 绑定校验器扩展 |

## types

泛型数值约束，全部带 `~`，支持底层类型为数值的自定义类型（如 `type UserID int64`、`type Amount float64`）。

```go
type SignedNumber interface { ~int | ~int8 | ~int16 | ~int32 | ~int64 }
type UnsignedNumber interface { ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 }
type Integer interface { SignedNumber | UnsignedNumber }               // 全部整型
type Number interface { SignedNumber | UnsignedNumber | ~float32 | ~float64 } // 全量数值
```

与标准库 `cmp.Ordered` 的分工：需要比较大小或包含 `string` 时用 `cmp.Ordered`；纯数值算术/格式化用 `types.Number`。

## xutils（根包）

深拷贝，基于反射递归复制（切片、Map、结构体、指针、interface），`time.Time` 按值处理：

```go
func Copy(src interface{}) interface{} // 深拷贝，返回值需类型断言
func Iface(src interface{}) interface{} // Copy 的兼容别名
```

自定义拷贝行为：类型实现 `Interface`（`DeepCopy() interface{}`）即可被递归调用。

## xstrings

### 默认值

```go
func Default(s, d string) string          // s 为空时返回 d
func DefaultWithTrim(s, d string) string  // 先判断 s 是否为空（含空白）
func DefaultAfterTrim(s, d string) string // 返回前对 s 做 trim
```

### 数字转换

```go
func MustInt(s string) int // s 非数字时 panic
func ToInt(s string) int   // s 非数字时忽略错误，返回 0
```

### 摘要

```go
func MD5(v string) string
func SHA1(v string) string
func SHA256(v string) string
func SHA512(v string) string
func SHA512384(v string) string
func CRC32(v string) uint32
func SHA256WithHMAC(v, k string) ([]byte, error)
func SHA256WithHMAC2Hex(v, k string) (string, error)
func SHA256WithHMAC2Base64Str(v, k string) (string, error)
```

### 字符串功能

```go
func Lcfirst(s string) string    // 首字母小写
func Ucfirst(s string) string    // 首字母大写
func StartsWith(s, p string) bool
func EndsWith(s, p string) bool
```

### 转切片

```go
func ToIntSlice(str, separation string) []int
func ToFloat64Slice(str, separation string) []float64
func ToStringSlice(str, separation string) []string

// 泛型版：T 为任意数值类型；非法片段自动跳过
func ToNumberSlice[T types.Number](str, separation string, to T) []T
```

## xslice

### 基础操作

```go
func Clone[T any](data []T) []T                       // 浅拷贝（新底层数组），nil 输入返回空切片
func Sort[T cmp.Ordered](data []T)                    // 原地升序排序
func CloneSort[T cmp.Ordered](data []T) []T           // 排序后的新副本，不改原切片
func Walk[T any](data []T, f func(v T, k int) T)      // 遍历并原位替换元素
func Map[T any](f func(v T) T, data ...[]T) []T       // 多个切片合并映射
func ConvertTo[T any, O any](from []T, f func(T) O) []O // 逐元素转换
func PluckToMap[T any, K comparable, V any](from []T, f func(T) (K, V)) map[K, V] // 键冲突时后者覆盖
```

### 查找与判断

```go
func Contains[T comparable](data []T, ct T) bool     // 是否包含元素
func ContainsSlice[T cmp.Ordered](data, ct []T) bool // ct 中每个元素是否都存在于 data（双方先去重，集合语义，不考虑重复次数）
func Filter[T any](data []T, f func(v T, k int) bool) []T // 剔除语义：f 返回 true 的元素被剔除
```

### 二分查找

```go
// 注意：查找前会先对 list 原地排序（修改调用方传入的切片），返回的下标基于排序后的顺序；未找到返回 -1
func BisectionSearch[T cmp.Ordered](list []T, target T) int
```

### 集合运算

```go
func Unique[T comparable](elements []T) []T        // 去重（保序）
func Distinct[T comparable](elements []T) []T      // Unique 别名
func Intersection[T comparable](A, B []T) []T      // 交集
func Union[T comparable](A, B []T) []T             // 并集（去重）
func CloneDiff[T cmp.Ordered](A, B []T) []T        // A 中不在 B 里的元素（不改入参）
func Diff[T cmp.Ordered](A, B []T) []T             // 同 CloneDiff，但会原地排序 A 和 B
```

### 极值

```go
func Min[T cmp.Ordered](data []T) *T // 返回指向最小值的指针，空切片返回 nil
func Max[T cmp.Ordered](data []T) *T
```

### 数值转换

```go
func Join[T types.Number](t []T, sep string) string // 数值切片连接为字符串；float 按自身精度最短表示输出，小数为 0 时舍弃

func ToNumberSlice[T types.Integer, U types.SignedNumber](t []T, u U) []U // 逐元素转换，u 仅用于类型推断（传零值）
func ToInt64Slice[T types.Integer](t []T) []int64
func ToInt32Slice[T types.Integer](t []T) []int32
func ToInt16Slice[T types.Integer](t []T) []int16
func ToInt8Slice[T types.Integer](t []T) []int8
func ToIntSlice[T types.Integer](t []T) []int
```

转换遵循 Go 语义：整数溢出回绕（`300 → int8 44`），float 小数截断；空切片/nil 返回 nil。

### 指针切片互转

```go
func PtrSlice[V any](vals []V) []*V   // []V → []*V，每个指针指向元素的独立副本
func ValueSlice[V any](vals []*V) []V // []*V → []V，nil 元素取零值
```

### 范围

```go
func Range(from, to, delta int) []int // [from, to] 内以 delta 为步长（含两端）；from > to 返回 nil
```

## xptr

### 取指针与解引用

```go
func Ptr[V any](val V) *V                        // 返回指向 val 副本的指针
func PtrValueOrEmpty[V any, T *V](val T) V       // 解引用，nil 时返回零值
func PtrValueOrDefault[V any, T *V](val T, def V) V // 解引用，nil 时返回默认值
func ClonePtr[V any](val *V) *V                  // 复制值到新指针，nil 透传
```

### 数值指针互转（nil 透传；溢出回绕、float 截断）

```go
func PtrToIntPtr[T canConvertIntPointerTypes](dt *T) *int // *T → *int（T 为内置数值类型）
func IntPtrTo[T canConvertIntPointerTypes](dt *int) *T    // *int → *T
func PtrTo[T, O canConvertIntPointerTypes](dt *T) *O      // 任意两种数值类型指针互转
```

### nil-safe 比较与合并

```go
func EqualPtr[V comparable](a, b *V) bool // 比较指向的值：双 nil 相等，单 nil 不等
func Coalesce[T any](vals ...*T) *T       // 返回第一个非 nil 的原指针，全 nil 返回 nil
func CoalesceValue[T any](def T, vals ...*T) T // 第一个非 nil 指针的值，全 nil 返回 def
```

## xcalc

### 百分比（占比：numerator / denominator × 100）

```go
func Percentage[T types.Number](numerator, denominator T) (float64, error)          // 2 位小数，分母为 0 报错
func PercentageNoDigit[T types.Number](numerator, denominator T) (int, error)       // 不带小数，报错
func PercentagePrecision[T types.Number](numerator, denominator T, precision int) (float64, error) // 自定义小数位，报错
func PercentageWithTwoDigits[T types.Number](numerator, denominator T) float64      // 2 位小数，忽略错误
func PercentageWithSuppression[T types.Number](numerator, denominator T) float64    // 2 位小数，错误/NaN/Inf 抑制为 0
func PercentageNoDigitWithSuppression[T types.Number](numerator, denominator T) int
func PercentagePrecisionWithSuppression[T types.Number](numerator, denominator T, precision int) float64
```

### 辅助

```go
func NumberPrecision[T types.Number](numerator, denominator T, precision int) float64 // 普通除法保留 precision 位小数，分母为 0 返回 0
func PromisePercentWithTwoDigits(n float64) float64 // 百分比数值保留（最多）2 位小数
func PromiseWithTwoDigits(n float64) float64        // 非百分比 ×100 后保留 2 位小数
func Round(val float64, precision int) float64      // ROUND_HALF_UP 四舍五入，precision 可为负/零
```

### 增长率（now 相对 pre 的变化百分比）

```go
func Ratio(now, pre float64, precision int) float64 // (now/pre-1)*100 保留 precision 位
func PercentageRatio(now, pre float64) float64      // Ratio 的 2 位小数版
```

特殊值约定：`now == 0 && pre == 0 → 0`；`now == 0 → -100`；`pre == 0 → +Inf`。

### 金额

```go
func FormatAmountToYuan[T AmountType](amount T) float64 // 非 float 视为分 ÷100 转元；float 直接保留 2 位小数
```

### 数值比较（钳制与判断）

```go
func Min[T types.Number](z ...T) T // 空变参返回 T(0)，单元素返回自身
func Max[T types.Number](z ...T) T

func PromiseGte[T types.Number](x, y T) T  // 下限钳制：x < y 时返回 y
func PromiseGte0[T types.Number](x T) T    // 承诺 >= 0
func PromiseLte[T types.Number](x, y T) T  // 上限钳制：x > y 时返回 y
func PromiseLte100[T types.Number](x T) T  // 承诺 <= 100

func IsGte[T types.Number](x, y T) bool // x >= y
func IsLte[T types.Number](x, y T) bool
func IsGt[T types.Number](x, y T) bool
func IsLt[T types.Number](x, y T) bool
```

> `Gte` / `Gte0` / `Lte` / `Lte100` 为旧版，已标记 Deprecated，请使用 `Promise*` 系列。

### 字节转换

```go
func Int16ToBytes(i int16) []byte
func Int32ToBytes(i int32) []byte
func Int64ToBytes(i int64) []byte
func BytesToInt16(b []byte) (int16, error)
func BytesToInt32(b []byte) (int32, error)
func BytesToInt64(b []byte) (int64, error)
```

## xtime

### DateRange

创建一个 time.Time 的时间范围

```go
type DateRange struct {
	from time.Time
	to   time.Time
}
```

另提供 `ComparableDateRange` / `WeekRange` / `ComparableWeekRange` / `MonthRange` / `ComparableMonthRange` 范围类型，以及 `Date` / `DateTime` / `JsonTimestamp` 自定义时间类型。

### 快捷方法

```go
func ISOYearWeek(t time.Time) int
func YearMonth(t time.Time) int
func YearQuarter(t time.Time) int
func ShortDate(t time.Time) string
func LongDate(t time.Time) string
func YmdDate(t time.Time) string
func WithTime(t time.Time, hour, minute, second int) time.Time
func MonthFirstDay(t time.Time) time.Time
func MonthLastDay(t time.Time) time.Time
func IsLeap(year int) bool
func DaysIn(t time.Time) int
func DaysInYearMonth(y, m int) int
func LastSecondOf(t time.Time) time.Time
func BuildFormat(template string) string
func MustOf(ts string) time.Time
func Parse(ts string) (time.Time, error)

func ThisYearRange(includeToday ...bool) *DateRange
func LastYearRange() *DateRange
func Today() time.Time
func TodayLongDate() string
func TodayShortDate() string
func TodayYmdDate() string
func Yesterday() time.Time
func YesterdayLongDate() string
func YesterdayShortDate() string
func YesterdayYmdDate() string
func Tomorrow() time.Time
func TomorrowLongDate() string
func TomorrowShortDate() string
func TomorrowYmdDate() string
func ISOLastWeek() int
func FirstDayOfThisMonth() time.Time
func LastDayOfThisMonth() time.Time
func Latest30Days() *DateRange
func Passed30Days() *DateRange
func LatestNDays(n int) *DateRange
func PassedNDays(n int) *DateRange
func Lastest1Year() *DateRange
func Passed1Year() *DateRange
func Latest12Months() *MonthRange
func Passed12Months() *MonthRange
func LatestNMonths(n int) *MonthRange
func PassedNMonths(n int) *MonthRange
```

### TimeSpan

按固定分钟跨度（`span`）将一天（00:00–24:00）切分为时间段：

```go
ts, err := xtime.NewTimeSpan(60) // 每 60 分钟一段
```

```go
func (ts *timeSpan) Segs() []string                                // 全部时间段，如 00:00、01:00 …
func (ts *timeSpan) TodaySegs() []string                           // 截至当前时间的段
func (ts *timeSpan) Index(i int) int                               // 分钟偏移 → 段下标
func (ts *timeSpan) IndexByTime(t string) (int, error)             // 时刻 → 段下标
func (ts *timeSpan) Of(i int) string                               // 段下标 → 时刻
func (ts *timeSpan) OfNextSeg(i int) string                        // 下一段时刻
func (ts *timeSpan) OfPrevSeg(i int) string                        // 上一段时刻
func (ts *timeSpan) OfTime(t string) (string, error)               // 时刻 → 所在段
func (ts *timeSpan) AlignToNextSeg(t string) (string, error)       // 对齐到下一段
func (ts *timeSpan) AlignToPrevSeg(t string) (string, error)       // 对齐到上一段
func (ts *timeSpan) TimeOfNextAlignSeg(t time.Time) (time.Time, error)
func (ts *timeSpan) TimeOfPrevAlignSeg(t time.Time) (time.Time, error)
func (ts *timeSpan) RangeSegs(from, to string) []string            // [from, to] 区间内的段
func (ts *timeSpan) Fill(m map[PointOfTimeSlice]int, pots ...PointOfTimeSlice) map[PointOfTimeSlice]int
func (ts *timeSpan) FillForStrKey(m map[string]int, pots ...string) map[string]int
```

`Fill` / `FillForStrKey`：从 00:00 起按段补全缺失数据（沿用最后一个已知值），`pots[0]` 指定截止时刻（缺省为当前时间）；传 `"24:00"` 表示全天（输出 00:00–23:00，不含 24:00 段）。

## xmath

```go
func Round(x float64) int // 四舍五入取整
```

## xrand

```go
func Int(min, max int) int // [min, max] 内的随机整数
func ULID() string         // 生成 ULID 字符串

func Chars(l int) string             // 全字符集随机串
func CustomizeChars(l, cs int) string // 指定字符集：cs 为 Numbers / LowerCaseChars / UpperCaseChars / SpecialChars 的位标志组合
func LowerChars(l int) string
func UpperChars(l int) string
func LowerCharNumbers(l int) string
func UpperCharNumbers(l int) string
```

## xserializer

```go
// 内嵌 time.Duration 并实现 UnmarshalText 的结构体，
// 用于 JSON 等文本序列化场景，文本格式如 "1h30m"、"300ms"
type Duration struct {
	time.Duration
}
```

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
