# chatlog_alpha

原 [chatlog](https://github.com/sjzar/chatlog) 项目的二开版本，导入自 [xiaofeng2042 的分支](https://github.com/xiaofeng2042/chatlog)，以防止上游删库后分支被自动删除。

未经修改的源代码在 [main 分支](https://github.com/CJYKK/chatlog_backup/tree/main)，本人不对代码中的任何内容负责。

# 提示
获取key请前往（https://github.com/ycccccccy/wx_key）下载

更新日志：

1.2025年12月13日同步了wx_key项目相关图片解密处理逻辑

2.http://127.0.0.1:5030/默认页面增加了清除图片缓存功能，方便测试

3.对ffmpeg未安装导致的dat转换失败进行显示提示

#重要提示
1.对dat转换一定要安装ffmpeg，并且在系统变量设置bin目录的path，否则会显式报错

2.图片解密aes key请转为hex，如：16位aes key -》 34000000386000006538323730000000
