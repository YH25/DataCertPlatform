package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"strings"
	"os"
	"io"
	"crypto/sha256"
	"io/ioutil"
	"encoding/hex"
)

/**
 * 该控制器结构体用于处理文件上传的功能
 */
type UploadFileController struct {
	beego.Controller
}

/**
 * 该post方法用于处理用户在客户端提交的文件
 */
func (u *UploadFileController) Post() {
	title := u.Ctx.Request.PostFormValue("upload_title") //用户输入的标题
	fmt.Println("电子数据标签：", title)
	file, header, err := u.GetFile("yuhongwei")
	if err != nil { //解析客户端提交的文件出现错误
		u.Ctx.WriteString("抱歉，文件解析失败，请重试！")
		return
	}

	defer file.Close() //延迟执行 空指针错误：invalid memory or nil pointer dereferenece

	//使用io包提供的方法保存文件
	saveFilePath := "static/upload/" + header.Filename
	saveFile, err := os.OpenFile(saveFilePath, os.O_CREATE|os.O_RDWR, 777)
	if err != nil {
		u.Ctx.WriteString("抱歉, 电子数据认证失败, 请重试!")
		return
	}

	_, err = io.Copy(saveFile, file)
	if err != nil {
		u.Ctx.WriteString("抱歉, 电子数据认证失败, 请重新尝试!")
		return
	}

	//2、计算文件的SHA256值
	hash256 := sha256.New()
	fileBytes, _ := ioutil.ReadAll(file)
	hash256.Write(fileBytes)
	hashBytes := hash256.Sum(nil)
	fmt.Println(hex.EncodeToString(hashBytes))

	//先查询用户id

	//把上传的文件作为记录保存到数据库当中

	u.Ctx.WriteString("恭喜，已接收到上传文件!")
}

/**
 * 该post方法用于处理用户在客户端提交的认证文件
 */
func (u *UploadFileController) Post1() {
	//1、解析用户上传的数据及文件内容
	//用户上传的自定义的标题
	title := u.Ctx.Request.PostFormValue("upload_title") //用户输入的标题

	//用户上传的文件
	file, header, err := u.GetFile("yuhongwei")
	if err != nil { //解析客户端提交的文件出现错误
		u.Ctx.WriteString("抱歉，文件解析失败，请重试！")
		return
	}
	defer file.Close()

	fmt.Println("自定义的标题：", title)
	//获得到了上传的文件
	fmt.Println("上传的文件名称:", header.Filename)
	//eg：支持jpg,png类型，不支持jpeg，gif类型
	//文件名： 文件名 + "." + 扩展名
	fileNameSlice := strings.Split(header.Filename, ".")
	fileType := fileNameSlice[1]
	fmt.Println(fileNameSlice)
	fmt.Println(":", strings.TrimSpace(fileType))
	isJpg := strings.HasSuffix(header.Filename, ".jpg")
	isPng := strings.HasSuffix(header.Filename, ".png")
	if !isJpg && !isPng {
		//文件类型不支持
		u.Ctx.WriteString("抱歉，文件类型不符合, 请上传符合格式的文件")
		return
	}

	//if fileType != " jpg" || fileType != "png" {
	//	//文件类型不支持
	//	u.Ctx.WriteString("抱歉，文件类型不符合, 请上传符合格式的文件")
	//	return
	//}

	//文件的大小 200kb
	config := beego.AppConfig
	fileSize, err := config.Int64("file_size")

	if header.Size/1024 > fileSize {
		u.Ctx.WriteString("抱歉，文件大小超出范围，请上传符合要求的文件")
		return
	}

	fmt.Println("上传的文件的大小:", header.Size) //字节大小

	//perm: permission 权限
	//权限的组成： a+b+c
	// a: 文件所有者对文件的操作权限， 读4、写2、执行1
	// b: 文件所有者所在组的用户的操作权限， 读4、写2、执行1
	// c: 其他用户的操作权限，读4、写2、执行1

	//eg：m文件，权限是：651。
	//判断题：文件所有者对该m文件有写权限（对）
	//文件的所有者所在组用户对该文件有写权限(错误)
	//有一个文件n，该文件的权限是985(错误)
	saveDir := "static/upload"
	//①先尝试打开文件夹
	_, err = os.Open(saveDir)

	if err != nil { //打开失败：文件夹不存在
		//②自己动手,创建文件夹
		err = os.Mkdir(saveDir, 777)
		if err != nil { // file exists
			fmt.Println(err.Error())
			u.Ctx.WriteString("抱歉，文件认证遇到错误，请重试!")
			return
		}
	}

	//文件名： 文件路径 + 文件名 + "." + 文件扩展名Filename
	saveName := saveDir + "/" + header.Filename
	fmt.Println("要保存的文件名", saveName)
	//fromFile: 文件，
	//toFile: 要保存的文件路径
	err = u.SaveToFile("yuhongwei", saveName)
	if err != nil {
		fmt.Println(err.Error())
		u.Ctx.WriteString("抱歉，文件认证失败，请重试!")
		return
	}

	fmt.Println("上传的文件:", file)
	u.Ctx.WriteString("已获取到上传文件。")
}
