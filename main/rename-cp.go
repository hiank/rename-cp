package main

// import (
// 	"os"
// 	"fmt"
// 	"flag"
// 	"github.com/hiank/rename-cp"
// )

// func main() {

// 	srcDir := flag.String("s", "", "src dir")
// 	dstDir := flag.String("d", "", "dst dir")
// 	// mixLen := flag.Int("l", 300, "the num of the mix byte")

// 	flag.Parse()

// 	switch {
// 	case *srcDir == "": 
// 		fmt.Println("should define srcDir with -s") 
// 		fallthrough
// 	case *dstDir == "": 
// 		fmt.Println("should define dstDir with -d") 
// 		return
// 	}
// 	os.RemoveAll(*dstDir)
// 	os.Mkdir(*dstDir, 0755)
// 	rc.DuplicateDir(*srcDir, *dstDir)
// }


import (
	"github.com/hiank/rename-cp"
	"fmt"
	"time"
	"crypto/aes"
	"crypto/cipher"
	"bytes"
	"encoding/base64"
	"flag"
	"os"
	"io/ioutil"
	"strings"
	"path"
	"io"
	"path/filepath"
	"crypto/md5"
	"encoding/hex"
	"crypto/rand"
)

var inputDir = flag.String("i", "", "input Dir")
var outputDir = flag.String("o", "", "output Dir")
var key = flag.String("k", "", "Encrypt Key")
var isUpdate = flag.Int("u", 0, "For update pacage: 1 else: 0")

var rename = flag.Int("r", 0, "Rename file: 1 else: 0")

var myKey []byte
var mapList map[string]string

func main() {

	flag.Parse()

	fmt.Println("encode begin ==")

	time1 := time.Now()

	//file, _ := exec.LookPath(os.Args[0])
	//path, _ := filepath.Abs(file)
	//println(path)

	os.RemoveAll(*outputDir)
	os.Mkdir(*outputDir, 0777)

	//fmt.Println("encode test begin 2==")
	myKey = []byte(*key)
	if len(myKey) > 16 {
		myKey = myKey[:16]
	}
	if len(myKey) < 16 {
		//addStr :=  strings.Repeat("#", 16-len(key))
		padtext := bytes.Repeat([]byte{0}, 16-len(myKey))
		myKey = append(myKey, padtext...)
	}

	mapList = make(map[string]string)

	relPath := ""
	processDir(inputDir, outputDir, &relPath)
	generateKeyFile(outputDir)
	//testAes()

	ds := time.Since(time1)
	fmt.Println(ds, "encode end ====== ")
}

func processFile(filePath *string, filePathOut *string)  {
	//file, error := os.Open( *filePath)
	//if error != nil {
	//	fmt.Println("error read file !", error)
	//	return
	//}
	//
	//defer file.Close()
	EncodeFile(*filePath, *filePathOut)
}


func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func UniqueId() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(b))
}


func processDir(subDir *string, outSubDir *string, relPath *string){
	dir, err := ioutil.ReadDir(*subDir)
	if err != nil {
		fmt.Printf("error read dir")
		return
	}

	for _, fi := range dir {
		//		fmt.Printf(fi.Name() + "\n")
		if fi.IsDir() { // 忽略目录
			//continue
			path := *subDir + "/" + fi.Name()
			rPath := *relPath + "/" + fi.Name()
			pathOut := *outSubDir + "/" + fi.Name()
			processDir(&path, &pathOut, &rPath)
			continue
		}

		if 1==1 || strings.HasSuffix(strings.ToLower(fi.Name()), "xlsx") { //匹配文件
			//files = append(files, dirPth+PthSep+fi.Name())
			//name := Substr(fi.Name(), 0, len(fi.Name()) - 5)
			lowName := strings.ToLower(fi.Name())
			switch {
				case strings.HasSuffix(lowName, "mp3"): fallthrough
				case strings.HasSuffix(lowName, "wav"): 
					filePath := *subDir + "/" + fi.Name()
					filePathOut := *outSubDir + "/" + fi.Name()

					//fmt.Printf(filePath + "-->" + filePathOut + "\n" )
					CopyFile(filePath, filePathOut)
				case Substr(fi.Name(), 0, 1) != ".":
					filePath := *subDir + "/" + fi.Name()
					var filePathOut string

					if *rename == 1 {
						// newName := UniqueId()
						newName := rc.RandName(fi.Name())
						filePathOut = *outputDir + "/data/" + newName

						strPath := *relPath + "/" + fi.Name()

						//if strings.HasSuffix(strings.ToLower(strPath), "jsc"){
						//	strPath = strPath[0: len(strPath) -1]
						//}
						mapList[strPath] = newName
					}else{
						filePathOut = *outSubDir + "/" + fi.Name()
					}
					//fmt.Printf(filePath + "-->" + filePathOut + "\n" )
					processFile(&filePath, &filePathOut)		
			}

		}
	}
}

