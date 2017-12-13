package entity
//package main

import (
//"fmt"
	"strconv"
)
type Date struct {
    Year, Month, Day, Hour, Minute int
}

func IsValid(d Date) bool {
    if d.Year < 1000 || d.Year > 9999 || d.Month < 1 ||
        d.Month > 12 || d.Day < 1 || d.Hour < 0 ||
        d.Hour >= 24 || d.Minute < 0 || d.Minute >= 60 {
        return false
    }
    if d.Month == 1 || d.Month == 3 || d.Month == 5 || d.Month == 7 ||
        d.Month == 8 || d.Month == 10 || d.Month == 12 {
        if d.Day > 31 {
            return false
        }
    } else if d.Month == 4 || d.Month == 6 ||
        d.Month == 9 || d.Month == 11 {
        if d.Day > 30 {
            return false
        }
    } else {
        if (d.Year%4 == 0 && d.Year%100 != 0) || (d.Year%400 == 0) {
            if d.Day > 29 {
                return false
            }
        } else {
            if d.Day > 28 {
                return false
            }
        }
    }
    return true
}

func StringToDate(s string) Date {
    var i_Date = Date{
        Year:   0,
        Month:  0,
        Day:    0,
        Hour:   0,
        Minute: 0,
    }
    if len(s) != 16 || s[4] != '-' || s[7] != '-' ||
        s[10] != '/' || s[13] != ':' {
        return i_Date
    }
    for i := 0; i < 16; i++ {
        if i == 4 || i == 7 || i == 10 || i == 13 {
            continue
        }
        if s[i] < '0' || s[i] > '9' {
            return i_Date
        }
    }
    i_Date.Year, _ = strconv.Atoi(s[0:4])
    i_Date.Month, _ = strconv.Atoi(s[5:7])
    i_Date.Day, _ = strconv.Atoi(s[8:10])
    i_Date.Hour, _ = strconv.Atoi(s[11:13])
    i_Date.Minute, _ = strconv.Atoi(s[14:16])
    return i_Date
}

func DateToString(d Date) string {
    var d_String string = ""
    var i_Time string = "0000-00-00/00:00"
    if !IsValid(d) {
        return i_Time
    }
    var s_Year string = strconv.Itoa(d.Year)
    var s_Month string = strconv.Itoa(d.Month)
    var s_Day string = strconv.Itoa(d.Day)
    var s_Hour string = strconv.Itoa(d.Hour)
    var s_Minute string = strconv.Itoa(d.Minute)
    if d.Month < 10 {
        s_Month = "0" + s_Month
    }
    if d.Day < 10 {
        s_Day = "0" + s_Day
    }
    if d.Hour < 10 {
        s_Hour = "0" + s_Hour
    }
    if d.Minute < 10 {
        s_Minute = "0" + s_Minute
    }
    d_String = s_Year + "-" + s_Month + "-" + s_Day + "/" +
        s_Hour + ":" + s_Minute
    return d_String
}

func IsOverlapping(s_d1 Date, e_d1 Date, s_d2 Date, e_d2 Date) bool {
    if (Date_MoreThan(s_d1, s_d2) && Date_LessThan(s_d1, e_d2)) ||
        (Date_MoreThan(e_d1, s_d2) && Date_LessThan(e_d1, e_d2)) ||
        (Date_MoreThan(s_d2, s_d1) && Date_LessThan(s_d2, e_d1)) ||
        (Date_MoreThan(e_d2, s_d1) && Date_LessThan(e_d2, e_d1)) {
        return true
    }
    if (Date_Equal(s_d1, s_d2)&&Date_Equal(e_d1, e_d2)) {
        return true
    }
    return false
}

func Date_LessThan(d1 Date, d2 Date) bool {
    if Date_MoreThan(d1, d2) == false && Date_Equal(d1, d2) == false {
        return true
    }
    return false
}

func Date_Equal(d1 Date, d2 Date) bool {
    if d1.Year == d2.Year && d1.Month == d2.Month &&
        d1.Day == d2.Day && d1.Hour == d2.Hour && d1.Minute == d2.Minute {
        return true
    }
    return false
}

func Date_MoreThan(d1 Date, d2 Date) bool {
    if d1.Year > d2.Year {
        return true
    }
    if d1.Year < d2.Year {
        return false
    }
    if d1.Month > d2.Month {
        return true
    }
    if d1.Month < d2.Month {
        return false
    }
    if d1.Day > d2.Day {
        return true
    }
    if d1.Day < d2.Day {
        return false
    }
    if d1.Hour > d2.Hour {
        return true
    }
    if d1.Hour < d2.Hour {
        return false
    }
    if d1.Minute > d2.Minute {
        return true
    }
    if d1.Minute < d2.Minute {
        return false
    }
    return false
}




func GetYear(a Date) int {
    return a.Year
}


func GetMonth(a Date) int {
    return a.Month
}

func GetDay(a Date) int {
    return a.Day
}
func GetHour(a Date) int {
    return a.Hour
}
func GetMinute(a Date) int {
    return a.Minute
}


/*
func main() {
    //s := "2017-13-12/12:00"
    s1 := "2017-12-212:00"
    var a = Date {
        Year : 2017,
        Month : 12,
        Day : 12,
        Hour : 11,
        Minute : 3,
    }
    var b = Date {
        Year : 2017,
        Month : 12,
        Day : 12,
        Hour : 11,
        Minute : 4,
    }
    fmt.Println(Date_LessThan(a, b))
    fmt.Println(IsValid(b))
    fmt.Println(DateToString(a))
    fmt.Println(StringToDate(s1))
}*/