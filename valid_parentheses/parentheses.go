package parentheses

/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']',
determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.

Example 1:

Input: s = "()"
Output: true

Example 2:

Input: s = "()[]{}"
Output: true

Example 3:

Input: s = "(]"
Output: false

*/

/*
	parenthesis: ()
	bracket: []
	brace: {}
*/
// ASCII mini-table
const (
	OPENING_PARENTHESIS = 0x28
	CLOSING_PARENTHESIS = 0x29
	OPENING_BRACKET     = 0x5B
	CLOSING_BRACKET     = 0x5D
	OPENING_BRACE       = 0x7B
	CLOSING_BRACE       = 0x7D
)

type lifo struct {
	queue []byte
}

func NewLifo() lifo {
	return lifo{
		queue: make([]byte, 0),
	}
}

func (l *lifo) Push(v byte) {
	l.queue = append(l.queue, v)
}

func (l *lifo) Pull() {
	l.queue = l.queue[:len(l.queue)-1]
}

func (l *lifo) GetPrevious() (byte, bool) {
	qLen := len(l.queue)
	if qLen == 0 {
		return 0, false
	}
	return l.queue[qLen-1], true
}

func isValidMyAlgo(s string) bool {
	stack := NewLifo()
	parenCount := 0
	bracketCount := 0
	braceCount := 0

	if s == "" {
		return false
	}

	for _, ch := range s {
		switch ch {
		case OPENING_PARENTHESIS:
			stack.Push(OPENING_PARENTHESIS)
			parenCount++
		case OPENING_BRACKET:
			stack.Push(OPENING_BRACKET)
			bracketCount++
		case OPENING_BRACE:
			stack.Push(OPENING_BRACE)
			braceCount++
		case CLOSING_PARENTHESIS:
			prev, exists := stack.GetPrevious()
			if !exists || prev != OPENING_PARENTHESIS {
				return false
			}
			stack.Pull()
			parenCount--
		case CLOSING_BRACKET:
			prev, exists := stack.GetPrevious()
			if !exists || prev != OPENING_BRACKET {
				return false
			}
			stack.Pull()
			bracketCount--
		case CLOSING_BRACE:
			prev, exists := stack.GetPrevious()
			if !exists || prev != OPENING_BRACE {
				return false
			}
			stack.Pull()
			braceCount--
		}
	}

	return parenCount == 0 && bracketCount == 0 && braceCount == 0
}
