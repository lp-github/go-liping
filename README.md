                                      实验报告-go开发环境安装
  1.github的基础使用方法：
      确保有一个github的账号。在登录了github之后新建一个叫做go-user（user取自己的名字）的仓库，勾选使用readme文档初始化。
     
      在本地目录github.com/user/下初始化仓库
      $:git init
      
      将修改添加到暂存区:
      $:git add
      
      提交修改
      $：git commit -m “comment”
      
      将提交推送到远程仓库
      $: git push
      在第一次推送时会让你设置仓库名和url，通过git push 仓库名的方式实现提交
      这个地方和windows就有所不同，dos下直接在第一次输入账号密码后就不需要再次输入账号密码，不过centos好像比较麻烦
      之后我们就讲本地代码和远程代码成功同步啦
2.go安装和环境配置
     安装go环境：
     通过教程给的：sudo yum install golang安装
     
     配置环境变量：
      $ mkdir $HOME/gowork
      
      新建~/.profile文件并添加内容
      export GOPATH=$HOME/gowork
      export PATH=$PATH:$GOPATH/bin
      将go工作环境加入环境变量
      
 3.go程序运行，测试与安装
      go run xxx.go
      
      go test
      
      go install
