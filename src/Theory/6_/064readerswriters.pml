#define wait(s) atomic { s > 0 -> s-- }
#define signal(s) s++

byte readers = 0
byte mutex = 1
byte roomEmpty = 1

active proctype Writer() {
	do
	::
		wait(roomEmpty)
		printf("Writer doing his thing\n")
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
		
		printf("Reader doing his thing\n")
		
		wait(mutex)
		readers--
		if
		:: (readers == 0) -> signal(roomEmpty)
		:: else ->
		fi
		signal(mutex)
	od
}