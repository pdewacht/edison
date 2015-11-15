{
#include "edipp.h"
#include <stdio.h>

#define PRINT fmt_print();
#define TAB fmt_tab();
#define NEWLINE fmt_newline();
#define JOIN fmt_join();
#define INC fmt_inc_indent();
#define DEC fmt_dec_indent();
}

%start parse_program, program;
%lexical token_get;
%token also1, and1, array1, asterisk1, becomes1, begin1, cobegin1,
    colon1, comma1, comment1, const1, div1, do1, else1, end1, enum1,
    equal1, greater1, if1, in1, lbracket1, less1, lib1, lparanth1,
    minus1, mod1, module1, name1, newline1, not1, notequal1,
    notgreater1, notless1, numeral1, or1, period1, plus1, pre1, post1,
    proc1, rbracket1, record1, rparanth1, semicolon1, set1, skip1,
    string1, val1, var1, when1, while1;

program:
    ws
    [ initial_declaration ]*
    complete_procedure
;

initial_declaration:
    constant_declaration_list
    | type_declaration
;

declaration:
    constant_declaration_list
    | type_declaration
    | variable_declaration_list
    | procedure_declaration
    | module_declaration
;

constant_declaration_list:
    const { INC } constant_declaration
    [ { JOIN } semicolon constant_declaration ]*
    { DEC }
;

constant_declaration:
    name equal constant_symbol
;

constant_symbol:  /* cheating */
    numeral
    | string
    | name [ procedure_call_tail ]?
;

type_declaration:
    enumeration_type
    | record_type
    | array_type
    | set_type
;

enumeration_type:
    enum { INC } name
    lparanth { JOIN } enumeration_symbol_list { JOIN } rparanth
    { DEC }
;

enumeration_symbol_list:
    name [ { JOIN } comma name ]*
;

record_type:
    record { INC } name
    lparanth { JOIN } field_list { JOIN } rparanth
    { DEC }
;

field_list:
    variable_list
;

array_type:
    array { INC } name
    lbracket { JOIN } range_symbol { JOIN } rbracket
    lparanth { JOIN } name { JOIN } rparanth
    { DEC }
;

range_symbol:
    constant_symbol { JOIN } colon { JOIN } constant_symbol
;

set_type:
    set { INC } name
    lparanth { JOIN } name { JOIN } rparanth
    { DEC }
;

variable_declaration_list:
    var { INC } variable_list { DEC }
;

variable_list:
    variable_group
    [ { JOIN } semicolon variable_group ]*
;

variable_group:
    name [ { JOIN } comma name ]*
    { JOIN } colon name
;

procedure_declaration:
    complete_procedure
    | preprocedure
    | postprocedure
    | library_procedure
;

complete_procedure:
    procedure_heading
    procedure_body
;

procedure_heading:
    proc { INC } name
    [ { JOIN } lparanth { JOIN } parameter_list { JOIN } rparanth ]?
    [ { JOIN } colon name ]?
    { DEC }
;

parameter_list:
    parameter_group
    [ { JOIN } semicolon parameter_group ]*
;

parameter_group:
    [ var ]? variable_group
    | procedure_heading
;

procedure_body:
    [ declaration ]*
    begin { INC }
    statement_list
    { DEC } end
;

preprocedure:
    pre procedure_heading
;

postprocedure:
    post complete_procedure
;

library_procedure:
    lib procedure_heading
    lbracket expression rbracket
;

module_declaration:
    module
    [ [ asterisk ]? { TAB INC } declaration { DEC } ]*
    begin { INC }
    statement_list
    { DEC } end
;

statement_list:
    statement
    [ { JOIN } semicolon statement ]*
;

statement:
    skip
    | assignment_or_procedure_call
    | if_statement
    | while_statement
    | when_statement
    | concurrent_statement
;

assignment_or_procedure_call:  /* hack hack hack */
    [ val ]? name { INC }
    [ { JOIN } selector ]*
    [ becomes expression ]?
    { DEC }
;

if_statement:
    if conditional_statement_list end
;

while_statement:
    while conditional_statement_list end
;

when_statement:
    when conditional_statement_list end
;

conditional_statement_list:
    condititional_statement
    [ else condititional_statement ]*
;

condititional_statement:
    { INC } expression { DEC } do
    { INC } statement_list { DEC }
