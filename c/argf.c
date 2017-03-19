#include <stdio.h>
#include <stdlib.h>
#include "argf.h"

// ----------------------------------------------------------------
void argf_init(argf_state_t* pstate, char** file_names, int num_file_names) {
	if (num_file_names == 0) {
		pstate->fp             = stdin;
		pstate->is_stdin       = 1;
		pstate->fidx           = 0;
		pstate->file_names     = NULL;
		pstate->num_file_names = 0;
	} else {
		pstate->fp             = NULL;
		pstate->is_stdin       = 0;
		pstate->fidx           = -1;
		pstate->file_names     = file_names;
		pstate->num_file_names = num_file_names;
	}
}

// ----------------------------------------------------------------
ssize_t argf_getline(argf_state_t* pstate, char** pline, size_t* plinecap) {

	if (!pstate->is_stdin) {
		if (pstate->fidx == -1) { // first call
			pstate->fidx = 0;
			pstate->fp = fopen(pstate->file_names[pstate->fidx], "r");
			if (pstate->fp == NULL) {
				perror("fopen");
				fprintf(stderr, "Could not open \"%s\" for read.\n",
					pstate->file_names[pstate->fidx]);
				exit(1);
			}
		}
	}

	ssize_t linelen = 0;
	while (1) {
		linelen = getline(pline, plinecap, pstate->fp);
		if (linelen > 0) {
			char* line = *pline;
			// chomp
			if (linelen >= 1 && line[linelen-1] == '\n') {
				line[linelen-1] = 0;
				linelen--;
			}
			if (linelen >= 1 && line[linelen-1] == '\r') {
				line[linelen-1] = 0;
				linelen--;
			}
			break;

		} else if (pstate->is_stdin) {
			break;

		} else if ((pstate->fidx + 1) >= pstate->num_file_names) {
			break;

		} else {
			pstate->fidx++;
			pstate->fp = fopen(pstate->file_names[pstate->fidx], "r");
			if (pstate->fp == NULL) {
				perror("fopen");
				fprintf(stderr, "Could not open \"%s\" for read.\n",
					pstate->file_names[pstate->fidx]);
				exit(1);
			}
		}
	}

	return linelen;
}