func testAes() {
	// AES-128。key长度：16, 24, 32 bytes 对应 AES-128, AES-192, AES-256
	key := []byte("test123456789012")

	if len(key) > 16 {
		key = key[:16]
	}
	if len(key) < 16 {
		//addStr :=  strings.Repeat("#", 16-len(key))
		padtext := bytes.Repeat([]byte{0}, 16-len(key))
		key = append(key, padtext...)
	}

	result, err := AesEncrypt([]byte("pol32Fs23s@studygolan##%$^HDSD测试g"), key)
	if err != nil {
		panic(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(result))

	origData, err := AesDecrypt(result, key)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(origData))
}

func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func Substr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}

//
func CopyFile(src, des string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println(err)
	}
	defer srcFile.Close()

	//fmt.Printf("check path ==" + path.Dir(des))
	os.MkdirAll(path.Dir(des), 0777)
	desFile, err := os.Create(des)
	if err != nil {
		fmt.Println(err)
	}
	defer desFile.Close()

	return io.Copy(desFile, srcFile)
}

func generateKeyFile(des *string)  {
	keyFile := *des + "/" + "sh.mk"

	os.MkdirAll(path.Dir(keyFile), 0777)
	desFile, err := os.Create(keyFile)
	if err != nil {
		fmt.Println(err)
	}
	defer desFile.Close()

	//result, err := AesEncrypt([]byte("pol32Fs23s@studygolan##%$^HDSD测试g"), myKey)

	keyStr := string(myKey)
	if *rename == 1 {
		var mapString string
		for k, v := range mapList {
			mapString += k + ":" + v + "#"
			fmt.Printf("check === %v === %v \n ", k, v)
		}
		//fmt.Printf("check == =%v \n ", mapString)
		//desFile.WriteString(mapString)
		//desFile.Write()
		keyStr += mapString
	}

	result, err := AesEncrypt([]byte(keyStr), []byte("ftmain@gxjnZ&3a+"))
	//result, err := AesDecrypt(plaintext, myKey)
	if err != nil {
		panic(err)
	}

	buf := make([]byte, 4)
	buf[0] = byte(6)
	buf[1] = byte(7)
	buf[2] = byte(70)
	buf[3] = byte(84)

	desFile.Write(buf)
	desFile.Write(result)
}

func EncodeFile(src, des string) {
	//fmt.Printf("check path ==" + path.Dir(des))
	os.MkdirAll(path.Dir(des), 0777)
	desFile, err := os.Create(des)
	if err != nil {
		fmt.Println(err)
	}
	defer desFile.Close()


	plaintext, err := ioutil.ReadFile(src)
	if err != nil {
		fmt.Printf("error read file " + src)
		return
	}

	fileName := filepath.Base(src)
	fmt.Printf( "encode -->" + fileName + "\n" )

	//result, err := AesEncrypt([]byte("pol32Fs23s@studygolan##%$^HDSD测试g"), myKey)

	var result []byte
	buf := make([]byte, 4)
	if *isUpdate == 1 || fileName == "version.manifest" || fileName == "project.manifest"{
		result, err = AesEncrypt(plaintext, []byte("ftmain@gxjnZ&3a+"))
		buf[1] = byte(7)
	}else {
		result, err = AesEncrypt(plaintext, myKey)
		buf[1] = byte(6)
		//result, err := AesDecrypt(plaintext, myKey)
	}

	if err != nil {
		panic(err)
	}


	buf[0] = byte(6)
	buf[2] = byte(70)
	buf[3] = byte(84)

	desFile.Write(buf)
	desFile.Write(result)
}

