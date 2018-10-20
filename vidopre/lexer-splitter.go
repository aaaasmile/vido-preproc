package vidopre

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type itemType int

const (
	h2Title            = "h2."
	pdata              = "p(data)."
	eof                = -1
	itemError itemType = iota
	itemEOF
	itemH2Title
	itemDateLine
	itemText
)

var (
	typeName = map[itemType]string{
		itemH2Title:  "itemH2Title",
		itemDateLine: "itemDateLine",
		itemText:     "itemText",
	}
)

type item struct {
	typ itemType

	val string
}

type lexer struct {
	name  string
	input string
	start int
	pos   int
	width int
	state stateFn
	items chan item
}

type stateFn func(*lexer) stateFn

func (l *lexer) run() {
	for state := lexStart; state != nil; {
		state(l)
	}
	close(l.items)
}

func (l *lexer) emit(t itemType) {
	l.items <- item{t, l.input[l.start:l.pos]}
	l.start = l.pos
}

func (l *lexer) next() rune {
	if l.pos >= len(l.input) {
		l.width = 0
		return eof
	}

	ru, s := utf8.DecodeRuneInString(l.input[l.pos:])
	l.width = s
	l.pos += l.width
	return ru
}

func (l *lexer) ignore() {
	l.start = l.pos
}

func (l *lexer) backup() {
	l.pos -= l.width
}

func (l *lexer) errorf(format string, args ...interface{}) stateFn {
	l.items <- item{
		itemError,
		fmt.Sprintf(format, args...),
	}
	return nil
}

func (l *lexer) peek() rune {
	ru := l.next()
	l.backup()
	return ru
}

func (l *lexer) nextItem() item {
	for {
		select {
		case item := <-l.items:
			return item
		default:
			l.state = l.state(l)
		}
	}
	panic("not reached")
}

// Lex States

func lexPostContent(l *lexer) stateFn {
	for {
		if strings.HasPrefix(l.input[l.pos:], h2Title) {
			l.pos -= len(h2Title)
			if l.pos > l.start {
				l.emit(itemText)
			}
			return lexStart
		}
		if l.next() == eof {
			break
		}
	}
	if l.pos > l.start {
		l.emit(itemText)
	}
	return nil
}

func lexInsideDate(l *lexer) stateFn {
	for {
		switch r := l.next(); {
		case r == eof:
			return l.errorf("Date in post is invalid")
		case r == '\n':
			l.emit(itemDateLine)
			return lexPostContent
		}
	}
	if l.pos > l.start {
		l.emit(itemText)
	}
	return nil
}

func lexInsidePost(l *lexer) stateFn {
	for {
		if strings.HasPrefix(l.input[l.pos:], pdata) {
			if l.pos > l.start {
				l.emit(itemText)
			}
			return lexInsideDate
		}
		if l.next() == eof {
			break
		}
	}
	if l.pos > l.start {
		l.emit(itemText)
	}
	return nil
}

func lexStart(l *lexer) stateFn {
	for {
		if strings.HasPrefix(l.input[l.pos:], h2Title) {
			l.pos -= len(h2Title)
			l.emit(itemH2Title)
			return lexInsidePost
		}
		if l.next() == eof {
			break
		}
	}
	if l.pos > l.start {
		l.emit(itemText)
	}
	return nil
}

func lexCtor(name, input string) *lexer {
	l := &lexer{
		name:  name,
		input: input,
		state: lexStart,
		items: make(chan item, 2),
	}
	return l
}

type PostInfo struct {
	Content string
	DateTxt string
}

func GetSplittedPosts(str string) []*PostInfo {
	pi := &PostInfo{}
	res := []*PostInfo{}
	l := lexCtor("Text lex", str)
	for {
		item := l.nextItem()
		fmt.Printf("**** EMIT *** type %v, val:%v\n", typeName[item.typ], item.val)
		switch item.typ {
		case itemH2Title:
			if len(pi.Content) > 0 {
				fmt.Println("*** Post is", pi)
				fmt.Println("Append to res")
				res = append(res, pi)
				pi = &PostInfo{}
			}
			break
		case itemDateLine:
			pi.Content += item.val
			pi.DateTxt = item.val
			break
		case itemText:
			pi.Content += item.val
			break
		}
		if l.state == nil {
			break
		}
	}
	return res
}
