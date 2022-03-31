package home

import (
	"strings"
	"fmt"
)

// vars := "\"my\""

func DoubleQuote(text string) string {
	brks := strings.Split(text, " ")


	for i, k := range brks {
		brks[i] = strings.Replace(strings.ToLower(k), `"`, "-----", -1)
	}

	for i, k := range brks {
		brks[i] = strings.Replace(strings.ToLower(k), "\n", "......", -1)
	}

	fmt.Println(brks)

	return strings.Join(brks, " ")
}

func MoldVariable(text string) string {

	newline := strings.ReplaceAll(text, "\n", "<br/>")
	space := strings.ReplaceAll(newline, " ", "&nbsp;")
	brks := strings.Split(space, " ")

	for i, k := range brks {
		brks[i] = strings.Replace(strings.ToLower(k), "-----", `"`, -1)
	}

	for i, k := range brks {
		brks[i] = strings.Replace(strings.ToLower(k), ">>>", "<br/>", -1)
	}
	
	for i, k := range brks {
		brks[i] = strings.Replace(strings.ToLower(k), "print", " <span style='color:green;' class='log'>print</span> ", -1)
	}

	for i, k := range brks {
		brks[i] = strings.Replace(strings.ToLower(k), "bool", " <span style='color:green;' class='log'>bool</span> ", -1)
	}

	for i, k := range brks {
		brks[i] = strings.Replace(strings.ToLower(k), "var", " <span style='color:blue;'  class='params'>var</span> ", -1)
	}

	for i, k := range brks {
		brks[i] = strings.Replace(strings.ToLower(k), "string", " <span style='color:blue;'  class='params'>string</span> ", -1)
	}


	for i, k := range brks {
	   brks[i] = strings.Replace(strings.ToLower(k), "function", " <span style='color:red;' class='function'>function</span> ", -1)
	}

	
	for i, k := range brks {
		brks[i] = strings.Replace(strings.ToLower(k), "range", " <span style='color:red;' class='function'>range</span> ", -1)
	}

	for i, k := range brks {
		brks[i] = strings.Replace(strings.ToLower(k), "append", " <span style='color:red;' class='function'>append</span> ", -1)
	 }

	for i, k := range brks {
		brks[i] = strings.Replace(strings.ToLower(k), "map", " <span style='color:red;' class='function'>map</span> ", -1)
	 }

	for i, k := range brks {
	   brks[i] = strings.Replace(strings.ToLower(k), "def", " <span style='color:red;' class='function'>def</span> ", -1)
	}


	for i, k := range brks {
		brks[i] = strings.Replace(strings.ToLower(k), "return", " <span style='color:blue;'  class='params'>return</span> ", -1)
	}

	for i, k := range brks {
		brks[i] = strings.Replace(strings.ToLower(k), "if", " <span style='color:purple;' class='select'>if</span> ", -1)
	}

	for i, k := range brks {
		brks[i] = strings.Replace(strings.ToLower(k), "else", " <span style='color:purple;'  class='select'>else</span> ", -1)
	}


	for i, k := range brks {
		brks[i] = strings.Replace(strings.ToLower(k), "for", " <span style='color:purple;'  class='select'>for</span> ", -1)
	}

	for i, k := range brks {
		brks[i] = strings.Replace(strings.ToLower(k), "while", " <span style='color:purple;'  class='select'>while</span> ", -1)
	}

	for i, k := range brks {
		brks[i] = strings.Replace(strings.ToLower(k), "console.log", " <span style='color:blue;'  class='params'>console.log</span> ", -1)
	}

	return "<div class='code-box'><span style='color:white;'>"+strings.Join(brks, " ")+"</span></div>"

}