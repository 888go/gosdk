# 以下是配置例子
# 目录可以留空, 留空就是整个项目所有文件都替换
# [github.com/gin-gonic/gin] 
# 目录=\
# 新地址=github.com/888go/gin
# 
# 节名称可以使用部分,如github.com/gin-gonic
# [github.com/gin-contrib]
# 目录=\gin-contrib
# 新地址=github.com/888go/gin

#internal
[internal/]
目录=\exec
新地址=github.com/888go/gosdk/exec/internal/

[os/exec/internal/]
目录=\exec
新地址=github.com/888go/gosdk/exec/exec/internal/