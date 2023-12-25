#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("File63");
	const int chars = 80;
	int K, length;
	cin >> K;
	cin.ignore();
	char C, fileName[100], fileName1[100], fileName2[100], str[chars];
	cin.getline(fileName, 100);
	cin.getline(fileName1, 100);
	cin.getline(fileName2, 100);
	ifstream ifs(fileName, ios::binary);
	ofstream ofs1(fileName1, ios::binary);
	ofstream ofs2(fileName2, ios::binary);
	while (true)
	{
		ifs.read(reinterpret_cast<char*>(&str), sizeof(str));
		if (ifs.eof())
		{
			ifs.clear();
			break;
		}
		
		length = strlen(str);
		if (K <= length)
		{
			length = K;
			str[length] = '\0';
			C = str[length-1];
		}
		else
			C = ' ';

		ofs1.write(reinterpret_cast<const char*>(&str), sizeof(str));
		ofs2.write(reinterpret_cast<const char*>(&C), sizeof(C));
	}
	ifs.close();
	ofs1.close();
	ofs2.close();
	
	return 0;
}
