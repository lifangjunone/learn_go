package main

import (
	"fmt"
	"regexp"
	"time"
)

func test() {
	now := time.Now()
	fmt.Println(now)
	// 当前时间戳
	fmt.Println(now.Unix())
	// 纳秒级时间戳
	fmt.Println(now.UnixNano())
	// 时间戳小数部分
	fmt.Println(now.Nanosecond())
	year, month, day := now.Date()
	fmt.Printf("year:%d, month:%d, day:%d", year, month, day)
	fmt.Println()
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	hour, minute, second := now.Clock()
	fmt.Printf("hour: %d, minute:%d, second: %d", hour, minute, second)
	fmt.Println()
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Weekday())
	fmt.Println(now.YearDay())
	fmt.Println(now.Location())

	// 时间格式化
	now2 := time.Now()
	fmt.Println(now2.Format("2006-01-02 15:04:05"))
	// 时间戳和日期字符串相互转化, 需要先将时间戳转成　time.Time　类型　再格式化成日期格式
	now3 := time.Now()
	layout := "2006-01-02 15:04:05"
	t := time.Unix(now3.Unix(), 0)
	fmt.Println(t.Format(layout))

	now4 := time.Now()

	// 根据指定时间返回　time.Time 类型
	t2 := time.Date(2011, time.Month(3), 12, 15, 30, 20, 0, now4.Location())
	fmt.Println(t2.Format(layout))
	// 日期字符串解析成　time.Time 类型

	t3, _ := time.ParseInLocation("2006-01-02 15:04:05", "2023-01-22 22:23:34", time.Local)
	fmt.Println(t3)
	fmt.Println(now.Local())
	// 需要注意时区的问题
	t4, _ := time.Parse("2006-01-02 15:04:05", "2022-01-02 15:33:33")
	fmt.Println(t4)

	t11, _ := time.ParseDuration("1h1m1s")
	fmt.Println(t11)
	t12 := now.Add(t11)
	fmt.Println(t12)
	t13, _ := time.ParseDuration("-2h")
	t14 := now.Add(t13)
	fmt.Println(t14)

	sub := now.Sub(t14)
	fmt.Println(sub)
	fmt.Println(sub.Hours())

	now5 := time.Now()
	t44, _ := time.ParseDuration("4h")
	fmt.Println(t44)
	t33 := now5.Add(t44)
	fmt.Println(t33)

	// type Duration int64

	xx := time.Until(t33)
	fmt.Println(xx)

	t55, _ := time.ParseDuration("-5h")
	t77 := now5.Add(t55)
	t66 := time.Since(t77)
	fmt.Println(t66)

	now6 := time.Now()
	xx2 := now6.AddDate(2, 3, 4)
	fmt.Println(xx2)

	xx3 := now6.AddDate(0, 0, -3)
	fmt.Println(xx3)

	now7 := time.Now()
	m1, _ := time.ParseDuration("3h")
	xxxx := now7.Add(m1)
	fmt.Println(now.Before(xxxx))
	fmt.Println(xxxx.After(now7))
	fmt.Println(now7.Equal(time.Now()))
}

// Str2timestamp 日期格式转时间戳
func Str2timestamp(fmtStr, valueStr, locStr string) int64 {
	loc := time.Local
	if locStr != "" {
		loc, _ = time.LoadLocation(locStr) // 设置时区
	}
	if fmtStr == "" {
		fmtStr = "2006-01-02 15:04:03"
	}
	t, _ := time.ParseInLocation(fmtStr, valueStr, loc)
	return t.Unix()
}

// GetCurrentFormatStr 获取当前时间字符串
func GetCurrentFormatStr(fmtStr string) string {
	if fmtStr == "" {
		fmtStr = "2006-01-02 15:04:05"
	}
	return time.Now().Format(fmtStr)
}

func TimeStampToDate(timeStamp int64, fmtStr string) string {
	if fmtStr == "" {
		fmtStr = "2006-01-02 15:04:05"
	}
	return time.Unix(timeStamp, 0).Format(fmtStr)
}

func ReTest() string {
	const text = "My emailisccmouse@gmail.com"
	compile := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z0-9]+`)
	match := compile.FindString(text)
	return match
}

func main() {
	fmt.Println("hello")
	res := Str2timestamp("", "2023-01-33 23:34:23", "")
	fmt.Println(res)
	res3 := GetCurrentFormatStr("")
	fmt.Println(res3)
	res4 := TimeStampToDate(time.Now().Unix(), "")
	fmt.Println(res4)
	res5 := ReTest()
	fmt.Println(res5)
}
