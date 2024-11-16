#define wait(s) atomic { s > 0 -> s-- }
#define signal(s) s++

#define N 5

byte fork[N] = { 1, 1, 1, 1, 1 }
byte room = 4

active[N] proctype Philosopher() {
	pid i = _pid
	do
	::
		printf("Pensando!\n")
		wait(room)
		wait(fork[i])
		wait(fork[(i+1) % N])
		printf("Comiendo!\n")
		signal(fork[i])
		signal(fork[(i+1) % N])
		signal(room)
	od
}