;

concurrent_statement:
    cobegin process_statement_list end
;

process_statement_list:
    process_statement
    [ also process_statement ]*
;

process_statement:
    { INC } constant_symbol { DEC } do
    { INC } statement_list { DEC }
;

expression:
    simple_expression
    [ relational_operator simple_expression ]?
;

relational_operator:
    equal | greater | in | less
    | notequal | notgreater | notless
;

simple_expression:
    [ sign_operator ]?
    term
    [ adding_operator term ]*
;

sign_operator:
    minus | plus
;

adding_operator:
    minus | or | plus
;

term:
    factor
    [ multiplying_operator factor ]*
;

multiplying_operator:
    and | asterisk | div | mod
;

factor:  /* cheating */
    factor_ [ %while(1) type_transfer ]*
;

factor_:
    numeral
    | string
    | variable_symbol
    | lparanth { JOIN } expression { JOIN } rparanth
    | not factor
;

variable_symbol:
    [ val ]?
    name
    [ %while(1) { JOIN } selector ]*
;

selector:
    field_selector
    | indexed_selector
    | type_transfer
    | procedure_call_tail  /* ugh */
;

field_selector:
    period { JOIN } name
;

indexed_selector:
    lbracket { JOIN } expression { JOIN } rbracket
;

type_transfer:
    colon { JOIN } name
;

procedure_call_tail:
    lparanth { JOIN } argument_list { JOIN } rparanth
;

argument_list:
    expression
    [ { JOIN } comma expression ]*
;


also: also1 { PRINT } ws;
and: and1 { PRINT } ws;
array: array1 { PRINT } ws;
asterisk: asterisk1 { PRINT } ws;
becomes: becomes1 { PRINT } ws;
begin: begin1 { PRINT } ws;
cobegin: cobegin1 { PRINT } ws;
colon: colon1 { PRINT } ws;
comma: comma1 { PRINT } ws;
const: const1 { PRINT } ws;
div: div1 { PRINT } ws;
do: do1 { PRINT } ws;
else: else1 { PRINT } ws;
end: end1 { PRINT } ws;
enum: enum1 { PRINT } ws;
equal: equal1 { PRINT } ws;
greater: greater1 { PRINT } ws;
if: if1 { PRINT } ws;
in: in1 { PRINT } ws;
lbracket: lbracket1 { PRINT } ws;
less: less1 { PRINT } ws;
lib: lib1 { PRINT } ws;
lparanth: lparanth1 { PRINT } ws;
minus: minus1 { PRINT } ws;
mod: mod1 { PRINT } ws;
module: module1 { PRINT } ws;
name: name1 { PRINT } ws;
not: not1 { PRINT } ws;
notequal: notequal1 { PRINT } ws;
notgreater: notgreater1 { PRINT } ws;
notless: notless1 { PRINT } ws;
numeral: numeral1 { PRINT } ws;
or: or1 { PRINT } ws;
period: period1 { PRINT } ws;
plus: plus1 { PRINT } ws;
pre: pre1 { PRINT } ws;
post: post1 { PRINT } ws;
proc: proc1 { PRINT } ws;
rbracket: rbracket1 { PRINT } ws;
record: record1 { PRINT } ws;
rparanth: rparanth1 { PRINT } ws;
semicolon: semicolon1 { PRINT } ws;
set: set1 { PRINT } ws;
skip: skip1 { PRINT } ws;
string: string1 { PRINT } ws;
val: val1 { PRINT } ws;
var: var1 { PRINT } ws;
when: when1 { PRINT } ws;
while: while1 { PRINT } ws;

ws: [ newline1 { NEWLINE } | comment1 { PRINT } ]*;


{
  #include <err.h>

#ifndef LLNONCORR
  #error "Re-run LLgen with the -n flag"
#endif

  void LLmessage(int flag) {
    fmt_error();
    if (flag < 0) {
      warnx("%d: End of file expected\n", lineno);
    }
    else if (flag == 0) {
      if (LLsymb == EOFILE) {
        warnx("%d: Unexpected end of file", lineno);
      } else {
        warnx("%d: \"%s\" not understood", lineno, yytext);
      }
    }
    else if (flag > 0) {
      token_push_back(LLsymb);
    }
  }
}
