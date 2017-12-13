package entity

import (
	"os"
	"fmt"
	"bufio"
	"io"
    "strings"
)




func User_ReadFromFile() []User{
	var tmp []User
    var user User
    sql := "select * from User"
    results, _ := engine.Query(sql)
    if (len(results) > 0) {
        if results[0]["Name"] != nil {
            for i := 0; i < len(results); i++ {
            user.Name = string(results[i]["Name"][:])
            user.Password = string(results[i]["Password"][:])
            user.Email = string(results[i]["Email"][:])
            user.Phone = string(results[i]["Phone"][:])
            tmp = append(tmp,user)
            }
        }
    }
    return tmp
}

func User_WriteToFile(My_User []User) {
    for i := 0; i < len(My_User); i++ {
        user := My_User[i]
        engine.Insert(&user)
	}
}


func Meeting_ReadFromFile() []Meeting{
	var tmp []Meeting
    var met Meeting
    sql := "select * from Met"
    results, _ := engine.Query(sql)
    if (len(results) > 0) {
        if results[0]["Title"] != nil {
            for i := 0; i < len(results); i++ {
                met.Sponsor = string(results[i]["Sponsor"][:])
                met.Title = string(results[i]["Title"][:])
                kk := StringToDate(string(results[i]["Start"][:]))
                met.Start = kk
                gg := StringToDate(string(results[i]["End"][:]))
                met.End = gg
                participator := string(results[i]["Participators"][:])
                tt := strings.Split(participator,",")
                for j:=0; j < len(tt); j++ {
                    met.Participator = append(met.Participator, tt[j])
                }
                tmp = append(tmp,met)
            }
        }
    }
    
    
    return tmp
}

func Meeting_WriteToFile(My_Meeting []Meeting) {
    var tt Met
    var ll string
	for i := 0; i < len(My_Meeting); i++ {
        ll = ""
        tt.Sponsor = My_Meeting[i].Sponsor
        tt.Title = My_Meeting[i].Title
        ee := DateToString(My_Meeting[i].Start)
        tt.Start = ee
        ii := DateToString(My_Meeting[i].End)
        tt.End = ii
        ll += My_Meeting[i].Participator[0]
        for j := 1; j < len(My_Meeting[i].Participator); j++ {
            ll += ","
            ll += My_Meeting[i].Participator[j]
        }
        tt.Participators = ll
        engine.Insert(&tt)
    }
}

func LN_ReadFromFile() []string{
	var tmp []string
	f, err := os.Open("entity/data/Host.txt")
    if err != nil {
        panic(err)
    }
    defer f.Close()
    rd := bufio.NewReader(f)
    for {
        line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
        if err != nil || io.EOF == err {
            break
        }
        tmp = append(tmp, line)
    }
    return tmp
}

func LN_WriteToFile(name string) {
	file, err := os.OpenFile("entity/data/Host.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	os.Truncate("entity/data/Host.txt", 0)
    if err != nil {
        fmt.Println("open file failed.", err.Error())
        os.Exit(1)
    }
    defer file.Close()
        file.WriteString(name)
        file.WriteString("\n")
}

func Empty_login() {
    os.Truncate("entity/data/Host.txt", 0)
}
