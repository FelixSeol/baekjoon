#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>

void test1(){
extern int global_var;

	pid_t pid1, pid2, pid3;
	pid1 = getpid();
	global_var = 0;
	pid2 = fork();
	if( pid2>0){
		globar_var = 1;
	}
	pid3 = getpid();
}
