package main

type operator = int

const ( /* operator enum */
	add2 = iota
	also2
	and2
	assign2
	blank2
	cobegin2
	constant2
	construct2
	difference2
	divide2
	do2
	else2
	endcode2
	endif2
	endlib2
	endproc2
	endwhen2
	equal2
	error2
	field2
	funcval2
	goto2
	greater2
	in2
	index2
	intersection2
	less2
	libproc2
	minus2
	modulo2
	multiply2
	newline2
	not2
	notequal2
	notgreater2
	notless2
	or2
	paramarg2
	paramcall2
	parameter2
	procarg2
	proccall2
	procedure2
	process2
	subtract2
	union2
	valspace2
	value2
	variable2
	wait2
	when2
	while2
	addr2
	halt2
	obtain2
	place2
	sense2
)

func pass3(next func(value *symbol), emit func(value operator), fail func(reason failure)) {
	const (
		maxlabel      = 1000
		maxname       = 750
		maxprocess    = 20
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
		elemlength    = 1
		liblength     = 1
		proclength    = 2
		setlength     = 8
		setlimit      = 127
		none          = 0
	)
	type namekind = int
	const ( /* namekind enum */
		univkind = iota
		constant
		univtype
		elemtype
		recordtype
		arraytype
		settype
		field
		variable
		valparam
		varparam
		procparam
		procedur
		standard
	)
	type nameattr = struct {
		kind namekind
		link int
		/* constattr */
		consttype, constvalue int
		/* typeattr */
		length, subtype1, subtype2, bound1, bound2 int
		/* varattr */
		varlevel, vardispl, vartype int
		/* procattr */
		proclevel, procaddr, proctype, param int
	}
	type nametable = [1 + maxname]nameattr
	type processattr = struct {
		procconst, proclabel int
	}
	type processtable = [1 + maxprocess]processattr
	type processset = Set
	var emit2 = func(a operator, b int) {
		emit(a)
		emit(operator(b))
	}
	var emit3 = func(a operator, b, c int) {
		emit2(a, b)
		emit(operator(c))
	}
	var emit4 = func(a operator, b, c, d int) {
		emit3(a, b, c)
		emit(operator(d))
	}
	var emit5 = func(a operator, b, c, d, e int) {
		emit4(a, b, c, d)
		emit(operator(e))
	}
	type symbols = Set
	var addsym, constsym, declsym, equalitysym, exprsym, hiddensym, initdeclsym, literalsym, multsym, ordersym, pairsym, procsym, relationsym, selectsym, signsym, statsym, typesym symbols
	addsym = Construct(minus1, or1, plus1)
	constsym = Construct(graphic1, name1, numeral1)
	declsym = Construct(array1, const1, enum1, lib1, module1, post1, pre1, proc1, record1, set1, var1)
	equalitysym = Construct(equal1, notequal1)
	exprsym = Construct(graphic1, lparanth1, minus1, not1, numeral1, plus1, val1)
	hiddensym = Construct(error1, newline1)
	initdeclsym = Construct(array1, const1, enum1, record1, set1)
	literalsym = Construct(graphic1, numeral1)
	multsym = Construct(and1, asterisk1, div1, mod1)
	ordersym = Construct(greater1, less1, notgreater1, notless1)
	pairsym = Construct(graphic1, name1, numeral1)
	procsym = Construct(lib1, post1, pre1, proc1)
	relationsym = Construct(equal1, greater1, in1, less1, notequal1, notgreater1, notless1)
	selectsym = Construct(colon1, lbracket1, period1)
	signsym = Construct(minus1, plus1)
	statsym = Construct(cobegin1, if1, skip1, when1, while1)
	typesym = Construct(array1, enum1, record1, set1)
	var sym symbol
	var x int
	var nextsym = func() {
		next(&sym)
		for hiddensym.Contains(sym) {
			next(&x)
			if sym == error1 {
				emit2(error2, x)
			} else if sym == newline1 {
				emit2(newline2, x)
			}
			next(&sym)
		}
		if pairsym.Contains(sym) {
			next(&x)
		}
	}
	nextsym()
	type namekinds = Set
	var names nametable
	var typekinds, varkinds, prockinds namekinds
	var standard_names = func() {
		names[int1] = nameattr{kind: elemtype, length: elemlength}
		names[bool1] = names[int1]
		names[char1] = names[int1]
		names[false1] = nameattr{kind: constant, consttype: bool1, constvalue: 0}
		names[true1] = nameattr{kind: constant, consttype: bool1, constvalue: 1}
		names[univname1] = nameattr{kind: univkind}
		names[addr1] = nameattr{kind: standard, proclevel: 0, procaddr: none, proctype: int1, param: none}
		names[halt1] = nameattr{kind: standard, proclevel: 0, procaddr: none, proctype: univtype1, param: none}
		names[obtain1] = names[halt1]
		names[place1] = names[halt1]
		names[sense1] = nameattr{kind: standard, proclevel: 0, procaddr: none, proctype: bool1, param: none}
	}
	var isfunction = func(x int) (val_isfunction bool) {
		if prockinds.Contains(names[x].kind) {
			val_isfunction = names[x].proctype != univtype1
		} else if true {
			val_isfunction = false
		}
		return
	}
	var i int
	i = 0
	for i < maxname {
		i = i + 1
		names[i] = nameattr{kind: univkind}
	}
	typekinds = Construct(univtype, elemtype, recordtype, arraytype, settype)
	varkinds = Construct(variable, valparam, varparam)
	prockinds = Construct(procparam, procedur, standard)
	var labelno int
	var newlabel = func(value *int) {
		if labelno == maxlabel {
			fail(labellimit)
		}
		labelno = labelno + 1
		*value = labelno
	}
	labelno = 0
	var error = func(kind errorkind) {
		emit2(error2, int(kind))
	}
	var syntax = func(succ symbols) {
		error(syntax3)
		for !(succ.Contains(sym)) {
			nextsym()
		}
	}
	var check = func(succ symbols) {
		if !(succ.Contains(sym)) {
			syntax(succ)
		}
	}
	var checksym = func(s symbol, succ symbols) {
		if sym == s {
			nextsym()
		} else if true {
			syntax(succ)
		}
	}
	var kinderror1 = func(name int) {
		if name != univname1 {
			error(type3)
		}
	}
	var kinderror2 = func(name int, typ *int) {
		kinderror1(name)
		*typ = univtype1
	}
	var typerror1 = func() {
		error(type3)
	}
	var typerror2 = func(typ *int) {
		if *typ != univtype1 {
			typerror1()
			*typ = univtype1
		}
	}
	var checkelem = func(typ *int) {
		if names[*typ].kind != elemtype {
			typerror2(typ)
		}
	}
	var checktype = func(typ1 *int, typ2 int) {
		if *typ1 != typ2 {
			if typ2 == univtype1 {
				*typ1 = univtype1
			} else if true {
				typerror2(typ1)
			}
		}
	}
	var variable_list func(kind namekind, level, displ int, first, length *int, succ symbols)
	var procedure_heading func(postx bool, level int, name, paramlength *int, succ symbols)
	var declaration func(level, displ int, varlength *int, succ symbols)
	var expression func(typ *int, succ symbols)
	var procedure_call func(succ symbols)
	var statement_list func(succ symbols)
	var control_symbol = func(typ, value *int, succ symbols) {
		nextsym()
		checksym(lparanth1, Construct(numeral1, rparanth1).Union(succ))
		if sym == numeral1 {
			*typ = char1
			*value = x
			nextsym()
		} else if true {
			*typ = univtype1
			*value = 0
			syntax(Singleton(rparanth1).Union(succ))
		}
		checksym(rparanth1, succ)
	}
	var constant_symbol = func(typ, value *int, succ symbols) {
		var c nameattr
		*typ = univtype1
		*value = 0
		if sym == numeral1 {
			*typ = int1
			*value = x
			nextsym()
		} else if sym == graphic1 {
			*typ = char1
			*value = x
			nextsym()
		} else if sym == name1 {
			if x == char1 {
				control_symbol(typ, value, succ)
			} else if names[x].kind == constant {
				c = names[x]
				*typ = c.consttype
				*value = c.constvalue
				nextsym()
			} else if true {
				kinderror1(x)
				nextsym()
			}
		} else if true {
			syntax(succ)
		}
		check(succ)
	}
	var constant_declaration = func(succ symbols) {
		var name, typ, value int
		if sym == name1 {
			name = x
			nextsym()
			checksym(equal1, constsym.Union(succ))
			constant_symbol(&typ, &value, succ)
			names[name] = nameattr{kind: constant, consttype: typ, constvalue: value}
		} else if true {
			syntax(succ)
		}
		check(succ)
	}
	var constant_declaration_list = func(succ symbols) {
		var enddecl symbols
		nextsym()
		enddecl = Singleton(semicolon1).Union(succ)
		constant_declaration(enddecl)
		for sym == semicolon1 {
			nextsym()
			constant_declaration(enddecl)
		}
	}
	var enumeration_symbol = func(typ, value int, succ symbols) {
		if sym == name1 {
			names[x] = nameattr{kind: constant, consttype: typ, constvalue: value}
			nextsym()
		} else if true {
			syntax(succ)
		}
		check(succ)
	}
	var enumeration_symbol_list = func(typ int, succ symbols) {
		var endsym symbols
		var value int
		endsym = Singleton(comma1).Union(succ)
		value = 0
		enumeration_symbol(typ, value, endsym)
		for sym == comma1 {
			nextsym()
			value = value + 1
			enumeration_symbol(typ, value, endsym)
		}
		check(succ)
	}
	var enumeration_type = func(succ symbols) {
		var typ int
		nextsym()
		if sym == name1 {
			typ = x
			nextsym()
			checksym(lparanth1, Construct(name1, rparanth1).Union(succ))
			enumeration_symbol_list(typ, Singleton(rparanth1).Union(succ))
			checksym(rparanth1, succ)
			names[typ] = nameattr{kind: elemtype, length: elemlength}
		} else if true {
			syntax(succ)
		}
	}
	var record_type = func(succ symbols) {
		var typ, first, length int
		nextsym()
		if sym == name1 {
			typ = x
			nextsym()
			checksym(lparanth1, Construct(name1, rparanth1).Union(succ))
			variable_list(field, 0, 0, &first, &length, Singleton(rparanth1).Union(succ))
			checksym(rparanth1, succ)
			names[typ] = nameattr{kind: recordtype, length: length, subtype1: first}
		} else if true {
			syntax(succ)
		}
	}
	var range_symbol = func(typ, lower, upper *int, succ symbols) {
		var typ2 int
		constant_symbol(typ, lower, Singleton(colon1).Union(constsym).Union(succ))
		checksym(colon1, constsym.Union(succ))
		constant_symbol(&typ2, upper, succ)
		checktype(typ, typ2)
		if *lower > *upper {
			error(range3)
			*lower = *upper
		}
	}
	var type_name = func(typ *int, succ symbols) {
		if sym == name1 {
			if typekinds.Contains(names[x].kind) {
				*typ = x
			} else if true {
				kinderror2(x, typ)
			}
			nextsym()
		} else if true {
			syntax(succ)
			*typ = univtype1
		}
	}
	var array_type = func(succ symbols) {
		var typ, rangetype, lower, upper, elemtype int
		nextsym()
		if sym == name1 {
			typ = x
			nextsym()
			checksym(lbracket1, constsym.Union(Construct(rbracket1, lparanth1, name1, rparanth1)).Union(succ))
			range_symbol(&rangetype, &lower, &upper, Construct(rbracket1, lparanth1, name1, rparanth1).Union(succ))
			checksym(rbracket1, Construct(lparanth1, name1, rparanth1).Union(succ))
			checksym(lparanth1, Construct(name1, rparanth1).Union(succ))
			type_name(&elemtype, Singleton(rparanth1).Union(succ))
			checksym(rparanth1, succ)
			names[typ] = nameattr{kind: arraytype, length: (upper - lower + 1) * names[elemtype].length, subtype1: rangetype, subtype2: elemtype, bound1: lower, bound2: upper}
		} else if true {
			syntax(succ)
		}
	}
	var set_type = func(succ symbols) {
		var typ, basetype int
		nextsym()
		if sym == name1 {
			typ = x
			nextsym()
			checksym(lparanth1, Construct(name1, rparanth1).Union(succ))
			type_name(&basetype, Singleton(rparanth1).Union(succ))
			checkelem(&basetype)
			checksym(rparanth1, succ)
			names[typ] = nameattr{kind: settype, length: setlength, subtype1: basetype}
		} else if true {
			syntax(succ)
		}
	}
	var type_declaration = func(succ symbols) {
		if sym == enum1 {
			enumeration_type(succ)
		} else if sym == record1 {
			record_type(succ)
		} else if sym == array1 {
			array_type(succ)
		} else if sym == set1 {
			set_type(succ)
		}
	}
	var variable_group = func(kind namekind, level, addr int, first, last, size *int, succ symbols) {
		var typ, varlength, nextvar, i int
		*first = none
		*last = none
		*size = 0
		if sym == name1 {
			*first = x
			*last = x
			nextsym()
			for sym == comma1 {
				nextsym()
				if sym == name1 {
					names[*last].link = x
					*last = x
					nextsym()
				} else if true {
					syntax(Construct(comma1, colon1).Union(succ))
				}
			}
			names[*last].link = none
			checksym(colon1, succ)
			type_name(&typ, succ)
			if kind == varparam {
				varlength = elemlength
			} else if true {
				varlength = names[typ].length
			}
			i = *first
			for i != none {
				nextvar = names[i].link
				names[i] = nameattr{kind: kind, link: nextvar, varlevel: level, vardispl: addr + *size, vartype: typ}
				*size = *size + varlength
				i = nextvar
			}
			check(succ)
		} else if true {
			syntax(succ)
		}
	}
	variable_list = func(kind namekind, level, displ int, first, length *int, succ symbols) {
		var last, first2, last2, length2 int
		var endgroup symbols
		endgroup = Singleton(semicolon1).Union(succ)
		variable_group(kind, level, displ, first, &last, length, endgroup)
		for sym == semicolon1 {
			nextsym()
			variable_group(kind, level, displ+*length, &first2, &last2, &length2, endgroup)
			if length2 > 0 {
				if *length == 0 {
					*first = first2
				} else if true {
					names[last].link = first2
				}
				last = last2
				*length = *length + length2
			}
		}
		check(succ)
	}
	var variable_declaration_list = func(level, displ int, length *int, succ symbols) {
		var first int
		nextsym()
		variable_list(variable, level, displ, &first, length, succ)
	}
	var parameter_group = func(level, displ int, first, last, length *int, succ symbols) {
		var name, paramlength int
		var varkind namekind
		if sym == proc1 {
			procedure_heading(false, level, &name, &paramlength, succ)
			if name != univname1 {
				names[name].kind = procparam
				names[name].procaddr = displ
				*first = name
				*last = name
				*length = proclength
			} else if true {
				*first = none
				*last = none
				*length = 0
			}
		} else if true {
			if sym == var1 {
				varkind = varparam
				nextsym()
			} else if true {
				varkind = valparam
			}
			variable_group(varkind, level, displ, first, last, length, succ)
		}
		check(succ)
	}
	var parameter_list = func(level int, first, length *int, succ symbols) {
		var last, first2, last2, length2 int
		var endgroup symbols
		endgroup = Singleton(semicolon1).Union(succ)
		parameter_group(level, 0, first, &last, length, endgroup)
		for sym == semicolon1 {
			nextsym()
			parameter_group(level, *length, &first2, &last2, &length2, endgroup)
			if length2 > 0 {
				if *length == 0 {
					*first = first2
				} else if true {
					names[last].link = first2
				}
				last = last2
				*length = *length + length2
			}
		}
		check(succ)
	}
	procedure_heading = func(postx bool, level int, name, paramlength *int, succ symbols) {
		var proclabel, firstparam, typ int
		checksym(proc1, Construct(name1, lparanth1, colon1).Union(succ))
		if sym == name1 {
			*name = x
			nextsym()
			if postx && (*name != univname1) {
				proclabel = names[*name].procaddr
			} else if true {
				newlabel(&proclabel)
			}
			if sym == lparanth1 {
				nextsym()
				parameter_list(level+1, &firstparam, paramlength, Construct(rparanth1, colon1).Union(succ))
				checksym(rparanth1, Singleton(colon1).Union(succ))
			} else if true {
				firstparam = none
				*paramlength = 0
			}
			if sym == colon1 {
				nextsym()
				type_name(&typ, succ)
			} else if true {
				typ = univtype1
			}
			names[*name] = nameattr{kind: procedur, proclevel: level, procaddr: proclabel, proctype: typ, param: firstparam}
		} else if true {
			*name = univname1
			*paramlength = 0
			syntax(succ)
		}
		check(succ)
	}
	var procedure_body = func(level int, varlength *int, succ symbols) {
		var sublength int
		*varlength = 0
		for declsym.Contains(sym) {
			declaration(level, *varlength, &sublength, declsym.Add(begin1).Add(end1).Union(statsym).Union(succ))
			*varlength = *varlength + sublength
		}
		checksym(begin1, statsym.Add(end1).Union(succ))
		statement_list(Singleton(end1).Union(succ))
		checksym(end1, succ)
	}
	var complete_procedure = func(postx bool, level int, succ symbols) {
		var name, proclabel, paramlength, endlabel, varlabel, templabel, varlength int
		procedure_heading(postx, level, &name, &paramlength, declsym.Add(begin1).Union(statsym).Union(succ))
		if name != univname1 {
			proclabel = names[name].procaddr
		} else if true {
			newlabel(&proclabel)
		}
		newlabel(&endlabel)
		newlabel(&varlabel)
		newlabel(&templabel)
		if level > 0 {
			emit2(goto2, endlabel)
		}
		emit5(procedure2, proclabel, paramlength, varlabel, templabel)
		procedure_body(level+1, &varlength, succ)
		emit5(endproc2, varlabel, varlength, templabel, endlabel)
	}
	var preprocedure = func(level int, succ symbols) {
		var name, paramlength int
		nextsym()
		procedure_heading(false, level, &name, &paramlength, succ)
	}
	var postprocedure = func(level int, succ symbols) {
		nextsym()
		complete_procedure(true, level, succ)
	}
	var library_procedure = func(level int, succ symbols) {
		var name, paramlength, proclabel, endlabel, templabel, exprtype int
		nextsym()
		procedure_heading(false, level, &name, &paramlength, Construct(lbracket1, rbracket1).Union(exprsym).Union(succ))
		if name != univname1 {
			proclabel = names[name].procaddr
		} else if true {
			newlabel(&proclabel)
		}
		newlabel(&endlabel)
		newlabel(&templabel)
		emit2(goto2, endlabel)
		emit4(libproc2, proclabel, paramlength, templabel)
		checksym(lbracket1, Singleton(rbracket1).Union(exprsym).Union(succ))
		expression(&exprtype, Singleton(rbracket1).Union(succ))
		checksym(rbracket1, succ)
		emit3(endlib2, templabel, endlabel)
	}
	var procedure_declaration = func(level int, succ symbols) {
		if sym == proc1 {
			complete_procedure(false, level, succ)
		} else if sym == pre1 {
			preprocedure(level, succ)
		} else if sym == post1 {
			postprocedure(level, succ)
		} else if sym == lib1 {
			library_procedure(level, succ)
		}
	}
	var module_declaration = func(level, displ int, varlength *int, succ symbols) {
		var sublength int
		nextsym()
		*varlength = 0
		for (Singleton(asterisk1).Union(declsym)).Contains(sym) {
			if sym == asterisk1 {
				nextsym()
			}
			declaration(level, displ+*varlength, &sublength, declsym.Union(Construct(asterisk1, begin1, end1)).Union(statsym).Union(succ))
			*varlength = *varlength + sublength
		}
		checksym(begin1, statsym.Add(end1).Union(succ))
		statement_list(Singleton(end1).Union(succ))
		checksym(end1, succ)
	}
	declaration = func(level, displ int, varlength *int, succ symbols) {
		*varlength = 0
		if sym == const1 {
			constant_declaration_list(succ)
		} else if typesym.Contains(sym) {
			type_declaration(succ)
		} else if sym == var1 {
			variable_declaration_list(level, displ, varlength, succ)
		} else if procsym.Contains(sym) {
			procedure_declaration(level, succ)
		} else if sym == module1 {
			module_declaration(level, displ, varlength, succ)
		} else if true {
			syntax(succ)
		}
		check(succ)
	}
	var function_variable = func(typ *int, succ symbols) {
		var p nameattr
		nextsym()
		if sym == name1 {
			if isfunction(x) {
				p = names[x]
				*typ = p.proctype
				emit3(funcval2, p.proclevel, names[*typ].length)
			} else if true {
				kinderror2(x, typ)
			}
			nextsym()
		} else if true {
			*typ = univtype1
			syntax(succ)
		}
	}
	var field_selector = func(typ *int, succ symbols) {
		var t nameattr
		var v nameattr
		var i int
		nextsym()
		if sym == name1 {
			t = names[*typ]
			if t.kind == recordtype {
				i = t.subtype1
				for (i != none) && (i != x) {
					i = names[i].link
				}
				if i == x {
					v = names[x]
					*typ = v.vartype
					emit2(field2, v.vardispl)
				} else if true {
					kinderror2(x, typ)
				}
			} else if true {
				typerror2(typ)
			}
			nextsym()
			check(succ)
		} else if true {
			*typ = univtype1
			syntax(succ)
		}
	}
	var indexed_selector = func(typ *int, succ symbols) {
		var t nameattr
		var exprtype int
		nextsym()
		t = names[*typ]
		expression(&exprtype, Singleton(rbracket1).Union(succ))
		if t.kind == arraytype {
			if exprtype == t.subtype1 {
				*typ = t.subtype2
				emit4(index2, t.bound1, t.bound2, names[*typ].length)
			} else if true {
				typerror2(typ)
			}
		} else if true {
			typerror2(typ)
		}
		checksym(rbracket1, succ)
		check(succ)
	}
	var type_transfer = func(typ *int, succ symbols) {
		var typ2 int
		nextsym()
		type_name(&typ2, succ)
		if names[*typ].length == names[typ2].length {
			*typ = typ2
		} else if true {
			typerror2(typ)
		}
		check(succ)
	}
	var variable_symbol = func(typ *int, succ symbols) {
		var v nameattr
		var endvar symbols
		endvar = selectsym.Union(succ)
		if sym == name1 {
			if varkinds.Contains(names[x].kind) {
				v = names[x]
				*typ = v.vartype
				if v.kind == variable {
					emit3(variable2, v.varlevel, v.vardispl)
				} else if v.kind == valparam {
					emit3(parameter2, v.varlevel, v.vardispl)
				} else if v.kind == varparam {
					emit3(parameter2, v.varlevel, v.vardispl)
					emit2(value2, elemlength)
				}
			} else if true {
				kinderror2(x, typ)
			}
			nextsym()
		} else if sym == val1 {
			function_variable(typ, endvar)
		} else if true {
			*typ = univtype1
			syntax(endvar)
		}
		for {
			if sym == period1 {
				field_selector(typ, endvar)
			} else if sym == lbracket1 {
				indexed_selector(typ, endvar)
			} else if sym == colon1 {
				type_transfer(typ, endvar)
			} else {
				break
			}
		}
		check(succ)
	}
	var constant_factor = func(typ *int, succ symbols) {
		var value int
		constant_symbol(typ, &value, succ)
		emit2(constant2, value)
	}
	var variable_factor = func(typ *int, succ symbols) {
		variable_symbol(typ, succ)
		emit2(value2, names[*typ].length)
	}
	var elementary_expression = func(typ int, succ symbols) {
		var exprtype int
		expression(&exprtype, succ)
		checkelem(&exprtype)
	}
	var field_expression = func(name *int, succ symbols) {
		var typ int
		var f nameattr
		expression(&typ, succ)
		if *name != none {
			f = names[*name]
			checktype(&typ, f.vartype)
			*name = f.link
		} else if true {
			error(constructor3)
		}
	}
	var field_expression_list = func(typ int, succ symbols) {
		var name int
		var endexpr symbols
		endexpr = Singleton(comma1).Union(exprsym).Union(succ)
		name = names[typ].subtype1
		field_expression(&name, endexpr)
		for sym == comma1 {
			nextsym()
			field_expression(&name, endexpr)
		}
		if name != none {
			error(constructor3)
		}
	}
	var element_expression_list = func(typ int, succ symbols) {
		var t nameattr
		var exprtype, no, max int
		var endexpr symbols
		endexpr = Singleton(comma1).Union(succ)
		t = names[typ]
		no = 1
		expression(&exprtype, endexpr)
		checktype(&exprtype, t.subtype2)
		for sym == comma1 {
			nextsym()
			no = no + 1
			expression(&exprtype, endexpr)
			checktype(&exprtype, t.subtype2)
		}
		max = t.bound2 - t.bound1 + 1
		if (t.subtype2 == char1) && (no < max) {
			emit2(blank2, max-no)
			no = max
		}
		if no != max {
			error(constructor3)
		}
	}
	var member_expression_list = func(typ int, succ symbols) {
		var t nameattr
		var exprtype, no int
		var endexpr symbols
		endexpr = Singleton(comma1).Union(succ)
		t = names[typ]
		no = 1
		expression(&exprtype, endexpr)
		checktype(&exprtype, t.subtype1)
		for sym == comma1 {
			nextsym()
			no = no + 1
			expression(&exprtype, endexpr)
			checktype(&exprtype, t.subtype1)
		}
		emit2(construct2, no)
	}
	var constructor = func(typ *int, succ symbols) {
		var mode namekind
		var endexpr symbols
		*typ = x
		mode = names[x].kind
		nextsym()
		if sym == lparanth1 {
			nextsym()
			endexpr = Singleton(rparanth1).Union(succ)
			if mode == elemtype {
				elementary_expression(*typ, endexpr)
			} else if mode == recordtype {
				field_expression_list(*typ, endexpr)
			} else if mode == arraytype {
				element_expression_list(*typ, endexpr)
			} else if mode == settype {
				member_expression_list(*typ, endexpr)
			}
			checksym(rparanth1, succ)
		} else if true {
			if mode == settype {
				emit2(construct2, 0)
			} else if true {
				error(constructor3)
			}
		}
		check(succ)
	}
	var function_call = func(typ *int, succ symbols) {
		if isfunction(x) {
			*typ = names[x].proctype
			emit2(valspace2, names[*typ].length)
		} else if true {
			kinderror2(x, typ)
		}
		procedure_call(succ)
	}
	var factor func(typ *int, succ symbols)
	factor = func(typ *int, succ symbols) {
		var mode namekind
		var endfactor symbols
		endfactor = Singleton(colon1).Union(succ)
		if literalsym.Contains(sym) {
			constant_factor(typ, endfactor)
		} else if sym == name1 {
			mode = names[x].kind
			if mode == constant {
				constant_factor(typ, endfactor)
			} else if typekinds.Contains(mode) {
				constructor(typ, endfactor)
			} else if varkinds.Contains(mode) {
				variable_factor(typ, endfactor)
			} else if prockinds.Contains(mode) {
				function_call(typ, endfactor)
			} else if true {
				kinderror2(x, typ)
				nextsym()
			}
		} else if sym == val1 {
			variable_factor(typ, endfactor)
		} else if sym == lparanth1 {
			nextsym()
			expression(typ, Singleton(rparanth1).Union(endfactor))
			checksym(rparanth1, endfactor)
		} else if sym == not1 {
			nextsym()
			factor(typ, endfactor)
			if *typ == bool1 {
				emit(not2)
			} else if true {
				typerror2(typ)
			}
		} else if true {
			*typ = univtype1
			syntax(endfactor)
		}
		for sym == colon1 {
			type_transfer(typ, endfactor)
		}
		check(succ)
	}
	var term = func(typ *int, succ symbols) {
		var op symbol
		var typ2 int
		var endfactor symbols
		endfactor = multsym.Union(succ)
		factor(typ, endfactor)
		for multsym.Contains(sym) {
			op = sym
			nextsym()
			factor(&typ2, endfactor)
			if *typ == int1 {
				checktype(typ, typ2)
				if op == asterisk1 {
					emit(multiply2)
				} else if op == div1 {
					emit(divide2)
				} else if op == mod1 {
					emit(modulo2)
				} else if true {
					typerror2(typ)
				}
			} else if *typ == bool1 {
				checktype(typ, typ2)
				if op == and1 {
					emit(and2)
				} else if true {
					typerror2(typ)
				}
			} else if names[*typ].kind == settype {
				checktype(typ, typ2)
				if op == asterisk1 {
					emit(intersection2)
				} else if true {
					typerror2(typ)
				}
			} else if true {
				typerror2(typ)
			}
		}
		check(succ)
	}
	var signed_term = func(typ *int, succ symbols) {
		var op symbol
		if signsym.Contains(sym) {
			op = sym
			nextsym()
			term(typ, succ)
			if *typ == int1 {
				if op == plus1 {
				} else if op == minus1 {
					emit(minus2)
				}
			} else if true {
				typerror2(typ)
			}
		} else if true {
			term(typ, succ)
		}
	}
	var simple_expression = func(typ *int, succ symbols) {
		var op symbol
		var typ2 int
		var endterm symbols
		endterm = addsym.Union(succ)
		signed_term(typ, endterm)
		for addsym.Contains(sym) {
			op = sym
			nextsym()
			term(&typ2, endterm)
			if *typ == int1 {
				checktype(typ, typ2)
				if op == plus1 {
					emit(add2)
				} else if op == minus1 {
					emit(subtract2)
				} else if true {
					typerror2(typ)
				}
			} else if *typ == bool1 {
				checktype(typ, typ2)
				if op == or1 {
					emit(or2)
				} else if true {
					typerror2(typ)
				}
			} else if names[*typ].kind == settype {
				checktype(typ, typ2)
				if op == plus1 {
					emit(union2)
				} else if op == minus1 {
					emit(difference2)
				} else if true {
					typerror2(typ)
				}
			} else if true {
				typerror2(typ)
			}
		}
		check(succ)
	}
	expression = func(typ *int, succ symbols) {
		var op symbol
		var typ2 int
		var t nameattr
		var endsimple symbols
		endsimple = relationsym.Union(succ)
		simple_expression(typ, endsimple)
		if relationsym.Contains(sym) {
			op = sym
			nextsym()
			simple_expression(&typ2, succ)
			t = names[typ2]
			if equalitysym.Contains(op) {
				checktype(typ, typ2)
				if op == equal1 {
					emit2(equal2, t.length)
				} else if true {
					emit2(notequal2, t.length)
				}
			} else if ordersym.Contains(op) {
				checktype(typ, typ2)
				checkelem(typ)
				if op == less1 {
					emit(less2)
				} else if op == greater1 {
					emit(greater2)
				} else if op == notless1 {
					emit(notless2)
				} else if op == notgreater1 {
					emit(notgreater2)
				}
			} else if op == in1 {
				if t.kind == settype {
					checktype(typ, t.subtype1)
					emit(in2)
				} else if true {
					typerror2(typ)
				}
			}
			if *typ != univtype1 {
				*typ = bool1
			}
		}
		check(succ)
	}
	var assignment_statement = func(succ symbols) {
		var vartype, exprtype int
		variable_symbol(&vartype, Singleton(becomes1).Union(exprsym).Union(succ))
		checksym(becomes1, exprsym.Union(succ))
		expression(&exprtype, succ)
		checktype(&vartype, exprtype)
		emit2(assign2, names[vartype].length)
	}
	var standard_call = func(succ symbols) {
		var endarg1, endarg2 symbols
		var typ int
		endarg2 = Singleton(rparanth1).Union(succ)
		endarg1 = Singleton(comma1).Union(endarg2)
		if x == addr1 {
			nextsym()
			checksym(lparanth1, endarg2)
			variable_symbol(&typ, endarg2)
			checksym(rparanth1, succ)
			emit(addr2)
		} else if x == halt1 {
			nextsym()
			emit(halt2)
		} else if x == obtain1 {
			nextsym()
			checksym(lparanth1, endarg1)
			expression(&typ, endarg1)
			checktype(&typ, int1)
			checksym(comma1, endarg2)
			variable_symbol(&typ, endarg2)
			checktype(&typ, int1)
			checksym(rparanth1, succ)
			emit(obtain2)
		} else if x == place1 {
			nextsym()
			checksym(lparanth1, endarg1)
			expression(&typ, endarg1)
			checktype(&typ, int1)
			checksym(comma1, endarg2)
			expression(&typ, endarg2)
			checktype(&typ, int1)
			checksym(rparanth1, succ)
			emit(place2)
		} else if x == sense1 {
			nextsym()
			checksym(lparanth1, endarg1)
			expression(&typ, endarg1)
			checktype(&typ, int1)
			checksym(comma1, endarg2)
			expression(&typ, endarg2)
			checktype(&typ, int1)
			checksym(rparanth1, succ)
			emit(sense2)
		}
	}
	var procedure_argument = func(succ symbols) {
		var p nameattr
		if sym == name1 {
			if prockinds.Contains(names[x].kind) {
				p = names[x]
				if p.kind == standard {
					kinderror1(x)
				} else if p.kind == procedur {
					emit3(procarg2, p.proclevel, p.procaddr)
				} else if p.kind == procparam {
					emit3(paramarg2, p.proclevel, p.procaddr)
				}
			} else if true {
				kinderror1(x)
			}
			nextsym()
		} else if true {
			syntax(succ)
		}
		check(succ)
	}
	var argument = func(param, size *int, succ symbols) {
		var typ int
		var n nameattr
		if *param != none {
			n = names[*param]
			if n.kind == valparam {
				expression(&typ, succ)
				checktype(&typ, n.vartype)
				*size = names[typ].length
			} else if n.kind == varparam {
				variable_symbol(&typ, succ)
				checktype(&typ, n.vartype)
				*size = elemlength
			} else if n.kind == procparam {
				procedure_argument(succ)
				*size = proclength
			}
			*param = n.link
		} else if true {
			expression(&typ, succ)
			*size = 0
			error(call3)
		}
	}
	var argument_list = func(name int, length *int, succ symbols) {
		var par, length2 int
		var endarg symbols
		endarg = Singleton(comma1).Union(succ)
		par = names[name].param
		argument(&par, length, endarg)
		for sym == comma1 {
			nextsym()
			argument(&par, &length2, endarg)
			*length = *length + length2
		}
		if par != none {
			error(call3)
		}
		check(succ)
	}
	procedure_call = func(succ symbols) {
		var name, length int
		var p nameattr
		name = x
		p = names[x]
		if p.kind == standard {
			standard_call(succ)
		} else if true {
			nextsym()
			if sym == lparanth1 {
				nextsym()
				argument_list(name, &length, Singleton(rparanth1).Union(succ))
				checksym(rparanth1, succ)
			} else if true {
				if p.param != none {
					error(call3)
				}
				length = 0
			}
			if p.kind == procedur {
				emit4(proccall2, p.proclevel, p.procaddr, length)
			} else if p.kind == procparam {
				emit4(paramcall2, p.proclevel, p.procaddr, length)
			}
		}
	}
	var conditional_statement = func(truelabel int, succ symbols) {
		var typ, falselabel int
		var enddo symbols
		enddo = statsym.Union(succ)
		expression(&typ, Singleton(do1).Union(enddo))
		if typ != bool1 {
			typerror2(&typ)
		}
		newlabel(&falselabel)
		checksym(do1, enddo)
		emit2(do2, falselabel)
		statement_list(succ)
		emit3(else2, truelabel, falselabel)
		check(succ)
	}
	var conditional_statement_list = func(truelabel int, succ symbols) {
		var endstat symbols
		endstat = Singleton(else1).Union(succ)
		conditional_statement(truelabel, endstat)
		for sym == else1 {
			nextsym()
			conditional_statement(truelabel, endstat)
		}
		check(succ)
	}
	var if_statement = func(succ symbols) {
		var truelabel int
		nextsym()
		newlabel(&truelabel)
		conditional_statement_list(truelabel, Singleton(end1).Union(succ))
		checksym(end1, succ)
		emit2(endif2, truelabel)
	}
	var while_statement = func(succ symbols) {
		var truelabel int
		nextsym()
		newlabel(&truelabel)
		emit2(while2, truelabel)
		conditional_statement_list(truelabel, Singleton(end1).Union(succ))
		checksym(end1, succ)
	}
	var when_statement = func(succ symbols) {
		var waitlabel, truelabel int
		nextsym()
		newlabel(&waitlabel)
		emit2(when2, waitlabel)
		newlabel(&truelabel)
		conditional_statement_list(truelabel, Singleton(end1).Union(succ))
		emit2(wait2, waitlabel)
		checksym(end1, succ)
		emit2(endwhen2, truelabel)
	}
	var process_statement = func(endlabel int, p *processattr, succ symbols) {
		var typ, procconst, proclabel, templabel int
		var enddo symbols
		enddo = statsym.Union(succ)
		constant_symbol(&typ, &procconst, Singleton(do1).Union(enddo))
		if (procconst < 0) || (procconst > setlimit) {
			error(cobegin3)
			procconst = 1
		}
		newlabel(&proclabel)
		newlabel(&templabel)
		checksym(do1, enddo)
		emit3(process2, proclabel, templabel)
		statement_list(succ)
		emit3(also2, endlabel, templabel)
		*p = processattr{procconst, proclabel}
		check(succ)
	}
	var process_statement_list = func(endlabel int, tasks *processtable, count *int, succ symbols) {
		var used processset
		var p processattr
		var endstat symbols
		endstat = Singleton(also1).Union(succ)
		process_statement(endlabel, &p, endstat)
		used = Singleton(p.procconst)
		*count = 1
		(*tasks)[1] = p
		for sym == also1 {
			nextsym()
			process_statement(endlabel, &p, endstat)
			if used.Contains(p.procconst) {
				error(cobegin3)
			} else if true {
				used = used.Add(p.procconst)
				if *count == maxprocess {
					fail(processlimit)
				}
				*count = *count + 1
				(*tasks)[*count] = p
			}
		}
		check(succ)
	}
	var concurrent_statement = func(succ symbols) {
		var beginlabel, endlabel, count, i int
		var tasks processtable
		nextsym()
		newlabel(&beginlabel)
		newlabel(&endlabel)
		emit2(goto2, beginlabel)
		process_statement_list(endlabel, &tasks, &count, Singleton(end1).Union(succ))
		checksym(end1, succ)
		emit4(cobegin2, beginlabel, endlabel, count)
		i = 0
		for i < count {
			i = i + 1
			emit(operator(tasks[i].procconst))
			emit(operator(tasks[i].proclabel))
		}
	}
	var statement = func(succ symbols) {
		if sym == skip1 {
			nextsym()
		} else if sym == val1 {
			assignment_statement(succ)
		} else if sym == name1 {
			if prockinds.Contains(names[x].kind) {
				if isfunction(x) {
					kinderror1(x)
				} else if true {
					procedure_call(succ)
				}
			} else if true {
				assignment_statement(succ)
			}
		} else if sym == if1 {
			if_statement(succ)
		} else if sym == while1 {
			while_statement(succ)
		} else if sym == when1 {
			when_statement(succ)
		} else if sym == cobegin1 {
			concurrent_statement(succ)
		} else if true {
			syntax(succ)
		}
		check(succ)
	}
	statement_list = func(succ symbols) {
		var semistat, endstat symbols
		semistat = Singleton(semicolon1).Union(statsym)
		endstat = semistat.Union(succ)
		statement(endstat)
		for semistat.Contains(sym) {
			checksym(semicolon1, statsym)
			statement(endstat)
		}
		check(succ)
	}
	var programx = func(succ symbols) {
		var enddecl symbols
		standard_names()
		enddecl = initdeclsym.Add(proc1).Union(succ)
		for initdeclsym.Contains(sym) {
			if sym == const1 {
				constant_declaration_list(enddecl)
			} else if typesym.Contains(sym) {
				type_declaration(enddecl)
			}
		}
		complete_procedure(false, 0, succ)
		check(succ)
	}
	programx(Singleton(endtext1))
	emit(endcode2)
}
