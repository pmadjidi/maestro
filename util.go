package main

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	"golang.org/x/crypto/bcrypt"

	"crypto/sha1"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"reflect"
	"strconv"
	. "maestro/api"
)


func Info(format string, a ...interface{}) {
	fmt.Printf("Info:\t"+format+"\n", a...)
}



func hashAndSalt(pwd []byte) []byte {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return hash
}

func RandomString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}


func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}

func CalcHash(name string) string {
	hasher := sha1.New()
	hasher.Write([]byte(name))
	return  fmt.Sprintf("%x", hasher.Sum(nil))
}

func Hextoint(hexnumber string) int {
	d, _ := strconv.ParseInt("0x" + hexnumber , 0, 64)
	return int(d)
}

func Upper_power_of_two(v int) int {
	v--;
	v |= v >> 1;
	v |= v >> 2;
	v |= v >> 4;
	v |= v >> 8;
	v |= v >> 16;
	v++;
	return v;
}

func CallMethod(i interface{}, methodName string) interface{} {
	var ptr reflect.Value
	var value reflect.Value
	var finalMethod reflect.Value

	value = reflect.ValueOf(i)

	// if we start with a pointer, we need to get value pointed to
	// if we start with a value, we need to get a pointer to that value
	if value.Type().Kind() == reflect.Ptr {
		ptr = value
		value = ptr.Elem()
	} else {
		ptr = reflect.New(reflect.TypeOf(i))
		temp := ptr.Elem()
		temp.Set(value)
	}

	// check for method on value
	method := value.MethodByName(methodName)
	if method.IsValid() {
		finalMethod = method
	}
	// check for method on pointer
	method = ptr.MethodByName(methodName)
	if method.IsValid() {
		finalMethod = method
	}

	if (finalMethod.IsValid()) {
		return finalMethod.Call([]reflect.Value{})[0].Interface()
	}

	// return or panic, method not found of either type
	return ""
}

func ExitWithErrors(err error) {
	msg := err.Error()
	log.Fatal(msg)
}



func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func randomUserForTest(size int) *RegisterReq{
		if size <= 12 {
			size = 12
		}
	    return &RegisterReq{UserName: RandomString(size),
		PassWord: []byte(RandomString(size)),
		FirstName: RandomString(size),
		LastName: RandomString(size),
		Email: RandomString(size - 10 ) +  "@gmail.com",
		Phone: RandomString(size),
		Address: &RegisterReq_Address{Street:RandomString(size),City: RandomString(size),State: RandomString(size),Zip: RandomString(size)},
		Device: RandomString(size)}
}

func randomMessageForTest(size int) *MsgReq{
	if size <= 12 {
		size = 12
	}

	return &MsgReq{
		Id: RandomString(size),
		Text: RandomString(size),
		Pic: []byte(RandomString(size)),
		ParentId: RandomString(size),
		Topic: RandomString(size),
		TimeName: &timestamp.Timestamp{},
		}
}
