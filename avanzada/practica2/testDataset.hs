import LatinSquare
import Data.List.Split (splitOn)
import System.IO
-- -- | The main function.
processPart :: String -> IO ()
processPart part = do
  let s = replaceAsterisks part -- Modify as needed
  print s
  let num = getLatinSquareDimensions s
  putStrLn $ show num
--   num <- getInt
--   let s = latinSquare_create num
  case resolver_latinSquare num s of
    [] -> do
      putStrLn "No solution."
    [h] -> do
      putStrLn "Unique solution:"
      putStr (mostrar_latinSquare h num)
    h:t -> do
      putStrLn "Non-unique solution:"
      putStr (mostrar_latinSquare h num)

filename = "test.txt"
main :: IO ()
main = do
  contents <- readFile filename
  let contents_parts = splitOn "#" contents
  mapM_ processPart contents_parts