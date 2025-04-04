# got-fuzzer
fuzzing tool for got overwrite.
## theorem
libc got is writable.
overwrite libc allow attacker to call arbitrary function with arbitrary arguments.

`fsop` is known and strong exploit technique to execute the shell.
but it needs little large space to set the disguised vtable.

this tool: `got-fuzzer` fuzz glibc got overwrite.
