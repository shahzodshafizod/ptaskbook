#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("File17");
	char fileName[100], fileName2[100];
	cin.getline(fileName, 100);
	cin.getline(fileName2, 100);
	ifstream ifs(fileName, ios::binary);
	ofstream ofs(fileName2, ios::binary);
	int a, b, length = 0;
	ifs.read(reinterpret_cast<char*>(&a), sizeof(a));
	length++;
	while (true)
	{
		ifs.read(reinterpret_cast<char*>(&b), sizeof(b));
		if (ifs.eof())
		{
			ifs.clear();
			break;
		}
		if (a == b)
			length++;
		else
		{
			ofs.write(reinterpret_cast<const char*>(&length), sizeof(length));
			length = 1;
		}
		a = b;
	}
	ofs.write(reinterpret_cast<const char*>(&length), sizeof(length));
	ifs.close();
	ofs.close();
	
	return 0;
}
