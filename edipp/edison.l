%option case-insensitive
%option noyywrap noinput nounput

%{

#include <err.h>
#include "Lpars.h"
#include "edipp.h"

int lineno = 1;

static void count_lines() {
  const char *s = strchr(yytext, '\n');
  while (s != NULL) {
    lineno += 1;
    s = strchr(s + 1, '\n');
  }
}

%}

%%

ALSO            return also1;
AND             return and1;
ARRAY           return array1;
BEGIN           return begin1;
COBEGIN         return cobegin1;
CONST           return const1;
DIV             return div1;
DO              return do1;
ELSE            return else1;
END             return end1;
ENUM            return enum1;
IF              return if1;
IN              return in1;
LIB             return lib1;
MOD             return mod1;
MODULE          return module1;
NOT             return not1;
OR              return or1;
PRE             return pre1;
POST            return post1;
PROC            return proc1;
RECORD          return record1;
SET             return set1;
SKIP            return skip1;
VAL             return val1;
VAR             return var1;
WHEN            return when1;
WHILE           return while1;

[ \t\r]+        /* ignore whitespace */
\n              lineno += 1; return newline1;
\"[^"]*\"       count_lines(); return comment1;
\"[^"]*         errx(1, "%d: unterminated comment", lineno);
[A-Z][A-Z0-9_]* return name1;
[0-9]+          return numeral1;
;               return semicolon1;
,               return comma1;
\.              return period1;
\(              return lparanth1;
\)              return rparanth1;
\[              return lbracket1;
\]              return rbracket1;
:=              return becomes1;
:               return colon1;
>=              return notless1;
>               return greater1;
\<>             return notequal1;
\<=             return notgreater1;
\<              return less1;
\+              return plus1;
-               return minus1;
\*              return asterisk1;
=               return equal1;
'[^'\n]+'       return string1;
'[^'\n]+        warnx("%d: unterminated string", lineno); return string1;
#[0-7]+         return numeral1;

.               return 1;

%%

static int saved_token;

int token_get(void) {
  int t;
  if (saved_token) {
    t = saved_token;
    saved_token = 0;
  } else {
    t = yylex();
  }
  return t;
}

void token_push_back(int token) {
  saved_token = token;
}
