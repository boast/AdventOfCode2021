package main

// By manual decompile and some help from reddit...

// Inputs are A;B;C;D;E;F;G;H;I;J;K;L;M;N
// z = new stack()

// z.push(A + 6)
// z.push(B + 14)
// z.push(C + 14)
// if D != z.pop() - 8
//    z.push(D + 10)
// z.push(E + 9)
// z.push(F + 12)
// if G != z.pop() - 11
//    z.push(G + 8)
// if H != z.pop() - 4
//    z.push(H + 13)
// if I != z.pop() - 15
//    z.push(I + 12)
// z.push(J + 6)
// z.push(K + 9)
// if L != z.pop() - 1
//    z.push(L + 15)
// if M != z.pop() - 8
//    z.push(M + 4)
// if N != z.pop() - 14
//    z.push(M + 10)

// As we need to prevent all conditional pushes so the stack is empty, we can replace the conditions

// D == C + 14 - 8
// G == F + 12 - 11
// H == E + 9 - 4
// I == B + 14 - 15
// L == K + 9 - 1
// M == J + 6 - 8
// N == A + 6 - 14

// Cleaning up

// D == C + 6
// G == F + 1
// H == E + 5
// I == B - 1
// L == K + 8
// M == J - 2
// N == A - 8

// No we solve by hand and make ordered by alphabet as large as possible

// N == A - 8 --> A = 9, N = 1
// I == B - 1 --> B = 9, I = 8
// D == C + 6 --> C = 3, D = 9
// H == E + 5 --> E = 4, H = 9
// G == F + 1 --> F = 8, G = 9
// M == J - 2 --> J = 9, M = 7
// L == K + 8 --> K = 1, L = 9

func Day24Part1() int64 {
	return int64(99394899891971)
}

// ...or as small as possible

// N == A - 8 --> A = 9, N = 1
// I == B - 1 --> B = 2, I = 1
// D == C + 6 --> C = 1, D = 7
// H == E + 5 --> E = 1, H = 6
// G == F + 1 --> F = 1, G = 2
// M == J - 2 --> J = 3, M = 1
// L == K + 8 --> K = 1, L = 9

func Day24Part2() int64 {
	return int64(92171126131911)
}
