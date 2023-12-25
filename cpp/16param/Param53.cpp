#include <iostream>
#include <fstream>
using namespace std;

void SplitIntFile(const char* S0, int K, const char* S1, const char* S2);

int main()
{
	//Task("Param53");
	char fileInName[100], fileOutName1[100], fileOutName2[100];
	cin.getline(fileInName, 100);
	int K;
	cin >> K;
	cin.ignore();
	cin.getline(fileOutName1, 100);
	cin.getline(fileOutName2, 100);
	SplitIntFile(fileInName, K, fileOutName1, fileOutName2);
	
	return 0;
}

void SplitIntFile(const char* S0, int K, const char* S1, const char* S2)
{
	ifstream In(S0, ios::binary);
	ofstream Out1(S1, ios::binary);
	ofstream Out2(S2, ios::binary);
	
	if (In.is_open())
	{
		int number, index = 0;
		while (In.peek() != -1)
		{
			In.read(reinterpret_cast<char*>(&number), sizeof(number));
			if (index < K)
				Out1.write(reinterpret_cast<const char*>(&number), sizeof(number));
			else
				Out2.write(reinterpret_cast<const char*>(&number), sizeof(number));
			index++;
		}
		In.close();
	}
	
	if (Out1.is_open())
		Out1.close();
	if (Out2.is_open())
		Out2.close();
}