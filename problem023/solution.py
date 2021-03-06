#!/usr/bin/python3

"""
Problem 23: Multiplication Tables

Challenge Description:

Print out the grade school multiplication table upto 12*12.

Input sample:

None

Output sample:

Print out the table in a matrix like fashion, each number formatted to a width of 4 (The numbers are right-aligned and strip out leadeing/trailing spaces on each line). The first 3 line will look like: 
e.g.

1   2   3   4   5   6   7   8   9  10  11  12
2   4   6   8  10  12  14  16  18  20  22  24
3   6   9  12  15  18  21  24  27  30  33  36
"""

def main():
    """Main function!

    Prints the multiplication table.

    """
    for i in range(1, 13):
        for multiplier in range(1, 13):
            result = str(i * multiplier)
            if multiplier == 1:
                print(result, end="")
            # elif multiplier == 2 and i > 9:
            #     print(" " + result, end="")
            elif len(result) < 2:
                print("   " + result, end="")
            elif len(result) < 3:
                print("  " + result, end="")
            else:
                print(" " + result, end="")
        print()

if __name__ == "__main__":
    main()
