/*
Package t443

Given an array of characters chars, compress it using the following algorithm:
Дан массив прописных символов, упакуй ее используя следующий алгоритм:

Begin with an empty string s. For each group of consecutive repeating characters in chars:
Начни с пустой строки s. Для каждой группы последовательно повторяющихся символов в строке:

If the group's length is 1, append the character to s.
Если длинна группы 1, добавь сивол в s.

Otherwise, append the character followed by the group's length.
В противном случае, добавьте символ за которым следует длина группы.

The compressed string s should not be returned separately, but instead, be stored in the input character array chars.
Note that group lengths that are 10 or longer will be split into multiple characters in chars.

Сжатаю строку s не следует возвращать отдельно, но вместо этого, хранить в входмно массиве символов.
Примечание: что группы длин которые составляют 10 или более, будут разделены на несколько сиволов в символах.

After you are done modifying the input array, return the new length of the array.
После того как ты модифицируешь входной массив, верни новую длинну массива.

You must write an algorithm that uses only constant extra space.
Ты должне написать алгоритм который использует только постоянное дополнительное пространство.

Note: The characters in the array beyond the returned length do not matter and should be ignored.
Примечание: Символы в массиве за пределами вовращаемой длины не имеют значения и должны быть игнорированы.
*/
package t443

import "strconv"

func compress(chars []byte) int {
	n := len(chars)
	write := 0
	read := 0

	for read < n {
		ch := chars[read]
		count := 0

		for read < n && chars[read] == ch {
			read++
			count++
		}

		chars[write] = ch
		write++

		if count > 1 {
			numStr := strconv.Itoa(count)
			for i := 0; i < len(numStr); i++ {
				chars[write] = numStr[i]
				write++
			}
		}
	}

	return write
}

