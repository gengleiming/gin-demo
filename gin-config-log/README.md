## 安装
```shell
# viper管理配置文件
go get -u github.com/spf13/viper 
# 日志库zap
go get -u go.uber.org/zap
# lumberjack日志切割
go get -u gopkg.in/natefinch/lumberjack.v2
```
## 使用
1. `viper`可将配置文件`config.yaml`对应的值映射到`global.App.Config`中  
2. 可通过 `v.AutomaticEnv()` 自动载入环境变量，负载配置文件中的值  
3. 使用日志库zap + 日志文件管理lumberjack

