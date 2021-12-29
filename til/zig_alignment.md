```yaml
- align()
- @alignOf()
- @alignCast()
- @ptrCast()
```

```zig
const Foo = struct {
  a: i32,
  b: i32,
};

pub fn main() {
  var array = []u8{1} ** 1024;
  const foo = @ptrCast(&Foo, &array[0]);
  foo.a += 1;
}
```

```sh
/home/andy/tmp/test.zig:8:17: error: cast increases pointer alignment
    const foo = @ptrCast(&Foo, &array[0]);
                ^
/home/andy/tmp/test.zig:8:38: note: '&u8' has alignment 1
    const foo = @ptrCast(&Foo, &array[0]);
                                     ^
/home/andy/tmp/test.zig:8:27: note: '&Foo' has alignment 4
    const foo = @ptrCast(&Foo, &array[0]);
                          ^
```

```diff
@@ -4,7 +4,7 @@
 };

 pub fn main() {
-    var array = []u8{1} ** 1024;
+    var array align(@alignOf(Foo)) = []u8{1} ** 1024;
     const foo = @ptrCast(&Foo, &array[0]);
     foo.a += 1;
 }
```
