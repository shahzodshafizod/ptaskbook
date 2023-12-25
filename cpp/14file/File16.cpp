#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("File16");
	char fileName[100];
	cin.getline(fileName, 100);
	ifstream ifs(fileName, ios::binary);
	int a, b, series = 0;
	ifs.read(reinterpret_cast<char*>(&a), sizeof(a));
	series++;
	while (true)
	{
		ifs.read(reinterpret_cast<char*>(&b), sizeof(b));
		if (ifs.eof())
		{
			ifs.clear();
			break;
		}
		if (b != a)
			series++;
		a = b;
	}
	cout << series;
	
	return 0;
}
