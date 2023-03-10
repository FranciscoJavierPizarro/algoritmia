module BinaryTree where
import Data.List
import Data.Bits

data Tree a = Branch a | Empty

simulameEsta :: Integer -> Integer -> Tree Integer -> Integer
simulameEsta 1 _ (Branch valor) = valor
simulameEsta p n (Branch valor) = simulameEsta p' n' subarbol where
    p' = p - 1
    n' =  ((n `div` 2) :: Integer) + (if n `mod` 2 == 1 then 1 else 0) 
    valorSiguienteNodo = if n `mod` 2 == 0 then 2*valor+1 else 2*valor where
    subarbol = Branch valorSiguienteNodo
simulameEsta p n _ = simulameEsta p' n' subarbol where
    p' = p - 1
    n' = ((n `div` 2) :: Integer) + (if n `mod` 2 == 1 then 1 else 0) 
    valorSiguienteNodo = if n `mod` 2 == 0 then 3 else 2 where
    subarbol = Branch valorSiguienteNodo

aiudaPuto :: Integer -> Integer -> Integer -> Integer
aiudaPuto x y i 
    | y - ((2 :: Integer) `shiftL` (fromInteger (x-4-i))) < 0   =  aiudaPuto x y (i+1)
    | otherwise                                                 = simulamelaDirecta x ( y - (2 :: Integer) `shiftL` fromInteger (x-4-i)) + ((2 :: Integer) `shiftL` fromInteger(i))



simulamelaDirecta :: Integer -> Integer -> Integer
simulamelaDirecta x 0 = ((2 :: Integer) `shiftL` (fromInteger (x-1)))  - 1
simulamelaDirecta x 1 = ((2 :: Integer) `shiftL` (fromInteger (x-2)))
simulamelaDirecta x y
    | y >= ((2 :: Integer) `shiftL` (fromInteger (x-2)))    = simulamelaDirecta x ( y `mod` ((2 :: Integer) `shiftL` (fromInteger (x-2))) )
    -- Cuando hay más bolas de las que caben en el árbol
    | y .&. (y-1) == 0                                      = ((2 :: Integer) `shiftL` (fromInteger (x-1))) - ((2 :: Integer) `shiftL` (fromInteger (x-2-(floor (sqrt (fromIntegral y))))))
    -- Cuando la bola es un número cuadrado de 2
    | (y+1) .&. y == 0                                      = simulamelaDirecta x (y+1) - ((2 :: Integer) `shiftL` (fromInteger (x-3)))
    -- Cuando la bola más 1 es un número cuadrado de 2
    | y > ((2 :: Integer) `shiftL` (fromInteger (x-3)))     = simulamelaDirecta x ( y - ((2 :: Integer) `shiftL` (fromInteger (x-3))) ) + 1
    -- Cuando la bola pasa de la mitad del número de hojas del árbol
    | y `mod` 2 == 0                                        = simulamelaDirecta x ( y-1 ) + ((2 :: Integer) `shiftL` (fromInteger (x-3)))
    -- Cuando la bola es un número par
    | (y-1) .&. (y-2) == 0                                  = simulamelaDirecta x 1 + ((2 :: Integer) `shiftL` (fromInteger (x-3-(fromIntegral (round (logBase 2 (fromIntegral (y-1))))))))
    -- Cuando la bola menos 1 es un número cuadrado de 2
    | otherwise                                             = aiudaPuto x y 0
    -- En cualqiuier otro caso





simulameEstaOptimizada :: Integer -> Integer -> Tree Integer -> Integer
simulameEstaOptimizada 1 _ (Branch valor) = valor
simulameEstaOptimizada p n (Branch valor) = simulameEstaOptimizada p' n' subarbol where
    p' = p - 1
    valorMod = (n .&. 1) == 1
    n' =  ((n `shiftR` 1) :: Integer) + (if valorMod then 1 else 0) 
    valorSiguienteNodo = if not valorMod then 2*valor+1 else 2*valor where
    subarbol = Branch valorSiguienteNodo
simulameEstaOptimizada p n _ = simulameEstaOptimizada p' n' subarbol where
    p' = p - 1
    valorMod = (n .&. 1) == 1
    n' = ((n `shiftR` 1) :: Integer) + (if valorMod then 1 else 0) 
    valorSiguienteNodo = if not valorMod then 3 else 2 where
    subarbol = Branch valorSiguienteNodo