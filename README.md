# Multimedia_Processing_Pipeline
使用了yt-dlp下载视频 然后使用openai-whisper生成视频 再使用translate-shell 翻译为双语字幕 最后合成为mkv外挂字幕的视频 以上这些功能会以流水线的形式处理多个视频

# TODO
- [ ] 以一个文件为单位分别写好`yt-dlp` `whisper` `trans` `merge` 部分 唯一标识符为文件的绝对路径
- [x] 新建环境变量`$merge=true|false`判断 是否同时进行字幕嵌入并转码
- [x] 如果在容器中运行需要提前结束 可以`touch /exit`