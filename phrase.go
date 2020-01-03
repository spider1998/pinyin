package pinyin

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/yanyiwu/gojieba"
)

var (
	DICT_DIR        string
	DICT_PATH       string
	HMM_PATH        string
	USER_DICT_PATH  string
	IDF_PATH        string
	STOP_WORDS_PATH string
)

func init() {
	currentPath, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	DICT_DIR = path.Join(currentPath, "dict")
	DICT_PATH = path.Join(DICT_DIR, "jieba.dict.utf8")
	HMM_PATH = path.Join(DICT_DIR, "hmm_model.utf8")
	USER_DICT_PATH = path.Join(DICT_DIR, "user.dict.utf8")
	IDF_PATH = path.Join(DICT_DIR, "idf.utf8")
	STOP_WORDS_PATH = path.Join(DICT_DIR, "stop_words.utf8")
}

func cutWords(s string) []string {
	var jieba *gojieba.Jieba
	if runtime.GOOS == "windows" {
		jieba = gojieba.NewJieba()
	} else {
		jieba = gojieba.NewJieba(DICT_PATH, HMM_PATH, USER_DICT_PATH, IDF_PATH, STOP_WORDS_PATH)
	}

	defer jieba.Free()

	return jieba.CutAll(s)
}

func pinyinPhrase(s string) string {
	words := cutWords(s)
	for _, word := range words {
		match := phraseDict[word]
		if match == "" {
			match = phraseDictAddition[word]
		}

		match = toFixed(match, paragraphOption)
		if match != "" {
			s = strings.Replace(s, word, " "+match+" ", 1)
		}
	}

	return s
}
