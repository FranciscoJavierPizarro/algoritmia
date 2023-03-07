module BinaryTree where
import Data.List

data Tree a = Branch a | Empty

simulameEsta :: Integer -> Integer -> Tree Integer -> Integer
simulameEsta 1 _ (Branch valor) = valor
simulameEsta p n (Branch valor) = simulameEsta p' n' subarbol where
    p' = p - 1
    n' =  ((n `div` 2) :: Integer) + (if n `mod` 2 == 0 then 1 else 0) 
    valorSiguienteNodo = if n `mod` 2 == 1 then 2*valor+1 else 2*valor where
    subarbol = Branch valorSiguienteNodo
simulameEsta p n _ = simulameEsta p' n' subarbol where
    p' = p - 1
    n' = ((n `div` 2) :: Integer) + (if n `mod` 2 == 0 then 1 else 0) 
    valorSiguienteNodo = if n `mod` 2 == 0 then 3 else 2 where
    subarbol = Branch valorSiguienteNodo
