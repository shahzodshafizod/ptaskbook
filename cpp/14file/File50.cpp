#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("File50");
	char S1[100], S2[100], S3[100];
	cin.getline(S1, 100);
	cin.getline(S2, 100);
	cin.getline(S3, 100);
	ifstream file1(S1, ios::binary);
	ifstream file2(S2, ios::binary);
	ofstream file3(S3, ios::binary);

	double a, b;
	file1.read(reinterpret_cast<char*>(&a), sizeof(a));
	file2.read(reinterpret_cast<char*>(&b), sizeof(b));

	while (!file1.eof() || !file2.eof())
	{
		if (file1.eof())
		{
			file1.clear();
			file3.write(reinterpret_cast<const char*>(&b), sizeof(b));
			while (true)
			{
				file2.read(reinterpret_cast<char*>(&b), sizeof(b));
				if (file2.eof())
				{
					file2.clear();
					break;
				}
				file3.write(reinterpret_cast<const char*>(&b), sizeof(b));
			}
		}
		else if (file2.eof())
		{
			file2.clear();
			file3.write(reinterpret_cast<const char*>(&a), sizeof(a));
			while (true)
			{
				file1.read(reinterpret_cast<char*>(&a), sizeof(a));
				if (file1.eof())
				{
					file1.clear();
					break;
				}
				file3.write(reinterpret_cast<const char*>(&a), sizeof(a));
			}
		}
		else if (a < b)
		{
			file3.write(reinterpret_cast<const char*>(&a), sizeof(a));
			file1.read(reinterpret_cast<char*>(&a), sizeof(a));
		}
		else
		{
			file3.write(reinterpret_cast<const char*>(&b), sizeof(b));
			file2.read(reinterpret_cast<char*>(&b), sizeof(b));
		}
	}

	file1.close();
	file2.close();
	file3.close();
	
	return 0;
}
