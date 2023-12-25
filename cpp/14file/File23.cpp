#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("File23");
	char fileInName[100], fileOutName[100];
	cin.getline(fileInName, 100);
	cin.getline(fileOutName, 100);
	ifstream ifs(fileInName, ios::binary);
	ofstream ofs(fileOutName, ios::binary);
	double a, b;
	ifs.read(reinterpret_cast<char*>(&a), sizeof(a));
	int length = 1;
	while (true)
	{
		ifs.read(reinterpret_cast<char*>(&b), sizeof(b));
		if (ifs.eof())
		{
			ifs.clear();
			break;
		}
		if (a > b)
			length++;
		else
		{
			if (length > 1)
				ofs.write(reinterpret_cast<const char*>(&length), sizeof(length));
			length = 1;
		}
		a = b;
	}
	if (length > 1)
		ofs.write(reinterpret_cast<const char*>(&length), sizeof(length));
	ifs.close();
	ofs.close();
	
	return 0;
}
