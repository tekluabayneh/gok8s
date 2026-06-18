Does this data live on the stack or the heap?

Am I passing a value, a pointer, or a reference—and who owns it?

Can multiple threads/goroutines touch this at the same time?

What is the source of truth—and where is the state machine?

When this loop/function panics or errors out, what state is left behind?

What is the absolute worst-case time and space complexity here?

What system call does this trigger under the hood?

What happens to the file descriptor, socket, or pipe if this process dies?

Why this specific data structure shape and not another?

What happens if the network latency spikes or a timeout occurs?
