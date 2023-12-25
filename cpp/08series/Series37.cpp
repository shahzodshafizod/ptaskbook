#include <iostream>
using namespace std;

int main()
{
	//Task("Series37");
	int K;
	cin >> K;
	int prev, curr, counts = 0;
	for (int i = 0; i < K; i++)
	{
		bool isProgress = true, isRegress = true;
		cin >> prev;
		do
		{
			cin >> curr;
			if (curr == 0)
				break;
			
			if (prev > curr)
				isProgress = false;
			
			if (prev < curr)
				isRegress = false;
			
			prev = curr;
		}
		while (true);
		counts += (int)(isProgress || isRegress);
	}
	cout << counts;
	
	return 0;
}
