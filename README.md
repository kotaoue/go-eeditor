# go-eeditor
Run an external editor and gets the inputs string like a behavior of 'git -commit'.  
Create and delete a temporary file for passing the inputs string.

## Usage
### Run vim
```Go
import (
	"fmt"

	"github.com/kotaoue/go-eeditor"
)

func main() {
	editor := eeditor.NewEditor()
	b, _ := editor.Open()
	fmt.Println(string(b))
}
```

### Run Emacs
```Go
editor := eeditor.NewEditor(eeditor.Command("emacs"))
```

### Change Path
```Go
editor := eeditor.NewEditor(eeditor.Path("./temporary"))
```

## Referenced links
* [GoでVimを開いて編集内容をパースする方法](https://qiita.com/lighttiger2505/items/d3b9ee9884c75a7819d8)
  * [lighttiger2505/launch-vim](https://github.com/lighttiger2505/launch-vim)