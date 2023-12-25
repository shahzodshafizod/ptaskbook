#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("Text2");
	char fileName[100];
	cin.getline(fileName, 100);
	int N;
	cin >> N;
	ofstream ofs(fileName);
	char C;
	for (int i = 0; i < N; i++)
	{
		C = 'a';
		for (int j = 0; j <= i; j++, C++)
			ofs << C;
		if (i < N-1)
			ofs << endl;
	}
	ofs.close();
	
	return 0;
}
