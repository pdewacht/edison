#include <assert.h>
#include <err.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include "edipp.h"
#include "Lpars.h"

static int indent_step = 2;
static int indent = 0;

static int column = 0;
static bool error_flag = false;
static bool join_flag = false;

void fmt_print(void) {
  if (!error_flag) {
    if (column == 0) {
      int i;
      for (i = 0; i < indent; ++i) {
        putchar(' ');
      }
      column = indent;
    }
    else if (!join_flag) {
      putchar(' ');
      column += 1;
    }
    fputs(yytext, stdout);
    column += strlen(yytext);
    join_flag = false;
  }
}

void fmt_newline(void) {
  if (!error_flag) {
    putchar('\n');
    column = 0;
  }
}

void fmt_join(void) {
  join_flag = true;
}

void fmt_tab(void) {
  while ((column % indent_step) != 0) {
    putchar(' ');
    column += 1;
  }
  join_flag = true;
}

void fmt_inc_indent(void) {
  indent += indent_step;
}

void fmt_dec_indent(void) {
  indent -= indent_step;
  assert(indent >= 0);
}

void fmt_error(void) {
  if (column != 0) {
    putchar('\n');
    column = 0;
  }
  error_flag = true;
}

static void usage(const char *argv0) {
  fprintf(stderr, "usage: %s [-i spaces] [file]\n", argv0);
  exit(1);
}

int main(int argc, char *argv[]) {
  int opt;
  while ((opt = getopt(argc, argv, "i:")) != -1) {
    switch (opt) {
    case 'i':
      indent_step = atoi(optarg);
      break;
    default:
      usage(argv[0]);
    }
  }
  if (optind + 1 < argc) {
    usage(argv[0]);
  }
  if (optind < argc) {
    yyin = fopen(argv[optind], "r");
    if (!yyin) {
      err(1, argv[optind]);
    }
  }

  parse_program();
  return error_flag;
}
