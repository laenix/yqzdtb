# 沈阳理工大学疫情自动填报
### 用法
把 main.go  中的学号和密码填入自己的学号和密码
### 安装golang
golang
你可以
>sudo apt-get install golang
来安装
### go的依赖
>export GO111MODULE=on

>export GOPROXY=https://goproxy.io

>go get github.com/robfig/cron

>go get github.com/PuerkitoBio/goquery
### 使用
>go init yqzdtd
>go build
>nohup ./yqzdtb &
