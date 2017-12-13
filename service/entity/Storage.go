package entity

import (
	"fmt"  
    "regexp"
    "strings"
) 







var mData []Meeting
var uData []User


func ListAllUser() []User{
	return uData
}

func ListAllMeeting() []Meeting{
	return mData
}
//read data from data.txt to mData and uData
func Init() {
	tmp_m := Meeting_ReadFromFile()
	tmp_u := User_ReadFromFile()
	for i := 0; i < len(tmp_u); i++ {
		uData = append(uData, tmp_u[i])
	}
	for i := 0; i < len(tmp_m); i++ {
		mData = append(mData, tmp_m[i])
	}
}


func IsEmail(str string) bool {  
    var b bool  
    b, _ = regexp.MatchString("^([a-z0-9_\\.-]+)@([\\da-z\\.-]+)\\.([a-z\\.]{2,6})$", str)  
    if false == b {  
        return b  
    }
    return b  
}

func IsCellphone(str string) bool {  
    var b bool  
    b, _ = regexp.MatchString("^1[0-9]{10}$", str)  
    if false == b {  
        return b  
    }  
    return b  
}  


func RegisterUser(name string, password string, email string, phone string) bool {
	var user User
	err := false
	if (IsEmail(email) == false) {
		fmt.Println("Email is error!")
		err = true
	}
	if (IsCellphone(phone) == false) {
		fmt.Println("Phone is error!")
		err = true
	}
	if (len(password) < 6) {
		fmt.Println("The length of password can't be less than 6!")
		err = true
	}
	_, isExit, _:= Query_user(name)
	if (isExit) {
		fmt.Println("This username exits, please use another username!")
		err = true
	}
	if (err) {
		return false
	}
	user.Name = name
	user.Password = password
	user.Email = email
	user.Phone = phone
	uData = append(uData,user)
	User_WriteToFile(uData)
	fmt.Println("Register successfully!")
	return true
}


func Query_meeting_by_title(title string) (Meeting, bool, int){
	for i := 0; i < len(mData); i++ {
		if mData[i].Title == title {
			return mData[i],true, i
		}
	}
	return mData[0], false, 0
}

func Query_meeting_by_participator(name string) []Meeting{
	var tmp_m []Meeting
	for i := 0; i < len(mData); i++ {
		for j := 0;j < len(mData[i].Participator); j++ {
			if mData[i].Participator[j] == name {
				tmp_m = append(tmp_m, mData[i])
			}
		}
	}
	return tmp_m
}


func Query_meeting_by_sponsor(sponsor string) []Meeting {
	var tmp_m []Meeting
	for i := 0; i < len(mData); i++ {
		if mData[i].Sponsor == sponsor {
			tmp_m = append(tmp_m, mData[i])
		}
	}
	return tmp_m
}


func Query_meeting_by_name(name string) []Meeting {
	tmp_m := Query_meeting_by_sponsor(name)
	tmp_tt := Query_meeting_by_participator(name)
	for i := 0; i < len(tmp_tt); i++ {
		tmp_m = append(tmp_m, tmp_tt[i])
	}
	return tmp_m
}
//time is overlap


func Cancell_meeting(title string, sponsor string) bool {
	_,_, i := Query_meeting_by_title(title)
	if (mData[i].Sponsor == sponsor) {
		index := i;
		mData = append(mData[:index],mData[index+1:]...)
		Meeting_WriteToFile(mData)
		return true
	} else {
		return false
	}
	
}


func Add_participator(name string, title string) bool {
	_,_, index := Query_meeting_by_title(title)
	tmp_mm := Query_meeting_by_name(name)
	for i := 0; i < len(tmp_mm); i++ {
		if IsOverlapping(tmp_mm[i].Start, tmp_mm[i].End, mData[index].Start, mData[index].End) {
			return false
		}
	}
	mData[index].Participator = append(mData[index].Participator,name)
	Meeting_WriteToFile(mData)
	return true
}

