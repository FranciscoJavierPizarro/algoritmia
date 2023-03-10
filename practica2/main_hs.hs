import BinaryTree
import System.Environment
import System.CPUTime
import Control.Exception
import Text.Printf
-- compilar ghc main_hs.hs
-- ejecutar ./main_hs 100000 100003

main :: IO ()
main = do
  args <- getArgs
  start <- getCPUTime
  -- let resultado = simulamelaDirecta (read (head args) :: Integer) (read (args !! 1) :: Integer)
  -- let resultado = simulameEsta(read (head args) :: Integer) (read (args !! 1) :: Integer) Empty
  
  end <- getCPUTime
  print resultado
  let time = (fromIntegral (end - start) / (10^12)) * 1000 -- seconds to milliseconds
  printf "Time: %.3f ms\n" (time :: Double)