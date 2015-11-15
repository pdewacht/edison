%option outfile="ediul.c"
%option case-insensitive
%option noyywrap noinput nounput

%{
#include <stdio.h>
#include <unistd.h>

static void fputs_ul(const char *s, FILE *stream);
static FILE *stream;

%}

%%

ALSO            |
AND             |
ARRAY           |
BEGIN           |
COBEGIN         |
CONST           |
DIV             |
DO              |
ELSE            |
END             |
ENUM            |
IF              |
IN              |
LIB             |
MOD             |
MODULE          |
NOT             |
OR              |
PRE             |
POST            |
PROC            |
RECORD          |
SET             |
SKIP            |
VAL             |
VAR             |
WHEN            |
WHILE           fputs_ul(yytext, stream);

\"[^"]*\"?      |
'[^'\n]*'?      |
[A-Z][A-Z0-9_]* |
[0-9]+          |
[^A-Z0-9"']+    fputs(yytext, stream);

.               abort();  /* not reached */

%%

static void fputs_ul(const char *s, FILE *stream) {
  int i;
  for (i = 0; s[i]; i += 1) {
    putc('_', stream);
    putc('\b', stream);
    putc(yytext[i], stream);
  }
}

int main(void) {
  stream = stdout;
  if (isatty(STDOUT_FILENO)) {
    FILE *ul = popen("/usr/bin/ul", "w");
    if (ul) {
      stream = ul;
    }
  }
  yylex();
  fclose(stream);
  return 0;
}
