package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"io"
	"log"
	. "maestro/api"
	mathRand "math/rand"
	"os"
	"reflect"
	"strconv"
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
		b[i] = letter[mathRand.Intn(len(letter))]
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

func randomUserForTest(size int) []*RegisterReq{
		if size == 0 {
			size = 1
		}

		reqArray := make([]*RegisterReq,0)

		for i := 0;i< size ; i++ {
		ps := strconv.Itoa(i)
		reqArray =append(reqArray,&RegisterReq{UserName: "kalle" + ps,
				PassWord:  []byte(RandomString(13)),
				FirstName: "Kalle" + ps,
				LastName:  "Svensson" + ps,
				Email:    "kalle" + ps  + "@gmail.com",
				Phone:     "08-12 18 " + ps,
				Address:   &RegisterReq_Address{Street: "Tomtebogatan " + ps, City: "Stockholm", State: "Sweden", Zip: "113 " + ps},
				Device:    "device-" + ps,
				AppName:   "Test" + ps,})
		}
		return reqArray
}


func randomUsersForTests(size,systems int) [][]*RegisterReq{
	if size == 0 {
		size = 1
	}

	if systems == 0 {
		systems = 1
	}

	sys := make([][]*RegisterReq,0)


	for i := 0;i< systems ; i++ {
		as := strconv.Itoa(i)
		reqArray := make([]*RegisterReq,0)
		for j := 0; j < size; j++ {
			ps := strconv.Itoa(j)
			pass := []byte(RandomString(13))
			if i < 5 && j < 5 {
				ps = "testpass" + ps
			}
			reqArray = append(reqArray, &RegisterReq{UserName: "kalle" + ps,
				PassWord:  pass,
				FirstName: "Kalle" + ps,
				LastName:  "Svensson" + ps,
				Email:     "kalle" + ps + "@gmail.com",
				Phone:     "08-12 18 " + ps,
				Address:   &RegisterReq_Address{Street: "Tomtebogatan " + ps, City: "Stockholm", State: "Sweden", Zip: "113 " + ps},
				Device:    "device-" + ps,
				AppName:   "Test" + as,})
		}
		sys = append(sys,reqArray)
	}
	return sys
}



func randomMessageForTest(size int,topic int) *MsgReq{
	if size <= 12 {
		size = 12
	}

	return &MsgReq{
		Uuid:uuid.New().String(),
		Text: RandomString(size),
		Pic: []byte(RandomString(size)),
		ParentId: RandomString(size),
		Topic: strconv.Itoa(topic),
		TimeName: &timestamp.Timestamp{},
		}
}

func decodeToken(token,appSecret string) (map[string]interface{},error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(appSecret), nil
	})
	if err != nil {
		return nil,err
	}
	// ... error handling

	// do something with decoded claims
	for key, val := range claims {
		fmt.Printf("Key: %v, value: %v\n", key, val)
	}
	return claims,nil
}


func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

func decrypt(data []byte, passphrase string) []byte {
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}



func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	if err == nil {
		return !info.Mode().IsDir()
	} else {
		return false
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func rangeRand(min,max int) int {
	return mathRand.Intn(max - min) + min
}
