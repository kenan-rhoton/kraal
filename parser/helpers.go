package parser

var alpha = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
var num = "0123456789"

func isText(r rune) bool {
	return r != '<'
}

func is(r rune, list []rune) bool {
	for _, l := range list {
		if r == l {
			return true
		}
	}
	return false
}

func isTagName(r rune) bool {
	return is(r, []rune(":"+alpha+num))
}

func isAttrName(r rune) bool {
	return is(r, []rune(alpha))
}

func isAttrValue(r rune) bool {
	return r != '"'
}
