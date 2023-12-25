#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("Text1");
	char fileName[100];
	cin.getline(fileName, 100);
	int N, K;
	cin >> N >> K;
	ofstream ofs(fileName);
	for (int i = 0; i < N; i++)
	{
		for (int j = 0; j < K; j++)
			ofs << '*';
		if (i < N-1)
			ofs << endl;
	}
	ofs.close();
	
	return 0;
}
