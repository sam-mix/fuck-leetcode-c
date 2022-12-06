package lib

import "bytes"

// func convert(s string, numRows int) string {
// 	if numRows < 2 {
// 		return s
// 	}
// 	sLen := len(s)
// 	result := make([]byte, sLen)
// 	sepMap := make(map[int]int)
// 	for i := sLen - numRows; i <= numRows; i++ {
// 		if i < 0 {
// 			continue
// 		}
// 		sep := int(math.Ceil(float64(i) / float64(numRows)))
// 		fmt.Println(i, sep)
// 		sepMap[i%numRows] = sep
// 	}
// 	for i, v := range []byte(s) {
// 		row := i % numRows
// 		sep, _ := sepMap[row]
// 		result[row*sep+i/numRows] = v
// 	}
// 	return string(result)
// }

// func convert(s string, numRows int) string {
// 	if numRows < 2 {
// 		return s
// 	}
// 	m := make(map[int][]byte)
// 	for i, v := range []byte(s) {
// 		row := i % numRows
// 		if l, ok := m[row]; ok {
// 			l = append(l, v)
// 			m[row] = l
// 		} else {
// 			l := []byte{v}
// 			m[row] = l
// 		}
// 	}
// 	result := make([]byte, 0)
// 	for i := 0; i < numRows; i++ {
// 		if l, ok := m[i]; ok {
// 			result = append(result, l...)
// 		}
// 	}
// 	return string(result)
// }

// func convert(s string, numRows int) string {
// 	len := len(s)
// 	if numRows < 2 || len < numRows {
// 		return s
// 	}
// 	rows := make([]string, numRows)
// 	n := 2*numRows - 2
// 	for i, char := range s {
// 		x := i % n
// 		rows[min(x, n-x)] += string(char)
// 	}
// 	return strings.Join(rows, "")
// }

func convert(s string, numRows int) string {
	r := numRows
	if r < 2 || r >= len(s) {
		return s
	}
	mat := make([][]byte, r)
	t, x := r*2-2, 0
	for i, ch := range s {
		mat[x] = append(mat[x], byte(ch))
		if i%t < r-1 {
			x++
		} else {
			x--
		}
	}
	return string(bytes.Join(mat, nil))
}

/*
==================================
my_col = index % 2

sep = (n + 1 / 2) = 4

list : 			1 2 3 4 5 6 7
col  : 			2
index: 			0 1 2 3 4 5 6
tl   : 			1 3 5 7 2 4 6
old_list_index: 0 4 1 5 2 6 3
old_list_col:   0 1 0 1 0 1 0

matrix_index:
0   2   4   6
  1   3   5

==================================
ni: new index
i: 	old index
sep: num per col
len: all member number
sep = len / col
ni = (i%col) * sep + floor(i/col)

==================================

sep = 7 / 2 = 3
col = 2
ni list: 0 4 1 5 2 6 3
0 = (0 % 2) * 4 + floor(0 / 2)
4 = (1 % 2) * 4 + floor(1 / 2)
1 = (2 % 2) * 4 + floor(2 / 2)
5 = (3 % 2) * 4 + floor(3 / 2)
2 = (4 % 2) * 4 + floor(4 / 2)
6 = (5 % 2) * 4 + floor(5 / 2)
3 = (6 % 2) * 4 + floor(6 / 2)

==================================


==================================

*/
