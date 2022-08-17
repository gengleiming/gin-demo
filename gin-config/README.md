## 安装
```shell
go get -u github.com/spf13/viper 
```
## 使用
`viper`可将配置文件`config.yaml`对应的值映射到`global.App.Config`中  
可通过 `v.AutomaticEnv()` 自动载入环境变量，负载配置文件中的值  

