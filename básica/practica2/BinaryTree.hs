--------------------------------------------------------------------------------
--                                                                            --
--    Archivo: BinaryTree.hs                                                  --
--    Fecha de última revisión: 14/03/2023                                    --
--    Autores: Francisco Javier Pizarro 821259                                --
--             Jorge Solán Morote   	816259                                --
--    Comms:                                                                  --
--          Este archivo contiene los algoritmos implementados en haskell     --
--          El primero de ellos simula el recorrido solo de la última bola a  --
--          lo largo de P, el segundo emplea propiedades matematicas del      --
--          problema para obtener una solución mucho mas eficiente.           --
--                                                                            --
--------------------------------------------------------------------------------
module BinaryTree where
import Data.List
import Data.Sequence
import Data.Bits

data Tree a = Branch a | Empty

--Simula el recorrido de la última bola a lo largo de P, para entender como funciona tratar de ver sobre
--papel como resulta el tirar N bolas por un nodo de cara a sus dos nodos hijos
simularConRecorridoParcial :: Integer -> Integer -> Tree Integer -> Integer
simularConRecorridoParcial 1 _ (Branch valor) = valor -- caso base
simularConRecorridoParcial p n (Branch valor) = simularConRecorridoParcial p' n' subarbol where -- caso recursivo
    p' = p - 1
    valorMod = (n `mod` 2) == 1
    nuevoValor = 2*valor
    n' =  ((n `div` 2) :: Integer) + (if valorMod then 1 else 0) 
    valorSiguienteNodo = if not valorMod then nuevoValor+1 else nuevoValor where
    subarbol = Branch valorSiguienteNodo
simularConRecorridoParcial p n _ = simularConRecorridoParcial p' n' subarbol where -- caso base de partida (_ == Empty)
    p' = p - 1
    valorMod = (n `mod` 2) == 1
    n' = ((n `div` 2) :: Integer) + (if valorMod then 1 else 0) 
    valorSiguienteNodo = if not valorMod then 3 else 2 where
    subarbol = Branch valorSiguienteNodo


--Aprovecha propiedades matemáticas del problema para resolverlo de forma mucho mas eficiente
simulacionDirecta :: Integer -> Integer -> Integer
simulacionDirecta x 0 = ((2 :: Integer) `shiftL` (fromInteger (x-1)))  - 1
simulacionDirecta x 1 = ((2 :: Integer) `shiftL` (fromInteger (x-2)))
simulacionDirecta x y
    | y >= ((2 :: Integer) `shiftL` (fromInteger (x-2)))    = simulacionDirecta x ( y `mod` ((2 :: Integer) `shiftL` (fromInteger (x-2))) )
    -- Cuando hay más bolas de las que caben en el árbol
    | y .&. (y-1) == 0                                      = ((2 :: Integer) `shiftL` (fromInteger (x-1))) - ((2 :: Integer) `shiftL` (fromInteger (x-2-(fromIntegral (floor (logBase 2 (fromIntegral y)))))))
    -- Cuando la bola es un número cuadrado de 2
    | (y+1) .&. y == 0                                      = simulacionDirecta x (y+1) - ((2 :: Integer) `shiftL` (fromInteger (x-3)))
    -- Cuando la bola más 1 es un número cuadrado de 2
    | y > ((2 :: Integer) `shiftL` (fromInteger (x-3)))     = simulacionDirecta x ( y - ((2 :: Integer) `shiftL` (fromInteger (x-3)))) + 1
    -- Cuando la bola pasa de la mitad del número de hojas del árbol
    | y `mod` 2 == 0                                        = simulacionDirecta x ( y-1 ) + ((2 :: Integer) `shiftL` (fromInteger (x-3)))
    -- Cuando la bola es un número par
    | (y-1) .&. (y-2) == 0                                  = simulacionDirecta x 1 + ((2 :: Integer) `shiftL` (fromInteger (x-3-(fromIntegral (floor (logBase 2 (fromIntegral (y-1))))))))
    -- Cuando la bola menos 1 es un número cuadrado de 2
    | otherwise                                             = simulacionDirecta x (y - (2 :: Integer) `shiftL`(fromIntegral (floor (logBase 2 (fromIntegral y)))-1)) + ((2 :: Integer) `shiftL` (fromInteger(x - 3 - (fromIntegral (floor (logBase 2 (fromIntegral y)))))))
    -- En cualqiuier otro caso