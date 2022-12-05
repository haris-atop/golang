package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"reflect"
)
type user struct{
	Name string
	Grade int
}
func add(asd *int){
	*asd+=1;
	fmt.Println("inside func",*asd)
}
func main(){
	var asd int = 1
	add(&asd)
	fmt.Println(asd)
	var object=map[string]int{"asd":1}
	var hasilStruct=user{"bambang",10}
	fmt.Println(reflect.TypeOf(hasilStruct))
	fmt.Println(reflect.TypeOf(object))
	fmt.Println(reflect.TypeOf(1))
	mux:=http.NewServeMux()
	mux.HandleFunc("/",func (res http.ResponseWriter,req *http.Request){
		switch req.Method{
		case "GET":
			inputUser:=[]user{
				user{"bambang",69},
				user{"get",99},
			}
			res.Header().Set("Content-Type","application/json")
			result,_:=json.Marshal(inputUser)
			res.Write(result)
		case "POST":
			fmt.Println("body",req.Body)
			decoder:=json.NewDecoder(req.Body)
			fmt.Println("decoder",decoder,reflect.ValueOf(decoder).Kind())
			inputUser:=[]user{
				user{"bambang",69},
				user{"post",99},
			}
			res.Header().Set("Content-Type","application/json")
			result,_:=json.Marshal(inputUser)
			res.Write(result)
		default:
			http.Error(res,"asdaaa",http.StatusBadRequest)
		}
	})
	mux.HandleFunc("/asd",func (res http.ResponseWriter,req *http.Request){
		inputUser:=[]user{
			user{"udin",99},
		}
		res.Header().Set("Content-Type","application/json")
		result,_:=json.Marshal(inputUser)
		res.Write(result)
	})
	server:=http.Server{
		Addr:"localhost:6969",
		Handler:mux,
	}
	fmt.Println("server starts")
	err:=server.ListenAndServe()
	if err !=nil{
		panic(err)
	}
}