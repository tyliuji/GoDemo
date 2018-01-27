package main

import(
	"fmt"
	"net/http"
	"html/template"
	"io"
	"time"
	"strconv"
	"github.com/go-redis/redis"
)

const (
	redisKey_account = "account"
)

var redisClient *redis.Client

func init(){
redisClient = redis.NewClient(&redis.Options{
	Addr:"192.168.137.128:6379",
	Password:"",
})

_,err:= redisClient.Ping().Result()

if err!=nil{
	fmt.Println(err.Error())
}else{
	fmt.Println("connect redis ok.")
}
}

func main(){

	http.HandleFunc("/register",register)
	http.HandleFunc("/login",login)

	http.ListenAndServe("localhost:9000",nil)
	
}

func login(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET"{
		t,err:= template.ParseFiles("login/login.gtpl")
		
		if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		t.Execute(w, nil)
	}else{

		username := r.FormValue("username")
		pwd:= r.FormValue("pwd")

		if len(username) == 0 || len(pwd) == 0{
			io.WriteString(w, "用户名或密码错误")
			return
		}

		value, err := redisClient.HGet(redisKey_account,username).Result()
		if err != nil || pwd != value{
			io.WriteString(w, "用户名或密码错误")
			return
		}

		io.WriteString(w, "登录成功！")
	}
}

func register(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET"{
		t,err:= template.ParseFiles("register/register.gtpl")
		
		if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		t.Execute(w, nil)
	}else{

		userName:= r.FormValue("username")
		pwd1 := r.FormValue("pwd")
		pwd2 := r.FormValue("pwd2")

		if len(userName) == 0{
			io.WriteString(w, "用户名不能为空")
			fmt.Println("Username is empty!")
			return
		}

		if len(pwd1) == 0 || len(pwd2) == 0 || pwd1!= pwd2{
			http.Error(w, "密码有误", http.StatusOK)
			fmt.Println("Password is invalid!")
			return
		}

		value, err := redisClient.HGet(redisKey_account,userName).Result()
	 	if err == nil && len(value) != 0{
			io.WriteString(w, "用户名已存在!")
			fmt.Println("Username is exist!")
			return
		}

		isOk, err := redisClient.HSet(redisKey_account, userName, pwd1).Result()
		if err!= nil || !isOk{
			io.WriteString(w,"注册失败")
			fmt.Println("Register is failed! err:"+ err.Error())
			return
		}

		fmt.Println("Register is ok.")
		io.WriteString(w,"注册成功！")
		ticker := time.NewTicker(time.Second)
		for i := 3; i > 0; i-- {
			<-ticker.C
			io.WriteString(w, strconv.Itoa(i)+" 秒后跳转到登录界面")
		}
	}
}