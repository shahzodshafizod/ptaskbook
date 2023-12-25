#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("File52");
	char S0[100], S[100];
	cin.getline(S0, 100);
	int N;
	cin >> N;
	cin.ignore();

	ofstream fileOut(S0, ios::binary);
	ifstream* filesIn = new ifstream [N];
	fileOut.write(reinterpret_cast<const char*>(&N), sizeof(N));
	int fileSize;
	for (int i = 0; i < N; i++)
	{
		cin.getline(S, 100);
		filesIn[i].open(S, ios::binary);
		filesIn[i].seekg(0, ios::end);
		fileSize = filesIn[i].tellg() / sizeof(int);
		fileOut.write(reinterpret_cast<const char*>(&fileSize), sizeof(fileSize));
	}
	int number;
	for (int i = 0; i < N; i++)
	{
		filesIn[i].seekg(0);
		while (true)
		{
			filesIn[i].read(reinterpret_cast<char*>(&number), sizeof(number));
			if (filesIn[i].eof())
			{
				filesIn[i].clear();
				break;
			}
			fileOut.write(reinterpret_cast<const char*>(&number), sizeof(number));
		}
		filesIn[i].close();
	}
	delete [] filesIn;
	filesIn = NULL;
	
	return 0;
}
