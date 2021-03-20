#include <iostream>
#include <fstream>

using namespace std;

int main() {

    ifstream iS;
    ofstream oS;

    oS.open("wallets.txt",ios_base::app);

    iS.open("wallets.json");


    if(!iS.is_open() || !oS.is_open()) {

        cout << "ERROR OPENING WALLETS";
        return -1;
    }

    string line;

    while(getline(iS, line)) {


        oS << line << endl;

    }

    oS << endl;

    iS.close();
    oS.close();

}