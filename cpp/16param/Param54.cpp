#include <iostream>
#include <fstream>
using namespace std;

void SplitText(const char* S0, int K, const char* S1, const char* S2);

int main()
{
	//Task("Param54");
	char InName[100], OutName1[100], OutName2[100];
	cin.getline(InName, 100);
	int K;
	cin >> K;
	cin.ignore();
	cin.getline(OutName1, 100);
	cin.getline(OutName2, 100);
	SplitText(InName, K, OutName1, OutName2);
	
	return 0;
}

void SplitText(const char* S0, int K, const char* S1, const char* S2)
{
	ifstream In(S0);
	ofstream Out1(S1);
	ofstream Out2(S2);

	if (In.is_open())
	{
		const int stringSize = 100;
		char str[stringSize+1];
		int index = 0;
		while (In.peek() != -1)
		{
			In.getline(str, stringSize+1);
			if (index < K)
				Out1 << str << endl;
			else
				Out2 << str << endl;
			
			index++;
		}
		In.close();
	}

	if (Out1.is_open())
		Out1.close();
	if (Out2.is_open())
		Out2.close();
}