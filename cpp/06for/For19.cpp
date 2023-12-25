#include <iostream>
using namespace std;

int main()
{
	//Task("For19");
	int N;
	cin >> N;
	double fact = 1;
	for (int i = 1; i <= N; i++)
		fact *= i;
	cout << fact;
	
	return 0;
}
