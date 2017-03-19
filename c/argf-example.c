#include <stdio.h>
#include <stdlib.h>
#include "argf.h"

// A simple line-oriented cat using argf.
int main(int argc, char** argv) {
	argf_state_t argf_state;
	argf_init(&argf_state, argv+1, argc-1);

	char *line = NULL;
	size_t linecap = 0;
	ssize_t linelen;

	// In this example, the line buffer is reused on each call.
	// If the 2nd argument to argf_getline were NULL each time, it
	// would allocate a new buffer each time, which should then be
	// freed within this while-loop.
	while ((linelen = argf_getline(&argf_state, &line, &linecap)) >= 0) {
		printf("%s\n", line);
	}

	if (line != NULL) {
		free(line);
	}

	return 0;
}
