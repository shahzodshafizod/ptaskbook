#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("Text3");
	char fileName[100];
	cin.getline(fileName, 100);
	int N;
	cin >> N;
	ofstream ofs(fileName);
	char C;
	for (int i = 0; i < N; i++)
	{
		C = 'A';
		int j = 0;
		for (; j <= i; j++, C++)
			ofs << C;
		for (; j < N; j++)
			ofs << '*';
		if (i < N-1)
			ofs << endl;
	}
	ofs.close();
	
	return 0;
}
