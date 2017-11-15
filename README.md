# confspace

to write config files under `~/.myproject` directory.

```go
	space, _ := New("myproject")
	f, _ := space.Open("config.json")

	fmt.Println(f.Name())
	// /Users/otiai10/.myproject/config.json
```
