#include <iostream>
using namespace std;

void IncTime(int& H, int& M, int& S, int T);

int main()
{
	//Task("Proc51");
	int h, m, s, t;
	cin >> h >> m >> s >> t;
	IncTime(h, m, s, t);
	cout << h << m << s;
	
	return 0;
}

void IncTime(int& H, int& M, int& S, int T)
{
	S += T % 60;

	M += T % 3600 / 60 + S / 60;
	S %= 60;

	H += T / 3600 + M / 60;
	M %= 60;
}