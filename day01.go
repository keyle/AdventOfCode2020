package main

import "fmt"

/**
--- Day 1: Report Repair ---

After saving Christmas five years in a row, you've decided to take a vacation at a nice resort on a tropical island. Surely, Christmas will go on without you.

The tropical island has its own currency and is entirely cash-only. The gold coins used there have a little picture of a starfish; the locals just call them stars. None of the currency exchanges seem to have heard of them, but somehow, you'll need to find fifty of these coins by the time you arrive so you can pay the deposit on your room.

To save your vacation, you need to get all fifty stars by December 25th.

Collect stars by solving puzzles. Two puzzles will be made available on each day in the Advent calendar; the second puzzle is unlocked when you complete the first. Each puzzle grants one star. Good luck!

Before you leave, the Elves in accounting just need you to fix your expense report (your puzzle input); apparently, something isn't quite adding up.

Specifically, they need you to find the two entries that sum to 2020 and then multiply those two numbers together.

For example, suppose your expense report contained the following:

1721
979
366
299
675
1456
In this list, the two entries that sum to 2020 are 1721 and 299. Multiplying them together produces 1721 * 299 = 514579, so the correct answer is 514579.

Of course, your expense report is much larger. Find the two entries that sum to 2020; what do you get if you multiply them together?

Your puzzle answer was 703131.
*/

var numto2020 = []int{1956, 1994, 457, 1654, 2003, 1902, 1741, 1494, 1597, 1129, 1146, 1589, 1989, 1093, 1881, 1288, 1848, 1371, 1508, 1035, 1813, 1335, 1634, 1102, 1262, 1637, 1048, 1807, 1270, 1528, 1670, 1803, 1202, 1294, 1570, 1640, 1484, 1872, 1140, 1207, 1485, 1781, 1778, 1772, 1334, 1267, 1045, 1194, 1873, 1441, 1557, 1414, 1123, 1980, 1527, 1591, 1665, 1916, 1662, 1139, 1973, 1258, 1041, 1134, 1609, 1554, 1455, 1124, 1478, 1938, 1759, 1281, 1410, 1511, 930, 1319, 1302, 1827, 1216, 1404, 1460, 2002, 1590, 1817, 1341, 1631, 1608, 1382, 1158, 1594, 1049, 1804, 1555, 1753, 447, 1021, 1079, 609, 1766, 1327, 1851, 1052, 1737, 1175, 1043, 1945, 1573, 1113, 1724, 1203, 1856, 1682, 1623, 1135, 1015, 1423, 1412, 1315, 1375, 1895, 1351, 1530, 1758, 1445, 1518, 1819, 1567, 1305, 1919, 1952, 1432, 1099, 1476, 1883, 1871, 1900, 1442, 1393, 1214, 1283, 1538, 1391, 1008, 1109, 1621, 1876, 1998, 1032, 1324, 1927, 481, 1732, 1370, 1683, 1199, 1465, 1882, 1293, 1671, 1456, 1197, 1506, 1381, 1469, 1830, 1957, 1850, 1184, 1564, 1170, 1943, 1131, 1867, 1208, 1788, 1337, 1722, 1760, 1651, 1069, 1574, 1959, 1770, 66, 1190, 1606, 1899, 1054, 980, 1693, 1173, 1479, 1333, 1579, 1720, 1782, 1971, 1438, 1178, 1306}

func day1_part1() {
	for _, a := range numto2020 {
		for _, b := range numto2020 {
			if a+b == 2020 {
				fmt.Println(a, "+", b, "= 2020", a*b)
			}
		}
	}
}

/**
--- Part Two ---

The Elves in accounting are thankful for your help; one of them even offers you a starfish coin they had left over from a past vacation. They offer you a second one if you can find three numbers in your expense report that meet the same criteria.

Using the above example again, the three entries that sum to 2020 are 979, 366, and 675. Multiplying them together produces the answer, 241861950.

In your expense report, what is the product of the three entries that sum to 2020?

Your puzzle answer was 272423970.

Both parts of this puzzle are complete! They provide two gold stars: **
*/

func day1_part2() {
	for _, a := range numto2020 {
		for _, b := range numto2020 {
			for _, c := range numto2020 {
				if a+b+c == 2020 {
					fmt.Println(a, "+", b, "+", c, "= 2020", a*b*c)
				}
			}
		}
	}
}
