## 安装
go get -u github.com/nicksnyder/go-i18n/goi18n, 并且配置goi18n到PATH

## 验证
goi18n -help

## 使用步骤
1. 以中文开发为例，在zh.all.json中加入一个新的string
```
{
    "id": "test_hello",
    "translation": "你好世界"
}
    ...
```
2. 进入Project/config/i18n/目录，运行命令
```
goi18n -sourceLanguage=zh  *.all.json, 该命令会在config/i18n目录下生成国际化文件和untranslated文件
```
3. 翻译untranslated里的内容
4. 再次运行goi18n命令去合并untranslated文件内容 
``` 
goi18n ./*.json ./*.untranslated.json
```