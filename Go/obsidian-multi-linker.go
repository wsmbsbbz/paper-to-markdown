package main

// 依照目录中的文件名，建立一个Tire字典树，逐字符遍历input文章内容，
// 在字典树中查找其后的字符串，查找到即可记录起始和结尾下标，并添加
// "[["和"]]"

// 几个实用的正则表达式
// 1. 去除前缀
// (?<=\w)\[\[(\w*)\]\]    ---> $1
// a ex[[ample]] here      ---> a example here
// 2. 去除后缀长度>=3的
// \[\[(\w*)\]\](?=\w{3,}) ---> $1
// a [[exam]]ple here      ---> a example here
// 3. 去除所有标记
// \[\[\(\w*)]\]           ---> $1
// [[a example]] here      ---> a example here

import (
	"bufio"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type strIdx struct {
	str  string
	l, r int
}

var dicTrie *Trie
func init()  {
    dicTrie = &Trie{}
}

func main() {
	target := os.Args[1]
	dir := os.Args[2]

	file, err := os.Open(target)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	output, err := os.Create("output-" + file.Name())
	if err != nil {
		log.Fatalln(err)
	}
	defer output.Close()

	dirEntrys, err := os.ReadDir(dir)
	if err != nil {
		log.Fatalln(err)
	}
	parseDirs(dicTrie, dirEntrys)

	fReader := bufio.NewReader(file)
    for brk := false; !brk; {
        s, err := fReader.ReadString('\n')
        if err == io.EOF {
            brk = true
        }
        idxs := findWordFromTrie(dicTrie, s)
        if len(idxs) == 0 {
            output.WriteString(s)
            continue
        }
        p := 0
        for _, idx := range idxs {
            if p > idx[0] {
                log.Fatalf("p > idx[0]: %v, %v, %v", s, p, idx)
            }
            output.WriteString(s[p:idx[0]])
            output.WriteString("[[")
            output.WriteString(s[idx[0]:idx[1]])
            output.WriteString("]]")
            p = idx[1]
        }
        output.WriteString(s[p:])
    }
}

func parseDirs(t *Trie, dirEntrys []os.DirEntry) {
	for _, de := range dirEntrys {
		fullname := strings.ToLower(de.Name())
		extension := filepath.Ext(fullname)
		name := strings.TrimSuffix(fullname, extension)
        // log.Printf("filename: %v\n", name)
		t.Insert(name)
	}
}

func findWordFromTrie(t *Trie, s string) [][2]int {
    idx := [][2]int{}
    var i int
    for i = 0; i < len(s); {
        t, n := t.Prefix(s[i:])
        if t.isEnd {
            idx = append(idx, [2]int{i, i + n})
            i += n
        } else {
            i++
        }
    }
    return idx
}

type Trie struct {
    chars [30]*Trie
    str string
    isEnd bool
}

func (t *Trie) Prefix(s string) (*Trie, int) {
    n := 0
    rt := t
    rs := strings.ToLower(s)

    for true {
        if rt == nil {
            log.Fatalf("rt-fatal: t, rt : %v, %v", t, rt)
        }
        // log.Printf("n: %v\n", n)
        if len(rs) == 0 {
            break
        }
        idx := getIdx(rs[0])
        if idx < 0 || rt.chars[idx] == nil {
            break
        }
        rs = rs[1:]
        rt = rt.chars[idx]
        n++
    }
    
    return rt, n
}

func (t *Trie) Search(s string) bool {
    end, n := t.Prefix(s)
    if n != len(s) {
        // log.Printf("search: n, len(s) : %v, %v", n, len(s))
        return false
    }
    return end.isEnd
}

func (t *Trie) Insert(s string) {
    if len(s) == 0 {
        t.isEnd = true
        // log.Println(t.isEnd, t.str)
        return
    }
    idx := getIdx(s[0])
    if idx < 0 {
        return
    }
    if t.chars[idx] == nil {
        t.chars[idx] = &Trie{}
        t.chars[idx].str = t.str + string(s[0])
    }
    t.chars[idx].Insert(s[1:])
}

func getIdx(c byte) int {
    if 'a' <= c && c <= 'z' {
        return int(c - 'a')
    } else if c == '-' {
        return 26
    } else if c == '\'' {
        return 27
    } else if c == ' ' {
        return 28
    } else {
        return -1
    }
}
