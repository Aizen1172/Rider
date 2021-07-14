package service

import (
	"errors"
	"fmt"
	"github.com/zhangrui-git/pinyin"
	"gorm.io/gorm"
	"io"
	"mime/multipart"
	"os"
	"path"
	"riderStyemProject/global"
	"riderStyemProject/model"
	"riderStyemProject/model/request"
	"strings"
	"unicode"
)

// Creat 添加文件
func Creat(file *request.File) (err error) {
	//查询文件名字是否重复
	if errors.Is(global.DB.Where("name = ?", file.Name).First(&model.File{}).Error, gorm.ErrRecordNotFound) {
		_file := file.Creat()
		err = global.DB.Create(&_file).Error
		if err != nil {
			return errors.New("添加失败。")
		}
		//查询文件地址是否发生改变
	} else if errors.Is(global.DB.Where("url = ?", file.Url).First(&model.File{}).Error, gorm.ErrRecordNotFound) {
		_file := file.Update()
		err := global.DB.Where("name = ?", file.Name).Updates(&_file).Error
		if err != nil {
			return errors.New("更新失败。")
		}
	} else {
		return errors.New("图片重复。")
	}
	return
}

func Delete(fileId []int) (err error) {
	_file := model.File{}
	//判断文件地址是否在本地地址
	if strings.Contains(_file.Url, global.Path) {
		fmt.Println(_file.Url)
		err = os.Remove(_file.Url)	//删除文件
		if err != nil {
			return errors.New("删除失败。")
		}
	}
	return global.DB.Where("id in ?", fileId).Delete(&_file).Error
}

// Upload 文件上传
func Upload(file *multipart.FileHeader) (err error) {
	ext := path.Ext(file.Filename)                 //获得后缀名
	name := strings.TrimSuffix(file.Filename, ext) //获得文件名,不包括后缀名
	for _, i := range name {
		//判断字符串是否为中文
		if unicode.Is(unicode.Han, i) {
			//对中文进行处理，转化为拼音
			py := pinyin.New(name)
			name = strings.Join(py.FullPinyin, "")
		}
	}
	filename := name + ext                   //拼接文件名
	filePath := global.Path + "/" + filename //拼接文件路径
	//判断文件是否存在
	if global.IsExist(global.Path) {
		fmt.Println(filename + "已存在。")
	}
	//创建路径目录
	if err := os.MkdirAll(global.Path, os.ModePerm); err != nil {
		fmt.Println(err)
	}
	//打开文件
	_file, err := file.Open()
	if err != nil {
		fmt.Println(err)
	}
	//创建新的文件
	dst, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
	}
	//复制文件
	_, err = io.Copy(dst, _file)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(filePath)
	//把文件添加到数据库
	ext = strings.Trim(ext, ".") //移除指定字符
	ext = strings.ToUpper(ext)   //转成大小写
	creatFile := &request.File{
		Url:  filePath,
		Name: file.Filename,
		Type: ext,
	}
	if err := Creat(creatFile); err != nil {
		return err
	}
	return
}
