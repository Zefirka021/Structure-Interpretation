#include <iostream>
using namespace std;

class Point {
private:
    int x, y;

public:
    Point() {
        x = 0;
        y = 0;
    }

    Point(int x, int y) {
        this->x = x;
        this->y = y;
    }

    Point operator+(const Point& p) {
        Point sum;
        sum.x = this->x + p.x;
        sum.y = this->y + p.y;
        return sum;
    }

    void display() {
        cout << "x: " << x << ", y: " << y << endl;
    }
};

int main() {
    Point p1(3, 4);
    Point p2(1, 2);
    Point p3 = p1 + p2;

    p1.display();
    p2.display();
    p3.display();

    return 0;
}