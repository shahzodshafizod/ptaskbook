#include <iostream>
using namespace std;

int main()
{
	//Task("While27");
	int N;
	cin >> N;
	int prev, curr, next;
	prev = curr = 1;
	int index = 2;
	do
	{
		index++;
		next = prev + curr;
		prev = curr;
		curr = next;
	}
	while (curr < N);
	cout << index;
	
	return 0;
}
