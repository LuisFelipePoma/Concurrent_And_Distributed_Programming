mtype = { red, yellow, green }
mtype light = green;

active proctype P() {
	byte c = 0
	do
	:: c < 20 ->
		c++
		if
		:: light == red -> 		light = green
		:: light == yellow -> 	light = red
		:: light == green -> 	light = yellow
		fi
		printf("The light is now %e\n", light)
	:: else -> break
	od
}