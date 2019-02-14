# Tree for ordering parent/child relationships

```go
//    A  B
//   C  D
//  E  F
// G H  I
t := tree.New()
t.AddParents(StringItem("G"), StringItem("E"))
t.AddParents(StringItem("H"), StringItem("E"), StringItem("F"))
t.AddParents(StringItem("I"), StringItem("F"))
t.AddParents(StringItem("E"), StringItem("C"))
t.AddParents(StringItem("F"), StringItem("D"))
t.AddParents(StringItem("C"), StringItem("A"))
t.AddParents(StringItem("D"), StringItem("B"))
fmt.Println(t.Sorted()) // A, B, C, D, E, F, G, H, I
```