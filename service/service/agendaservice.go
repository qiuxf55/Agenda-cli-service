package service

import (
	"fmt"  
    "os"
    "strings"
    "log"
    "agenda-cli-service/service/entity"
)

var my_name, my_password string
var Login_flag bool 
var All_name []string

var log_file *os.File

func GetmyName() string{
	return my_name
}

func GetFlag() bool {
	return Login_flag
}

func Init() {
	entity.Init()


    logFile,err := os.OpenFile("service/agenda.log",os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
    log_file = logFile
    if err != nil {
        log.Fatalln("open file error !")
    }
    //debugLog := log.New(log_file,"[Operation]",log.LstdFlags)
	tmp := entity.LN_ReadFromFile()
	if (len(tmp)==0) {
		Login_flag = false
	} else {
		Login_flag = true
		my_name = strings.Replace(tmp[0],"\n","",-1)
	}
	
}

func RegisterUser(name string, password string, email string, phone string) {
	debugLog := log.New(log_file,"[Operation]",log.LstdFlags)
	var err = false
	
	if (len(password) < 6) {
		debugLog.Println("The length of password can't be less than 6!")
		err = true
	}
	if (err) {
		os.Exit(0)
	}
	i := entity.RegisterUser(name, password, email, phone)
	if (i) {
		debugLog.Println(name, " register successfully!")
	} else {
		debugLog.Println(name, " register failed!")
	}
	defer log_file.Close()
}

func ListAllUser() []entity.User{
	return entity.ListAllUser()
}

func ListAllMeeting() []entity.Meeting{
	return entity.ListAllMeeting()
}
func Cancell_meeting(title string) {
	debugLog := log.New(log_file,"[Operation]",log.LstdFlags)
	_,flag, _ := entity.Query_meeting_by_title(title)
    if flag == false {
    	fmt.Println(title, "doesn't exists!")
    	debugLog.Println(my_name, " cancell meeting ", title, " failed!")
    	os.Exit(1)
    } else {
    	is_flag := entity.Cancell_meeting(title, my_name)
    	if (is_flag == false) {
    		fmt.Println(title, " is not sponsor by you, so you can't cancell the meeting!")
    	debugLog.Println(my_name, " cancell meeting ", title, " failed!")
    		os.Exit(1)
    	}
    	debugLog.Println(my_name, " cancell meeting ", title, " successfully!")
    	fmt.Println("Cancell meeting successfully!")
    }
    defer log_file.Close()
}

func Add_participator(name string, title string) {
	debugLog := log.New(log_file,"[Operation]",log.LstdFlags)
	tmp_m := entity.Query_meeting_by_sponsor(my_name)
	if (len(tmp_m) != 0) {
		for i := 0; i < len(tmp_m); i++ {
			if entity.GetTitle(tmp_m[i]) == title {
				if entity.Add_participator(name, title) {
					debugLog.Println(my_name, " add participator ", name, " to ", title, " successfully!")
					fmt.Println("Add participator to ",title, " successfully!")
				} else {
					debugLog.Println(my_name, " add participator ", name, " to ", title, " failed!")
					fmt.Println("Time is overlap!")  //time overlap
				}
			}
		}
	} else {
		debugLog.Println(my_name, " add participator ", name, " to ", title, " failed!")
		fmt.Println("You didn't initiate the meeting!")
	}
	defer log_file.Close()
}

func Create_meeting(title string, start string, end string, participator string) {
	debugLog := log.New(log_file,"[Operation]",log.LstdFlags)
	if !entity.IsValid(entity.StringToDate(start))|| !entity.IsValid(entity.StringToDate(end))||
	entity.Date_LessThan(entity.StringToDate(end), entity.StringToDate(start)) {
		debugLog.Println(my_name, " create meeting ", title, " failed!")
		fmt.Println("Time is error!")
	} else {
		if entity.Create_meeting(title, entity.StringToDate(start), entity.StringToDate(end), my_name, participator) {
			debugLog.Println(my_name, " create meeting ", title, " successfully!")
		} else {
			debugLog.Println(my_name, " create meeting ", title, " failed!")
		}
	}
	defer log_file.Close()
}

func QueryMeetingByTitle(title string) (entity.Meeting, bool){
	tt,flag,_ := entity.Query_meeting_by_title(title)
	if flag {
		return tt, true
	}
	return tt,false
}


func Delete_user() {
	debugLog := log.New(log_file,"[Operation]",log.LstdFlags)
	debugLog.Println(my_name, " log off account successfully!")
	entity.Empty_login()
	entity.Delete_user(my_name)
	fmt.Println("log off successfully!")
	defer log_file.Close()
}

func Empty_meeting() {
	debugLog := log.New(log_file,"[Operation]",log.LstdFlags)
	i := entity.Empty_meeting(my_name)
	if i == 0 {
		debugLog.Println(my_name, " empty meeting failed!")
		fmt.Println("You didn't sponsor any meeting!")
	} else {
		debugLog.Println(my_name, " empty meeting successfully!")
		fmt.Println("Empty meeting successfully!")
	}
	defer log_file.Close()
}

func Exit_meeting(title string) {
	debugLog := log.New(log_file,"[Operation]",log.LstdFlags)
	a := entity.Exit_meeting(my_name, title)
	if a == 1 {
		debugLog.Println(my_name, " exit ", title, " meeting failed!")
		fmt.Println("It doesn't exists this meeting!")
	} else if a == 2 {
		debugLog.Println(my_name, " exit ", title, " meeting successfully!")
		fmt.Println("Exit meeting successfully!")
	} else if a == 3 { 
		debugLog.Println(my_name, " exit ", title, " meeting failed!")
		fmt.Println("You are the sponsor, you can't exit meeting!")
    }else {
    	debugLog.Println(my_name, " exit ", title, " meeting failed!")
		fmt.Println("You don't attend the meeting!")
	}
	defer log_file.Close()
}

func Log_in(name string, password string) bool {
	tmp_u, flag, _:= entity.Query_user(name)
	if flag == true {
		my_name = name
		my_password = password
		if (entity.GetPassword(tmp_u) != password) {
			fmt.Println("Password is wrong!")
			return false
		} else {
			entity.LN_WriteToFile(name)
			fmt.Println("Log in successfully!\nWelcome to Agenda!")
			return true
		}
	} else {
		fmt.Println("You don't register!")
		return false
	}
	
}

func Log_out() {
	debugLog := log.New(log_file,"[Operation]",log.LstdFlags)
	debugLog.Println(my_name, " log out successfully!")
	fmt.Println("Log out successfully!")
	entity.Empty_login()
	defer log_file.Close()
}

func Query_meeting(start string, end string) {
	debugLog := log.New(log_file,"[Operation]",log.LstdFlags)
	if !entity.IsValid(entity.StringToDate(start))|| !entity.IsValid(entity.StringToDate(end))||
	entity.Date_LessThan(entity.StringToDate(end), entity.StringToDate(start)) {
		debugLog.Println(my_name, " query meeting failed!")
		fmt.Println("Time is error!")
		os.Exit(1)
	} else {
		tmp_m := entity.Query_meeting(entity.StringToDate(start), entity.StringToDate(end), my_name)
		if len(tmp_m) == 0 {
			debugLog.Println(my_name, " query meeting failed!")
			fmt.Println("You don't have any meeting during this period!")
		} else {
			debugLog.Println(my_name, " query meeting successfully!")
			for i := 0; i< len(tmp_m); i++ {
				fmt.Println("Start time : ",entity.DateToString(entity.GetStart(tmp_m[i])), "\nEnd time : ", entity.DateToString(entity.GetEnd(tmp_m[i])))
				fmt.Println("Title : ", entity.GetTitle(tmp_m[i]))
				fmt.Println("Sponsor : ", entity.GetSponsor(tmp_m[i]), "\nParticipator : ", entity.GetParticipator(tmp_m[i]))
			}
		}
	}
	defer log_file.Close()
}

func Query_user(name string) (entity.User, bool){
	debugLog := log.New(log_file,"[Operation]",log.LstdFlags)
	tmp_u, flag, _ := entity.Query_user(name)
	if !flag {
		debugLog.Println(my_name, " query user ", name, " failed!")
		fmt.Println(name," doesn't exists!")
	} else {
		debugLog.Println(my_name, " query user ", name, " successfully!")
		fmt.Println("Name : ", entity.GetName(tmp_u))
		fmt.Println("Email : ", entity.GetEmail(tmp_u))
		fmt.Println("Phone : ", entity.GetPhone(tmp_u))
	}
	defer log_file.Close()
	return tmp_u, flag
}

func Rm_participator(name string, title string) {
	debugLog := log.New(log_file,"[Operation]",log.LstdFlags)
	a := entity.Rm_participator(name, title)
	if a == 1 {
		debugLog.Println(my_name, " remove participator ", name, " from meeting ", title, " failed!")
		fmt.Println("It doesn't exists this meeting!")
	} else if a == 2 {
		debugLog.Println(my_name, " remove participator ", name, " from meeting ", title, " successfully!")
		fmt.Println("Remove participator successfully!")
	} else {
		debugLog.Println(my_name, " remove participator ", name, " from meeting ", title, " failed!")
		fmt.Println("It doesn't exists the participator!")
	}
	defer log_file.Close()
}
/*
func main() {
	Log_in("qq", "123456")
	RegisterUser("ee","123456","1097810144@qq.com", "13719330671")
	entity.Show()
}*/
/*
func main() {
	RegisterUser("xx","1245678","10254@qq.com", "13719330671")
	Log_in("xx", "1245678")
	Create_meeting("Golang", "2017-12-12/12:00","2017-12-13/12:00")
	Create_meeting("Golang", "2017-12-12/12:00","2017-12-11/12:00")
	Query_user("tt")
	Query_user("re")
	Query_meeting("2017-12-12/11:00", "2017-12-12/23:00")
	Rm_participator("tt", "Golang")
	Add_participator("aa", "Golang")
	Empty_meeting()
	Delete_user()
	entity.Show()
}*/