# No More `if err != nil` in main code?

| Pros              | Cons |
| :---------------- | :------: |
|  Concise error handling	     |   Panics are disruptive and final   |
| Useful for prototyping           |   Not idiomatic for production Go   |
| Works with any type    |  No error recovery mechanism   |
