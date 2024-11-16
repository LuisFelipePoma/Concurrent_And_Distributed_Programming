byte turn = 1

active proctype P() {
    do
    ::
        printf("NCS 1\n")
        printf("NCS 2\n")
        printf("NCS 3\n")

        (turn == 1) ->

        printf("Critical 1\n")
        printf("Critical 2\n")
        printf("Critical 3\n")

        turn = 2
    od
}

active proctype Q() {
    do
    ::
        printf("NCS 1\n")
        printf("NCS 2\n")
        printf("NCS 3\n")

        (turn == 2) ->

        printf("Critical 1\n")
        printf("Critical 2\n")
        printf("Critical 3\n")

        turn = 1
    od
}
