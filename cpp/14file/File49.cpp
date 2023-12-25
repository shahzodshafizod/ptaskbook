#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("File49");
	char S[4][100], SE[100];
	ifstream fileIn[4];
	for (int i = 0; i < 4; i++)
	{
		cin.getline(S[i], 100);
		fileIn[i].open(S[i], ios::binary);
	}
	cin.getline(SE, 100);
	ofstream fileOut(SE, ios::binary);
	int numbers[4];
	bool exit = false;
	while (true)
	{
		for (int i = 0; i < 4; i++)
		{
			fileIn[i].read(reinterpret_cast<char*>(&numbers[i]), sizeof(numbers[i]));
			if (fileIn[i].eof())
			{
				fileIn[i].clear();
				exit = true;
				break;
			}
		}
		if (exit)
			break;
		for (int i = 0; i < 4; i++)
			fileOut.write(reinterpret_cast<const char*>(&numbers[i]), sizeof(numbers[i]));
	}
	for (int i = 0; i < 4; i++)
		fileIn[i].close();
	fileOut.close();
	
	return 0;
}
