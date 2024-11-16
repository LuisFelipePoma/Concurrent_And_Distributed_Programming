#define wait(s) atomic { s > 0 -> s-- }
#define signal(s) s++

#define N 5

byte fork[N] = { 1, 1, 1, 1, 1 }

active[N-1] proctype Philosopher() {
	pid i = _pid
	do
	::
		printf("Pensando!\n")
		wait(fork[i])
		wait(fork[(i+1) % N])
		printf("Comiendo!\n")
		signal(fork[i])
		signal(fork[(i+1) % N])
		
	od
}

active proctype PhilosopherLeftie() {
	pid i = _pid
	do
	::
		printf("Pensando!\n")
		wait(fork[0])
		wait(fork[4])
		printf("Comiendo!\n")
		signal(fork[0])
		signal(fork[4])
		
	od
}