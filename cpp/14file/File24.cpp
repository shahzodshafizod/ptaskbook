#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("File24");
	char fileInName[100], fileOutName[100];
	cin.getline(fileInName, 100);
	cin.getline(fileOutName, 100);
	ifstream ifs(fileInName, ios::binary);
	ofstream ofs(fileOutName, ios::binary);
	double a, b;
	ifs.read(reinterpret_cast<char*>(&a), sizeof(a));
	int afz = 1, kam = 1;
	while (true)
	{
		ifs.read(reinterpret_cast<char*>(&b), sizeof(b));
		if (ifs.eof())
		{
			ifs.clear();
			break;
		}
		if (a < b)
		{
			afz++;
			if (kam > 1)
			{
				ofs.write(reinterpret_cast<const char*>(&kam), sizeof(kam));
				kam = 1;
			}
		}
		else if (a > b)
		{
			kam++;
			if (afz > 1)
			{
				ofs.write(reinterpret_cast<const char*>(&afz), sizeof(afz));
				afz = 1;
			}
		}
		a = b;
	}
	if (afz > 1)
		ofs.write(reinterpret_cast<const char*>(&afz), sizeof(afz));
	if (kam > 1)
		ofs.write(reinterpret_cast<const char*>(&kam), sizeof(kam));
	ifs.close();
	ofs.close();
	
	return 0;
}
