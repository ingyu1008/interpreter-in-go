# My Language (G)

The name of the programming language here is named after my first name In`G`yu

# Syntax of G

Basic syntax of G is somewhat similar to C++ but still have some unique features.

For instance, heres few lines of code in G
```
var x = -1
var b = true

fn gcd(x, y){
    if(x == 0){
        return y
    }
    return gcd(y%x, x)
}
```

From this code snippet, we see following properties of G

- Define variables with keyword `var` and no explicit type is used.
- Does not have semicolon in the end of the line.
- Keyword `fn` is used to define functions. Here too, no explicit return value is needed.
- Recursive functions can be used.


This part of the project can be revised as I build my interpreter.