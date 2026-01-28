module Test exposing (..)

estVide lst = case lst of
   [] -> True
   (x :: xs) -> False

addElemInList elem n lst =
    if n <= 0 then 
        lst
    else 
        addElemInList elem (n - 1) (elem :: lst)

dupli lst = case lst of 
    [] -> 
        []
    x :: xs -> 
        x :: x :: dupli xs 

compress lst = case lst of 
    [] -> 
        []
    [x] ->
        [x]
    x :: y :: xs -> 
        if x == y then
            x :: compress xs
        else 
            x :: y :: compress xs

addElemInListnonrec elem n lst =
    List.repeat n elem ++ lst

compressHelper x partialRes = case partialRes of
  [] -> [x]
  (y :: ys) -> if x == y
               then partialRes
               else x :: partialRes 

compressnonrec lst = 
    List.foldr compressHelper [] lst


        