# 南信大校园网登录工具

一个可以操作南信大 i-NUIST 校园网的小工具，使用 Go 编写。支持列出可用渠道、登录、展示在线状态、注销。

## 获取

Windows 64位：直接在 [Release](https://github.com/XeroAlpha/NUISTLanLogin/releases/) 页面下载

其他系统：下载或 clone 源代码后使用 `go build` 手动构建

## 快捷使用方法（以 Windows 用户为例）

1. 获取可执行文件
2. 在可执行文件的同一目录下创建一个批处理文件，内容如下
```
@echo off
NUISTLanLogin login <用户名> <密码> <渠道名>
pause
```
3. 创建到这个批处理文件的快捷方式（按住 Alt 键拖动）
4. 将快捷方式放到你想放置的地方

## 命令行用法

```
NUISTLanLogin {list|login|status|logout} <用户名> <密码> [<渠道名>]

list   - 列出所有可用的渠道
login  - 登录校园网，需要提供 <渠道名>
status - 展示当前在线状态
logout - 注销

例如：NUISTLanLogin login 12345678901 123456 中国电信
```