func Create_meeting(title string, start Date, end Date,sponsor string,participator string) bool {
	is_flag := false
	//meeting title is exist
	var m_meet Meeting
	for i := 0; i < len(mData); i++ {
		if mData[i].Title == title {
			fmt.Println("Meeting is exists!")  //error
			is_flag = true
		}
	}
	tmp_m := Query_meeting_by_name(sponsor)
	if (len(tmp_m) !=0) {
		for j := 0; j < len(tmp_m); j++ {
			if IsOverlapping(tmp_m[j].Start, tmp_m[j].End, start, end) {
				fmt.Println("Your meeting named ",tmp_m[j].Title, " is overlap with the meeting!")
				is_flag = true
			}
		}
	}
	m_meet.Sponsor = sponsor
	m_meet.Start = start
	m_meet.End = end
	m_meet.Title = title
	var m_participator []string

	tt := strings.Split(participator,",")
    for j:=0; j < len(tt); j++ {
        m_participator = append(m_participator, tt[j])
    }

	count := 0
	for k := 0; k < len(m_participator); k++ {
		flag1 := true
		_,flag, _:= Query_user(m_participator[k])
		if (!flag)  {
			fmt.Println(m_participator[k], " is not registered!")
			is_flag = true
		} else {
			tmp_mm := Query_meeting_by_name(m_participator[k])
			for i := 0; i < len(tmp_mm); i++ {
				if IsOverlapping(tmp_mm[i].Start, tmp_mm[i].End, start, end) {
					flag1 = false
				}
			}
			if !flag1 {
				//time overlap
				fmt.Println(m_participator[k], " can't attend the meeting because of timeoverlap")
				is_flag = true
			} else {
				m_meet.Participator = append(m_meet.Participator, m_participator[k])
				count++;
			}
		}
	}
	if is_flag {
		fmt.Println("Create meeting failed!")
		return false
	}
	if count != 0 {
		mData = append(mData, m_meet);
		Meeting_WriteToFile(mData)
		fmt.Println("Create meeting successfully!")
		return true
	} else {
		fmt.Println("No participator! you can't create it")
		return false
	}

}

func Delete_user(name string) {
	_,_, i := Query_user(name)
	index := i;
	uData = append(uData[:index],uData[index+1:]...)
	User_WriteToFile(uData)
	//empty meeting sponsor
	Empty_meeting(name)

	tmp_m := Query_meeting_by_name(name)
	for i := 0; i< len(tmp_m); i++ {
		Rm_participator(name, tmp_m[i].Title)
	}

}



func Empty_meeting(sponsor string) int {
	count := 0
	for i := 0; i < len(mData); i++ {
        if mData[i].Sponsor == sponsor {
        	count++
            mData = append(mData[:i], mData[i+1:]...)
        }
    }
    Meeting_WriteToFile(mData)
    if count == 0 {
    	return 0
    } else {
    	return 1
    }  
}

func Exit_meeting(name string,title string) int {
	return Rm_participator(name, title)
}



func Query_meeting(start Date, end Date, name string) []Meeting{
	var tmp []Meeting
	tmp_m := Query_meeting_by_participator(name)
	tmp_mm := Query_meeting_by_sponsor(name)
	for i := 0; i < len(tmp_m); i++ {
		if IsOverlapping(start, end, tmp_m[i].Start, tmp_m[i].End) {
			tmp = append(tmp, tmp_m[i])
		}
	}
	for i := 0; i < len(tmp_mm); i++ {
		if IsOverlapping(start, end, tmp_mm[i].Start, tmp_mm[i].End) {
			tmp = append(tmp, tmp_mm[i])
		}
	}
	return tmp
}


func Query_user(name string) (User,bool, int){
	for i := 0; i< len(uData); i++ {
		if uData[i].Name == name {
			return uData[i], true, i
		}
	}
	return User{"1","2","3","4"}, false, 0
}

func Rm_participator(name string, title string) int {
	_,flag,index := Query_meeting_by_title(title)
	if  !flag {
		//fmt.Println("It doesn't exists this meeting!")
		return 1
	} else {
		for i := 0; i < len(mData[index].Participator); i++ {
			if mData[index].Participator[i] == name {
				mData[index].Participator = append(mData[index].Participator[:i], mData[index].Participator[i+1:]...)
				Meeting_WriteToFile(mData)
				//fmt.Println("Remove participator successfully!")
				if len(mData[index].Participator) == 0 {
					Cancell_meeting(mData[index].Title, mData[index].Sponsor)    //0 cancell meeting
				}
				return 2
			}
		}
		if (mData[index].Sponsor == name) {
			return 3
		}
		return 4
		//fmt.Println("It doesn't exists the participator!")
	}
}