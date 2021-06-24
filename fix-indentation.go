package main

import (
  "flag"
  "fmt"
  "io/ioutil"
  "strings"
)

var fileSuffix string     // 存储匹配的文件后缀名
var indentation int       // 一个\t代表的空格数量
var processFiles []string // 进行处理的文件列表

func main() {

  dirpath, indent, suffix := getInput()

  if len(suffix) == 0 {
    fmt.Println("程序退出：没有指定文件后缀，可以输入参数 -h 查看帮助")
    return
  }

  fileSuffix = suffix
  indentation = indent

  var input string
  fmt.Printf(`处理路径为"%s"，用%d个空格替代一个制表符，是否确认执行？(y/n) `, dirpath, indentation)
  fmt.Scanf("%s", &input)

  if input != "y" && input != "yes" {
    fmt.Println("用户取消，程序退出")
  } else {
    if strings.HasSuffix(dirpath, "/") {
      dirpath = (dirpath)[:len(dirpath)-1]
    }

    traversal(dirpath, "")
  }

  fmt.Println("执行完毕，有下列文件进行了Tab替换：")

  for _, item := range processFiles {
    fmt.Println(item)
  }
}

func traversal(dirpath string, fileName string) {

  currentPath := dirpath + fileName + "/"

  infos, _ := ioutil.ReadDir(currentPath)

  for _, item := range infos {
    itemName := item.Name()

    if item.IsDir() {
      traversal(currentPath, itemName)
    } else {
      if strings.HasSuffix(itemName, fileSuffix) {
        currentFilePath := currentPath + itemName
        replaceTab(currentFilePath)
        processFiles = append(processFiles, currentFilePath)
      }
    }
  }
}

func getInput() (string, int, string) {
  dirpath := flag.String("p", `.`, "请输入起始目录路径(绝对路径或相对路径)")
  indentation := flag.Int("n", 2, "请输入一个制表符代表的空格数量，默认为2个")
  suffix := flag.String("s", "", "请输入匹配的文件后缀，如 .txt")

  flag.Parse()

  return *dirpath, *indentation, *suffix
}

func replaceTab(filePath string) {

  blank := ""
  for i := 0; i < indentation; i++ {
    blank += " "
  }

  contentByte, err := ioutil.ReadFile(filePath)

  if err != nil {
    panic("读取文件失败：err: " + err.Error())
  }

  contentString := string(contentByte)

  newContent := strings.ReplaceAll(contentString, "\t", blank)

  ioutil.WriteFile(filePath, []byte(newContent), 0666)
}
