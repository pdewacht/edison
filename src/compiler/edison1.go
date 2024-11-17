package main

type symbol = int

const ( /* symbol enum */
	also1 = iota
	and1
	array1
	asterisk1
	becomes1
	begin1
	cobegin1
	colon1
	comma1
	const1
	div1
	do1
	else1
	end1
	endtext1
	enum1
	equal1
	error1
	graphic1
	greater1
	if1
	in1
	lbracket1
	less1
	lib1
	lparanth1
	minus1
	mod1
	module1
	name1
	newline1
	not1
	notequal1
	notgreater1
	notless1
	numeral1
	or1
	period1
	plus1
	pre1
	post1
	proc1
	rbracket1
	record1
	rparanth1
	semicolon1
	set1
	skip1
	val1
	var1
	when1
	while1
)

func pass1(next func(value *rune), emit func(value symbol), fail func(reason failure)) {
	const (
		maxchar = 6500
		maxname = 750
		maxword = 809
	)
	var emit2 = func(sym symbol, arg int) {
		emit(sym)
		emit(symbol(arg))
	}
	var error = func(kind errorkind) {
		emit2(error1, int(kind))
	}
	const (
		bool1         = 1
		char1         = 2
		false1        = 3
		int1          = 4
		true1         = 5
		univname1     = 6
		univtype1     = 7
		addr1         = 8
		halt1         = 9
		obtain1       = 10
		place1        = 11
		sense1        = 12
		last_standard = 20
		none          = maxword
	)
	var table = map[string]int{}
	var name int
	var convert = func(word string, length int, value *int) {
		val, exists := table[word]
		if !exists {
			val = -name
			name++
			table[word] = val
		}
		*value = val
	}
	table["ALSO"] = int(also1)
	table["AND"] = int(and1)
	table["ARRAY"] = int(array1)
	table["BEGIN"] = int(begin1)
	table["COBEGIN"] = int(cobegin1)
	table["CONST"] = int(const1)
	table["DIV"] = int(div1)
	table["DO"] = int(do1)
	table["ELSE"] = int(else1)
	table["END"] = int(end1)
	table["ENUM"] = int(enum1)
	table["IF"] = int(if1)
	table["IN"] = int(in1)
	table["LIB"] = int(lib1)
	table["MOD"] = int(mod1)
	table["MODULE"] = int(module1)
	table["NOT"] = int(not1)
	table["OR"] = int(or1)
	table["PRE"] = int(pre1)
	table["POST"] = int(post1)
	table["PROC"] = int(proc1)
	table["RECORD"] = int(record1)
	table["SET"] = int(set1)
	table["SKIP"] = int(skip1)
	table["VAL"] = int(val1)
	table["VAR"] = int(var1)
	table["WHEN"] = int(when1)
	table["WHILE"] = int(while1)
	table["BOOL"] = -bool1
	table["CHAR"] = -char1
	table["FALSE"] = -false1
	table["INT"] = -int1
	table["TRUE"] = -true1
	table["ADDR"] = -addr1
	table["HALT"] = -halt1
	table["OBTAIN"] = -obtain1
	table["PLACE"] = -place1
	table["SENSE"] = -sense1
	name = last_standard + 1
	const (
		newline    = rune(10)
		endmedium  = rune(25)
		space      = ' '
		quote      = '"'
		apostrophy = rune(39)
		maxint     = 32767
	)
	type charset = Set
	var alphanum, comment, composite, capital_letters, digits, graphic, letters, octals, parantheses, punctuation, single, small_letters, special, stringchar charset
	var lineno int
	var ch rune
	var nextsym = func() {
		var word []rune
		var value, i int
		for {
			if ch == space {
				next(&ch)
			} else if ch == newline {
				lineno = lineno + 1
				emit2(newline1, lineno)
				next(&ch)
			} else if ch == quote {
				next(&ch)
				for comment.ContainsChar(ch) {
					if ch == newline {
						lineno = lineno + 1
					}
					next(&ch)
				}
				if ch == quote {
					next(&ch)
				} else if true {
					error(syntax3)
				}
			} else {
				break
			}
		}
		if letters.ContainsChar(ch) {
			i = 0
			for alphanum.ContainsChar(ch) {
				if small_letters.ContainsChar(ch) {
					ch = rune(int(ch) - 32)
				}
				i = i + 1
				word = append(word, ch)
				next(&ch)
			}
			convert(string(word), i, &value)
			if value < 0 {
				emit2(name1, -value)
			} else if true {
				emit(symbol(value))
			}
		} else if digits.ContainsChar(ch) {
			value = 0
			for digits.ContainsChar(ch) {
				i = int(ch) - int('0')
				if value <= (maxint-i)/10 {
					value = 10*value + i
					next(&ch)
				} else if true {
					error(numeral3)
					for digits.ContainsChar(ch) {
						next(&ch)
					}
				}
			}
			emit2(numeral1, value)
		} else if punctuation.ContainsChar(ch) {
			if ch == ';' {
				emit(semicolon1)
			} else if ch == ',' {
				emit(comma1)
			} else if ch == '.' {
				emit(period1)
			}
			next(&ch)
		} else if parantheses.ContainsChar(ch) {
			if ch == '(' {
				emit(lparanth1)
			} else if ch == ')' {
				emit(rparanth1)
			} else if ch == '[' {
				emit(lbracket1)
			} else if ch == ']' {
				emit(rbracket1)
			}
			next(&ch)
		} else if composite.ContainsChar(ch) {
			if ch == ':' {
				next(&ch)
				if ch == '=' {
					emit(becomes1)
					next(&ch)
				} else if true {
					emit(colon1)
				}
			} else if ch == '>' {
				next(&ch)
				if ch == '=' {
					emit(notless1)
					next(&ch)
				} else if true {
					emit(greater1)
				}
			} else if ch == '<' {
				next(&ch)
				if ch == '>' {
					emit(notequal1)
					next(&ch)
				} else if ch == '=' {
					emit(notgreater1)
					next(&ch)
				} else if true {
					emit(less1)
				}
			}
		} else if single.ContainsChar(ch) {
			if ch == '+' {
				emit(plus1)
			} else if ch == '-' {
				emit(minus1)
			} else if ch == '*' {
				emit(asterisk1)
			} else if ch == '=' {
				emit(equal1)
			}
			next(&ch)
		} else if ch == apostrophy {
			next(&ch)
			if stringchar.ContainsChar(ch) {
				emit2(graphic1, int(ch))
				next(&ch)
				for stringchar.ContainsChar(ch) {
					emit(comma1)
					emit2(graphic1, int(ch))
					next(&ch)
				}
				if ch == apostrophy {
					next(&ch)
				} else if true {
					error(syntax3)
				}
			} else if true {
				if ch == apostrophy {
					next(&ch)
				}
				error(syntax3)
			}
		} else if ch == '#' {
			value = 0
			next(&ch)
			if octals.ContainsChar(ch) {
				for (octals.ContainsChar(ch)) && (value <= 4095) {
					value = 8*value + (int(ch) - int('0'))
					next(&ch)
				}
				if (octals.ContainsChar(ch)) && (value <= 8191) {
					value = 32768 + 8*(value-4096) + (int(ch) - int('0'))
					next(&ch)
				}
				if octals.ContainsChar(ch) {
					error(numeral3)
					for octals.ContainsChar(ch) {
						next(&ch)
					}
				}
			} else if true {
				error(syntax3)
			}
			emit2(numeral1, value)
		} else if ch == endmedium {
		} else if true {
			error(syntax3)
			next(&ch)
		}
	}
	capital_letters = ConstructFromString("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	small_letters = ConstructFromString("abcdefghijklmnopqrstuvwxyz")
	letters = capital_letters.Union(small_letters)
	digits = ConstructFromString("0123456789")
	alphanum = letters.Union(digits).AddChar('_')
	special = ConstructFromString("'!\"#$%()*+,-./:;<=>?@[]_")
	graphic = alphanum.Union(special).AddChar(space)
	comment = graphic.RemoveChar(quote).AddChar(newline)
	composite = ConstructFromString(":<>")
	octals = ConstructFromString("01234567")
	parantheses = ConstructFromString("()[]")
	punctuation = ConstructFromString(";,.")
	single = ConstructFromString("+-*=")
	stringchar = graphic.RemoveChar(apostrophy)
	lineno = 1
	emit2(newline1, 1)
	next(&ch)
	for ch != endmedium {
		nextsym()
	}
	emit(endtext1)
}
