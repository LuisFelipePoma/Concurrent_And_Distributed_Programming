#define wait(s) atomic { s > 0 -> s-- }
#define signal(s) s++

byte S = 1
byte cc = 0

active[2] proctype P() {
    do
    ::
        wait(S)

        cc++
        assert(cc < 2)
        cc--

        signal(S)
    od
}
