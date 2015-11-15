CFLAGS = -O2 -Wall -Wextra

all: edipp ediul

edipp: edison.l edison.g edipp.c edipp.h
	flex edison.l
	LLgen -n edison.g
	$(CC) $(CFLAGS) -o edipp edipp.c edison.c Lpars.c Lncor.c lex.yy.c

ediul: ediul.l
	flex ediul.l
	$(CC) $(CFLAGS) -o ediul ediul.c

clean:
	rm -f edipp edison.c Lpars.h Lpars.c Lncor.c lex.yy.c
	rm -f ediul ediul.c
