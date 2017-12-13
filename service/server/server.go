package server

import (
	"net/http"
	"agenda-cli-service/service/service"
	"agenda-cli-service/service/entity"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {

	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx, formatter)

	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/v1/users/login", loginHandler(formatter)).Methods("POST")
	mx.HandleFunc("/v1/users/logout", logoutHandler(formatter)).Methods("GET")
	mx.HandleFunc("/v1/users/{name}", getUserByNameHandler(formatter)).Methods("GET")
	mx.HandleFunc("/v1/users", UserRegisterHandler(formatter)).Methods("POST")
	mx.HandleFunc("/v1/users", ListAllUserHandler(formatter)).Methods("GET")
	mx.HandleFunc("/v1/meetings", CreateMeetingHandler(formatter)).Methods("POST")
	mx.HandleFunc("/v1/meetings/{title}", getMeetingByTitleHandler(formatter)).Methods("GET")
	mx.HandleFunc("/v1/meetings", ListAllMeetingHandler(formatter)).Methods("GET")
	/*mx.HandleFunc("/v1/users", userRegisterHandler(formatter)).Methods("POST")
	mx.HandleFunc("/v1/users/{name}", deleteUserByNameHandler(formatter)).Methods("DELETE")
	mx.HandleFunc("/v1/users/{name}", updateUserByNameHandler(formatter)).Methods("PATCH")

	mx.HandleFunc("/v1/meetings/{title}", getMeetingByTitleHandler(formatter)).Methods("GET")
	mx.HandleFunc("/v1/meetings", createMeetingHandler(formatter)).Methods("POST")
	mx.HandleFunc("/v1/meetings/{title}", deleteMeetingByTitleHandler(formatter)).Methods("DELETE")*/
}

/*func getAllUsersHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		ulist := entities.AgendaService.ListAllUsers()
		formatter.JSON(w, http.StatusOK, ulist)
	}
}*/

func loginHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		flag := true
		info := req.PostForm
		if info[`name`] == nil || info[`password`] == nil {
			flag = false
	    }
		if flag == true {
			tt := service.Log_in(info[`name`][0], info[`password`][0])
			if (!tt) {
				formatter.JSON(w, http.StatusBadRequest, struct{ ErrorIndo string }{"login failed!"})
            	return
			}
		} else {
			formatter.JSON(w, http.StatusBadRequest, struct{ ErrorIndo string }{"login failed!"})
            return
		}
	}
}


func logoutHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		service.Log_out()
		formatter.JSON(w, http.StatusOK, struct{ successIndo string }{"logout successful!"})
	}
}

func ListAllUserHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		u := service.ListAllUser()
		formatter.JSON(w, http.StatusOK, u)
	}
}

func ListAllMeetingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		u := service.ListAllMeeting()
		formatter.JSON(w, http.StatusOK, u)
	}
}


func CreateMeetingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		flag := true
		info := req.PostForm
		if info[`title`] == nil || info[`start`] == nil || info[`end`] == nil||info["participators"] == nil {
			flag = false
	    }
	    var u entity.Met
	    u.Title = info[`title`][0]
	    u.Start = info[`start`][0]
	    u.End = info[`end`][0]
	    u.Participators = info[`participators`][0]
		if flag == true {
			u.Sponsor = service.GetmyName()
			service.Create_meeting(info[`title`][0], info[`start`][0], info[`end`][0],info[`participators`][0])
			formatter.JSON(w,201,u) // expected a user id
		} else {
			formatter.JSON(w,404,nil)
		}
	}
}



func UserRegisterHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		flag := true
		info := req.PostForm
		if info[`name`] == nil || info[`password`] == nil || info[`email`] == nil || info[`phone`] == nil {
			flag = false
	    }
	    var u entity.User
	    u.Name = info[`name`][0]
	    u.Password = info[`password`][0]
	    u.Email = info[`email`][0]
	    u.Phone = info[`phone`][0]
		if flag == true {
			service.RegisterUser(info[`name`][0], info[`password`][0], info[`email`][0], info[`phone`][0])
			formatter.JSON(w,201,u) // expected a user id
		} else {
			formatter.JSON(w,404,nil)
		}
	}
}

func getUserByNameHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		name := vars["name"]

		u,flag := service.Query_user(name)
		u.Password = "******"
		if flag {
			formatter.JSON(w, http.StatusOK, u)
		} else {
			formatter.JSON(w, http.StatusBadRequest, struct{ ErrorIndo string }{"No Exist!"})
            return
		}
	}
}

func getMeetingByTitleHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		title := vars["title"]

		u, flag:= service.QueryMeetingByTitle(title)
		if flag {
			formatter.JSON(w, http.StatusOK, u)
		} else {
			formatter.JSON(w, http.StatusBadRequest, struct{ ErrorIndo string }{"No Exist!"})
            return
		}
 		
	}
}