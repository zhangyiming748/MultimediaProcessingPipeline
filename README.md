# Multimedia_Processing_Pipeline

使用了yt-dlp下载视频然后使用openai-whisper生成视频再使用translate-shell翻译为双语字幕最后合成为mkv外挂字幕的视频以上这些功能会以流水线的形式处理多个视频

# TODO

 -[x] 以一个文件为单位分别写好`yt-dlp``whisper``trans``merge`部分唯一标识符为文件的绝对路径
 -[x] 新建环境变量`$merge=true|false`判断是否同时进行字幕嵌入并转码
 -[x] 如果在容器中运行需要提前结束可以`touch/exit`
 -[ ] 想办法整合成流水线模式

## Availablemodelsandlanguages

|Size|Parameters|English-onlymodel|Multilingualmodel|RequiredVRAM|Relativespeed|
|:---:|:---:|:---:|:---:|:---:|:---:|
|tiny|39M|[tiny.en](https://openaipublic.azureedge.net/main/whisper/models/d3dd57d32accea0b295c96e26691aa14d8822fac7d9d27d5dc00b4ca2826dd03/tiny.en.pt)|[tiny](https://openaipublic.azureedge.net/main/whisper/models/65147644a518d12f04e32d6f3b26facc3f8dd46e5390956a9424a650c0ce22b9/tiny.pt)|~1GB|~32x|
|base|74M|[base.en](https://openaipublic.azureedge.net/main/whisper/models/25a8566e1d0c1e2231d1c762132cd20e0f96a85d16145c3a00adf5d1ac670ead/base.en.pt)|[base](https://openaipublic.azureedge.net/main/whisper/models/ed3a0b6b1c0edf879ad9b11b1af5a0e6ab5db9205f891f668f8b0e6c6326e34e/base.pt)|~1GB|~16x|
|small|244M|[small.en](https://openaipublic.azureedge.net/main/whisper/models/f953ad0fd29cacd07d5a9eda5624af0f6bcf2258be67c92b79389873d91e0872/small.en.pt)|[small](https://openaipublic.azureedge.net/main/whisper/models/9ecf779972d90ba49c06d968637d720dd632c55bbf19d441fb42bf17a411e794/small.pt)|~2GB|~6x|
|medium|769M|[medium.en](https://openaipublic.azureedge.net/main/whisper/models/d7440d1dc186f76616474e0ff0b3b6b879abc9d1a4926b7adfa41db2d497ab4f/medium.en.pt)|[medium](https://openaipublic.azureedge.net/main/whisper/models/345ae4da62f9b3d59415adc60127b97c714f32e89e936602e85993674d08dcb1/medium.pt)|~5GB|~2x|
|large|1550M|N/A|[largev1](https://openaipublic.azureedge.net/main/whisper/models/e4b87e7e0bf463eb8e6956e646f1e277e901512310def2c24bf0e11bd3c28e9a/large-v1.pt)<br>[largev2](https://openaipublic.azureedge.net/main/whisper/models/81f7c96c852ee8fc832187b0132e569d6c3065a3252ed18e56effd0b6a73e524/large-v2.pt)<br>[largev3](https://openaipublic.azureedge.net/main/whisper/models/e5b1a55b89c1367dacf97e3e19bfd829a01529dbfdeefa8caeb59b3f1b81dadb/large-v3.pt)|~10GB|1x|

# 获取token

程序运行前需要[获取](https://connect.linux.do/)deeplx的token
并且在程序运行前设置为环境变量
如 `export TOKEN=token`
