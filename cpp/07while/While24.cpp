#include <iostream>
using namespace std;

int main()
{
	//Task("While24");
	int N;
	cin >> N;
	int next, prev = 1, curr = 1;
	do
	{
		next = prev + curr;
		prev = curr;
		curr = next;
	}
	while (curr < N);
	cout << (curr == N);
	
	return 0;
}
