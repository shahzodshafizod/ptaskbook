#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("File5");
	char fileName[100];
	cin.getline(fileName, 100);
	ifstream ifs(fileName, ios::binary);
	int numbers = -1;
	if (ifs.is_open())
	{
		numbers++;
		int number;
		while (true)
		{
			ifs.read( reinterpret_cast<char*>(&number), sizeof(number) );
			if (ifs.eof())
			{
				ifs.clear();
				break;
			}
			numbers++;
		}
		ifs.close();
	}
	cout << numbers;
	
	return 0;
}
