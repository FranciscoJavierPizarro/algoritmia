import BinaryTree
import System.Environment
-- compilar ghc main_hs.hs
-- ejecutar ./main_hs 100000 100003

main :: IO ()
main = do
  args <- getArgs
  let resultado = simulameEsta (read (head args) :: Integer) (read (args !! 1) :: Integer) Empty
  print resultado