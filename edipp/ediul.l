%option outfile="ediul.c"
%option case-insensitive
%option noyywrap noinput nounput

%{
#include <err.h>
#include <stdbool.h>
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

static bool opt_underline, opt_bold;

static void fputs_ul(const char *s, FILE *stream) {
  int i;
  for (i = 0; s[i]; i += 1) {
    if (opt_underline) {
      putc('_', stream);
      putc('\b', stream);
    }
    if (opt_bold) {
      putc(yytext[i], stream);
      putc('\b', stream);
    }
    putc(yytext[i], stream);
  }
}

static void usage(const char *argv0) {
  fprintf(stderr, "usage: %s [-u] [-b] [file ...]\n", argv0);
}

int main(int argc, char *argv[]) {
  int opt;
  while ((opt = getopt(argc, argv, "ub")) != -1) {
    switch (opt) {
    case 'u':
      opt_underline = true;
      break;
    case 'b':
      opt_bold = true;
      break;
    default:
      usage(argv[0]);
    }
  }
  if (!opt_underline && !opt_bold) {
    opt_underline = true;
  }

  stream = stdout;
  if (isatty(STDOUT_FILENO)) {
    FILE *ul = popen("/usr/bin/ul", "w");
    if (ul) {
      stream = ul;
    }
  }

  if (optind == argc) {
    yylex();
  } else {
    int i;
    for (i = optind; i < argc; i += 1) {
      yyin = fopen(argv[i], "r");
      if (!yyin) {
        err(1, "%s", argv[i]);
      }
      yylex();
      fclose(yyin);
    }
  }

  fclose(stream);
  return 0;
}
