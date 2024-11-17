package main

type line = [80 + 1]rune
type failure = int

const ( /* failure enum */
	charlimit = iota
	namelimit
)

type errorkind = int

const ( /* errorkind enum */
	ambiguous3 = iota
	declaration3
	kind3
	padding3
	range3
	syntax3
	trap3
	undeclared3
)

func alva(next func(value *rune), real_emit func(value int), rewind func(), report func(lineno int, error errorkind), fail func(reason failure)) {
	const (
		maxchar = 3500
		maxname = 400
		maxword = 503
	)
	type symbol = int
	const ( /* symbol enum */
		undefined_ = iota
		address_
		register_
		constant_
		char_
		instr_monadic_
		instr_dyadic_
		instr_register_
		instr_branch_
		instr_call_
		instr_return_
		instr_repeat_
		instr_push_
		instr_pop_
		addr_
		array_
		const_
		do_
		instr_
		pad_
		reg_
		st_
		text_
		trap_
		word_
		plus_
		minus_
		equal_
		comma_
		colon_
		lparanth_
		rparanth_
		lbracket_
		rbracket_
		eof_
	)
	type symbolset = Set
	type chartable = [maxchar + 1]rune
	type wordattr struct {
		wordlength, lastchar int
	}
	type wordtable = [maxword + 1]wordattr
	var heap chartable
	var top int
	var table wordtable
	var size int
	var key = func(word line, length int) (val_key int) {
		const (
			span = 26
		)
		var hash, i int
		hash = 1
		i = 0
		for i < length {
			i = i + 1
			hash = hash * (int(word[i])%span + 1) % maxword
		}
		val_key = hash + 1
		return
	}
	var insert = func(word line, length, index int) {
		var m, n int
		top = top + length
		if top > maxchar {
			fail(charlimit)
		}
		m = length
		n = top - m
		for m > 0 {
			heap[m+n] = word[m]
			m = m - 1
		}
		table[index] = wordattr{length, top}
		size = size + 1
		if size == maxname {
			fail(namelimit)
		}
	}
	var found = func(word line, length, index int) (val_found bool) {
		var same bool
		var m, n int
		if table[index].wordlength != length {
			same = false
		} else if true {
			same = true
			m = length
			n = table[index].lastchar - m
			for same && (m > 0) {
				same = word[m] == heap[m+n]
				m = m - 1
			}
		}
		val_found = same
		return
	}
	var convert = func(word line, length int) (val_convert int) {
		var i int
		var more bool
		i = key(word, length)
		more = true
		for more {
			if table[i].wordlength == 0 {
				insert(word, length, i)
				val_convert = i
				more = false
			} else if found(word, length, i) {
				val_convert = i
				more = false
			} else if true {
				i = i%maxword + 1
			}
		}
		return
	}
	top = 0
	size = maxword
	for size > 0 {
		table[size] = wordattr{0, 0}
		size = size - 1
	}
	const (
		newline    = rune(10)
		endmedium  = rune(25)
		space      = ' '
		quote      = '"'
		apostrophy = rune(39)
		maxint     = 32767
	)
	type nameattr struct {
		namesym   symbol
		namevalue int
	}
	type nametable = [maxword + 1]nameattr
	var names nametable
	const (
		noname = 0
	)
	var sym symbol
	var value, name int
	var lineno int
	type charset = Set
	var alphanum, digits, graphic, letters, small_letters, octals, special, stringchar charset
	var ch rune
	var in_string bool
	var error = func(e errorkind) {
		report(lineno, e)
	}
	var decimal = func() {
		var i int
		value = 0
		for digits.ContainsChar(ch) {
			i = int(ch) - int('0')
			if value <= (maxint-i)/10 {
				value = 10*value + i
				next(&ch)
			} else if true {
				error(range3)
				for digits.ContainsChar(ch) {
					next(&ch)
				}
			}
		}
	}
	var octal = func() {
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
				error(range3)
				for octals.ContainsChar(ch) {
					next(&ch)
				}
			}
		} else if true {
			error(syntax3)
		}
	}
	var nextstring = func() {
		if !(stringchar.ContainsChar(ch)) {
			in_string = false
			error(syntax3)
		}
		if ch != '[' {
			value = int(ch)
			next(&ch)
		} else if true {
			next(&ch)
			decimal()
			if ch == ']' {
				next(&ch)
			} else if true {
				error(syntax3)
			}
		}
		if ch == apostrophy {
			in_string = false
			next(&ch)
		}
	}
	var nextsym = func() {
		var word line
		var i int
		if in_string {
			nextstring()
		} else if true {
			for {
				if ch == space {
					next(&ch)
				} else if ch == newline {
					lineno = lineno + 1
					next(&ch)
				} else if ch == '$' {
					for int(ch) >= 32 {
						next(&ch)
					}
				} else {
					break
				}
			}
			name = noname
			if letters.ContainsChar(ch) {
				i = 0
				for alphanum.ContainsChar(ch) {
					if small_letters.ContainsChar(ch) {
						ch = rune(int(ch) - 32)
					}
					i = i + 1
					word[i] = ch
					next(&ch)
				}
				name = convert(word, i)
				sym = names[name].namesym
				value = names[name].namevalue
			} else if ch == '(' {
				sym = lparanth_
				next(&ch)
			} else if ch == ')' {
				sym = rparanth_
				next(&ch)
			} else if ch == ',' {
				sym = comma_
				next(&ch)
			} else if ch == '[' {
				sym = lbracket_
				next(&ch)
			} else if ch == ']' {
				sym = rbracket_
				next(&ch)
			} else if ch == ':' {
				sym = colon_
				next(&ch)
			} else if digits.ContainsChar(ch) {
				sym = constant_
				decimal()
			} else if ch == '+' {
				sym = plus_
				next(&ch)
			} else if ch == '-' {
				sym = minus_
				next(&ch)
			} else if ch == '=' {
				sym = equal_
				next(&ch)
			} else if ch == '#' {
				sym = constant_
				octal()
			} else if ch == apostrophy {
				sym = char_
				in_string = true
				next(&ch)
				nextstring()
			} else if ch == endmedium {
				sym = eof_
			} else if true {
				error(syntax3)
				next(&ch)
			}
		}
	}
	var firstsym = func() {
		rewind()
		lineno = 1
		next(&ch)
		nextsym()
	}
	var define = func(name int, sym symbol, value int) {
		names[name] = nameattr{sym, value}
	}
	name = maxword
	for name > 0 {
		names[name] = nameattr{undefined_, 0}
		name = name - 1
	}
	small_letters = ConstructFromString("abcdefghijklmnopqrstuvwxyz")
	letters = ConstructFromString("ABCDEFGHIJKLMNOPQRSTUVWXYZ").Union(small_letters)
	digits = ConstructFromString("0123456789")
	alphanum = letters.Union(digits).AddChar('_')
	special = ConstructFromString("'!\"#$%()*+,-./:;<=>?@[]_")
	graphic = alphanum.Union(special).AddChar(space)
	octals = ConstructFromString("01234567")
	stringchar = graphic.RemoveChar(apostrophy)
	var sp_name int
	var enter = func(word string, length int, sym symbol, value int) {
		tmp := line{}
		for i, ch := range word {
			tmp[i+1] = rune(ch)
	}
		define(convert(tmp, length), sym, value)
	}
	var enter2 = func(word string, length int, sym symbol, value int) {
		enter(word, length, sym, value)
		enter(word+"_BYTE", length+5, sym, value+32768)
		}
	enter("ADDR", 4, addr_, 0)
	enter("ARRAY", 5, array_, 0)
	enter("CONST", 5, const_, 0)
	enter("DO", 2, do_, 0)
	enter("INSTR", 5, instr_, 0)
	enter("PAD", 3, pad_, 0)
	enter("REG", 3, reg_, 0)
	enter("ST", 2, st_, 0)
	enter("TEXT", 4, text_, 0)
	enter("TRAP", 4, trap_, 0)
	enter("WORD", 4, word_, 0)
	enter("EXTENDSIGN", 10, instr_monadic_, 3520)
	enter("SWAPBYTES", 9, instr_monadic_, 192)
	enter2("ADDCARRY", 8, instr_monadic_, 2880)
	enter2("CLEAR", 5, instr_monadic_, 2560)
	enter2("DECREMENT", 9, instr_monadic_, 2752)
	enter2("DOUBLE", 6, instr_monadic_, 3264)
	enter2("HALVE", 5, instr_monadic_, 3200)
	enter2("INCREMENT", 9, instr_monadic_, 2688)
	enter2("NEGATE", 6, instr_monadic_, 2816)
	enter2("NOT", 3, instr_monadic_, 2624)
	enter2("ROTATELEFT", 10, instr_monadic_, 3136)
	enter2("ROTATERIGHT", 11, instr_monadic_, 3072)
	enter2("SUBTRACTCARRY", 13, instr_monadic_, 2944)
	enter2("TEST", 4, instr_monadic_, 3008)
	enter("ADD", 3, instr_dyadic_, 24576)
	enter("SUBTRACT", 8, instr_dyadic_, 57344)
	enter2("ANDNOT", 6, instr_dyadic_, 16384)
	enter2("COMPARE", 7, instr_dyadic_, 8192)
	enter2("MOVE", 4, instr_dyadic_, 4096)
	enter2("OR", 2, instr_dyadic_, 20480)
	enter2("TESTBIT", 7, instr_dyadic_, 12288)
	enter("DIVIDE", 6, instr_register_, 29184)
	enter("DOUBLESHIFT", 11, instr_register_, 30208)
	enter("MULTIPLY", 8, instr_register_, 28672)
	enter("SHIFT", 5, instr_register_, 29696)
	enter("BRANCH", 6, instr_branch_, 256)
	enter("IFCARRY", 7, instr_branch_, 34560)
	enter("IFEQUAL", 7, instr_branch_, 768)
	enter("IFGREATER", 9, instr_branch_, 1536)
	enter("IFHIGHER", 8, instr_branch_, 33280)
	enter("IFLESS", 6, instr_branch_, 1280)
	enter("IFLOWER", 7, instr_branch_, 34560)
	enter("IFNOTCARRY", 10, instr_branch_, 34304)
	enter("IFNOTEQUAL", 10, instr_branch_, 512)
	enter("IFNOTGREATER", 12, instr_branch_, 1792)
	enter("IFNOTHIGHER", 11, instr_branch_, 33536)
	enter("IFNOTLESS", 9, instr_branch_, 1024)
	enter("IFNOTLOWER", 10, instr_branch_, 34304)
	enter("IFNOTOVERFLOW", 13, instr_branch_, 33792)
	enter("IFOVERFLOW", 10, instr_branch_, 34048)
	enter("CALL", 4, instr_call_, 0)
	enter("RETURN", 6, instr_return_, 0)
	enter("REPEAT", 6, instr_repeat_, 0)
	enter("PUSH", 4, instr_push_, 0)
	enter("POP", 3, instr_pop_, 0)
	sp_name = convert(line{' ', 'S', 'P'}, 2)
	var restart_symbols symbolset
	var first_pass bool
	var sp_reg int
	var ptr int
	var emit = func(w int) {
		ptr = ptr + 2
		if !first_pass {
			real_emit(w)
		}
	}
	var name_error = func() {
		if sym == undefined_ {
			error(undeclared3)
			nextsym()
		} else if name != noname {
			error(kind3)
			nextsym()
		} else if true {
			error(syntax3)
		}
	}
	var checksym = func(s symbol) {
		if sym == s {
			nextsym()
		} else if true {
			error(syntax3)
		}
	}
	var new_name = func() (val_new_name int) {
		if (sym == undefined_) || !first_pass {
			val_new_name = name
			nextsym()
		} else if true {
			val_new_name = noname
			if name != noname {
				error(ambiguous3)
				nextsym()
			} else if true {
				error(syntax3)
			}
		}
		return
	}
	var try_address_name = func(address *int) (val_try_address_name bool) {
		if (sym == address_) || (first_pass && (sym == undefined_)) {
			*address = value
			val_try_address_name = true
			nextsym()
		} else if true {
			val_try_address_name = false
		}
		return
	}
	var address_name = func() (val_address_name int) {
		if !try_address_name(&val_address_name) {
			name_error()
			val_address_name = 0
		}
		return
	}
	var try_constant_symbol = func(v *int) (val_try_constant_symbol bool) {
		if (sym == constant_) || (sym == char_) {
			*v = value
			val_try_constant_symbol = true
			nextsym()
		} else if true {
			val_try_constant_symbol = false
		}
		return
	}
	var constant_symbol = func() (val_constant_symbol int) {
		if !try_constant_symbol(&val_constant_symbol) {
			name_error()
			val_constant_symbol = 0
		}
		return
	}
	var constant_declaration = func() {
		var n, v int
		nextsym()
		n = new_name()
		checksym(equal_)
		v = constant_symbol()
		if first_pass {
			define(n, constant_, v)
		}
	}
	var register_declaration = func() {
		var n, v int
		nextsym()
		n = new_name()
		checksym(lparanth_)
		v = constant_symbol()
		if (v < 0) || (v > 7) {
			error(declaration3)
			v = 0
		}
		checksym(rparanth_)
		if first_pass {
			define(n, register_, v)
			if n == sp_name {
				sp_reg = v
			}
		}
	}
	var text_declaration = func() {
		nextsym()
		define(new_name(), address_, ptr)
		checksym(equal_)
		if sym != char_ {
			error(syntax3)
		}
		for sym == char_ {
			emit(value)
			nextsym()
		}
	}
	var word_declaration = func() {
		nextsym()
		define(new_name(), address_, ptr)
		emit(0)
	}
	var array_declaration = func() {
		var n int
		nextsym()
		define(new_name(), address_, ptr)
		checksym(lbracket_)
		n = constant_symbol()
		if n < 1 {
			error(declaration3)
		}
		checksym(rbracket_)
		for n > 0 {
			emit(0)
			n = n - 1
		}
	}
	var address_list_declaration = func() {
		nextsym()
		define(new_name(), address_, ptr)
		checksym(equal_)
		emit(address_name())
		for sym == comma_ {
			nextsym()
			emit(address_name())
		}
	}
	var pad_sentence = func() {
		var n int
		nextsym()
		n = constant_symbol()
		if (n < ptr) || (n%2 != 0) {
			error(padding3)
		}
		for ptr < n {
			emit(0)
		}
	}
	var trap_sentence = func() {
		var n int
		nextsym()
		n = constant_symbol()
		if (n < ptr) || (n%4 != 0) {
			error(trap3)
		}
		for ptr < n {
			emit(0)
		}
		checksym(colon_)
		emit(address_name())
		checksym(comma_)
		emit(constant_symbol())
	}
	type operand struct {
		modebits int
		has_imm  bool
		imm      int
	}
	var register = func() (val_register int) {
		if sym == register_ {
			val_register = value
			nextsym()
		} else if true {
			name_error()
			val_register = 0
		}
		return
	}
	var try_constant_operand = func(v *int) (val_try_constant_operand bool) {
		val_try_constant_operand = try_address_name(v) || try_constant_symbol(v)
		return
	}
	var apply_sign = func(v *int, s symbol) {
		if s == minus_ {
			if *v == 32768 {
				error(range3)
			} else if true {
				*v = -*v
			}
		}
	}
	var composite_address = func(op *operand, r int) {
		var imm int
		var sign symbol
		sign = sym
		nextsym()
		if (sign != minus_) && (sign != plus_) {
			error(syntax3)
		}
		if try_constant_operand(&imm) {
			apply_sign(&imm, sign)
			*op = operand{48 + r, true, imm}
		} else if sign == minus_ {
			*op = operand{32 + r, false, 0}
		} else if true {
			*op = operand{16 + r, false, 0}
		}
	}
	var indirect_address = func(op *operand) {
		nextsym()
		checksym(lbracket_)
		composite_address(op, register())
		op.modebits = op.modebits + 8
		checksym(rbracket_)
	}
	var register_address = func(op *operand, r int) {
		nextsym()
		if (sym == plus_) || (sym == minus_) {
			composite_address(op, r)
		} else if true {
			*op = operand{8 + r, false, 0}
		}
	}
	var direct_address = func(op *operand) {
		var imm int
		if sym == register_ {
			register_address(op, value)
		} else if try_constant_operand(&imm) {
			*op = operand{31, true, imm}
		} else if true {
			name_error()
			*op = operand{0, false, 0}
		}
	}
	var location_symbol = func(op *operand) {
		nextsym()
		checksym(lbracket_)
		if sym == st_ {
			indirect_address(op)
		} else if true {
			direct_address(op)
		}
		checksym(rbracket_)
	}
	var variable_symbol = func(op *operand) {
		if sym == register_ {
			*op = operand{value, false, 0}
			nextsym()
		} else if sym == st_ {
			location_symbol(op)
		} else if true {
			name_error()
			*op = operand{0, false, 0}
		}
	}
	var value_operand = func(op *operand) {
		var v int
		var sign symbol
		if (sym == register_) || (sym == st_) {
			variable_symbol(op)
		} else if true {
			if (sym == plus_) || (sym == minus_) {
				sign = sym
				nextsym()
			} else if true {
				sign = plus_
			}
			if !try_constant_operand(&v) {
				name_error()
			}
			apply_sign(&v, sign)
			*op = operand{23, true, v}
		}
	}
	var emit_imm = func(op *operand) {
		if op.has_imm {
			emit(op.imm)
		}
	}
	var instr_monadic = func(instr int) {
		var op operand
		nextsym()
		checksym(lparanth_)
		variable_symbol(&op)
		checksym(rparanth_)
		emit(instr + op.modebits)
		emit_imm(&op)
	}
	var instr_dyadic = func(instr int) {
		var op1, op2 operand
		nextsym()
		checksym(lparanth_)
		value_operand(&op1)
		checksym(comma_)
		variable_symbol(&op2)
		checksym(rparanth_)
		emit(instr + 64*op1.modebits + op2.modebits)
		emit_imm(&op1)
		emit_imm(&op2)
	}
	var instr_register = func(instr int) {
		var op operand
		nextsym()
		checksym(lparanth_)
		value_operand(&op)
		checksym(comma_)
		emit(instr + 64*register() + op.modebits)
		checksym(rparanth_)
		emit_imm(&op)
	}
	var instr_branch = func(instr int) {
		var a, ofs int
		nextsym()
		checksym(lparanth_)
		a = address_name()
		if first_pass {
			ofs = 0
		} else if true {
			ofs = (a - ptr - 2) / 2
			if (ofs < -128) || (ofs > 127) {
				error(range3)
				ofs = 0
			}
			ofs = (ofs + 256) % 256
		}
		checksym(rparanth_)
		emit(instr + ofs)
	}
	var instr_call = func() {
		nextsym()
		checksym(lparanth_)
		emit(2527)
		emit(address_name())
		checksym(rparanth_)
	}
	var instr_return = func() {
		nextsym()
		emit(135)
	}
	var instr_repeat = func() {
		var a, r, ofs int
		nextsym()
		checksym(lparanth_)
		a = address_name()
		if first_pass {
			ofs = 0
		} else if true {
			ofs = (ptr + 2 - a) / 2
			if (ofs < 0) || (ofs >= 64) {
				error(range3)
				ofs = 0
			}
		}
		checksym(comma_)
		r = register()
		checksym(rparanth_)
		emit(32256 + r*64 + ofs)
	}
	var instr_push = func() {
		var op operand
		nextsym()
		if sp_reg < 0 {
			error(undeclared3)
		}
		checksym(lparanth_)
		value_operand(&op)
		checksym(rparanth_)
		emit(4128 + 64*op.modebits + sp_reg)
		emit_imm(&op)
	}
	var instr_pop = func() {
		var n int
		nextsym()
		if sp_reg < 0 {
			error(undeclared3)
		}
		checksym(lparanth_)
		n = constant_symbol()
		checksym(rparanth_)
		if n == 1 {
			emit(3024 + sp_reg)
		} else if n > 1 {
			emit(26048 + sp_reg)
			emit(2 * n)
		} else if true {
			error(range3)
		}
	}
	var encoded_instruction = func() {
		nextsym()
		checksym(lparanth_)
		emit(constant_symbol())
		checksym(rparanth_)
	}
	var instruction_sentence = func() {
		nextsym()
		define(new_name(), address_, ptr)
		checksym(colon_)
		for {
			if sym == instr_monadic_ {
				instr_monadic(value)
			} else if sym == instr_dyadic_ {
				instr_dyadic(value)
			} else if sym == instr_register_ {
				instr_register(value)
			} else if sym == instr_branch_ {
				instr_branch(value)
			} else if sym == instr_call_ {
				instr_call()
			} else if sym == instr_return_ {
				instr_return()
			} else if sym == instr_repeat_ {
				instr_repeat()
			} else if sym == instr_push_ {
				instr_push()
			} else if sym == instr_pop_ {
				instr_pop()
			} else if sym == instr_ {
				encoded_instruction()
			} else if !(restart_symbols.Contains(sym)) {
				error(syntax3)
				for !(restart_symbols.Contains(sym)) {
					nextsym()
				}
			} else {
				break
			}
		}
	}
	var program = func() {
		firstsym()
		ptr = 0
		for {
			if sym == const_ {
				constant_declaration()
			} else if sym == reg_ {
				register_declaration()
			} else if sym == text_ {
				text_declaration()
			} else if sym == word_ {
				word_declaration()
			} else if sym == array_ {
				array_declaration()
			} else if sym == addr_ {
				address_list_declaration()
			} else if sym == do_ {
				instruction_sentence()
			} else if sym == pad_ {
				pad_sentence()
			} else if sym == trap_ {
				trap_sentence()
			} else if sym != eof_ {
				error(syntax3)
				nextsym()
				for !(restart_symbols.Contains(sym)) {
					nextsym()
				}
			} else {
				break
			}
		}
	}
	restart_symbols = Construct(const_, reg_, text_, word_, array_, addr_, do_, pad_, trap_, instr_monadic_, instr_dyadic_, instr_branch_, instr_call_, instr_return_, instr_repeat_, instr_push_, instr_pop_, instr_)
	sp_reg = -1
	first_pass = true
	program()
	first_pass = false
	program()
}
