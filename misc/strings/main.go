package main


import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Contains \t", strings.Contains("abcdefgh", "de"))
	fmt.Println("Count \t", strings.Count("abcdefghabcdefghabc", "abc"))
	fmt.Println("HasPrefix \t", strings.HasPrefix("abcdefgh", "abc"))
	fmt.Println("HasSuffix \t", strings.HasSuffix("abcdefgh", "fgh"))
	fmt.Println("Index \t", strings.Index("abcdefgh", "d"))
	fmt.Println("Join \t", strings.Join([]string{"abc", "def", "ghi"}, "-"))
	fmt.Println("Repeat \t", strings.Repeat("abc", 5))
	fmt.Println("Replace \t", strings.Replace("foo","o", "0", -1))
	fmt.Println("Replace \t", strings.Replace("foo","o", "0", 1))
	fmt.Println("Split \t", strings.Split("abc|def|ghi", "|"))
	fmt.Println("ToLower \t", strings.ToLower("AbCdEfGhI"))
	fmt.Println("ToUpper \t", strings.ToUpper("abcdefghi"))
	fmt.Println("Trim \t", strings.Trim("! abc def ghi !", " !"))
	fmt.Println("TrimSpace \t", strings.TrimSpace("   abcd efgh    "))
	fmt.Println("Fields \t", strings.Fields("dwdwd   dwqdwqd qdqd    dwqdqw"))
	fmt.Println(string(rune("abcdefgh"[5])))


	builder := &strings.Builder{}

	builder.WriteString("abc");
	builder.WriteString("def");
	builder.WriteString("ghi");

	fmt.Println(builder.String())

}