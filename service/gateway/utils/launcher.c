#include <assert.h>
#include <stdio.h>
#include <sys/types.h>
#include <unistd.h>

int main(int argc, char* argv[]) {
    const char* daemon = argv[1];
    assert(daemon != NULL);

    pid_t pid = fork();
    if (pid == 0) {
        execl(daemon, daemon, NULL);
    }

    return 0;
}
