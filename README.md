# Agenda-cli-service


## 1、概述


利用命令行 或 web 客户端调用远端服务是服务开发的重要内容。其中，要点是如何实现 API First 开发，使得团队协作变得更有效率。


### 任务目标


1. 熟悉 API 设计工具，实现从资源（领域）建模，到 API 设计的过程
1. 使用 Github ，通过 API 文档，实现 agenda 命令行项目 与 RESTful 服务项目同步开发
1. 使用 API 设计工具提供 Mock 服务，两个团队独立测试 API
1. 使用 travis 测试相关模块
1. 利用 dockerfile 在 docker hub 上构建一个镜像，同时包含 agenda cli 和 agenda service， 如果 mysql 包含 服务器 和 客户端一样




## 2、agenda 开发项目


重构、或新建 agenda 项目，根目录必须包含
cli 目录
service 目录
.travis
apiary.apib
dockerfile
LICENSE
README.md
README-yourid.md 记录你的工作摘要（个人评分依据）




### - API 开发
- 使用 API Blueprint 设计 API
- 资源 URL 命名符合 RESTful 设计标准
- 资源 CRUD 基本完整
- 
### - API 客户端开发
- 可用命令 5 个以上
- 必须有 XXX-test.go 文件
- 
### - 服务端开发
- 使用 sqlite3 作为数据库
- 建议使用课程提供的服务端框架
- 必须有 XXX-test.go 文件
### - 容器镜像制作
- 在 docker hub 上生成镜像
- base 镜像 go-1.8
- 需要加载 sqlite3
- 同时包含客户端与服务器


#### API界面
![这里写图片描述](http://img.blog.csdn.net/20171214200915400?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvSE9NRVJVTklU/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)


### 镜像


#### 1. 下载镜像


```
sudo docker pull qiuxf/agenda-cli-service
```

![这里写图片描述](http://img.blog.csdn.net/20171214200950894?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvSE9NRVJVTklU/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)


#### 2. 启动服务器




```
sudo docker run -dit -v $GOPATH:/data qiuxf/agenda-cli-service service


```
![这里写图片描述](http://img.blog.csdn.net/20171214201013707?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvSE9NRVJVTklU/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)


#### 3. 启动客户端


```
sudo docker run --rm --network host qiuxf/agenda-cli-service cli -h


```
![这里写图片描述](http://img.blog.csdn.net/20171214201034753?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvSE9NRVJVTklU/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)






## 测试


#### 注册
![这里写图片描述](http://img.blog.csdn.net/20171214201052485?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvSE9NRVJVTklU/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)



#### 登陆
![这里写图片描述](http://img.blog.csdn.net/20171214201112727?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvSE9NRVJVTklU/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

####  列出所有用户
![这里写图片描述](http://img.blog.csdn.net/20171214201224590?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvSE9NRVJVTklU/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

#### 查找用户
![这里写图片描述](http://img.blog.csdn.net/20171214201247928?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvSE9NRVJVTklU/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)


####创建会议
![这里写图片描述](http://img.blog.csdn.net/20171214202015072?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvSE9NRVJVTklU/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)
####列出所有会议
![这里写图片描述](http://img.blog.csdn.net/20171214202042772?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvSE9NRVJVTklU/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)
#### 查询会议
![这里写图片描述](http://img.blog.csdn.net/20171214202055911?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvSE9NRVJVTklU/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)


###服务端
![这里写图片描述](http://img.blog.csdn.net/20171214202243195?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvSE9NRVJVTklU/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

![这里写图片描述](http://img.blog.csdn.net/20171214202257842?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvSE9NRVJVTklU/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

![这里写图片描述](http://img.blog.csdn.net/20171214202310780?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvSE9NRVJVTklU/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

![这里写图片描述](http://img.blog.csdn.net/20171214202321746?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvSE9NRVJVTklU/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)
