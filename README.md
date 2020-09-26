<h1 align="center">
    - 东北大学体育场馆预约工具 -
</h1>
<p align="center">
    <img src="img/logo.png">
</p>
<p align="center">
<img src="https://img.shields.io/github/go-mod/go-version/amtoaer/iwsp?longCache=true&style=for-the-badge">
<img src="https://goreportcard.com/badge/github.com/amtoaer/iwsp?longCache=true&style=for-the-badge">
<img src="https://img.shields.io/github/workflow/status/amtoaer/iwsp/Go?longCache=true&style=for-the-badge&color=%23ea7070">
<img src="https://img.shields.io/github/license/amtoaer/iwsp?longCache=true&style=for-the-badge&color=%23e59572">
</p>



## 介绍

> 该项目处于早期开发阶段，暂时仅支持风雨操场的预约。

IWSP<del>（I wanna do sports!）</del>是一款东北大学体育场馆预约工具，使用[neugo](https://github.com/neucn/neugo)进行一网通登录。

## 截图

![](img/screenshot.png)

## 安装

有两种安装途径：

1. 前往[release界面](https://github.com/amtoaer/iwsp/releases)下载对应平台的最新版本。

2. 克隆仓库后自行编译（需要安装`golang`环境）。

   ```bash
   git clone https://github.com/amtoaer/iwsp
   cd iwsp
   go build -v .
   ```

## 用法

如`iwsp --help`所示：

```
iwsp 东北大学场馆预约工具

        -u 一网通学号
        -p 一网通密码
        -s 保存学号密码到配置文件
        -v 使用webVPN，默认不使用
        -o 输出历史预约列表
        -l 预约地点，可选值fycc
        -t 预约时段，可选值
                07：00-10：00
                10：40-12：30
                12：30-14：00
                14：00-16：00
                16：00-18：00
                18：00-19：30
                19：30-21：00
        -d 启用debug模式
        -h 打印该帮助信息

```

基本的用法示例：

1. 查看历史预约列表

   ```bash
   iwsp -u username -p password -o
   ```

2. 保存帐号并查看历史预约列表（保存帐号后再进行操作则无需`-u/-p`参数）

   ```bash
   iwsp -u username -p password -s -o
   ```

3. 预约风雨操场的`7:00-10:00`场（假设已经使用`-s`参数保存过帐号）

   ```bash
   iwsp -l fycc -t 07:00-10:00
   ```

## 依赖

使用到的第三方库：

1. [neugo](https://github.com/neucn/neugo):东北大学一网通操作库
2. [go-homedir](https://github.com/mitchellh/go-homedir):跨平台获取用户目录
