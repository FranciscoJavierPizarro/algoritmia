--------------------------------------------------------------------------------
--                                                                            --
--    Archivo: simulation.hs                                                  --
--    Fecha de última revisión: 14/03/2023                                    --
--    Autores: Francisco Javier Pizarro 821259                                --
--             Jorge Solán Morote   	816259                                  --
--    Comms:                                                                  --
--          Este archivo contiene el core de ejecución de los algoritmos      --
--          implementados en haskell pasandoles los atributos de              --
--          invocación como argumentos, además cronometra sus tiempos de      --
--          ejecución                                                         --
--    Use: ./simulation P N [0]                                               --
--          siendo P y N los enteros de profundidad y nbolas y 0 un parámetro --
--          opcional para determinar que algoritmo emplear, en caso de usarlo --
--          se emplea el algoritmo de recorrido parcial                       --
--    Compilation: ghc -O2 simulation.hs                                      --
--                                                                            --
--------------------------------------------------------------------------------

{-# LANGUAGE BangPatterns #-}
import BinaryTree
import System.Environment
import System.CPUTime
import Control.Exception
import Text.Printf
import Data.Time.Clock
import Data.Sequence
import Data.Bits
-- compilar ghc -O2 simulation.hs 
-- ejecutar time ./simulation 10000000 10000 [0]

main :: IO ()
main = do
  --extraemos los argumentos de invocación
  args <- getArgs
  let n = read (head args) :: Integer
      m = read (args !! 1) :: Integer
      algoritmo = if Prelude.length args == 3 then read (args !! 2) :: Int else -1

  --calculamos el valor del resultado midieendo los timepos
  start <- getCurrentTime
  let !resultado = if not (n==1) then (if algoritmo == 0 then simularConRecorridoParcial n m BinaryTree.Empty else simulacionDirecta n m) else 1
  end <- getCurrentTime
  let diff = (realToFrac (diffUTCTime end start) :: Double) * (10^3)

  --mostramos los resultados deseados
  --printf "resultado recorrido:%i\n" resultado
  printf "%0.9f " diff --tiempo en MS
  print resultado