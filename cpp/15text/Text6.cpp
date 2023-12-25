#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("Text6");
	const int stringSize = 80;
	char fileName1[100], fileName2[100], str[stringSize];
	cin.getline(fileName1, 100);
	cin.getline(fileName2, 100);
	ofstream firstFile(fileName1, ios::app);
	ifstream secondFile(fileName2, ios::app);
	while (true)
	{
		secondFile.getline(str, stringSize);
		if (secondFile.eof())
		{
			secondFile.clear();
			break;
		}
		firstFile << str << endl;
	}
	firstFile.close();
	secondFile.close();
	
	return 0;
}
