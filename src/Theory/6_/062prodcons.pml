#define wait(s) atomic { s > 0 -> s-- }
#define signal(s) s++

#define N 5

byte notEmpty = 0
byte notFull = N
byte buffer[N]
byte top = 0
byte mutex = 1

active proctype Producer() {
	byte d
	do
	::
		d++
		wait(notFull)
		wait(mutex)
		buffer[top] = d
		top++
		signal(mutex)
		signal(notEmpty)
	od
}
active proctype Consumer() {
	byte d
	do
	::
		wait(notEmpty)
		wait(mutex)
		top--
		d = buffer[top]
		signal(mutex)
		signal(notFull)
		printf("Consumiendo producto %d\n", d)
	od
}