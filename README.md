# GoStudy
新手一步一步学习 Go 的过程 

- go下载模块的方法( 使用go modules )

  - 修改环境变量

    - ```
      export GOPROXY="https://goproxy.cn,direct"
      export GO111MODULE=on  // 开启go modules
      配置 GoPath路径
      ```
    
  - 初始化mod
  
    - ```
      go mod init ...
      ```
  
      
  
  - go下载第三方模块
  
    - ```bash
      go get -u github.com/go-sql-driver/mysql
      ```
  
  
    ​      

