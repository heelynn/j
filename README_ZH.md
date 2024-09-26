# j

[English wiki](https://github.com/heelynn/j/blob/main/README.md)

**注意：**`master`分支可能处于开发之中并**非稳定版本**，请通过 tag 下载稳定版本的源代码，或通过[release](https://github.com/heelynn/j/releases)下载已编译的二进制可执行文件。

`j`是一个 Linux、macOS、Windows 下的命令行工具，可以提供一个便捷的多版本 Java 环境的管理和切换。

## 手工安装

- windows 
  
 ```shell
 # 使用powershell
 mkdir $HOME/.j/bin,$HOME/.j/version
 # 选择windows版本下载安装包，下载地址：https://github.com/heelynn/j/releases/latest
 # 将j.exe文件解压到$HOME/.j/bin目录下 
 ```
配置环境变量,建议`使用系统界面`设置环境变量，使环境永久生效。
 ```shell
 $env:JAVA_HOME="$HOME/.j/java"
 $env:PATH="$HOME/.j/bin:$JAVA_HOME/bin;$env:PATH;"
 ```

 
- macOS/Linux

 ```shell
 # 使用shell
 $ mkdir -p $HOME/.j/bin $HOME/.j/version
 # 选择liunx/macos版本下载安装包，下载地址：https://github.com/heelynn/j/releases/latest
 
 # 将j文件解压到$HOME/.j/bin目录下 
 # 给j文件添加可执行权限
 $ chmod +x $HOME/.j/bin/j
 
 # 配置环境变量（适用于 bash、zsh）
 $ export JAVA_HOME="$HOME/.j/java"
 $ export PATH="$HOME/.j/bin:$JAVA_HOME/bin:$PATH"
 ```

## 使用方法

### 下载并安装Java
将下载好的Java安装包`Redhat、openJDk、OracleJDK等`解压到`$HOME/.j/version`目录下。
建议修改为简单文件夹名，如`Java8、Java11、Java17`，方便管理。
 ```shell
 # 列出已安装的Java版本
 $ ls $HOME/.j/version
 Java8
 Java11
 Java17
 ```

### 查看当前Java版本
`使用j命令`查看当前Java版本，与`$HOME/.j/version`目录下Java版本对应
 ```shell
 $ j ls
  java8
  java11
 ```

`使用j命令`切换Java版本，`j ls` 命令查看当前已安装的Java版本
 ```shell
 $ j use java8
 ```
### 验证Java版本
 ```shell
 $ java -version
 ```


