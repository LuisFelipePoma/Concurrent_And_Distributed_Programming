#define wait(s) atomic { s > 0 -> s-- }
#define signal(s) atomic { s++ }

byte S = 1;  // Semáforo
bool B = false; // Variable compartida

proctype P() {
    wait(S);               // Espera el semáforo
    B = true;             // Modifica B
    assert(B == true);    // Aserción para validar que B es true
    signal(S);            // Libera el semáforo
}

proctype Q() {
    wait(S);              // Espera el semáforo
    do
    :: !B -> printf("*\n"); // Imprime un asterisco si B es falso
    :: B -> break;        // Sale del ciclo si B es verdadero
    od;
    signal(S);            // Libera el semáforo
}

init {
    run P();              // Inicia proceso P
    run Q();              // Inicia proceso Q
}
