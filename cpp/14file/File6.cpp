#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("File6");
	int K;
	cin >> K;
	cin.ignore();
	char fileName[100];
	cin.getline(fileName, 100);
	int number, numbers = 0;
	ifstream ifs(fileName, ios::binary);
	bool found = false;
	while (true)
	{
		ifs.read( reinterpret_cast<char*>(&number), sizeof(number) );
		if (ifs.eof())
		{
			ifs.clear();
			break;
		}
		
		numbers++;
		if (numbers == K)
		{
			found = true;
			break;
		}
	}
	ifs.close();

	cout << (found ? number : -1);
	
	return 0;
}
