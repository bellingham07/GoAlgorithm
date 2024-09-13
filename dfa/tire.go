package dfa

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Result struct {
	CharStart  int // 1-based, left closed right closed
	CharEnd    int
	ByteStart  int // zero-based, left closed right open
	ByteEnd    int
	HitWord    string
	MatchedStr string
}

type TrieNode struct {
	End      bool
	Children map[rune]*TrieNode
}

type Trie struct {
	Root       *TrieNode
	IgnoreCase bool
	Noises     map[rune]struct{}
}

func (t *Trie) Insert(word string) {
	if word == "" {
		return
	}

	node := t.Root
	for _, r := range word {
		if t.IgnoreCase {
			r = unicode.ToLower(r)
		}
		if child, ok := node.Children[r]; ok {
			node = child
		} else {
			child := newNode()
			node.Children[r] = child
			node = child
		}
	}

	node.End = true
}

func (t *Trie) Check(text string) bool {
	if len(t.Match(text)) > 0 {
		fmt.Println("res:", t.Match(text))
		return true
	}
	return false
}

func (t *Trie) Filter(text string, replace ...string) string {
	results := t.Match(text)
	if len(results) == 0 {
		return text
	}

	var (
		sb     = strings.Builder{}
		l      = len(text)
		merged = []Result{
			{
				CharStart: results[0].CharStart,
				CharEnd:   results[0].CharEnd,
				ByteStart: results[0].ByteStart,
				ByteEnd:   results[0].ByteEnd,
			},
		}
		rep = "*"
	)
	if len(replace) > 0 {
		rep = replace[0]
	}

	for cur, j := 0, 0; cur < len(results)-1; cur++ {
		next := cur + 1
		if results[cur].ByteEnd >= results[next].ByteStart && results[cur].ByteEnd <= results[next].ByteEnd {
			merged[j].CharEnd = results[next].CharEnd
			merged[j].ByteEnd = results[next].ByteEnd
		} else {
			merged = append(merged, Result{
				CharStart: results[next].CharStart,
				CharEnd:   results[next].CharEnd,
				ByteStart: results[next].ByteStart,
				ByteEnd:   results[next].ByteEnd,
			})
			j++
		}
	}
	var pos int
	for _, res := range merged {
		sb.WriteString(text[pos:res.ByteStart])
		sb.WriteString(strings.Repeat(rep, res.CharEnd-res.CharStart+1))
		pos = res.ByteEnd
	}
	if pos < l {
		sb.WriteString(text[pos:])
	}

	return sb.String()
}

func (t *Trie) Search(text string) []string {
	results := t.Match(text)

	// unique
	m := make(map[string]struct{}, len(results))
	words := make([]string, 0, len(results))
	for _, res := range results {
		if _, ok := m[res.HitWord]; !ok {
			m[res.HitWord] = struct{}{}
			words = append(words, res.HitWord)
		}
	}

	return words
}

func (t *Trie) Match(text string) (results []Result) {
	var (
		node                 = t.Root            // 根节点
		sb                   = strings.Builder{} // 记录铭感词
		start                = 0                 // 临时指针，指向第一次命中的位置
		pos                  = 0                 // 所遍历字符所在的位置
		cstart               = 0                 // utf-8 char start pos
		ncmatched            = 0                 // matched char counter，已命中多少个字符
		l                    = len(text)
		firstMatchedRuneSize int
	)

	for pos < l {
		// 截取一段字符串
		r, size := utf8.DecodeRuneInString(text[pos:])
		// 是否忽略大小写
		if t.IgnoreCase {
			r = unicode.ToLower(r)
		}

		// 判断当前字符是否在敏感词树中
		if child, ok := node.Children[r]; ok {
			// 添加到builder中
			sb.WriteRune(r)
			// 记录已匹配字符数量
			ncmatched++
			// 记录是否匹配过，方便后续判断干扰词
			if firstMatchedRuneSize == 0 {
				firstMatchedRuneSize = size
			}
			// 如果是匹配到铭感词最后一个字符，构建返回结果
			if child.End {
				result := Result{
					CharStart: cstart + 1,         // 开始位置
					CharEnd:   cstart + ncmatched, // 结束位置
					ByteStart: start,              // 字符索引开始位置
					ByteEnd:   pos + size,         // 索引结束位置
					HitWord:   sb.String(),        // 出现的敏感词
				}
				// 含有敏感词的字符串
				result.MatchedStr = text[result.ByteStart:result.ByteEnd]
				results = append(results, result)
			}
			// 遍历下一个字符
			pos += size
			// 获取下一个节点
			node = child
		} else {
			// 当前字符不在tire树上
			// 查看之前是否有匹配上的字符，查看是否是干扰词
			if firstMatchedRuneSize > 0 {
				// 判断当前字符是否为干扰词
				if _, ok = t.Noises[r]; ok {
					pos += size
					ncmatched++
					continue // 是干扰词，结束此次循环，进行下一次循环，即开始遍历下一个字符
				}
			}
			// 当前字符不在tire树上，并且不是干扰词，回到根节点，重置变量
			node = t.Root
			if firstMatchedRuneSize > 0 {
				start += firstMatchedRuneSize // 如果之前已经命中了，那么start将停留在第一次命中的位置
				firstMatchedRuneSize = 0
				ncmatched = 0
			} else {
				start += size // 没有命中，start一直跟着size在加，所以只需要加size
			}
			pos = start
			cstart++
			sb.Reset()
		}
	}

	return results
}

func (t *Trie) GetAllWords() (words []string) {
	stack := []*TrieNode{t.Root}
	prefixStack := []string{""}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		prefix := prefixStack[len(prefixStack)-1]
		prefixStack = prefixStack[:len(prefixStack)-1]

		if node.End {
			words = append(words, prefix)
		}

		for r, child := range node.Children {
			stack = append(stack, child)
			prefixStack = append(prefixStack, prefix+string(r))
		}
	}

	return
}

func (t *Trie) Delete(word string) {
	var (
		node  = t.Root
		stack []*TrieNode
	)

	for _, r := range word {
		if _, ok := node.Children[r]; !ok {
			return
		}
		stack = append(stack, node)
		node = node.Children[r]
	}
	node.End = false
	if len(node.Children) == 0 {
		for len(stack) > 0 {
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			delete(node.Children, []rune(word)[len(stack)])
			if len(node.Children) > 0 || node.End {
				break
			}
		}
	}
}

func newNode() *TrieNode {
	node := new(TrieNode)
	node.Children = make(map[rune]*TrieNode)
	return node
}

func NewTrie() *Trie {
	trie := new(Trie)
	trie.Root = newNode()
	trie.Noises = make(map[rune]struct{})
	return trie
}
