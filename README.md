### 功能

该程序会遍历指定目录下的所有文件（包括子文件夹中的文件），将特定的文件中的制表符`\t`替换为指定的空格数量

### 使用场景

在代码编写的过程中，会经常使用`Tab`键进行缩进，但由于不同的编辑器或网站对`\t`的显示方式不同，所以显示效果不同，比如`Github`的默认缩进长度是8个空格，显示效果较差

### 用法

```shell
# 克隆项目到本地
git clone git@github.com:kuonz/fix-indentation.git

# 编译为可执行文件
go build fix-indentation.go

# 运行可执行文件
./fix-indentation <参数列表>
```

参数列表

| 参数 | 说明                         | 默认值                        | 是否必须 | 缩写含义     |
| ---- | ---------------------------- | ----------------------------- | -------- | ------------ |
| -p   | 起始目录路径                 | fix-indentation程序所在的路径 | 不必须   | p ==> path   |
| -n   | 一个制表符代表的空格数量     | 2                             | 不必须   | n ==> number |
| -s   | 匹配进行制表符替换的文件类型 | 无                            | 必须     | s ==> suffix |

### 使用示例

```shell
# 查看参数列表提示
./fix-indentation -h

# 将 fix-indentation 目录下所有的 .go 文件中每个制表符替换为2个空格
./fix-indentation -s=".go"

# 将 fix-indentation 目录下所有的 .txt 文件中每个制表符替换为4个空格
./fix-indentation -n=4 -s=".txt"

# 将 /home/user/project 目录下所有的 .js 文件中每个制表符替换为8个空格
./fix-indentation -p="/home/user/project" -n=8 -s=".js"
```

技巧：如果不想输入起始目录路径，可以将`fix-indentation`拷贝到特定目录路径，然后再执行程序