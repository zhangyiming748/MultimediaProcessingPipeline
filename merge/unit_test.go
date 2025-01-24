package merge

import (
	"Multimedia_Processing_Pipeline/replace"
	"Multimedia_Processing_Pipeline/util"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func init() {

}

// go test -v -run TestMerge
func TestMerge(t *testing.T) {
	root := "C:\\Users\\zen\\Videos\\export\\KAGP-116"
	mp4s, _ := getMP4Files(root)
	for _, mp4 := range mp4s {
		Mp4WithSrt(mp4)
	}
}

// go test -timeout 2000m -v -run TestInsideMerge
func TestInsideMerge(t *testing.T) {
	root := "C:\\Users\\zen\\Videos\\export\\sdde\\KAGP-116 素人少女撒尿百科4 13人在镜头前撒尿"
	dirPath := path.Dir(root)
	mp4s, _ := getMP4Files(root)
	cmds := []string{}
	for _, mp4 := range mp4s {
		cmd := Mp4WithSrtHard(mp4)
		if strings.Contains(cmd, "C:\\Program Files\\ffmpeg\\bin\\") {
			cmd = strings.Replace(cmd, "C:\\Program Files\\ffmpeg\\bin\\ffmpeg", "ffmpeg", -1)
		}
		if strings.Contains(cmd, root) {
			cmd = strings.Replace(cmd, root, "", 1)
		}
		cmds = append(cmds, cmd)
	}
	fp := ""
	if runtime.GOOS == "windows" {
		fp = filepath.Join(root, "mergeInside.ps1")
		log.Println(fp)
		util.WriteByLine(fp, cmds)
	} else {
		fp = filepath.Join(dirPath, "mergeInside.sh")
		util.WriteByLine(fp, cmds)
	}

}

// getMP4Files 遍历指定目录，返回所有 mp4 文件的路径
func getMP4Files(dir string) ([]string, error) {
	var mp4Files []string

	// 使用 Walk 函数遍历目录
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 检查文件是否是 mp4 文件
		if !info.IsDir() && strings.HasSuffix(strings.ToLower(info.Name()), ".mp4") {
			mp4Files = append(mp4Files, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return mp4Files, nil
}

func TestReplaceEnglish(t *testing.T) {
	input := "这是一个测试字符串 [包含方括号内容],请忽略这部分内容。"
	ret := replace.ReplaceEnglishSquareBrackets(input)
	t.Log(ret)
}

func TestRename(t *testing.T) {
	srt := "C:\\Users\\zen\\Github\\MultimediaProcessingPipeline\\test\\Impregnation, No Strings Attached [674e2d97be0f7].srt"
	dst := renameSrt(srt)
	t.Log(dst)
}
func TestMp4Inside(t *testing.T) {
	mp4 := "C:\\Users\\zen\\Github\\MultimediaProcessingPipeline\\test\\Impregnation, No Strings Attached [674e2d97be0f7].mp4"
	srt := "C:\\Users\\zen\\Github\\MultimediaProcessingPipeline\\test\\Impregnation, No Strings Attached [674e2d97be0f7].srt"
	Mp4Inside(mp4, srt)
}
