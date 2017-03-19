#ifndef ARGF_H
#define ARGF_H

// ARGF-style processing as in Ruby, adhering closely to POSIX getline()
// (in fact, a wrapper around the latter).
typedef struct _argf_state_t {
	FILE* fp;
	int is_stdin;
	int fidx;
	char** file_names;
	int num_file_names;
} argf_state_t;

void argf_init(argf_state_t* pstate, char** file_names, int num_file_names);

// Trailing "\n" or "\r\n" are chomped.
ssize_t argf_getline(argf_state_t* pstate, char** pline, size_t* plinecap);

#endif // ARGF_H
