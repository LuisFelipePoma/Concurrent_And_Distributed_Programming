byte turn = 1
byte cc = 0

active proctype P() {
    do
    ::
        (turn == 1) ->
        cc++
        assert(cc < 2)
        cc--
        turn = 2
    od
}

active proctype Q() {
    do
    ::
        (turn == 2) ->
        cc++
        assert(cc < 2)
        cc--
        turn = 1
    od
}
