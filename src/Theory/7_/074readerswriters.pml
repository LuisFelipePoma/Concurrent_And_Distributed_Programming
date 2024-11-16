#define wait(s) atomic { s > 0 -> s-- }
#define signal(s) s++

byte readers = 0
byte mutex = 1
byte roomEmpty = 1

byte countW = 0
byte countR = 0

active proctype Writer() {
    do
    ::
        wait(roomEmpty)

        countW++
        assert(countW == 1 && countR == 0)
        countW--

        signal(roomEmpty)
    od
}

active[2] proctype Reader() {
    do
    ::
        wait(mutex)
        readers++
        if
        :: (readers == 1) -> wait(roomEmpty)
        :: else ->
        fi
        signal(mutex)

        countR++
        assert(countW == 0 && countR > 0)
        countR--

        wait(mutex)
        readers--
        if
        :: (readers == 0) -> signal(roomEmpty)
        :: else ->
        fi
        signal(mutex)
    od
}
