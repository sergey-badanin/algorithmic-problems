package sequentialpages

import "strconv"

/*
 * There is a book with pages numbered from 1 to some unknown number n.
 * It turns out that you can write down the page numbers, without any spaces between them,
 * to make a seemingly meaningless long number that nonetheless comes directly from the page numbering.
 */

/*
 * Suppose you didn't know how many pages were in the book, but you did know how long the number was — it has 990 digits.
 * Find out n — the number of pages in the book
 */
func digitsNumberToPage(num int) int {
	page := 1
	digits := 1
	for digits < num {
		page++
		str := strconv.Itoa(page)
		digits += len(str)
	}
	return page
}

/*
 * You don't know anything else about the book, but you do know that when you look at the number,
 * it has 260 '1's in it. What's the maximum number of pages in the book?
 */
func maxEncountersOfDigit(digit byte, encounter int) int {
	page := 1
	digits := 1
	for digits <= encounter {
		page++
		str := strconv.Itoa(page)
		for i := 0; i < len(str); i++ {
			if str[i] == digit {
				digits++
			}
		}
	}

	if digits > encounter {
		return page - 1
	}
	return page
}
