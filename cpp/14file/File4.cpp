#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("File4");
	char fileName[100];
	int counts = 0;
	ifstream ifs;
	for (int i = 0; i < 4; i++)
	{
		cin.getline(fileName, 100);
		ifs.open(fileName);
		
		if (ifs.is_open())
		{
			ifs.close();
			counts++;
		}
	}
	cout << counts;
	
	return 0;
}
