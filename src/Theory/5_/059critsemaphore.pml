#define wait(s) atomic { s > 0 -> s-- }
#define signal(s) s++

byte S = 1

active[2] proctype P() {
    do
    ::  printf("NCS 1\n")
        printf("NCS 2\n")
        printf("NCS 3\n")

        wait(S)

        printf("Critical 1\n")
        printf("Critical 2\n")
        printf("Critical 3\n")

        signal(S)
    od
}
