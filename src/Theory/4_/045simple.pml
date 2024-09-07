int n = 0

proctype P() {
	byte k1 = 1
	n = k1
}

proctype Q() {
	byte k2 = 2
	n = k2
}

init {
	atomic {
		run P()
		run Q()
	}
	(_nr_pr == 1) -> printf("n = %d\n", n)
	assert(n == 1)
}