#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("File15");
	char fileName[100];
	cin.getline(fileName, 100);
	ifstream ifs(fileName, ios::binary);
	double number, sum = 0;
	int raqam = 0;
	while (true)
	{
		ifs.read(reinterpret_cast<char*>(&number), sizeof(number));
		if (ifs.eof())
		{
			ifs.clear();
			break;
		}
		raqam++;
		if (raqam % 2 == 0)
			sum += number;
	}
	cout << sum;
	
	return 0;
}
