package main

func pass2(next func(value *symbol), emit func(value symbol), fail func(reason failure)) {
	const (
		maxname = 750
	)
	var emit2 = func(sym symbol, arg int) {
		emit(sym)
		emit(symbol(arg))
	}
	type symbols = Set
	var addsym, constsym, declsym, exprsym, hiddensym, initdeclsym, literalsym, multsym, pairsym, procsym, relationsym, selectsym, signsym, statsym, typesym symbols
	addsym = Construct(minus1, or1, plus1)
	constsym = Construct(graphic1, name1, numeral1)
	declsym = Construct(array1, const1, enum1, lib1, module1, post1, pre1, proc1, record1, set1, var1)
	exprsym = Construct(graphic1, lparanth1, minus1, not1, numeral1, plus1, val1)
	hiddensym = Construct(error1, newline1)
	initdeclsym = Construct(array1, const1, enum1, record1, set1)
	literalsym = Construct(graphic1, numeral1)
	multsym = Construct(and1, asterisk1, div1, mod1)
	pairsym = Construct(graphic1, name1, numeral1)
	procsym = Construct(lib1, post1, pre1, proc1)
	relationsym = Construct(equal1, greater1, in1, less1, notequal1, notgreater1, notless1)
	selectsym = Construct(colon1, lbracket1, period1)
	signsym = Construct(minus1, plus1)
	statsym = Construct(cobegin1, if1, skip1, when1, while1)
	typesym = Construct(array1, enum1, record1, set1)
	var sym symbol
	var x int
	var skipsym = func() {
		next(&sym)
		for hiddensym.Contains(sym) {
			next(&x)
			emit2(sym, x)
			next(&sym)
		}
		if pairsym.Contains(sym) {
			next(&x)
		}
	}
	var nextsym = func() {
		emit(sym)
		if pairsym.Contains(sym) {
			emit(x)
		}
		skipsym()
	}
	skipsym()
	var error = func(kind errorkind) {
		emit2(error1, int(kind))
	}
	var syntax = func(succ symbols) {
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
		noname        = 0
	)
	type namekind = int
	const ( /* namekind enum */
		undeclared = iota
		incomplete
		universal
		constant
		type_
		field
		variable
		split
		partial
		complete
	)
	type namekinds = Set
	type nameattr = struct {
		kind                             namekind
		minlevel, maxlevel, originalname int
	}
	type nametable = [1 + maxname]nameattr
	type maptable = [1 + maxname]int
	var names nametable
	var nameno int
	var map_ maptable
	var blocklevel int
	var predeclare = func(name int, kind namekind) {
		map_[name] = name
		names[name] = nameattr{kind, blocklevel, blocklevel, name}
	}
	var beginblock = func() {
		blocklevel = blocklevel + 1
	}
	var endblock = func() {
		var finalname int
		var n nameattr
		finalname = 0
		for finalname < nameno {
			finalname = finalname + 1
			n = names[finalname]
			if n.kind != undeclared {
				if n.maxlevel < blocklevel {
					map_[n.originalname] = finalname
				} else if true {
					if n.kind == split {
						error(split3)
					}
					if n.maxlevel == n.minlevel {
						if map_[n.originalname] == finalname {
							map_[n.originalname] = noname
						}
						n = nameattr{undeclared, 0, 0, 0}
					} else if true {
						n.maxlevel = n.minlevel
						map_[n.originalname] = finalname
					}
					names[finalname] = n
				}
			}
		}
		blocklevel = blocklevel - 1
	}
	var newname = func(export bool, mode namekind) {
		var origin, scope int
		var n nameattr
		if nameno == maxname {
			fail(namelimit)
		}
		nameno = nameno + 1
		origin = blocklevel
		if export {
			scope = origin - 1
		} else if true {
			scope = origin
		}
		n = names[map_[x]]
		if !(Construct(undeclared, universal).Contains(n.kind)) && (n.maxlevel >= scope) {
			error(ambiguous3)
		}
		names[nameno] = nameattr{mode, scope, origin, x}
		map_[x] = nameno
		emit2(name1, nameno)
		skipsym()
	}
	var change = func(name int, newkind namekind) {
		if name != univname1 {
			names[map_[name]].kind = newkind
		}
	}
	var postname = func() {
		var finalname int
		var n nameattr
		finalname = map_[x]
		n = names[finalname]
		if (n.kind == split) && (n.maxlevel == blocklevel) {
			names[finalname].kind = partial
			emit2(name1, finalname)
			skipsym()
		} else if true {
			error(split3)
			newname(false, partial)
		}
	}
	var ischar = func(name int) (val_ischar bool) {
		val_ischar = map_[name] == char1
		return
	}
	var isproc = func(name int) (val_isproc bool) {
		val_isproc = Construct(split, partial, complete).Contains(names[map_[name]].kind)
		return
	}
	var kindof = func(name int) (val_kindof namekind) {
		val_kindof = names[map_[name]].kind
		return
	}
	var oldname = func() {
		var n nameattr
		n = names[map_[x]]
		if n.kind == undeclared {
			error(undeclared3)
			map_[x] = univname1
		} else if n.kind == incomplete {
			error(incomplete3)
			map_[x] = univname1
		}
		emit2(name1, map_[x])
		skipsym()
	}
	var valname = func() {
		var n nameattr
		n = names[map_[x]]
		if Construct(split, complete).Contains(n.kind) {
			error(funcval3)
			emit2(name1, univname1)
			skipsym()
		} else if true {
			oldname()
		}
	}
	names[noname] = nameattr{undeclared, 0, 0, 0}
	nameno = 0
	for nameno < maxname {
		nameno = nameno + 1
		map_[nameno] = noname
		names[nameno] = nameattr{undeclared, 0, 0, 0}
	}
	blocklevel = 0
	predeclare(bool1, type_)
	predeclare(char1, type_)
	predeclare(false1, constant)
	predeclare(int1, type_)
	predeclare(true1, constant)
	predeclare(univname1, universal)
	predeclare(addr1, complete)
	predeclare(halt1, complete)
	predeclare(obtain1, complete)
	predeclare(place1, complete)
	predeclare(sense1, complete)
	nameno = last_standard
	var variable_list func(export bool, kind namekind, succ symbols)
	var procedure_heading func(export, postx bool, name *int, succ symbols)
	var declaration func(export bool, succ symbols)
	var expression func(succ symbols)
	var procedure_call func(succ symbols)
	var statement_list func(succ symbols)
	var control_symbol = func(succ symbols) {
		nextsym()
		checksym(lparanth1, Construct(numeral1, rparanth1).Union(succ))
		if sym == numeral1 {
			nextsym()
		} else if true {
			syntax(Singleton(rparanth1).Union(succ))
		}
		checksym(rparanth1, succ)
	}
	var constant_symbol = func(succ symbols) {
		if sym == numeral1 {
			nextsym()
		} else if sym == graphic1 {
			nextsym()
		} else if sym == name1 {
			if ischar(x) {
				control_symbol(succ)
			} else if true {
				oldname()
			}
		} else if true {
			syntax(succ)
		}
		check(succ)
	}
	var constant_declaration = func(export bool, succ symbols) {
		var name int
		if sym == name1 {
			name = x
			newname(export, incomplete)
			checksym(equal1, constsym.Union(succ))
			constant_symbol(succ)
			change(name, constant)
		} else if true {
			syntax(succ)
		}
		check(succ)
	}
	var constant_declaration_list = func(export bool, succ symbols) {
		var enddecl symbols
		nextsym()
		enddecl = Singleton(semicolon1).Union(succ)
		constant_declaration(export, enddecl)
		for sym == semicolon1 {
			nextsym()
			constant_declaration(export, enddecl)
		}
	}
	var enumeration_symbol = func(export bool, succ symbols) {
		if sym == name1 {
			newname(export, constant)
		} else if true {
			syntax(succ)
		}
		check(succ)
	}
	var enumeration_symbol_list = func(export bool, succ symbols) {
		var endsym symbols
		endsym = Singleton(comma1).Union(succ)
		enumeration_symbol(export, endsym)
		for sym == comma1 {
			nextsym()
			enumeration_symbol(export, endsym)
		}
		check(succ)
	}
	var enumeration_type = func(export bool, succ symbols) {
		nextsym()
		if sym == name1 {
			newname(export, type_)
			checksym(lparanth1, Construct(name1, rparanth1).Union(succ))
			enumeration_symbol_list(export, Singleton(rparanth1).Union(succ))
			checksym(rparanth1, succ)
		} else if true {
			syntax(succ)
		}
	}
	var record_type = func(export bool, succ symbols) {
		var name int
		nextsym()
		if sym == name1 {
			name = x
			newname(export, incomplete)
			checksym(lparanth1, Construct(name1, rparanth1).Union(succ))
			variable_list(false, field, Singleton(rparanth1).Union(succ))
			checksym(rparanth1, succ)
			change(name, type_)
		} else if true {
			syntax(succ)
		}
	}
	var range_symbol = func(succ symbols) {
		constant_symbol(Singleton(colon1).Union(constsym).Union(succ))
		checksym(colon1, constsym.Union(succ))
		constant_symbol(succ)
	}
	var type_name = func(succ symbols) {
		if sym == name1 {
			oldname()
		} else if true {
			syntax(succ)
		}
	}
	var array_type = func(export bool, succ symbols) {
		var name int
		nextsym()
		if sym == name1 {
			name = x
			newname(export, incomplete)
			checksym(lbracket1, constsym.Union(Construct(rbracket1, lparanth1, name1, rparanth1)).Union(succ))
			range_symbol(Construct(rbracket1, lparanth1, name1, rparanth1).Union(succ))
			checksym(rbracket1, Construct(lparanth1, name1, rparanth1).Union(succ))
			checksym(lparanth1, Construct(name1, rparanth1).Union(succ))
			type_name(Singleton(rparanth1).Union(succ))
			checksym(rparanth1, succ)
			change(name, type_)
		} else if true {
			syntax(succ)
		}
	}
	var set_type = func(export bool, succ symbols) {
		var name int
		nextsym()
		if sym == name1 {
			name = x
			newname(export, incomplete)
			checksym(lparanth1, Construct(name1, rparanth1).Union(succ))
			type_name(Singleton(rparanth1).Union(succ))
			checksym(rparanth1, succ)
			change(name, type_)
		} else if true {
			syntax(succ)
		}
	}
	var type_declaration = func(export bool, succ symbols) {
		if sym == enum1 {
			enumeration_type(export, succ)
		} else if sym == record1 {
			record_type(export, succ)
		} else if sym == array1 {
			array_type(export, succ)
		} else if sym == set1 {
			set_type(export, succ)
		}
	}
	var variable_group = func(export bool, kind namekind, succ symbols) {
		if sym == name1 {
			newname(export, kind)
			for sym == comma1 {
				nextsym()
				if sym == name1 {
					newname(export, kind)
				} else if true {
					syntax(Construct(comma1, colon1).Union(succ))
				}
			}
			checksym(colon1, succ)
			type_name(succ)
			check(succ)
		} else if true {
			syntax(succ)
		}
	}
	variable_list = func(export bool, kind namekind, succ symbols) {
		var endgroup symbols
		endgroup = Singleton(semicolon1).Union(succ)
		variable_group(export, kind, endgroup)
		for sym == semicolon1 {
			nextsym()
			variable_group(export, kind, endgroup)
		}
		check(succ)
	}
	var variable_declaration_list = func(export bool, succ symbols) {
		nextsym()
		variable_list(export, variable, succ)
	}
	var parameter_group = func(succ symbols) {
		var name int
		if sym == proc1 {
			procedure_heading(false, false, &name, succ)
			endblock()
			change(name, complete)
		} else if true {
			if sym == var1 {
				nextsym()
			}
			variable_group(false, variable, succ)
		}
		check(succ)
	}
	var parameter_list = func(succ symbols) {
		var endgroup symbols
		endgroup = Singleton(semicolon1).Union(succ)
		parameter_group(endgroup)
		for sym == semicolon1 {
			nextsym()
			parameter_group(endgroup)
		}
		check(succ)
	}
	procedure_heading = func(export, postx bool, name *int, succ symbols) {
		checksym(proc1, Construct(name1, lparanth1, colon1).Union(succ))
		if sym == name1 {
			*name = x
			if postx {
				postname()
			} else if true {
				newname(export, partial)
			}
			beginblock()
			if sym == lparanth1 {
				nextsym()
				parameter_list(Construct(rparanth1, colon1).Union(succ))
				checksym(rparanth1, Singleton(colon1).Union(succ))
			}
			if sym == colon1 {
				nextsym()
				type_name(succ)
			}
		} else if true {
			*name = univname1
			beginblock()
			syntax(succ)
		}
		check(succ)
	}
	var procedure_body = func(succ symbols) {
		for declsym.Contains(sym) {
			declaration(false, declsym.Add(begin1).Add(end1).Union(statsym).Union(succ))
		}
		checksym(begin1, statsym.Add(end1).Union(succ))
		statement_list(Singleton(end1).Union(succ))
		checksym(end1, succ)
	}
	var complete_procedure = func(export, postx bool, succ symbols) {
		var name int
		procedure_heading(export, postx, &name, declsym.Add(begin1).Union(statsym).Union(succ))
		procedure_body(succ)
		endblock()
		change(name, complete)
	}
	var preprocedure = func(export bool, succ symbols) {
		var name int
		nextsym()
		procedure_heading(export, false, &name, succ)
		endblock()
		change(name, split)
	}
	var postprocedure = func(export bool, succ symbols) {
		nextsym()
		complete_procedure(false, true, succ)
	}
	var library_procedure = func(export bool, succ symbols) {
		var name int
		nextsym()
		procedure_heading(export, false, &name, Construct(lbracket1, rbracket1).Union(exprsym).Union(succ))
		checksym(lbracket1, exprsym.Add(rbracket1).Union(succ))
		expression(Singleton(rbracket1).Union(succ))
		checksym(rbracket1, succ)
		endblock()
		change(name, complete)
	}
	var procedure_declaration = func(export bool, succ symbols) {
		if sym == proc1 {
			complete_procedure(export, false, succ)
		} else if sym == pre1 {
			preprocedure(export, succ)
		} else if sym == post1 {
			postprocedure(export, succ)
		} else if sym == lib1 {
			library_procedure(export, succ)
		}
	}
	var module_declaration = func(succ symbols) {
		var export bool
		nextsym()
		beginblock()
		for (Singleton(asterisk1).Union(declsym)).Contains(sym) {
			if sym == asterisk1 {
				export = true
				nextsym()
			} else if true {
				export = false
			}
			declaration(export, declsym.Union(Construct(asterisk1, begin1, end1)).Union(statsym).Union(succ))
		}
		checksym(begin1, statsym.Add(end1).Union(succ))
		statement_list(Singleton(end1).Union(succ))
		checksym(end1, succ)
		endblock()
	}
	declaration = func(export bool, succ symbols) {
		if sym == const1 {
			constant_declaration_list(export, succ)
		} else if typesym.Contains(sym) {
			type_declaration(export, succ)
		} else if sym == var1 {
			variable_declaration_list(export, succ)
		} else if procsym.Contains(sym) {
			procedure_declaration(export, succ)
		} else if sym == module1 {
			module_declaration(succ)
		} else if true {
			syntax(succ)
		}
		check(succ)
	}
	var function_variable = func(succ symbols) {
		nextsym()
		if sym == name1 {
			valname()
		} else if true {
			syntax(succ)
		}
	}
	var field_selector = func(succ symbols) {
		nextsym()
		if sym == name1 {
			oldname()
			check(succ)
		} else if true {
			syntax(succ)
		}
	}
	var indexed_selector = func(succ symbols) {
		nextsym()
		expression(Singleton(rbracket1).Union(succ))
		checksym(rbracket1, succ)
		check(succ)
	}
	var type_transfer = func(succ symbols) {
		nextsym()
		type_name(succ)
		check(succ)
	}
	var variable_symbol = func(succ symbols) {
		var endvar symbols
		endvar = selectsym.Union(succ)
		if sym == name1 {
			oldname()
		} else if sym == val1 {
			function_variable(endvar)
		} else if true {
			syntax(endvar)
		}
		for {
			if sym == period1 {
				field_selector(endvar)
			} else if sym == lbracket1 {
				indexed_selector(endvar)
			} else if sym == colon1 {
				type_transfer(endvar)
			} else {
				break
			}
		}
		check(succ)
	}
	var constructor = func(succ symbols) {
		var endexpr symbols
		oldname()
		if sym == lparanth1 {
			nextsym()
			endexpr = Construct(comma1, rparanth1).Union(succ)
			expression(endexpr)
			for sym == comma1 {
				nextsym()
				expression(endexpr)
			}
			checksym(rparanth1, succ)
		}
		check(succ)
	}
	var factor func(succ symbols)
	factor = func(succ symbols) {
		var endfactor symbols
		var kind namekind
		endfactor = Singleton(colon1).Union(succ)
		if literalsym.Contains(sym) {
			constant_symbol(endfactor)
		} else if sym == name1 {
			kind = kindof(x)
			if kind == constant {
				constant_symbol(endfactor)
			} else if kind == type_ {
				constructor(endfactor)
			} else if kind == variable {
				variable_symbol(endfactor)
			} else if isproc(x) {
				procedure_call(endfactor)
			} else if true {
				oldname()
			}
		} else if sym == val1 {
			variable_symbol(endfactor)
		} else if sym == lparanth1 {
			nextsym()
			expression(Singleton(rparanth1).Union(succ))
			checksym(rparanth1, endfactor)
		} else if sym == not1 {
			nextsym()
			factor(endfactor)
		} else if true {
			syntax(endfactor)
		}
		for sym == colon1 {
			type_transfer(endfactor)
		}
		check(succ)
	}
	var term = func(succ symbols) {
		var endfactor symbols
		endfactor = multsym.Union(succ)
		factor(endfactor)
		for multsym.Contains(sym) {
			nextsym()
			factor(endfactor)
		}
		check(succ)
	}
	var simple_expression = func(succ symbols) {
		var endterm symbols
		endterm = addsym.Union(succ)
		if signsym.Contains(sym) {
			nextsym()
		}
		term(endterm)
		for addsym.Contains(sym) {
			nextsym()
			term(endterm)
		}
		check(succ)
	}
	expression = func(succ symbols) {
		var endsimple symbols
		endsimple = relationsym.Union(succ)
		simple_expression(endsimple)
		if relationsym.Contains(sym) {
			nextsym()
			simple_expression(succ)
		}
		check(succ)
	}
	var assignment_statement = func(succ symbols) {
		variable_symbol(Singleton(becomes1).Union(exprsym).Union(succ))
		checksym(becomes1, exprsym.Union(succ))
		expression(succ)
	}
	var argument_list = func(succ symbols) {
		var endexpr symbols
		endexpr = Singleton(comma1).Union(succ)
		expression(endexpr)
		for sym == comma1 {
			nextsym()
			expression(endexpr)
		}
		check(succ)
	}
	procedure_call = func(succ symbols) {
		oldname()
		if sym == lparanth1 {
			nextsym()
			argument_list(Singleton(rparanth1).Union(succ))
			checksym(rparanth1, succ)
		}
	}
	var conditional_statement = func(succ symbols) {
		var enddo symbols
		enddo = statsym.Union(succ)
		expression(Singleton(do1).Union(enddo))
		checksym(do1, enddo)
		statement_list(succ)
		check(succ)
	}
	var conditional_statement_list = func(succ symbols) {
		var endstat symbols
		endstat = Singleton(else1).Union(succ)
		conditional_statement(endstat)
		for sym == else1 {
			nextsym()
			conditional_statement(endstat)
		}
		check(succ)
	}
	var if_statement = func(succ symbols) {
		nextsym()
		conditional_statement_list(Singleton(end1).Union(succ))
		checksym(end1, succ)
	}
	var while_statement = func(succ symbols) {
		nextsym()
		conditional_statement_list(Singleton(end1).Union(succ))
		checksym(end1, succ)
	}
	var when_statement = func(succ symbols) {
		nextsym()
		conditional_statement_list(Singleton(end1).Union(succ))
		checksym(end1, succ)
	}
	var process_statement = func(succ symbols) {
		var enddo symbols
		enddo = statsym.Union(succ)
		constant_symbol(Singleton(do1).Union(enddo))
		checksym(do1, enddo)
		statement_list(succ)
		check(succ)
	}
	var process_statement_list = func(succ symbols) {
		var endstat symbols
		endstat = Singleton(also1).Union(succ)
		process_statement(endstat)
		for sym == also1 {
			nextsym()
			process_statement(endstat)
		}
		check(succ)
	}
	var concurrent_statement = func(succ symbols) {
		nextsym()
		process_statement_list(Singleton(end1).Union(succ))
		checksym(end1, succ)
	}
	var statement = func(succ symbols) {
		if sym == skip1 {
			nextsym()
		} else if sym == val1 {
			assignment_statement(succ)
		} else if sym == name1 {
			if isproc(x) {
				procedure_call(succ)
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
	var program = func(succ symbols) {
		var enddecl symbols
		enddecl = initdeclsym.Add(proc1).Union(succ)
		for initdeclsym.Contains(sym) {
			if sym == const1 {
				constant_declaration_list(false, enddecl)
			} else if typesym.Contains(sym) {
				type_declaration(false, enddecl)
			}
		}
		complete_procedure(false, false, succ)
		check(succ)
	}
	program(Singleton(endtext1))
	emit(endtext1)
}
