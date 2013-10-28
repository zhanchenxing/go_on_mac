package main


import (
	"fmt"
	"net/http"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore( []byte("something-very-secret") )

func LoginHandler( w http.ResponseWriter, r *http.Request ){
	
	fmt.Println("\n\nLoginHandler")
	fmt.Println("form:", r.Form )

	fmt.Println(r.URL.Path)

	fmt.Println("r=", r)
	err := r.ParseForm()
	if err != nil {
		fmt.Println("ParseForm failed:", err )
	}

	uid := r.FormValue("uid")
	username := r.FormValue("username")
	fmt.Println("uid=", uid)
	fmt.Println("username=", username)
	if len(uid) == 0 {
		fmt.Println("uid doesn't exist!")
		http.Redirect(w,r, "/html/login.html", 307)
	} else{
		session, _ := store.Get(r, "login")
		session.Values["userid"] = uid
		session.Options = &sessions.Options{
			MaxAge: 30,
		}

		session.Save(r, w)

		http.Redirect(w, r, "/", 307 )
        
	}
    
}


func DataReqHandler( w http.ResponseWriter, r *http.Request ){
	fmt.Println("\n\nDataReqHandler")
	fmt.Println(r.URL.Path)
	session, _ := store.Get(r, "login")
	fmt.Println("session=", session)
	uid := session.Values["userid"]
	fmt.Println("uid=", uid)
	switch uid.(type){
	case string:
		fmt.Fprintf(w, "Hello %v! You have login!", uid)
	default:
		http.Redirect(w,r, "/html/login.html", 307)
	}


	// when redirecting, use 307!!! other code doesn't work!
}


func main(){
	http.Handle( "/html/", http.StripPrefix( "/html/", http.FileServer( http.Dir("html") ) ) )
	// http.Handle( "/html/", http.FileServer( http.Dir("html") ) )

	// 注意：写不写/区别是很大很大的哦！
	http.HandleFunc("/dologin", LoginHandler )
	http.HandleFunc("/", DataReqHandler )
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Print("ListenAndServe failed with err:[%v]\n", err)
		return
	}
	fmt.Println("exiting...")
    
    
}

