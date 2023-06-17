#include <stdio.h>
#include <unistd.h>
#include <seccomp.h>

int main() {
    // Initialize the libseccomp context
    scmp_filter_ctx ctx = seccomp_init(SCMP_ACT_KILL); // default action: kill

    // Allow certain system calls
    seccomp_rule_add(ctx, SCMP_ACT_ALLOW, SCMP_SYS(read), 0);
    seccomp_rule_add(ctx, SCMP_ACT_ALLOW, SCMP_SYS(write), 0);
    seccomp_rule_add(ctx, SCMP_ACT_ALLOW, SCMP_SYS(exit), 0);

    // Load the seccomp rules into the kernel
    seccomp_load(ctx);

    // Execute a system call that is not allowed
    printf("Executing an illegal system call...\n");
    unlink("nonexistent");

    printf("You will never see this line\n");
    return 0;
}

