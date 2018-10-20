package vidopre

import (
	"fmt"
	"log"
	"strings"
	"time"
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
		case r == '\n':
			l.emit(itemDateLine)
			return lexPostContent
		}
	}
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
	Date    time.Time
	Year    string
	Month   string
	Day     string
}

func (p *PostInfo) parsePostDate(dateTxt string) {
	//dateTxt something similar to "venerdì, 15 ottobre 2010"
	// Month field is in Italian
	//fmt.Println(dateTxt)

	arr := strings.Split(dateTxt, ",")
	dds := arr[len(arr)-1]
	dds = strings.Trim(dds, " ")
	//dds = strings.Replace(dds, "\n", "", -1)
	items := strings.Split(dds, " ")
	//fmt.Println(dds)
	//fmt.Println(items)
	if len(items) < 3 {
		log.Fatalln("Invalid date format for parsing (4 items expected)", len(items), items, dateTxt)
	}
	mm := ""
	mese := items[len(items)-2]
	//fmt.Println("mese is ", mese)
	switch strings.ToLower(mese) {
	case "gennaio":
		mm = "01"
		break
	case "febbraio":
		mm = "02"
		break
	case "marzo":
		mm = "03"
		break
	case "aprile":
		mm = "04"
		break
	case "maggio":
		mm = "05"
		break
	case "giugno":
		mm = "06"
		break
	case "luglio":
		mm = "07"
		break
	case "agosto":
		mm = "08"
		break
	case "settembre":
		mm = "09"
		break
	case "ottobre":
		mm = "10"
		break
	case "novembre":
		mm = "11"
		break
	case "dicembre":
		mm = "12"
		break
	default:
		log.Fatal("Month not recognized", mese)
	}
	day := items[len(items)-3]
	if len(day) < 2 {
		day = "0" + day
	}
	yyyy := items[len(items)-1]
	yyyy = strings.TrimSuffix(yyyy, "\r\n") // Attenzione al formato windows
	//fmt.Printf("Year is %s type is %T\n", yyyy, yyyy)
	//fmt.Println("Day is ", reflect.TypeOf(day))
	//fmt.Println("Month is ", reflect.TypeOf(mm))
	dateFormatted := fmt.Sprintf("%s-%s-%s", yyyy, mm, day)
	//fmt.Printf("Formatted date is %s\n", dateFormatted)
	const layout = "2006-01-02"
	t, err := time.Parse(layout, dateFormatted)
	if err != nil {
		log.Fatalln("Invalid date: ", err)
	}
	//fmt.Println(t)

	p.Year = yyyy
	p.Month = mm
	p.Day = day
	p.Date = t
}

func GetSplittedPosts(str string) []*PostInfo {
	pi := &PostInfo{}
	res := []*PostInfo{}
	l := lexCtor("Text lex", str)
	for {
		item := l.nextItem()
		//fmt.Printf("**** EMIT *** type %v, val:%v\n", typeName[item.typ], item.val)
		switch item.typ {
		case itemH2Title:
			if len(pi.Content) > 0 {
				log.Println("Post created on ", pi.DateTxt, pi.Day, pi.Month, pi.Year)
				res = append(res, pi)
				pi = &PostInfo{}
			}
			break
		case itemDateLine:
			pi.Content += item.val
			pi.DateTxt = strings.Replace(item.val, pdata, "", 1)
			pi.DateTxt = strings.TrimSuffix(pi.DateTxt, "\r\n")
			pi.parsePostDate(pi.DateTxt)
			break
		case itemText:
			pi.Content += item.val
			break
		}
		if l.state == nil {
			break
		}
	}
	if len(pi.Content) > 0 {
		log.Println("Post createad on ", pi.DateTxt, pi.Day, pi.Month, pi.Year)
		res = append(res, pi)
	}
	return res
}
