package main

func pass4(trim bool, next func(op *operator), emit func(value int), report func(lineno int, error errorkind), rerun func(), fail func(reason failure)) {
	const (
		maxblock   = 10
		maxlabel   = 1000
		elemlength = 1
		liblength  = 1
		linklength = 5
		proclength = 2
		setlength  = 8
		setlimit   = 127
		none       = 0
	)
	type opcode = int
	const ( /* opcode enum */
		add4 = iota
		also4
		and4
		assign4
		blank4
		cobegin4
		constant4
		construct4
		difference4
		divide4
		do4
		else4
		endcode4
		endlib4
		endproc4
		endwhen4
		equal4
		field4
		goto4
		greater4
		in4
		index4
		instance4
		intersection4
		less4
		libproc4
		minus4
		modulo4
		multiply4
		newline4
		not4
		notequal4
		notgreater4
		notless4
		or4
		paramarg4
		paramcall4
		procarg4
		proccall4
		procedure4
		process4
		subtract4
		union4
		valspace4
		value4
		variable4
		wait4
		when4
		addr4
		halt4
		obtain4
		place4
		sense4
		elemassign4
		elemvalue4
		localcase4
		localset4
		localvalue4
		localvar4
		outercall4
		outercase4
		outerparam4
		outerset4
		outervalue4
		outervar4
		setconst4
		singleton4
		stringconst4
	)
	var final bool
	type operators = Set
	var no_arguments, one_argument, two_arguments, three_arguments, four_arguments operators
	var op operator
	var a, b, c, d int
	var lineno int
	var nextop = func() {
		next(&op)
		for op == newline2 {
			next(&lineno)
			next(&op)
		}
		if no_arguments.Contains(op) {
		} else if one_argument.Contains(op) {
			next(&a)
		} else if two_arguments.Contains(op) {
			next(&a)
			next(&b)
		} else if three_arguments.Contains(op) {
			next(&a)
			next(&b)
			next(&c)
		} else if four_arguments.Contains(op) {
			next(&a)
			next(&b)
			next(&c)
			next(&d)
		}
	}
	no_arguments = Construct(add2, and2, difference2, divide2, endcode2, greater2, in2, intersection2, less2, minus2, modulo2, multiply2, not2, notgreater2, notless2, or2, subtract2, union2, addr2, halt2, obtain2, place2, sense2)
	one_argument = Construct(assign2, blank2, constant2, construct2, do2, endif2, endwhen2, equal2, error2, field2, goto2, newline2, notequal2, valspace2, value2, wait2, when2, while2)
	two_arguments = Construct(also2, else2, endlib2, funcval2, paramarg2, parameter2, procarg2, process2, variable2)
	three_arguments = Construct(cobegin2, index2, libproc2, paramcall2, proccall2)
	four_arguments = Construct(endproc2, procedure2)
	type labeltable = [1 + maxlabel]int
	var labels labeltable
	var i int
	var define = func(index, value int) {
		labels[index] = value
	}
	var valueof = func(index int) (val_valueof int) {
		val_valueof = labels[index]
		return
	}
	i = 0
	for i < maxlabel {
		i = i + 1
		labels[i] = none
	}
	const (
		opbase  = 256
		spacing = 2
	)
	var pointer, wordno int
	var codelength int
	var oper = func(op opcode) {
		pointer = wordno
		emit(opbase + spacing*int(op))
		wordno = wordno + 1
	}
	var offset = func(value int) {
		emit(spacing * value)
		wordno = wordno + 1
	}
	var literal = func(value int) {
		emit(value)
		wordno = wordno + 1
	}
	var label = func(index int) {
		emit(spacing * (valueof(index) - pointer))
		wordno = wordno + 1
	}
	var defaddr = func(label int) {
		define(label, wordno)
	}
	var out_again = func() {
		codelength = wordno - 1
		pointer = 1
		wordno = 1
	}
	codelength = 0
	pointer = 1
	wordno = 1
	var errorline int
	var error = func(kind errorkind) {
		if !final && (lineno != errorline) {
			report(lineno, kind)
			errorline = lineno
		}
		nextop()
	}
	errorline = none
	type paramtable = [1 + maxblock]int
	var procs paramtable
	var proclevel int
	var newproc = func(paramlength int) {
		if proclevel == maxblock {
			fail(blocklimit)
		}
		proclevel = proclevel + 1
		procs[proclevel] = paramlength
	}
	var thislevel = func() (val_thislevel int) {
		val_thislevel = proclevel
		return
	}
	var paramlength = func(proclevel int) (val_paramlength int) {
		val_paramlength = procs[proclevel]
		return
	}
	var endprocx = func() {
		proclevel = proclevel - 1
	}
	var initparam = func() {
		proclevel = 0
	}
	initparam()
	type tempattr = struct {
		temp, maxtemp int
	}
	type temptable = [1 + maxblock]tempattr
	var temps temptable
	var templevel int
	var newtemp = func() {
		if templevel == maxblock {
			fail(blocklimit)
		}
		templevel = templevel + 1
		temps[templevel] = tempattr{0, 0}
	}
	var push = func(length int) {
		var t tempattr
		t = temps[templevel]
		t.temp = t.temp + length
		if t.maxtemp < t.temp {
			t.maxtemp = t.temp
		}
		temps[templevel] = t
	}
	var pop = func(length int) {
		temps[templevel].temp = temps[templevel].temp - length
	}
	var endtemp = func(templength *int) {
		*templength = temps[templevel].maxtemp
		templevel = templevel - 1
	}
	var inittemp = func() {
		templevel = 0
	}
	inittemp()
	var again = func() {
		rerun()
		out_again()
		initparam()
		inittemp()
	}
	var in_setrange = func(value int) (val_in_setrange bool) {
		val_in_setrange = (0 <= value) && (value <= setlimit)
		return
	}
	var constlist func(value int)
	var nearby_case = func(steps, displ, value, falselabel int) {
		if steps == 0 {
			oper(localcase4)
		} else if true {
			oper(outercase4)
		}
		offset(displ)
		literal(value)
		label(falselabel)
		nextop()
	}
	var nearby_equal = func(steps, displ, value int) {
		nextop()
		if op == do2 {
			nearby_case(steps, displ, value, a)
		} else if true {
			if steps == 0 {
				oper(localvalue4)
			} else if true {
				oper(outervalue4)
			}
			offset(displ)
			push(1)
			oper(constant4)
			literal(value)
			push(1)
			oper(equal4)
			offset(1)
			pop(1)
		}
	}
	var nearby_elem_const = func(steps, displ, value int) {
		nextop()
		if (op == equal2) && (a == 1) {
			nearby_equal(steps, displ, value)
		} else if true {
			if steps == 0 {
				oper(localvalue4)
			} else if true {
				oper(outervalue4)
			}
			offset(displ)
			push(1)
			if in_setrange(value) && (op == constant2) {
				constlist(value)
			} else if in_setrange(value) && (op == construct2) && (a == 1) {
				oper(singleton4)
				literal(value)
				push(setlength)
				nextop()
			} else if true {
				oper(constant4)
				literal(value)
				push(1)
			}
		}
	}
	var nearby_elem = func(steps, displ int) {
		nextop()
		if op == constant2 {
			nearby_elem_const(steps, displ, a)
		} else if true {
			if steps == 0 {
				oper(localvalue4)
			} else if true {
				oper(outervalue4)
			}
			offset(displ)
			push(1)
		}
	}
	var nearby_set = func(steps, displ int) {
		nextop()
		if steps == 0 {
			oper(localset4)
		} else if true {
			oper(outerset4)
		}
		offset(displ)
		push(setlength)
	}
	var nearby_variable = func(steps, displ int) {
		if (op == value2) && (a == 1) {
			nearby_elem(steps, displ)
		} else if (op == value2) && (a == setlength) {
			nearby_set(steps, displ)
		} else if true {
			if steps == 0 {
				oper(localvar4)
			} else if true {
				oper(outervar4)
			}
			offset(displ)
			push(1)
		}
	}
	constlist = func(value1 int) {
		const (
			maxn = 80
		)
		type table = [1 + maxn]int
		var list table
		var n, i int
		n = 1
		list[1] = value1
		for (op == constant2) && in_setrange(a) && (n < maxn) {
			n = n + 1
			list[n] = a
			nextop()
		}
		if (op == construct2) && (a <= n) {
			i = 0
			for i < n-a {
				i = i + 1
				oper(constant4)
				literal(list[i])
			}
			push(n - a)
			if a == 1 {
				oper(singleton4)
				literal(list[n])
			} else if true {
				oper(setconst4)
				literal(a)
				for i < n {
					i = i + 1
					literal(list[i])
				}
			}
			push(setlength)
			nextop()
		} else if true {
			oper(stringconst4)
			literal(n)
			i = 0
			for i < n {
				i = i + 1
				literal(list[i])
			}
			push(n)
		}
	}
	var singleton = func(value int) {
		oper(singleton4)
		literal(value)
		push(setlength)
		nextop()
	}
	var one = func() {
		if op == do2 {
			nextop()
		} else if true {
			oper(constant4)
			literal(1)
			push(1)
		}
	}
	var smallconst = func(value int) {
		nextop()
		if op == constant2 {
			constlist(value)
		} else if (op == construct2) && (a == 1) {
			singleton(value)
		} else if value == 1 {
			one()
		} else if true {
			oper(constant4)
			literal(value)
			push(1)
		}
	}
	var elemvalue = func() {
		oper(elemvalue4)
		nextop()
	}
	var elemassign = func() {
		oper(elemassign4)
		pop(2)
		nextop()
	}
	var outercall = func(proclabel, arglength int) {
		oper(outercall4)
		label(proclabel)
		push(linklength)
		pop(arglength + linklength)
		nextop()
	}
	var outerparam = func(displ, arglength int) {
		oper(outerparam4)
		offset(displ)
		push(linklength)
		pop(arglength + linklength)
		nextop()
	}
	var goto_ = func(endlabel int) {
		oper(goto4)
		label(endlabel)
		nextop()
	}
	var libproc = func(proclabel, paramlength, templabel int) {
		defaddr(proclabel)
		oper(libproc4)
		offset(paramlength)
		offset(valueof(templabel))
		literal(lineno)
		newproc(paramlength)
		newtemp()
		nextop()
	}
	var endlib = func(templabel, endlabel int) {
		var templength int
		endtemp(&templength)
		endprocx()
		define(templabel, templength)
		oper(endlib4)
		literal(lineno)
		defaddr(endlabel)
		nextop()
	}
	var procedure = func(proclabel, paramlength, varlabel, templabel int) {
		defaddr(proclabel)
		oper(procedure4)
		offset(paramlength)
		offset(valueof(varlabel))
		offset(valueof(templabel))
		literal(lineno)
		newproc(paramlength)
		newtemp()
		nextop()
	}
	var endproc = func(varlabel, varlength, templabel, endlabel int) {
		var templength int
		endtemp(&templength)
		endprocx()
		define(templabel, templength)
		define(varlabel, varlength)
		oper(endproc4)
		defaddr(endlabel)
		nextop()
	}
	var field = func(displ int) {
		nextop()
		for op == field2 {
			displ = displ + a
			nextop()
		}
		if displ != 0 {
			oper(field4)
			offset(displ)
		}
	}
	var index = func(lower, upper, length int) {
		oper(index4)
		literal(lower)
		literal(upper)
		offset(length)
		literal(lineno)
		pop(1)
		nextop()
	}
	var whole_variable = func(level, displ int) {
		var steps int
		steps = thislevel() - level
		nextop()
		for op == field2 {
			displ = displ + a
			nextop()
		}
		if trim && (steps <= 1) {
			nearby_variable(steps, displ)
		} else if true {
			oper(instance4)
			literal(steps)
			oper(variable4)
			offset(displ)
			push(1)
		}
	}
	var variable = func(level, displ int) {
		whole_variable(level, linklength+displ)
	}
	var parameter = func(level, displ int) {
		whole_variable(level, -paramlength(level)+displ)
	}
	var funcval = func(level, length int) {
		whole_variable(level+1, -paramlength(level+1)-length)
	}
	var blank = func(number int) {
		oper(blank4)
		literal(number)
		push(number)
		nextop()
	}
	var construct = func(number int) {
		oper(construct4)
		literal(number)
		literal(lineno)
		pop(number)
		push(setlength)
		nextop()
	}
	var constant = func(value int) {
		if trim && in_setrange(value) {
			smallconst(value)
		} else if true {
			oper(constant4)
			literal(value)
			push(1)
			nextop()
		}
	}
	var value = func(length int) {
		if trim && (length == 1) {
			elemvalue()
		} else if true {
			oper(value4)
			offset(length)
			pop(1)
			push(length)
			nextop()
		}
	}
	var valspace = func(length int) {
		oper(valspace4)
		offset(length)
		push(length)
		nextop()
	}
	var notx = func() {
		oper(not4)
		nextop()
	}
	var multiply = func() {
		oper(multiply4)
		literal(lineno)
		pop(1)
		nextop()
	}
	var divide = func() {
		oper(divide4)
		literal(lineno)
		pop(1)
		nextop()
	}
	var modulo = func() {
		oper(modulo4)
		literal(lineno)
		pop(1)
		nextop()
	}
	var andx = func() {
		oper(and4)
		pop(1)
		nextop()
	}
	var intersection = func() {
		oper(intersection4)
		pop(setlength)
		nextop()
	}
	var minus = func() {
		oper(minus4)
		literal(lineno)
		nextop()
	}
	var add = func() {
		oper(add4)
		literal(lineno)
		pop(1)
		nextop()
	}
	var subtract = func() {
		oper(subtract4)
		literal(lineno)
		pop(1)
		nextop()
	}
	var orx = func() {
		oper(or4)
		pop(1)
		nextop()
	}
	var union = func() {
		oper(union4)
		pop(setlength)
		nextop()
	}
	var difference = func() {
		oper(difference4)
		pop(setlength)
		nextop()
	}
	var equal = func(length int) {
		oper(equal4)
		offset(length)
		pop(2 * length)
		push(1)
		nextop()
	}
	var notequal = func(length int) {
		oper(notequal4)
		offset(length)
		pop(2 * length)
		push(1)
		nextop()
	}
	var less = func() {
		oper(less4)
		pop(1)
		nextop()
	}
	var notless = func() {
		oper(notless4)
		pop(1)
		nextop()
	}
	var greater = func() {
		oper(greater4)
		pop(1)
		nextop()
	}
	var notgreater = func() {
		oper(notgreater4)
		pop(1)
		nextop()
	}
	var inx = func() {
		oper(in4)
		literal(lineno)
		pop(setlength)
		nextop()
	}
	var assign = func(length int) {
		if trim && (length == 1) {
			elemassign()
		} else if true {
			oper(assign4)
			offset(length)
			pop(1 + length)
			nextop()
		}
	}
	var addrx = func() {
		oper(addr4)
		pop(1)
		nextop()
	}
	var haltx = func() {
		oper(halt4)
		literal(lineno)
		nextop()
	}
	var obtainx = func() {
		oper(obtain4)
		pop(2)
		nextop()
	}
	var placex = func() {
		oper(place4)
		pop(2)
		nextop()
	}
	var sensex = func() {
		oper(sense4)
		pop(2)
		nextop()
	}
	var procarg = func(level, proclabel int) {
		oper(instance4)
		literal(thislevel() - level)
		oper(procarg4)
		label(proclabel)
		push(proclength)
		nextop()
	}
	var paramarg = func(level, displ int) {
		oper(instance4)
		literal(thislevel() - level)
		oper(paramarg4)
		offset(-paramlength(level) + displ)
		push(proclength)
		nextop()
	}
	var proccall = func(level, proclabel, arglength int) {
		var steps int
		steps = thislevel() - level
		if trim && (steps == 1) {
			outercall(proclabel, arglength)
		} else if true {
			oper(instance4)
			literal(steps)
			oper(proccall4)
			label(proclabel)
			push(linklength)
			pop(arglength + linklength)
			nextop()
		}
	}
	var paramcall = func(level, displ, arglength int) {
		var steps int
		steps = thislevel() - level
		displ = -paramlength(level) + displ
		if trim && (steps == 1) {
			outerparam(displ, arglength)
		} else if true {
			oper(instance4)
			literal(steps)
			oper(paramcall4)
			offset(displ)
			push(linklength)
			pop(arglength + linklength)
			nextop()
		}
	}
	var dox = func(falselabel int) {
		oper(do4)
		label(falselabel)
		pop(1)
		nextop()
	}
	var elsex = func(truelabel, falselabel int) {
		nextop()
		if op != endif2 {
			oper(else4)
			label(truelabel)
		}
		defaddr(falselabel)
	}
	var endif = func(truelabel int) {
		defaddr(truelabel)
		nextop()
	}
	var whilex = func(truelabel int) {
		defaddr(truelabel)
		nextop()
	}
	var whenx = func(waitlabel int) {
		oper(when4)
		defaddr(waitlabel)
		nextop()
	}
	var wait = func(waitlabel int) {
		oper(wait4)
		label(waitlabel)
		nextop()
	}
	var endwhen = func(truelabel int) {
		defaddr(truelabel)
		oper(endwhen4)
		nextop()
	}
	var process = func(proclabel, templabel int) {
		defaddr(proclabel)
		oper(process4)
		offset(valueof(templabel))
		literal(lineno)
		newtemp()
		nextop()
	}
	var alsox = func(endlabel, templabel int) {
		var templength int
		endtemp(&templength)
		define(templabel, templength)
		oper(also4)
		label(endlabel)
		nextop()
	}
	var cobeginx = func(beginlabel, endlabel, number int) {
		var procconst, proclabel, i int
		defaddr(beginlabel)
		oper(cobegin4)
		literal(number)
		literal(lineno)
		i = 0
		for i < number {
			next(&procconst)
			literal(procconst)
			next(&proclabel)
			label(proclabel)
			i = i + 1
		}
		defaddr(endlabel)
		nextop()
	}
	var endcode = func() {
		oper(endcode4)
		literal(lineno - 1)
	}
	var assemble = func(last_scan bool) {
		var more bool
		final = last_scan
		more = true
		offset(codelength)
		nextop()
		for more {
			if op <= construct2 {
				if op == add2 {
					add()
				} else if op == also2 {
					alsox(a, b)
				} else if op == and2 {
					andx()
				} else if op == assign2 {
					assign(a)
				} else if op == blank2 {
					blank(a)
				} else if op == cobegin2 {
					cobeginx(a, b, c)
				} else if op == constant2 {
					constant(a)
				} else if op == construct2 {
					construct(a)
				}
			} else if op <= endproc2 {
				if op == difference2 {
					difference()
				} else if op == divide2 {
					divide()
				} else if op == do2 {
					dox(a)
				} else if op == else2 {
					elsex(a, b)
				} else if op == endcode2 {
					endcode()
					more = false
				} else if op == endif2 {
					endif(a)
				} else if op == endlib2 {
					endlib(a, b)
				} else if op == endproc2 {
					endproc(a, b, c, d)
				}
			} else if op <= in2 {
				if op == endwhen2 {
					endwhen(a)
				} else if op == equal2 {
					equal(a)
				} else if op == error2 {
					error(errorkind(a))
				} else if op == field2 {
					field(a)
				} else if op == funcval2 {
					funcval(a, b)
				} else if op == goto2 {
					goto_(a)
				} else if op == greater2 {
					greater()
				} else if op == in2 {
					inx()
				}
			} else if op <= not2 {
				if op == index2 {
					index(a, b, c)
				} else if op == intersection2 {
					intersection()
				} else if op == less2 {
					less()
				} else if op == libproc2 {
					libproc(a, b, c)
				} else if op == minus2 {
					minus()
				} else if op == modulo2 {
					modulo()
				} else if op == multiply2 {
					multiply()
				} else if op == not2 {
					notx()
				}
			} else if op <= procarg2 {
				if op == notequal2 {
					notequal(a)
				} else if op == notgreater2 {
					notgreater()
				} else if op == notless2 {
					notless()
				} else if op == or2 {
					orx()
				} else if op == paramarg2 {
					paramarg(a, b)
				} else if op == paramcall2 {
					paramcall(a, b, c)
				} else if op == parameter2 {
					parameter(a, b)
				} else if op == procarg2 {
					procarg(a, b)
				}
			} else if op <= variable2 {
				if op == proccall2 {
					proccall(a, b, c)
				} else if op == procedure2 {
					procedure(a, b, c, d)
				} else if op == process2 {
					process(a, b)
				} else if op == subtract2 {
					subtract()
				} else if op == union2 {
					union()
				} else if op == valspace2 {
					valspace(a)
				} else if op == value2 {
					value(a)
				} else if op == variable2 {
					variable(a, b)
				}
			} else if op <= sense2 {
				if op == wait2 {
					wait(a)
				} else if op == when2 {
					whenx(a)
				} else if op == while2 {
					whilex(a)
				} else if op == addr2 {
					addrx()
				} else if op == halt2 {
					haltx()
				} else if op == obtain2 {
					obtainx()
				} else if op == place2 {
					placex()
				} else if op == sense2 {
					sensex()
				}
			}
		}
	}
	assemble(false)
	again()
	assemble(true)
}
