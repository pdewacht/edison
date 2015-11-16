#include <stdio.h>

/* edison.l */

extern int token_get(void);
extern void token_push_back(int token);
extern FILE *yyin;
extern char *yytext;
extern int lineno;

/* edison.g */

extern void parse_program(void);

/* edipp.c */

extern void fmt_print(void);
extern void fmt_newline(void);
extern void fmt_tab(void);
extern void fmt_join(void);
extern void fmt_inc_indent(void);
extern void fmt_dec_indent(void);
extern void fmt_error(void);
